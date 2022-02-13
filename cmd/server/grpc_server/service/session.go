package service

import (
	"context"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/session"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SessionServiceServer struct {
	config      *config.Config
	serverStore *store.ServerStore
	db          *database.Sqlite

	session.UnimplementedSessionServiceServer
}

func NewSessionServiceServer(
	db *database.Sqlite, config *config.Config,
	server *store.ServerStore,
) *SessionServiceServer {
	return &SessionServiceServer{
		config:      config,
		serverStore: server,
		db:          db,
	}
}

func (uss *SessionServiceServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (uss *SessionServiceServer) GetServerPublicKey(ctx context.Context, msg *emptypb.Empty) (*session.GetServerPublicKeyResponse, error) {
	pubicKey := uss.serverStore.GetPublicKey()

	now := time.Now().Add(24 * time.Hour)
	secs := int64(now.Second())
	nanos := int32(now.Nanosecond())
	expiresAt := &timestamp.Timestamp{Seconds: secs, Nanos: nanos}

	return &session.GetServerPublicKeyResponse{
		Key:       pubicKey,
		ExpiresAt: expiresAt,
	}, nil
}

func (uss *SessionServiceServer) Login(ctx context.Context, msg *session.LoginMessage) (*session.LoginMessage, error) {
	clientPubKey := msg.GetClientPublicKey()
	serverPubKey := msg.GetServerPublicKey()
	setupKey := msg.GetSetupKey()

	sessionUsecase := usecase.NewSessionUsecase(uss.db, uss.serverStore)

	_, err := sessionUsecase.CreatePeer(setupKey, clientPubKey, serverPubKey)
	if err != nil {
		return nil, err
	}

	return &session.LoginMessage{
		SetupKey:        setupKey,
		ServerPublicKey: serverPubKey,
		ClientPublicKey: clientPubKey,
	}, nil
}
