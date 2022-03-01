package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/url"
	"sync"
	"time"

	server "github.com/Notch-Technologies/wizy/cmd/server/grpc_server/service"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/negotiation"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/session"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/user"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Status string

const StreamConnected Status = "Connected"
const StreamDisconnected Status = "Disconnected"

type GrpcClientManager interface {
	GetServerPublicKey() (*wgtypes.Key, error)
	Login(setupKey, clientPubKey, serverPubKey string) (string, error)
}

type GrpcClient struct {
	privateKey           wgtypes.Key
	userServiceClient    user.UserServiceClient
	peerServiceClient    peer.PeerServiceClient
	sessionServiceClient session.SessionServiceClient
	negotiationClient    negotiation.NegotiationClient
	stream               negotiation.Negotiation_ConnectStreamClient

	ctx  context.Context
	conn *grpc.ClientConn
	mux  sync.Mutex

	connectedCh chan struct{}

	status Status
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
	nc := negotiation.NewNegotiationClient(conn)

	return &GrpcClient{
		privateKey:           privKey,
		userServiceClient:    usc,
		peerServiceClient:    psc,
		sessionServiceClient: sec,
		negotiationClient:    nc,

		ctx:    ctx,
		conn:   conn,
		mux:    sync.Mutex{},
		status: StreamDisconnected,
	}, nil
}

func (client *GrpcClient) isReady() bool {
	return client.conn.GetState() == connectivity.Ready || client.conn.GetState() == connectivity.Idle
}

func (client *GrpcClient) GetServerPublicKey() (string, error) {
	if !client.isReady() {
		return "", fmt.Errorf("no connection wics server")
	}

	usCtx, cancel := context.WithTimeout(client.ctx, 10*time.Second)
	defer cancel()

	res, err := client.sessionServiceClient.GetServerPublicKey(usCtx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}

	return res.Key, nil
}

func (client *GrpcClient) Login(setupKey, clientPubKey, serverPubKey, ip string, wgPrivateKey wgtypes.Key) (*session.LoginMessage, error) {
	if !client.isReady() {
		return nil, fmt.Errorf("no connection wics server")
	}

	usCtx, cancel := context.WithTimeout(client.ctx, 10*time.Second)
	defer cancel()

	msg, err := client.sessionServiceClient.Login(usCtx, &session.LoginMessage{
		SetupKey:        setupKey,
		ClientPublicKey: clientPubKey,
		ServerPublicKey: serverPubKey,
		WgPublicKey:     wgPrivateKey.PublicKey().String(),
		Ip:              ip,
	})
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
	if err != nil {
		return nil, err
	}
	client.stream = stream

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

	if !client.Ready() {
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

		err = msgHandler(msg)

		if err != nil {
			fmt.Printf("error while handling message of Peer [key: %s] error: [%s]\n", msg.ClientMachineKey, err.Error())
			return err
		}
	}
}

func (client *GrpcClient) Sync(clientMachineKey string, msgHandler func(msg *peer.SyncResponse) error) error {
	stream, err := client.peerServiceClient.Sync(client.ctx, &peer.SyncMessage{
		PrivateKey:       client.privateKey.String(),
		ClientMachineKey: clientMachineKey,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	for {
		update, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("recv error io")
			return err
		}

		if err != nil {
			fmt.Println("recv error")
			return err
		}

		err = msgHandler(update)
		if err != nil {
			return err
		}
	}
}

func (client *GrpcClient) Ready() bool {
	return client.conn.GetState() == connectivity.Ready || client.conn.GetState() == connectivity.Idle
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

func (client *GrpcClient) Send(msg *negotiation.Body) error {
	if !client.Ready() {
		return fmt.Errorf("no connection server stream")
	}

	//ctx, cancel := context.WithTimeout(client.ctx., 5*time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := client.negotiationClient.Send(ctx, msg)
	if err != nil {
		return err
	}
	fmt.Println(body)

	return nil
}

func (c *GrpcClient) SendToStream(msg *negotiation.Body) error {
	if !c.Ready() {
		return fmt.Errorf("no connection to signal")
	}
	if c.stream == nil {
		return fmt.Errorf("connection to the Signal Exchnage has not been established yet. Please call Client.Receive before sending messages")
	}

	err := c.stream.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
