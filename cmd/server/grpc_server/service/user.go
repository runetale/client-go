package service

import (
	"context"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/user"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
	"github.com/Notch-Technologies/wizy/types/key"
)

type UserServiceServer struct {
	db *database.Sqlite

	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(
	db *database.Sqlite,
) *UserServiceServer {
	return &UserServiceServer{
		db: db,
	}
}

func (uss *UserServiceServer) SetupKey(ctx context.Context, msg *user.SetupKeyRequest) (*user.SetupKeyResponse, error) {
	sub := getSub(ctx)
	networkID := msg.GetNetworkID()
	userGroupID := msg.GetUserGroupID()
	job := msg.GetJob()
	orgGroupID := msg.GetOrgGroupID()
	permission := msg.GetPermission()

	tx, err := uss.db.Begin()
	if err != nil {
		return nil, err
	}

	setupKeyUsecase := usecase.NewSetupKeyUsecase(tx)
	setupKey, err := setupKeyUsecase.CreateSetupKey(uint(networkID), uint(userGroupID), job, orgGroupID, key.PermissionType(permission), sub)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &user.SetupKeyResponse{
		SetupKey: setupKey.SetupKey(),
	}, nil
}
