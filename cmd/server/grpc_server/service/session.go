package service

import (
	"context"
	"log"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/session"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SessionServiceServer struct {
	config      *config.Config
	serverStore *store.ServerStore

	session.UnimplementedSessionServiceServer
}

func NewSessionServiceServer(
	db *database.Sqlite, config *config.Config,
	server *store.ServerStore,
) *SessionServiceServer {
	return &SessionServiceServer{
		config:      config,
		serverStore: server,
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

	log.Println("get server public key")

	return &session.GetServerPublicKeyResponse{
		Key:       pubicKey,
		ExpiresAt: expiresAt,
	}, nil
}
