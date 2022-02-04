package service

import (
	"context"
	"fmt"

	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/pb/user"
	"github.com/Notch-Technologies/wizy/cmd/wics/redis"
	"github.com/Notch-Technologies/wizy/store"
)

type UserServiceServer struct {
	redis        *redis.RedisClient
	config       *config.Config
	accountStore *redis.AccountStore
	serverStore  *store.ServerStore

	user.UnimplementedUserServiceServer
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

func (uss *UserServiceServer) Login(ctx context.Context, msg *user.LoginMessage) (*user.LoginMessage, error) {
	clientPubKey := msg.GetClientPublicKey()
	serverPubKey := msg.GetServerPublicKey()
	setupKey := msg.GetSetupKey()

	fmt.Println(clientPubKey)
	fmt.Println(serverPubKey)
	fmt.Println(setupKey)

	//_, err := uss.accountStore.GetPeer(clientPubKey)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//
	return &user.LoginMessage{
		SetupKey:        setupKey,
		ServerPublicKey: serverPubKey,
		ClientPublicKey: clientPubKey,
	}, nil
}
