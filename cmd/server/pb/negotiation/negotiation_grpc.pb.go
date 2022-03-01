// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package negotiation

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NegotiationClient is the client API for Negotiation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NegotiationClient interface {
	Send(ctx context.Context, in *Body, opts ...grpc.CallOption) (*Body, error)
	ConnectStream(ctx context.Context, opts ...grpc.CallOption) (Negotiation_ConnectStreamClient, error)
}

type negotiationClient struct {
	cc grpc.ClientConnInterface
}

func NewNegotiationClient(cc grpc.ClientConnInterface) NegotiationClient {
	return &negotiationClient{cc}
}

func (c *negotiationClient) Send(ctx context.Context, in *Body, opts ...grpc.CallOption) (*Body, error) {
	out := new(Body)
	err := c.cc.Invoke(ctx, "/protos.Negotiation/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *negotiationClient) ConnectStream(ctx context.Context, opts ...grpc.CallOption) (Negotiation_ConnectStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Negotiation_ServiceDesc.Streams[0], "/protos.Negotiation/ConnectStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &negotiationConnectStreamClient{stream}
	return x, nil
}

type Negotiation_ConnectStreamClient interface {
	Send(*Body) error
	Recv() (*Body, error)
	grpc.ClientStream
}

type negotiationConnectStreamClient struct {
	grpc.ClientStream
}

func (x *negotiationConnectStreamClient) Send(m *Body) error {
	return x.ClientStream.SendMsg(m)
}

func (x *negotiationConnectStreamClient) Recv() (*Body, error) {
	m := new(Body)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NegotiationServer is the server API for Negotiation service.
// All implementations must embed UnimplementedNegotiationServer
// for forward compatibility
type NegotiationServer interface {
	Send(context.Context, *Body) (*Body, error)
	ConnectStream(Negotiation_ConnectStreamServer) error
	mustEmbedUnimplementedNegotiationServer()
}

// UnimplementedNegotiationServer must be embedded to have forward compatible implementations.
type UnimplementedNegotiationServer struct {
}

func (UnimplementedNegotiationServer) Send(context.Context, *Body) (*Body, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedNegotiationServer) ConnectStream(Negotiation_ConnectStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectStream not implemented")
}
func (UnimplementedNegotiationServer) mustEmbedUnimplementedNegotiationServer() {}

// UnsafeNegotiationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NegotiationServer will
// result in compilation errors.
type UnsafeNegotiationServer interface {
	mustEmbedUnimplementedNegotiationServer()
}

func RegisterNegotiationServer(s grpc.ServiceRegistrar, srv NegotiationServer) {
	s.RegisterService(&Negotiation_ServiceDesc, srv)
}

func _Negotiation_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Body)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NegotiationServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Negotiation/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NegotiationServer).Send(ctx, req.(*Body))
	}
	return interceptor(ctx, in, info, handler)
}

func _Negotiation_ConnectStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NegotiationServer).ConnectStream(&negotiationConnectStreamServer{stream})
}

type Negotiation_ConnectStreamServer interface {
	Send(*Body) error
	Recv() (*Body, error)
	grpc.ServerStream
}

type negotiationConnectStreamServer struct {
	grpc.ServerStream
}

func (x *negotiationConnectStreamServer) Send(m *Body) error {
	return x.ServerStream.SendMsg(m)
}

func (x *negotiationConnectStreamServer) Recv() (*Body, error) {
	m := new(Body)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Negotiation_ServiceDesc is the grpc.ServiceDesc for Negotiation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Negotiation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Negotiation",
	HandlerType: (*NegotiationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Negotiation_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectStream",
			Handler:       _Negotiation_ConnectStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/negotiation.proto",
}
