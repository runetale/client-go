package service

import (
	"context"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/channel"
	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/session"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SessionServerServiceCaller interface {
	GetServerPublicKey(ctx context.Context, msg *emptypb.Empty) (*session.GetServerPublicKeyResponse, error)
	Login(ctx context.Context, msg *session.LoginRequest) (*session.LoginResponse, error)
}

type SessionServerService struct {
	config            *config.ServerConfig
	serverStore       *store.ServerStore
	db                *database.Sqlite
	peerUpdateManager *channel.PeersUpdateManager
}

func NewSessionServerService(
	db *database.Sqlite, config *config.ServerConfig,
	server *store.ServerStore, peerUpdateManager *channel.PeersUpdateManager,
) SessionServerServiceCaller {
	return &SessionServerService{
		config:            config,
		serverStore:       server,
		db:                db,
		peerUpdateManager: peerUpdateManager,
	}
}

func (s *SessionServerService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (s *SessionServerService) GetServerPublicKey(ctx context.Context, msg *emptypb.Empty) (*session.GetServerPublicKeyResponse, error) {
	pubicKey := s.serverStore.GetPublicKey()

	now := time.Now().Add(24 * time.Hour)
	secs := int64(now.Second())
	nanos := int32(now.Nanosecond())
	expiresAt := &timestamp.Timestamp{Seconds: secs, Nanos: nanos}

	return &session.GetServerPublicKeyResponse{
		Key:       pubicKey,
		ExpiresAt: expiresAt,
	}, nil
}

func (s *SessionServerService) Login(ctx context.Context, msg *session.LoginRequest) (*session.LoginResponse, error) {
	clientMachinePubKey := msg.GetClientPublicKey()
	serverMachinePubKey := msg.GetServerPublicKey()
	wgPubKey := msg.GetWgPublicKey()
	setupKey := msg.GetSetupKey()

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	sessionUsecase := usecase.NewSessionUsecase(tx, s.config, s.serverStore, s.peerUpdateManager)

	// validate setupkey
	//
	if setupKey != "" {
		err = sessionUsecase.ValidateSetupKey(setupKey)
		if err != nil {
			return nil, err
		}
	}

	// create peer
	//
	peer, err := sessionUsecase.CreatePeer(setupKey, clientMachinePubKey, serverMachinePubKey, wgPubKey)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &session.LoginResponse{
		SetupKey:        setupKey,
		ServerPublicKey: clientMachinePubKey,
		ClientPublicKey: serverMachinePubKey,
		Ip:              peer.IP,
		Cidr:            uint64(peer.CIDR),
		SignalingHost:   s.config.Signal.URL,
	}, nil
}
