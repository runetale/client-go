package server

import (
	"context"
	"fmt"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/protobuf/types/known/emptypb"
	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/proto"
	"github.com/Notch-Technologies/wizy/store"
)

type Server struct {
	config *config.Config
	account *store.Account
	privateKey wgtypes.Key

	proto.UnimplementedPeerServiceServer
	proto.UnimplementedUserServiceServer
}

func NewServer(config *config.Config, account *store.Account) (*Server, error) {
	key, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &Server{
		config: config,
		account: account,
		privateKey: key,

	}, nil
}

func (userServer *Server) Login(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	panic("not implement Login")
}

func (peerServer *Server) WSync(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	panic("not implement WSync")
}

