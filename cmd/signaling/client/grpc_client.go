package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/url"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/signaling/client/service"
	"github.com/Notch-Technologies/wizy/cmd/signaling/pb/negotiation"
	"github.com/Notch-Technologies/wizy/wislog"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type ClientCaller interface {
	IsReady() bool
	Send(msg *negotiation.Body) error
	Receive(wgPubKey string, msgHandler func(msg *negotiation.Body) error) error
	WaitStreamConnected()
}

type SignalingClient struct {
	negotiationClientService service.NegotiationClientServiceCaller
	stream                   negotiation.Negotiation_ConnectStreamClient

	ctx  context.Context
	conn *grpc.ClientConn

	wislog *wislog.WisLog
}

func NewSignalingClient(
	ctx context.Context, hostString string,
	privateKey wgtypes.Key, wislog *wislog.WisLog,
) (*SignalingClient, error) {
	clientCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	option := grpc.WithTransportCredentials(insecure.NewCredentials())

	url, err := url.Parse(hostString)
	if err != nil {
		return nil, err
	}

	if url.Scheme != "http" {
		option = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{}))
	}

	conn, err := grpc.DialContext(
		clientCtx,
		url.Host,
		option,
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    10 * time.Second,
			Timeout: 10 * time.Second,
		}))
	if err != nil {
		return nil, err
	}

	return &SignalingClient{
		negotiationClientService: service.NewNegotiationClientService(ctx, conn, privateKey),
		stream:                   nil,

		ctx:  ctx,
		conn: conn,

		wislog: wislog,
	}, nil
}

func (client *SignalingClient) IsReady() bool {
	return client.conn.GetState() == connectivity.Ready || client.conn.GetState() == connectivity.Idle
}

func (client *SignalingClient) Send(msg *negotiation.Body) error {
	if !client.IsReady() {
		return fmt.Errorf("no connection server stream")
	}

	err := client.negotiationClientService.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (client *SignalingClient) Receive(
	wgPubKey string,
	msgHandler func(msg *negotiation.Body) error,
) error {
	if !client.IsReady() {
		return fmt.Errorf("no connection server stream")
	}

	err := client.negotiationClientService.Receive(wgPubKey, msgHandler)
	if err != nil {
		return err
	}

	return nil
}

func (client *SignalingClient) WaitStreamConnected() {
	client.negotiationClientService.WaitStreamConnected()
}
