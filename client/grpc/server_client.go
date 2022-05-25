package grpc

import (
	"context"
	"fmt"

	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/login_session"
	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/machine"
	"github.com/Notch-Technologies/dotshake/system"
	"github.com/Notch-Technologies/dotshake/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerClientImpl interface {
	GetMachine(mk string) (*machine.GetMachineResponse, error)
	ConnectStreamPeerLoginSession(ctx context.Context, mk string) (*login_session.PeerLoginSessionResponse, error)
}

type ServerClient struct {
	machineClient      machine.MachineServiceClient
	loginSessionClient login_session.LoginSessionServiceClient
	conn               *grpc.ClientConn
	ctx                context.Context
}

func NewServerClient(ctx context.Context, conn *grpc.ClientConn) ServerClientImpl {
	return &ServerClient{
		machineClient:      machine.NewMachineServiceClient(conn),
		loginSessionClient: login_session.NewLoginSessionServiceClient(conn),
		conn:               conn,
		ctx:                ctx,
	}
}

func (c *ServerClient) GetMachine(mk string) (*machine.GetMachineResponse, error) {
	md := metadata.New(map[string]string{utils.MachineKey: mk})
	ctx := metadata.NewOutgoingContext(c.ctx, md)

	res, err := c.machineClient.GetMachine(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return &machine.GetMachineResponse{
		IsRegistered: res.IsRegistered,
		LoginUrl:     res.LoginUrl,
		Ip:           res.Ip,
		Cidr:         res.Cidr,
		SignalHost:   res.SignalHost,
		SignalPort:   res.SignalPort,
	}, nil
}

func (c *ServerClient) ConnectStreamPeerLoginSession(ctx context.Context, mk string) (*login_session.PeerLoginSessionResponse, error) {
	var (
		msg = &login_session.PeerLoginSessionResponse{}
	)

	sys := system.NewSysInfo()
	md := metadata.New(map[string]string{utils.MachineKey: mk, utils.HostName: sys.Hostname, utils.OS: sys.OS})
	newctx := metadata.NewOutgoingContext(ctx, md)

	stream, err := c.loginSessionClient.StreamPeerLoginSession(newctx, grpc.WaitForReady(true))
	if err != nil {
		return nil, err
	}

	header, err := stream.Header()
	if err != nil {
		return nil, err
	}

	sessionid := getLoginSessionID(header)
	fmt.Printf("sessionid: [%s]\n", sessionid)

	for {
		msg, err = stream.Recv()
		if err != nil {
			return nil, err
		}

		err = stream.Send(&emptypb.Empty{})
		if err != nil {
			return nil, err
		}
		break
	}

	return msg, nil
}
