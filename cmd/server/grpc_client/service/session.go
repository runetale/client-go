package service

import (
	"context"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/pb/session"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SessionServiceClientCaller interface {
	GetServerPublicKey() (string, error)
	Login(setupKey, clientPubKey, serverPubKey, ip string, wgPublicKey string) (*session.LoginMessage, error)
}

type SessionServiceClient struct {
	sessionServiceClient session.SessionServiceClient
	privateKey           wgtypes.Key

	ctx context.Context
}

func NewSessionServiceClient(ctx context.Context, conn *grpc.ClientConn, privateKey wgtypes.Key) *SessionServiceClient {
	return &SessionServiceClient{
		sessionServiceClient: session.NewSessionServiceClient(conn),
		privateKey:           privateKey,

		ctx: ctx,
	}
}

func (s *SessionServiceClient) GetServerPublicKey() (string, error) {
	usCtx, cancel := context.WithTimeout(s.ctx, 5*time.Second)
	defer cancel()

	res, err := s.sessionServiceClient.GetServerPublicKey(usCtx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}

	return res.Key, nil
}

func (s *SessionServiceClient) Login(setupKey, clientPubKey, serverPubKey, ip string, wgPublicKey string) (*session.LoginMessage, error) {
	usCtx, cancel := context.WithTimeout(s.ctx, 5*time.Second)
	defer cancel()

	msg, err := s.sessionServiceClient.Login(usCtx, &session.LoginMessage{
		SetupKey:        setupKey,
		ClientPublicKey: clientPubKey,
		ServerPublicKey: serverPubKey,
		WgPublicKey:     wgPublicKey,
		Ip:              ip,
	})
	if err != nil {
		return nil, err
	}

	return msg, nil
}
