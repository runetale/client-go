package server

import (
	"context"

	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/proto"
	"github.com/Notch-Technologies/wizy/store"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	config  *config.Config
	storeAccount *store.Account
	storeServer *store.Server

	// grpcServer
	UserServiceServer *UserServiceServer
	PeerServiceServer *PeerServiceServer
}

func NewServer(config *config.Config, account *store.Account, server *store.Server) (*Server, error) {
	return &Server{
		config:  config,
		storeAccount: account,
		storeServer: server,

		UserServiceServer: NewUserServiceServer(),
		PeerServiceServer: NewPeerServiceServer(),
	}, nil
}

// TODO:(shintard) create a service for each of the gRPC Servers.
type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
}

func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{}
}

func (uss *UserServiceServer) Login(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	panic("not implement Login")
}

type PeerServiceServer struct {
	proto.UnimplementedPeerServiceServer
}

func NewPeerServiceServer() *PeerServiceServer {
	return &PeerServiceServer{}
}

func (pss *PeerServiceServer) WSync(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	panic("not implement WSync")
}
