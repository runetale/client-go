package grpc

import (
	"context"

	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/machine"
	"github.com/Notch-Technologies/dotshake/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerClientImpl interface {
	GetMachine(mk string) (*machine.GetMachineResponse, error)
}

type ServerClient struct {
	client machine.MachineServiceClient
	conn   *grpc.ClientConn
	ctx    context.Context
}

func NewServerClient(ctx context.Context, conn *grpc.ClientConn) ServerClientImpl {
	return &ServerClient{
		client: machine.NewMachineServiceClient(conn),
		conn:   conn,
		ctx:    ctx,
	}
}

func (c *ServerClient) GetMachine(mk string) (*machine.GetMachineResponse, error) {
	md := metadata.New(map[string]string{utils.MachineKey: mk})
	ctx := metadata.NewOutgoingContext(c.ctx, md)

	res, err := c.client.GetMachine(ctx, &emptypb.Empty{})
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
