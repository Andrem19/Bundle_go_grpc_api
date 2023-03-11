package gapi

import (
	"context"
	"database/sql"

	db "github.com/Andrem19/bundle_go_grpc_api/db/sqlc"
	"github.com/Andrem19/bundle_go_grpc_api/pb"
	"github.com/Andrem19/bundle_go_grpc_api/helpers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server * Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserRsponse, error) {
	user, err := server.store.GetUser(ctx, req.GetEmail())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}

	err = helpers.CheckPassword(req.GetPassword(), user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect password")
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to create access token")
	}
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to create refresh token")
	}

	mtdt := server.extractMetadata(ctx)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           accessPayload.ID,
		Username:     user.Username,
		Token: 		  accessToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    accessPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to create session")
	}

	rsp := &pb.LoginUserRsponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		TokenExpiresAt:  	   timestamppb.New(accessPayload.ExpiredAt),
	}
	return rsp, nil

}