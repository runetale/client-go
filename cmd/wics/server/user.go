package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/proto"
	"github.com/Notch-Technologies/wizy/cmd/wics/server/redis"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServiceServer struct {
	redis        *redis.RedisClient
	config       *config.Config
	accountStore *redis.AccountStore
	serverStore  *store.ServerStore

	proto.UnimplementedUserServiceServer
}

func NewUserServiceServer(
	r *redis.RedisClient, config *config.Config, account *redis.AccountStore,
	server *store.ServerStore,
) *UserServiceServer {
	return &UserServiceServer{
		redis:        r,
		config:       config,
		accountStore: account,
		serverStore:  server,
	}
}

// UserService
//
func (uss *UserServiceServer) Login(ctx context.Context, msg *proto.LoginMessage) (*proto.LoginMessage, error) {
	a := msg.GetPublicMachineKey()
	peer, err := uss.accountStore.GetPeer(a)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(peer)

	return nil, err
}

func (uss *UserServiceServer) GetServerPublicKey(ctx context.Context, msg *emptypb.Empty) (*proto.GetServerPublicKeyResponse, error) {
	pubicKey := uss.serverStore.GetPublicKey()

	now := time.Now().Add(24 * time.Hour)
	secs := int64(now.Second())
	nanos := int32(now.Nanosecond())
	expiresAt := &timestamp.Timestamp{Seconds: secs, Nanos: nanos}

	log.Println("get server public key")

	return &proto.GetServerPublicKeyResponse{
		Key:       pubicKey,
		ExpiresAt: expiresAt,
	}, nil
}
