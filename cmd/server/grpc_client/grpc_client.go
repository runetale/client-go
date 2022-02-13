package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/url"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/session"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/user"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClientManager interface {
	GetServerPublicKey() (*wgtypes.Key, error)
	Login(setupKey, clientPubKey, serverPubKey string) (string, error)
}

type GrpcClient struct {
	privateKey        wgtypes.Key
	userServiceClient user.UserServiceClient
	peerServiceClient peer.PeerServiceClient
	sessionServiceClient session.SessionServiceClient
	ctx               context.Context
	conn              *grpc.ClientConn
}

func NewGrpcClient(ctx context.Context, url *url.URL, port int, privKey wgtypes.Key) (*GrpcClient, error) {
	clientCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
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
			Time:    15 * time.Second,
			Timeout: 10 * time.Second,
		}))
	if err != nil {
		return nil, err
	}

	usc := user.NewUserServiceClient(conn)
	psc := peer.NewPeerServiceClient(conn)
	sec := session.NewSessionServiceClient(conn)

	return &GrpcClient{
		privateKey:        privKey,
		userServiceClient: usc,
		peerServiceClient: psc,
		sessionServiceClient: sec,
		ctx:               ctx,
		conn:              conn,
	}, nil
}

func (client *GrpcClient) isReady() bool {
	return client.conn.GetState() == connectivity.Ready || client.conn.GetState() == connectivity.Idle
}

func (wc *GrpcClient) GetServerPublicKey() (string, error) {
	if !wc.isReady() {
		return "", fmt.Errorf("no connection wics server")
	}

	usCtx, cancel := context.WithTimeout(wc.ctx, 10*time.Second)
	defer cancel()

	res, err := wc.sessionServiceClient.GetServerPublicKey(usCtx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}

	pubKey, err := wgtypes.ParseKey(res.Key)
	if err != nil {
		return "", err
	}

	return pubKey.PublicKey().String(), nil
}

func (client *GrpcClient) Login(setupKey, clientPubKey, serverPubKey string) (*session.LoginMessage, error) {
	if !client.isReady() {
		return nil, fmt.Errorf("no connection wics server")
	}

	usCtx, cancel := context.WithTimeout(client.ctx, 10*time.Second)
	defer cancel()

	msg, err := client.sessionServiceClient.Login(usCtx, &session.LoginMessage{
		SetupKey:        setupKey,
		ClientPublicKey: clientPubKey,
		ServerPublicKey: serverPubKey,
	})
	if err != nil {
		return nil, err
	}

	return msg, nil
}
