// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.20.3
// source: runetale/runetale/v1/rtc.proto

package rtc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	RtcService_GetStunTurnConfig_FullMethodName = "/protos.RtcService/GetStunTurnConfig"
)

// RtcServiceClient is the client API for RtcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RtcServiceClient interface {
	GetStunTurnConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStunTurnConfigResponse, error)
}

type rtcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRtcServiceClient(cc grpc.ClientConnInterface) RtcServiceClient {
	return &rtcServiceClient{cc}
}

func (c *rtcServiceClient) GetStunTurnConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStunTurnConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStunTurnConfigResponse)
	err := c.cc.Invoke(ctx, RtcService_GetStunTurnConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RtcServiceServer is the server API for RtcService service.
// All implementations should embed UnimplementedRtcServiceServer
// for forward compatibility
type RtcServiceServer interface {
	GetStunTurnConfig(context.Context, *emptypb.Empty) (*GetStunTurnConfigResponse, error)
}

// UnimplementedRtcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRtcServiceServer struct {
}

func (UnimplementedRtcServiceServer) GetStunTurnConfig(context.Context, *emptypb.Empty) (*GetStunTurnConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStunTurnConfig not implemented")
}

// UnsafeRtcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RtcServiceServer will
// result in compilation errors.
type UnsafeRtcServiceServer interface {
	mustEmbedUnimplementedRtcServiceServer()
}

func RegisterRtcServiceServer(s grpc.ServiceRegistrar, srv RtcServiceServer) {
	s.RegisterService(&RtcService_ServiceDesc, srv)
}

func _RtcService_GetStunTurnConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).GetStunTurnConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RtcService_GetStunTurnConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).GetStunTurnConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RtcService_ServiceDesc is the grpc.ServiceDesc for RtcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RtcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.RtcService",
	HandlerType: (*RtcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStunTurnConfig",
			Handler:    _RtcService_GetStunTurnConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runetale/runetale/v1/rtc.proto",
}
