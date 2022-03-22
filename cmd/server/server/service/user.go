package service

import (
	"context"

	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/user"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
)

type UserServerServiceCaller interface {
	CreateSetupKey(ctx context.Context, msg *user.CreateSetupKeyRequest) (*user.CreateSetupKeyResponse, error)
}

type UserServerService struct {
	db     *database.Sqlite
	config *config.ServerConfig
}

func NewUserServerService(
	db *database.Sqlite, config *config.ServerConfig,
) UserServerServiceCaller {
	return &UserServerService{
		db:     db,
		config: config,
	}
}

func (u *UserServerService) CreateSetupKey(ctx context.Context, msg *user.CreateSetupKeyRequest) (*user.CreateSetupKeyResponse, error) {
	sub := getSub(ctx)

	userGroupID := msg.GetUserGroupID()
	orgGroupID := msg.GetOrgGroupID()
	networkID := msg.GetNetworkID()
	jobID := msg.GetJobID()
	roleID := msg.GetRoleID()
	email := msg.GetEmail()

	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	setupKeyUsecase := usecase.NewSetupKeyUsecase(tx, u.config)
	setupKey, err := setupKeyUsecase.CreateSetupKey(uint(networkID), uint(userGroupID), uint(jobID), uint(roleID), orgGroupID, sub, email)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &user.CreateSetupKeyResponse{
		SetupKey: setupKey.SetupKey(),
	}, nil
}
