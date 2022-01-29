package server

import (
	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/server/redis"
	"github.com/Notch-Technologies/wizy/store"
)

type Server struct {
	UserServiceServer *UserServiceServer
	PeerServiceServer *PeerServiceServer
}

func NewServer(config *config.Config, account *redis.AccountStore, server *store.ServerStore, r *redis.RedisClient) (*Server, error) {
	return &Server{
		UserServiceServer: NewUserServiceServer(r, config, account, server),
		PeerServiceServer: NewPeerServiceServer(r),
	}, nil
}
