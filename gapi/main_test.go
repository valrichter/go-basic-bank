package gapi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	db "github.com/valrichter/go-basic-bank/db/sqlc"
	"github.com/valrichter/go-basic-bank/util"
	"github.com/valrichter/go-basic-bank/worker"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)
	return server
}
