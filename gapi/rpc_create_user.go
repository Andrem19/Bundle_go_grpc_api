package gapi

import (
	"context"

	db "github.com/Andrem19/bundle_go_grpc_api/db/sqlc"
	"github.com/Andrem19/bundle_go_grpc_api/helpers"
	"github.com/Andrem19/bundle_go_grpc_api/pb"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserRsponse, error) {
	hashedPassword, err := helpers.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "failed to hash password")
	}

	arg := db.CreateUserParams{
		Username: req.GetUsername(),
		FirstName: req.GetFirstName(),
		SecondName: req.GetSecondName(),
		Email: req.GetEmail(),
		HashedPassword: hashedPassword,
		Roles: req.GetRoles(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exist %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user %s", err)
	}

	rsp := &pb.CreateUserRsponse{
		User: convertUser(user),
	}

	return rsp, nil
}