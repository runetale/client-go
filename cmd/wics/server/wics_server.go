package server

import (
	"context"

	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/proto"
	"github.com/Notch-Technologies/wizy/cmd/wics/server/redis"
	"github.com/Notch-Technologies/wizy/store"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	config  *config.Config
	accountStore *store.AccountStore
	serverStore *store.ServerStore

	// grpcServer
	UserServiceServer *UserServiceServer
	PeerServiceServer *PeerServiceServer
}

func NewServer(config *config.Config, account *store.AccountStore, server *store.ServerStore, r *redis.RedisClient) (*Server, error) {
	return &Server{
		config:  config,
		accountStore: account,
		serverStore: server,

		UserServiceServer: NewUserServiceServer(r),
		PeerServiceServer: NewPeerServiceServer(r),
	}, nil
}

// TODO:(shintard) create a service for each of the gRPC Servers.
type UserServiceServer struct {
	redis *redis.RedisClient
	proto.UnimplementedUserServiceServer
}

func NewUserServiceServer(r *redis.RedisClient) *UserServiceServer {
	return &UserServiceServer{
		redis: r,
	}
}

func (uss *UserServiceServer) Login(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	panic("not implement Login")
}

type PeerServiceServer struct {
	redis *redis.RedisClient
	proto.UnimplementedPeerServiceServer
}

func NewPeerServiceServer(r *redis.RedisClient) *PeerServiceServer {
	return &PeerServiceServer{
		redis: r,
	}
}

func (pss *PeerServiceServer) WSync(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	panic("not implement WSync")
}
