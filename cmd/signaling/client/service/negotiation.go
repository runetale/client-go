package service

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/signaling/key"
	"github.com/Notch-Technologies/wizy/cmd/signaling/pb/negotiation"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Status string

const StreamConnected Status = "Connected"
const StreamDisconnected Status = "Disconnected"

type NegotiationClientServiceCaller interface {
	Send(msg *negotiation.Body) error
	Receive(wgPubKey string, msgHandler func(msg *negotiation.Body) error) error
	WaitStreamConnected()
}

type NegotiationClientService struct {
	negotiationClient negotiation.NegotiationClient
	stream            negotiation.Negotiation_ConnectStreamClient

	ctx context.Context
	mux sync.Mutex

	status      Status
	connectedCh chan struct{}
}

func NewNegotiationClientService(
	ctx context.Context, conn *grpc.ClientConn,
	privateKey wgtypes.Key,
) *NegotiationClientService {
	return &NegotiationClientService{
		negotiationClient: negotiation.NewNegotiationClient(conn),
		stream:            nil,

		ctx:    ctx,
		mux:    sync.Mutex{},
		status: StreamDisconnected,
	}
}

func (n *NegotiationClientService) Send(msg *negotiation.Body) error {
	// ctx, cancel := context.WithTimeout(n.ctx, 5*time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := n.negotiationClient.Send(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}

func (n *NegotiationClientService) connectStream(wgPubKey string) (negotiation.Negotiation_ConnectStreamClient, error) {
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

func (n *NegotiationClientService) notifyStreamDisconnected() {
	n.mux.Lock()
	defer n.mux.Unlock()

	n.status = StreamDisconnected
}

func (n *NegotiationClientService) notifyStreamConnected() {
	n.mux.Lock()
	defer n.mux.Unlock()

	n.status = StreamConnected
	if n.connectedCh != nil {
		// there are goroutines waiting on this channel -> release them
		close(n.connectedCh)
		n.connectedCh = nil
	}
}

func (n *NegotiationClientService) Receive(
	wgPubKey string,
	msgHandler func(msg *negotiation.Body) error,
) error {
	n.notifyStreamDisconnected()

	stream, err := n.connectStream(wgPubKey)
	if err != nil {
		return err
	}

	n.notifyStreamConnected()

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

func (n *NegotiationClientService) getStreamStatusChan() <-chan struct{} {
	n.mux.Lock()
	defer n.mux.Unlock()

	if n.connectedCh == nil {
		n.connectedCh = make(chan struct{})
	}
	return n.connectedCh
}

func (n *NegotiationClientService) WaitStreamConnected() {
	if n.status == StreamConnected {
		return
	}

	ch := n.getStreamStatusChan()
	select {
	case <-n.ctx.Done():
	case <-ch:
	}
}
