package service

import (
	"context"
	"io"

	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
)

type PeerServiceClientCaller interface {
	Sync(clientMachineKey string, msgHandler func(msg *peer.SyncResponse) error) error
}

type PeerServiceClient struct {
	peerServiceClient peer.PeerServiceClient
	privateKey        wgtypes.Key

	ctx context.Context
}

func NewPeerServiceClient(ctx context.Context, conn *grpc.ClientConn, privateKey wgtypes.Key) *PeerServiceClient {
	return &PeerServiceClient{
		peerServiceClient: peer.NewPeerServiceClient(conn),
		privateKey:        privateKey,

		ctx: ctx,
	}
}

func (p *PeerServiceClient) Sync(
	clientMachineKey string,
	msgHandler func(msg *peer.SyncResponse) error,
) error {
	stream, err := p.peerServiceClient.Sync(p.ctx, &peer.SyncMessage{
		PrivateKey:       p.privateKey.String(),
		ClientMachineKey: clientMachineKey,
	})
	if err != nil {
		return err
	}

	for {
		update, err := stream.Recv()
		if err == io.EOF {
			return err
		}

		if err != nil {
			return err
		}

		err = msgHandler(update)
		if err != nil {
			return err
		}
	}
}
