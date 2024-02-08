package main

import (
	"context"
	"embed"
	"errors"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/valrichter/go-basic-bank/api"
	db "github.com/valrichter/go-basic-bank/db/sqlc"
	"github.com/valrichter/go-basic-bank/gapi"
	"github.com/valrichter/go-basic-bank/mail"
	"github.com/valrichter/go-basic-bank/pb"
	"github.com/valrichter/go-basic-bank/util"
	"github.com/valrichter/go-basic-bank/worker"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config:")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db:")
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(connPool)

	redisOtp := asynq.RedisClientOpt{
		Addr: config.RedisAdress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOtp)

	waitGroup, ctx := errgroup.WithContext(ctx)

	runTaskProcessor(ctx, waitGroup, config, redisOtp, store)
	runGatewayServer(ctx, waitGroup, config, store, taskDistributor)
	go runGinServer(config, store)
	runGRPCServer(ctx, waitGroup, config, store, taskDistributor)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")

	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance:")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("fialed to run migrate up:")
	}

	log.Info().Msg("db migrated successfully")
}

func runTaskProcessor(
	ctx context.Context, waitGroup *errgroup.Group,
	config util.Config, redisOpt asynq.RedisClientOpt, store db.Store) {

	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)

	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor:")
	}

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("grafull shutdown task processor")
		taskProcessor.Shutdown()
		log.Info().Msg("task processor is stoped")
		return nil
	})

}

func runGRPCServer(
	ctx context.Context, waitGroup *errgroup.Group,
	config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server:")
	}

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterBasicBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener:")
	}

	waitGroup.Go(func() error {
		log.Printf("start gRPC server on %s", listener.Addr().String())

		err = grpcServer.Serve(listener)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Error().Err(err).Msg("failed to serve gRPC server")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("gracefull shutting down gRPC server")

		grpcServer.GracefulStop()
		log.Info().Msg("gRPC server gracefully stopped")
		return nil
	})

}

var (
	//go:embed docs/rpc_swagger
	swagger embed.FS
	//TODO: fix swaggerFS se guarda en la memoeria loca el primer estado registrado, por lo que no se actualiza cuando se modifica la documentacion
	swaggerFS, _ = fs.Sub(swagger, "docs/rpc_swagger")
)

func runGatewayServer(
	ctx context.Context, waitGroup *errgroup.Group,
	config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server:")
	}

	grpcMux := runtime.NewServeMux()

	err = pb.RegisterBasicBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handler server:")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	swaggerFileServer := http.FileServer(http.FS(swaggerFS))
	swaggerHandler := http.StripPrefix("/swagger/", swaggerFileServer)
	mux.Handle("/swagger/", swaggerHandler)

	// fs.WalkDir(swaggerFS, ".", func(path string, d fs.DirEntry, err error) error {
	// 	fmt.Println(path)
	// 	return nil
	// })

	httpServer := &http.Server{
		Handler: gapi.HttpLogger(mux),
		Addr:    config.HTTPGatewayAddress,
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start HTTP gateway server on %s", httpServer.Addr)

		err = httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Fatal().Err(err).Msg("HTTP Gatewat server failed to start")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("gracefull shutting down HTTP gateway server")

		if err = httpServer.Shutdown(context.Background()); err != nil {
			log.Error().Err(err).Msg("failded to shutdown HTTP gateway server:")
			return err
		}

		log.Info().Msg("HTTP gateway server is stopped")
		return nil
	})

}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server:")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server:")
	}
}
