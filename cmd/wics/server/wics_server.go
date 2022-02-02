package server

import (
	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/redis"
	"github.com/Notch-Technologies/wizy/cmd/wics/server/service"
	"github.com/Notch-Technologies/wizy/store"
)

type Server struct {
	UserServiceServer *service.UserServiceServer
	PeerServiceServer *service.PeerServiceServer
}

func NewServer(
	config *config.Config, account *redis.AccountStore, 
	server *store.ServerStore, r *redis.RedisClient) (*Server, error) {
	return &Server{
		UserServiceServer: service.NewUserServiceServer(r, config, account, server),
		PeerServiceServer: service.NewPeerServiceServer(r),
	}, nil
}
