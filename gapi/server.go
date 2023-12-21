package gapi

import (
	"fmt"

	db "github.com/valrichter/go-basic-bank/db/sqlc"
	"github.com/valrichter/go-basic-bank/pb"
	"github.com/valrichter/go-basic-bank/token"
	"github.com/valrichter/go-basic-bank/util"
	"github.com/valrichter/go-basic-bank/worker"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedBasicBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
