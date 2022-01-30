package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/url"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/wics/proto"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WicsClientManager interface {
	GetServerPublicKey() (*wgtypes.Key, error)
}

type WicsClient struct {
	privateKey        wgtypes.Key
	userServiceClient proto.UserServiceClient
	peerServiceClient proto.PeerServiceClient
	ctx               context.Context
	conn              *grpc.ClientConn
}

func NewWicsClient(ctx context.Context, url *url.URL, port int, privKey wgtypes.Key) (*WicsClient, error) {
	wicsctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	option := grpc.WithTransportCredentials(insecure.NewCredentials())

	if url.Scheme != "http" {
		option = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{}))
	}

	fmt.Println("url host")
	fmt.Println(url.Host)

	conn, err := grpc.DialContext(
		wicsctx,
		url.Host,
		option,
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    15 * time.Second,
			Timeout: 10 * time.Second,
		}))
	if err != nil {
		return nil, err
	}

	usc := proto.NewUserServiceClient(conn)
	psc := proto.NewPeerServiceClient(conn)

	return &WicsClient{
		privateKey:        privKey,
		userServiceClient: usc,
		peerServiceClient: psc,
		ctx:               ctx,
		conn:              conn,
	}, nil
}

func (wc *WicsClient) isReady() bool {
	return wc.conn.GetState() == connectivity.Ready || wc.conn.GetState() == connectivity.Idle
}

func (wc *WicsClient) GetServerPublicKey() (*wgtypes.Key, error) {
	if !wc.isReady() {
		return nil, fmt.Errorf("no connection wics server")
	}

	usCtx, cancel := context.WithTimeout(wc.ctx, 10*time.Second)
	defer cancel()

	res, err := wc.userServiceClient.GetServerPublicKey(usCtx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	pubKey, err := wgtypes.ParseKey(res.Key)
	if err != nil {
		return nil, err
	}

	return &pubKey, nil
}
