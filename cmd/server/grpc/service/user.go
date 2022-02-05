package service

import (
	"context"

	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/user"
	"github.com/Notch-Technologies/wizy/cmd/server/redis"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
	"github.com/Notch-Technologies/wizy/store"
)

type UserServiceServer struct {
	redis              *redis.RedisClient
	config             *config.Config
	accountStore       *redis.AccountStore
	serverStore        *store.ServerStore
	setupKeyRepository *repository.SetupKeyRepository

	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(
	r *redis.RedisClient, config *config.Config, account *redis.AccountStore,
	server *store.ServerStore, user *redis.UserStore, network *redis.NetworkStore,
	group *redis.OrgGroupStore, setupKey *redis.SetupKeyStore,
) *UserServiceServer {
	setupKeyRepository := repository.NewSetupKeyRepository(r, config, account, server, user, network, group, setupKey)
	return &UserServiceServer{
		redis:              r,
		config:             config,
		accountStore:       account,
		serverStore:        server,
		setupKeyRepository: setupKeyRepository,
	}
}

func (uss *UserServiceServer) SetupKey(ctx context.Context, msg *user.SetupKeyMessage) (*user.SetupKeyMessage, error) {
	return &user.SetupKeyMessage{}, nil
}
