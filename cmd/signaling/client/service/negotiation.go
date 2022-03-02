package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/signaling/key"
	"github.com/Notch-Technologies/wizy/cmd/signaling/pb/negotiation"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type NegotiationClientServiceCaller interface {
	Send(msg *negotiation.Body) error
	ConnectStream(wgPubKey string) (negotiation.Negotiation_ConnectStreamClient, error)
}

type NegotiationClientService struct {
	negotiationClient negotiation.NegotiationClient
	stream            negotiation.Negotiation_ConnectStreamClient

	ctx context.Context
}

func NewNegotiationClientService(
	ctx context.Context, conn *grpc.ClientConn,
	privateKey wgtypes.Key,
) *NegotiationClientService {
	return &NegotiationClientService{
		negotiationClient: negotiation.NewNegotiationClient(conn),
		stream:            nil,

		ctx: ctx,
	}
}

func (n *NegotiationClientService) Send(msg *negotiation.Body) error {
	ctx, cancel := context.WithTimeout(n.ctx, 5*time.Second)
	defer cancel()

	_, err := n.negotiationClient.Send(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}

func (n *NegotiationClientService) ConnectStream(wgPubKey string) (negotiation.Negotiation_ConnectStreamClient, error) {
	n.stream = nil

	md := metadata.New(map[string]string{key.WgPubKey: wgPubKey})
	ctx := metadata.NewOutgoingContext(n.ctx, md)

	stream, err := n.negotiationClient.ConnectStream(ctx, grpc.WaitForReady(true))
	n.stream = stream
	if err != nil {
		return nil, err
	}

	header, err := n.stream.Header()
	if err != nil {
		return nil, err
	}

	registered := header.Get(key.HeaderRegisterd)
	if len(registered) == 0 {
		return nil, fmt.Errorf("didn't receive a registration header from the Signal server whille connecting to the streams")
	}

	return stream, nil
}
