package service

import (
	"context"
	"fmt"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/user"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
)

type UserServiceServer struct {
	setupKeyUsecase *usecase.SetupKeyUsecase

	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(
	db *database.Sqlite,
) *UserServiceServer {
	return &UserServiceServer{
		setupKeyUsecase: usecase.NewSetupKeyUsecase(db),
	}
}

func (uss *UserServiceServer) SetupKey(ctx context.Context, msg *user.SetupKeyMessage) (*user.SetupKeyMessage, error) {
	sub := getSub(ctx)
	fmt.Println(sub)

	return &user.SetupKeyMessage{}, nil
}

