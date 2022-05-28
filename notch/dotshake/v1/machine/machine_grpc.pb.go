// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: notch/dotshake/v1/machine.proto

package machine

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MachineServiceClient is the client API for MachineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MachineServiceClient interface {
	GetMachine(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMachineResponse, error)
	SyncMachines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (MachineService_SyncMachinesClient, error)
}

type machineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMachineServiceClient(cc grpc.ClientConnInterface) MachineServiceClient {
	return &machineServiceClient{cc}
}

func (c *machineServiceClient) GetMachine(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMachineResponse, error) {
	out := new(GetMachineResponse)
	err := c.cc.Invoke(ctx, "/protos.MachineService/GetMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) SyncMachines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (MachineService_SyncMachinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &MachineService_ServiceDesc.Streams[0], "/protos.MachineService/SyncMachines", opts...)
	if err != nil {
		return nil, err
	}
	x := &machineServiceSyncMachinesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MachineService_SyncMachinesClient interface {
	Recv() (*SyncMachinesResponse, error)
	grpc.ClientStream
}

type machineServiceSyncMachinesClient struct {
	grpc.ClientStream
}

func (x *machineServiceSyncMachinesClient) Recv() (*SyncMachinesResponse, error) {
	m := new(SyncMachinesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MachineServiceServer is the server API for MachineService service.
// All implementations should embed UnimplementedMachineServiceServer
// for forward compatibility
type MachineServiceServer interface {
	GetMachine(context.Context, *emptypb.Empty) (*GetMachineResponse, error)
	SyncMachines(*emptypb.Empty, MachineService_SyncMachinesServer) error
}

// UnimplementedMachineServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMachineServiceServer struct {
}

func (UnimplementedMachineServiceServer) GetMachine(context.Context, *emptypb.Empty) (*GetMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachine not implemented")
}
func (UnimplementedMachineServiceServer) SyncMachines(*emptypb.Empty, MachineService_SyncMachinesServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncMachines not implemented")
}

// UnsafeMachineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MachineServiceServer will
// result in compilation errors.
type UnsafeMachineServiceServer interface {
	mustEmbedUnimplementedMachineServiceServer()
}

func RegisterMachineServiceServer(s grpc.ServiceRegistrar, srv MachineServiceServer) {
	s.RegisterService(&MachineService_ServiceDesc, srv)
}

func _MachineService_GetMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).GetMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MachineService/GetMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).GetMachine(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_SyncMachines_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MachineServiceServer).SyncMachines(m, &machineServiceSyncMachinesServer{stream})
}

type MachineService_SyncMachinesServer interface {
	Send(*SyncMachinesResponse) error
	grpc.ServerStream
}

type machineServiceSyncMachinesServer struct {
	grpc.ServerStream
}

func (x *machineServiceSyncMachinesServer) Send(m *SyncMachinesResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MachineService_ServiceDesc is the grpc.ServiceDesc for MachineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MachineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.MachineService",
	HandlerType: (*MachineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMachine",
			Handler:    _MachineService_GetMachine_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SyncMachines",
			Handler:       _MachineService_SyncMachines_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "notch/dotshake/v1/machine.proto",
}
