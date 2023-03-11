package gapi

import (
	"fmt"

	db "github.com/Andrem19/bundle_go_grpc_api/db/sqlc"
	"github.com/Andrem19/bundle_go_grpc_api/pb"
	"github.com/Andrem19/bundle_go_grpc_api/token"
	"github.com/Andrem19/bundle_go_grpc_api/helpers"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedMyBundleProjServer
	config          helpers.Config
	store           db.Store
	tokenMaker      token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config helpers.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
	}

	return server, nil
}