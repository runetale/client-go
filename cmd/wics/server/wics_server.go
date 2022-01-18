package server

import (
	"context"
	"fmt"

	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/proto"
	"github.com/Notch-Technologies/wizy/state"
	"github.com/Notch-Technologies/wizy/store"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	config *config.Config
	account *store.Account
	privateKey wgtypes.Key

	// grpcServer
	UserServiceServer *UserServiceServer
	PeerServiceServer *PeerServiceServer
}

func NewServer(config *config.Config, account *store.Account) (*Server, error) {
	key, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return nil, err
	}

	s, err := state.NewServerPrivateKey()
	if err != nil {
		return nil, err
	}
	fmt.Println(s.MarshalText())

	return &Server{
		config: config,
		account: account,
		privateKey: key,

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

