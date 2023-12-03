package gapi

import (
	db "github.com/valrichter/go-basic-bank/db/sqlc"
	"github.com/valrichter/go-basic-bank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChagedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
