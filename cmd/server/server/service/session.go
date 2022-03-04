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

type SessionServiceServerCaller interface {
	GetServerPublicKey(ctx context.Context, msg *emptypb.Empty) (*session.GetServerPublicKeyResponse, error)
	Login(ctx context.Context, msg *session.LoginRequest) (*session.LoginResponse, error)
}

type SessionServerService struct {
	config            *config.Config
	serverStore       *store.ServerStore
	db                *database.Sqlite
	peerUpdateManager *channel.PeersUpdateManager
}

func NewSessionServerService(
	db *database.Sqlite, config *config.Config,
	server *store.ServerStore,
	peerUpdateManager *channel.PeersUpdateManager,
) *SessionServerService {
	return &SessionServerService{
		config:            config,
		serverStore:       server,
		db:                db,
		peerUpdateManager: peerUpdateManager,
	}
}

func (sss *SessionServerService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (sss *SessionServerService) GetServerPublicKey(ctx context.Context, msg *emptypb.Empty) (*session.GetServerPublicKeyResponse, error) {
	pubicKey := sss.serverStore.GetPublicKey()

	now := time.Now().Add(24 * time.Hour)
	secs := int64(now.Second())
	nanos := int32(now.Nanosecond())
	expiresAt := &timestamp.Timestamp{Seconds: secs, Nanos: nanos}

	return &session.GetServerPublicKeyResponse{
		Key:       pubicKey,
		ExpiresAt: expiresAt,
	}, nil
}

func (sss *SessionServerService) Login(ctx context.Context, msg *session.LoginRequest) (*session.LoginResponse, error) {
	clientMachinePubKey := msg.GetClientPublicKey()
	serverMachinePubKey := msg.GetServerPublicKey()
	wgPubKey := msg.GetWgPublicKey()
	setupKey := msg.GetSetupKey()

	tx, err := sss.db.Begin()
	if err != nil {
		return nil, err
	}

	sessionUsecase := usecase.NewSessionUsecase(tx, sss.serverStore, sss.peerUpdateManager)

	_, err = sessionUsecase.CreatePeer(setupKey, clientMachinePubKey, serverMachinePubKey, wgPubKey)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &session.LoginResponse{
		SetupKey:        setupKey,
		ServerPublicKey: clientMachinePubKey,
		ClientPublicKey: serverMachinePubKey,
	}, nil
}
