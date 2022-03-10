package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/client/service"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/session"
	"github.com/Notch-Technologies/wizy/wislog"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type Status string

const StreamConnected Status = "Connected"
const StreamDisconnected Status = "Disconnected"

type ClientCaller interface {
	IsReady() bool
	GetServerPublicKey() (string, error)
	Login(setupKey, clientPubKey, serverPubKey string, wgPrivateKey wgtypes.Key) (*session.LoginResponse, error)
	Sync(clientMachineKey string, msgHandler func(msg *peer.SyncResponse) error) error
}

type GrpcClient struct {
	peerClientService    service.PeerClientServiceCaller
	sessionClientService service.SessionClientServiceCaller

	ctx  context.Context
	conn *grpc.ClientConn
	mux  sync.Mutex

	status Status

	wislog *wislog.WisLog
}

func NewGrpcClient(
	ctx context.Context, url *url.URL,
	privateKey wgtypes.Key, wislog *wislog.WisLog,
) (*GrpcClient, error) {
	clientCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	option := grpc.WithTransportCredentials(insecure.NewCredentials())

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

	return &GrpcClient{
		peerClientService:    service.NewPeerClientService(ctx, conn, privateKey),
		sessionClientService: service.NewSessionClientService(ctx, conn, privateKey),

		ctx:    ctx,
		conn:   conn,
		mux:    sync.Mutex{},
		status: StreamDisconnected,

		wislog: wislog,
	}, nil
}

func (client *GrpcClient) IsReady() bool {
	return client.conn.GetState() == connectivity.Ready || client.conn.GetState() == connectivity.Idle
}

func (client *GrpcClient) GetServerPublicKey() (string, error) {
	if !client.IsReady() {
		return "", fmt.Errorf("no connection grpc server")
	}

	key, err := client.sessionClientService.GetServerPublicKey()
	if err != nil {
		return "", err
	}

	return key, nil
}

func (client *GrpcClient) Login(
	setupKey, clientPubKey, serverPubKey string,
	wgPrivateKey wgtypes.Key,
) (*session.LoginResponse, error) {
	if !client.IsReady() {
		return nil, fmt.Errorf("no connection grpc server")
	}

	msg, err := client.sessionClientService.Login(
		setupKey, clientPubKey, serverPubKey,
		wgPrivateKey.PublicKey().String(),
	)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (client *GrpcClient) Sync(clientMachineKey string, msgHandler func(msg *peer.SyncResponse) error) error {
	err := client.peerClientService.Sync(clientMachineKey, msgHandler)
	if err != nil {
		return err
	}
	return nil
}
