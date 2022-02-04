package service

import (
	"context"
	"log"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/session"
	"github.com/Notch-Technologies/wizy/cmd/server/redis"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SessionServiceServer struct {
	redis        *redis.RedisClient
	config       *config.Config
	accountStore *redis.AccountStore
	serverStore  *store.ServerStore

	session.UnimplementedSessionServiceServer
}

func NewSessionServiceServer(
	r *redis.RedisClient, config *config.Config, account *redis.AccountStore,
	server *store.ServerStore,
) *SessionServiceServer {
	return &SessionServiceServer{
		redis:        r,
		config:       config,
		accountStore: account,
		serverStore:  server,
	}
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
