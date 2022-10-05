// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: notch/dotshake/v1/daemon.proto

package daemon

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

// DaemonServiceClient is the client API for DaemonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaemonServiceClient interface {
	// connections
	//
	Connect(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetConnectionStatusResponse, error)
	Disconnect(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetConnectionStatusResponse, error)
	GetConnectionStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetConnectionStatusResponse, error)
}

type daemonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDaemonServiceClient(cc grpc.ClientConnInterface) DaemonServiceClient {
	return &daemonServiceClient{cc}
}

func (c *daemonServiceClient) Connect(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetConnectionStatusResponse, error) {
	out := new(GetConnectionStatusResponse)
	err := c.cc.Invoke(ctx, "/protos.DaemonService/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) Disconnect(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetConnectionStatusResponse, error) {
	out := new(GetConnectionStatusResponse)
	err := c.cc.Invoke(ctx, "/protos.DaemonService/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) GetConnectionStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetConnectionStatusResponse, error) {
	out := new(GetConnectionStatusResponse)
	err := c.cc.Invoke(ctx, "/protos.DaemonService/GetConnectionStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServiceServer is the server API for DaemonService service.
// All implementations should embed UnimplementedDaemonServiceServer
// for forward compatibility
type DaemonServiceServer interface {
	// connections
	//
	Connect(context.Context, *emptypb.Empty) (*GetConnectionStatusResponse, error)
	Disconnect(context.Context, *emptypb.Empty) (*GetConnectionStatusResponse, error)
	GetConnectionStatus(context.Context, *emptypb.Empty) (*GetConnectionStatusResponse, error)
}

// UnimplementedDaemonServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDaemonServiceServer struct {
}

func (UnimplementedDaemonServiceServer) Connect(context.Context, *emptypb.Empty) (*GetConnectionStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedDaemonServiceServer) Disconnect(context.Context, *emptypb.Empty) (*GetConnectionStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedDaemonServiceServer) GetConnectionStatus(context.Context, *emptypb.Empty) (*GetConnectionStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectionStatus not implemented")
}

// UnsafeDaemonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaemonServiceServer will
// result in compilation errors.
type UnsafeDaemonServiceServer interface {
	mustEmbedUnimplementedDaemonServiceServer()
}

func RegisterDaemonServiceServer(s grpc.ServiceRegistrar, srv DaemonServiceServer) {
	s.RegisterService(&DaemonService_ServiceDesc, srv)
}

func _DaemonService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.DaemonService/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).Connect(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.DaemonService/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).Disconnect(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_GetConnectionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetConnectionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.DaemonService/GetConnectionStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetConnectionStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DaemonService_ServiceDesc is the grpc.ServiceDesc for DaemonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DaemonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.DaemonService",
	HandlerType: (*DaemonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _DaemonService_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _DaemonService_Disconnect_Handler,
		},
		{
			MethodName: "GetConnectionStatus",
			Handler:    _DaemonService_GetConnectionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notch/dotshake/v1/daemon.proto",
}