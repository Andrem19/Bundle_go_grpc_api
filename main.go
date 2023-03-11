package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	db "github.com/Andrem19/bundle_go_grpc_api/db/sqlc"
	"github.com/Andrem19/bundle_go_grpc_api/gapi"
	"github.com/Andrem19/bundle_go_grpc_api/helpers"
	"github.com/Andrem19/bundle_go_grpc_api/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := helpers.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	go runGatewayServer(config, store)
	runGrpcServer(config, store)
}

func runGrpcServer(config helpers.Config, store db.Store) {

	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMyBundleProjServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener", err)
	}

	log.Printf("start fRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:", err)
	}
}

func runGatewayServer(config helpers.Config, store db.Store) {

	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcMux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterMyBundleProjHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatal("cannot create listener", err)
	}

	log.Printf("start HTTP server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start HTTP gateway server", err)
	}
}