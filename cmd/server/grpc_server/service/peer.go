package service

import (
	"context"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PeerServiceServer struct {
	peer.UnimplementedPeerServiceServer
}

func NewPeerServiceServer(db *database.Sqlite) *PeerServiceServer {
	return &PeerServiceServer{}
}

func (pss *PeerServiceServer) WSync(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	panic("not implement WSync")
}
