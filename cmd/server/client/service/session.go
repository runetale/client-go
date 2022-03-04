package service

import (
	"context"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/pb/session"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SessionClientServiceCaller interface {
	GetServerPublicKey() (string, error)
	Login(setupKey, clientPubKey, serverPubKey, wgPublicKey string) (*session.LoginResponse, error)
}

type SessionClientService struct {
	sessionClientService session.SessionServiceClient
	privateKey           wgtypes.Key

	ctx context.Context
}

func NewSessionClientService(ctx context.Context, conn *grpc.ClientConn, privateKey wgtypes.Key) *SessionClientService {
	return &SessionClientService{
		sessionClientService: session.NewSessionServiceClient(conn),
		privateKey:           privateKey,

		ctx: ctx,
	}
}

func (s *SessionClientService) GetServerPublicKey() (string, error) {
	usCtx, cancel := context.WithTimeout(s.ctx, 5*time.Second)
	defer cancel()

	res, err := s.sessionClientService.GetServerPublicKey(usCtx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}

	return res.Key, nil
}

func (s *SessionClientService) Login(
	setupKey, clientPubKey string,
	serverPubKey, wgPublicKey string,
) (*session.LoginResponse, error) {
	usCtx, cancel := context.WithTimeout(s.ctx, 5*time.Second)
	defer cancel()

	msg, err := s.sessionClientService.Login(usCtx, &session.LoginRequest{
		SetupKey:        setupKey,
		ClientPublicKey: clientPubKey,
		ServerPublicKey: serverPubKey,
		WgPublicKey:     wgPublicKey,
	})
	if err != nil {
		return nil, err
	}

	return msg, nil
}
