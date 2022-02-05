package server

import (
	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/grpc/service"
	"github.com/Notch-Technologies/wizy/cmd/server/redis"
	"github.com/Notch-Technologies/wizy/store"
)

type Server struct {
	UserServiceServer    *service.UserServiceServer
	PeerServiceServer    *service.PeerServiceServer
	SessionServiceServer *service.SessionServiceServer
}

func NewServer(
	config *config.Config, account *redis.AccountStore,
	server *store.ServerStore, r *redis.RedisClient,
	user *redis.UserStore, network *redis.NetworkStore,
	group *redis.OrgGroupStore, setupKey *redis.SetupKeyStore,
) (*Server, error) {
	return &Server{
		UserServiceServer:    service.NewUserServiceServer(r, config, account, server, user, network, group, setupKey),
		PeerServiceServer:    service.NewPeerServiceServer(r),
		SessionServiceServer: service.NewSessionServiceServer(r, config, account, server),
	}, nil
}
