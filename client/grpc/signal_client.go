package grpc

import (
	"context"
	"sync"

	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/negotiation"
	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/rtc"
	"github.com/Notch-Technologies/dotshake/conn"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SignalClientImpl interface {
	// TODO: (shintard) Offer, Answer, Candidate
	StartConnect(mk string, handler func(msg *negotiation.NegotiationResponse) error) error
	WaitStartConnect()
	IsReady() bool
	GetStunTurnConfig() (*rtc.GetStunTurnConfigResponse, error)
}

type SignalClient struct {
	negClient negotiation.NegotiationServiceClient
	rtcClient rtc.RtcServiceClient
	conn      *grpc.ClientConn
	ctx       context.Context

	mux sync.Mutex

	connState *conn.ConnectState

	dotlog *dotlog.DotLog
}

func NewSignalClient(
	ctx context.Context,
	conn *grpc.ClientConn,
	cs *conn.ConnectState,
	dotlog *dotlog.DotLog,
) SignalClientImpl {
	return &SignalClient{
		negClient: negotiation.NewNegotiationServiceClient(conn),
		rtcClient: rtc.NewRtcServiceClient(conn),
		conn:      conn,
		ctx:       ctx,

		mux: sync.Mutex{},

		// at this time, it is in an absolutely DISCONNECTED state
		connState: cs,

		dotlog: dotlog,
	}
}

// actually connected to grpc stream
func (c *SignalClient) connectStream(ctx context.Context) (negotiation.NegotiationService_StartConnectClient, error) {
	stream, err := c.negClient.StartConnect(ctx, grpc.WaitForReady(true))
	if err != nil {
		return nil, err
	}

	// set connState to Connected
	c.connState.Connected()

	return stream, nil
}

func (c *SignalClient) StartConnect(mk string, handler func(msg *negotiation.NegotiationResponse) error) error {
	md := metadata.New(map[string]string{utils.MachineKey: mk})
	ctx := metadata.NewOutgoingContext(c.ctx, md)

	stream, err := c.connectStream(ctx)
	if err != nil {
		return err
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			c.dotlog.Logger.Errorf("failed to get grpc client stream for machinek key: %s", msg.SrcPeerMachineKey)
			return err
		}

		err = handler(msg)
		if err != nil {
			c.dotlog.Logger.Errorf("failed to handle grpc client stream stream in machine key: %s", msg.SrcPeerMachineKey)
			return err
		}
	}
}

// connStateがConnectedになるまでできるまで待つ関数
func (c *SignalClient) WaitStartConnect() {
	if c.connState.IsConnected() {
		return
	}

	ch := c.connState.GetConnStatus()
	select {
	case <-c.ctx.Done():
	case <-ch:
	}
}

func (c *SignalClient) IsReady() bool {
	return c.conn.GetState() == connectivity.Ready || c.conn.GetState() == connectivity.Idle
}

func (c *SignalClient) GetStunTurnConfig() (*rtc.GetStunTurnConfigResponse, error) {
	conf, err := c.rtcClient.GetStunTurnConfig(c.ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return conf, nil
}
