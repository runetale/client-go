package service

import (
	"context"

	"github.com/Notch-Technologies/wizy/cmd/wics/proto"
	"github.com/Notch-Technologies/wizy/cmd/wics/server/redis"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PeerServiceServer struct {
	redis *redis.RedisClient
	proto.UnimplementedPeerServiceServer
}

func NewPeerServiceServer(r *redis.RedisClient) *PeerServiceServer {
	return &PeerServiceServer{
		redis: r,
	}
}

func (pss *PeerServiceServer) WSync(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	panic("not implement WSync")
}
