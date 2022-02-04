package service

import (
	"context"

	"github.com/Notch-Technologies/wizy/cmd/wics/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/wics/redis"
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
