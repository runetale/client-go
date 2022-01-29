package server

import (
	"context"
	"fmt"

	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/proto"
	"github.com/Notch-Technologies/wizy/cmd/wics/server/redis"
	"github.com/Notch-Technologies/wizy/store"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	// grpcServer
	UserServiceServer *UserServiceServer
	PeerServiceServer *PeerServiceServer
}

func NewServer(config *config.Config, account *redis.AccountStore, server *store.ServerStore, r *redis.RedisClient) (*Server, error) {
	return &Server{
		UserServiceServer: NewUserServiceServer(r, config, account, server),
		PeerServiceServer: NewPeerServiceServer(r),
	}, nil
}

// TODO:(shintard) create a service for each of the gRPC Servers.
type UserServiceServer struct {
	redis *redis.RedisClient
	config       *config.Config
	accountStore *redis.AccountStore
	serverStore  *store.ServerStore
	
	proto.UnimplementedUserServiceServer
}

func NewUserServiceServer(
	r *redis.RedisClient, config *config.Config, account *redis.AccountStore,
	server *store.ServerStore,
) *UserServiceServer {
	return &UserServiceServer{
		redis: r,
		config: config,
		accountStore: account,
		serverStore: server,
	}
}

// UserService
//
func (uss *UserServiceServer) Login(ctx context.Context, msg *proto.LoginMessage) (*proto.LoginMessage, error) {
	a := msg.GetPublicMachineKey()
	peer, err := uss.accountStore.GetPeer(a)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(peer)

	return nil, err
}

// PeerService
//
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
