package gapi

import (
	db "github.com/Andrem19/bundle_go_grpc_api/db/sqlc"
	"github.com/Andrem19/bundle_go_grpc_api/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FirstName: user.FirstName,
		SecondName: user.SecondName,
		Email:             user.Email,
		Roles: user.Roles,
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}