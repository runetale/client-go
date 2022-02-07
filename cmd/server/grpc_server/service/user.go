package service

import (
	"context"
	"fmt"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/user"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
	"github.com/Notch-Technologies/wizy/types/key"
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
	networkName := msg.GetNetworkID()
	userGroupName := msg.GetUserGroupID()
	job := msg.GetJob()
	orgGroupID := msg.GetOrgGroupID()
	permission := msg.GetPermission()

	setupKey, err := uss.setupKeyUsecase.CreateSetupKey(networkName, userGroupName, job, orgGroupID, key.PermissionType(permission), sub)
	if err != nil {
		return nil, err
	}
	fmt.Println(setupKey)

	return &user.SetupKeyMessage{}, nil
}

