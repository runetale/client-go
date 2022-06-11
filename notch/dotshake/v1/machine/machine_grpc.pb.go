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
	SyncRemoteMachinesConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SyncMachinesResponse, error)
	ConnectToHangoutMachines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (MachineService_ConnectToHangoutMachinesClient, error)
	JoinHangOutMachines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HangOutMachinesResponse, error)
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

func (c *machineServiceClient) SyncRemoteMachinesConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SyncMachinesResponse, error) {
	out := new(SyncMachinesResponse)
	err := c.cc.Invoke(ctx, "/protos.MachineService/SyncRemoteMachinesConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) ConnectToHangoutMachines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (MachineService_ConnectToHangoutMachinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &MachineService_ServiceDesc.Streams[0], "/protos.MachineService/ConnectToHangoutMachines", opts...)
	if err != nil {
		return nil, err
	}
	x := &machineServiceConnectToHangoutMachinesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MachineService_ConnectToHangoutMachinesClient interface {
	Recv() (*HangOutMachinesResponse, error)
	grpc.ClientStream
}

type machineServiceConnectToHangoutMachinesClient struct {
	grpc.ClientStream
}

func (x *machineServiceConnectToHangoutMachinesClient) Recv() (*HangOutMachinesResponse, error) {
	m := new(HangOutMachinesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *machineServiceClient) JoinHangOutMachines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HangOutMachinesResponse, error) {
	out := new(HangOutMachinesResponse)
	err := c.cc.Invoke(ctx, "/protos.MachineService/JoinHangOutMachines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MachineServiceServer is the server API for MachineService service.
// All implementations should embed UnimplementedMachineServiceServer
// for forward compatibility
type MachineServiceServer interface {
	GetMachine(context.Context, *emptypb.Empty) (*GetMachineResponse, error)
	SyncRemoteMachinesConfig(context.Context, *emptypb.Empty) (*SyncMachinesResponse, error)
	ConnectToHangoutMachines(*emptypb.Empty, MachineService_ConnectToHangoutMachinesServer) error
	JoinHangOutMachines(context.Context, *emptypb.Empty) (*HangOutMachinesResponse, error)
}

// UnimplementedMachineServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMachineServiceServer struct {
}

func (UnimplementedMachineServiceServer) GetMachine(context.Context, *emptypb.Empty) (*GetMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachine not implemented")
}
func (UnimplementedMachineServiceServer) SyncRemoteMachinesConfig(context.Context, *emptypb.Empty) (*SyncMachinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncRemoteMachinesConfig not implemented")
}
func (UnimplementedMachineServiceServer) ConnectToHangoutMachines(*emptypb.Empty, MachineService_ConnectToHangoutMachinesServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectToHangoutMachines not implemented")
}
func (UnimplementedMachineServiceServer) JoinHangOutMachines(context.Context, *emptypb.Empty) (*HangOutMachinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinHangOutMachines not implemented")
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

func _MachineService_SyncRemoteMachinesConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).SyncRemoteMachinesConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MachineService/SyncRemoteMachinesConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).SyncRemoteMachinesConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_ConnectToHangoutMachines_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MachineServiceServer).ConnectToHangoutMachines(m, &machineServiceConnectToHangoutMachinesServer{stream})
}

type MachineService_ConnectToHangoutMachinesServer interface {
	Send(*HangOutMachinesResponse) error
	grpc.ServerStream
}

type machineServiceConnectToHangoutMachinesServer struct {
	grpc.ServerStream
}

func (x *machineServiceConnectToHangoutMachinesServer) Send(m *HangOutMachinesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MachineService_JoinHangOutMachines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).JoinHangOutMachines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MachineService/JoinHangOutMachines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).JoinHangOutMachines(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
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
		{
			MethodName: "SyncRemoteMachinesConfig",
			Handler:    _MachineService_SyncRemoteMachinesConfig_Handler,
		},
		{
			MethodName: "JoinHangOutMachines",
			Handler:    _MachineService_JoinHangOutMachines_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectToHangoutMachines",
			Handler:       _MachineService_ConnectToHangoutMachines_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "notch/dotshake/v1/machine.proto",
}
