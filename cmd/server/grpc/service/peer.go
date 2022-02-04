package service

import (
	"context"
	"fmt"

	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/redis"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PeerServiceServer struct {
	redis *redis.RedisClient
	peer.UnimplementedPeerServiceServer
}

func NewPeerServiceServer(r *redis.RedisClient) *PeerServiceServer {
	return &PeerServiceServer{
		redis: r,
	}
}

func (pss *PeerServiceServer) WSync(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	panic("not implement WSync")
}

func (uss *PeerServiceServer) Login(ctx context.Context, msg *peer.PeerLoginMessage) (*peer.PeerLoginMessage, error) {
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
	return &peer.PeerLoginMessage{
		SetupKey:        setupKey,
		ServerPublicKey: serverPubKey,
		ClientPublicKey: clientPubKey,
	}, nil
}
