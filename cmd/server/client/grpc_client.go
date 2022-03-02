package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/url"
	"sync"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/client/service"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/negotiation"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/session"
	server "github.com/Notch-Technologies/wizy/cmd/server/server/service"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Status string

const StreamConnected Status = "Connected"
const StreamDisconnected Status = "Disconnected"

type ClientCaller interface {
	GetServerPublicKey() (string, error)
	IsReady() bool
	Login(setupKey, clientPubKey, serverPubKey, ip string, wgPrivateKey wgtypes.Key) (*session.LoginMessage, error)
	ConnectStream(wgPubKey string) (negotiation.Negotiation_ConnectStreamClient, error)
	Receive(wgPubKey string, msgHandler func(msg *negotiation.Body) error) error
	Sync(clientMachineKey string, msgHandler func(msg *peer.SyncResponse) error) error
	WaitStreamConnected()
	Send(msg *negotiation.Body) error
	GetStatus() Status
	StreamConnected() bool
}

type GrpcClient struct {
	peerClientService    service.PeerClientServiceCaller
	sessionClientService service.SessionClientServiceCaller
	negotiationClient    negotiation.NegotiationClient
	stream               negotiation.Negotiation_ConnectStreamClient

	ctx  context.Context
	conn *grpc.ClientConn
	mux  sync.Mutex

	connectedCh chan struct{}

	status Status
}

func NewGrpcClient(
	ctx context.Context, url *url.URL, port int,
	privateKey wgtypes.Key,
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
			Time:    15 * time.Second,
			Timeout: 10 * time.Second,
		}))
	if err != nil {
		return nil, err
	}

	nc := negotiation.NewNegotiationClient(conn)

	return &GrpcClient{
		peerClientService:    service.NewPeerClientService(ctx, conn, privateKey),
		sessionClientService: service.NewSessionClientService(ctx, conn, privateKey),
		negotiationClient:    nc,

		ctx:    ctx,
		conn:   conn,
		mux:    sync.Mutex{},
		status: StreamDisconnected,
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
	setupKey, clientPubKey, serverPubKey, ip string,
	wgPrivateKey wgtypes.Key,
) (*session.LoginMessage, error) {
	if !client.IsReady() {
		return nil, fmt.Errorf("no connection grpc server")
	}

	msg, err := client.sessionClientService.Login(
		setupKey, clientPubKey, serverPubKey, ip,
		wgPrivateKey.PublicKey().String(),
	)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (client *GrpcClient) ConnectStream(wgPubKey string) (negotiation.Negotiation_ConnectStreamClient, error) {
	client.stream = nil

	md := metadata.New(map[string]string{server.WgPubKey: wgPubKey})
	ctx := metadata.NewOutgoingContext(client.ctx, md)

	stream, err := client.negotiationClient.ConnectStream(ctx, grpc.WaitForReady(true))
	client.stream = stream
	if err != nil {
		return nil, err
	}

	header, err := client.stream.Header()
	if err != nil {
		return nil, err
	}

	registered := header.Get(server.HeaderRegisterd)
	if len(registered) == 0 {
		return nil, fmt.Errorf("didn't receive a registration header from the Signal server whille connecting to the streams")
	}

	return stream, nil
}

func (client *GrpcClient) Receive(
	wgPubKey string,
	msgHandler func(msg *negotiation.Body) error,
) error {
	client.notifyStreamDisconnected()

	if !client.IsReady() {
		return fmt.Errorf("no conection grpc client")
	}

	stream, err := client.ConnectStream(wgPubKey)
	if err != nil {
		return err
	}

	client.notifyStreamConnected()

	for {
		msg, err := stream.Recv()
		if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
			fmt.Println("stream canceled (usually indicates shutdown)")
			return err
		} else if s.Code() == codes.Unavailable {
			fmt.Println("Signal Service is unavailable")
			return err
		} else if err == io.EOF {
			fmt.Println("Signal Service stream closed by server")
			return err
		} else if err != nil {
			return err
		}

		fmt.Printf("received a new message from Peer [fingerprint: %s]\n", msg.ClientMachineKey)

		// CreatePeerの時に他のPeerのClientMachineKeyは送信されているのは確認できた
		err = msgHandler(msg)

		if err != nil {
			fmt.Printf("error while handling message of Peer [key: %s] error: [%s]\n", msg.ClientMachineKey, err.Error())
			return err
		}
	}
}

func (client *GrpcClient) Sync(clientMachineKey string, msgHandler func(msg *peer.SyncResponse) error) error {
	err := client.peerClientService.Sync(clientMachineKey, msgHandler)
	if err != nil {
		return err
	}
	return nil
}

func (client *GrpcClient) WaitStreamConnected() {
	if client.status == StreamConnected {
		fmt.Println("Stream Connected")
		return
	}

	ch := client.getStreamStatusChan()
	select {
	case <-client.ctx.Done():
	case <-ch:
	}
}

func (client *GrpcClient) StreamConnected() bool {
	return client.status == StreamConnected
}

func (client *GrpcClient) GetStatus() Status {
	return client.status
}

func (client *GrpcClient) Send(msg *negotiation.Body) error {
	if !client.IsReady() {
		return fmt.Errorf("no connection server stream")
	}

	ctx, cancel := context.WithTimeout(client.ctx, 5*time.Second)
	defer cancel()

	_, err := client.negotiationClient.Send(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}

func (client *GrpcClient) getStreamStatusChan() <-chan struct{} {
	client.mux.Lock()
	defer client.mux.Unlock()

	if client.connectedCh == nil {
		client.connectedCh = make(chan struct{})
	}
	return client.connectedCh
}

func (client *GrpcClient) notifyStreamDisconnected() {
	client.mux.Lock()
	defer client.mux.Unlock()

	client.status = StreamDisconnected
}

func (client *GrpcClient) notifyStreamConnected() {
	client.mux.Lock()
	defer client.mux.Unlock()

	client.status = StreamConnected
	if client.connectedCh != nil {
		// there are goroutines waiting on this channel -> release them
		close(client.connectedCh)
		client.connectedCh = nil
	}
}
