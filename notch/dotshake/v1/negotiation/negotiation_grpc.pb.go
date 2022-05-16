// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: notch/dotshake/v1/negotiation.proto

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
	Offer(ctx context.Context, in *NegotiationRequest, opts ...grpc.CallOption) (*NegotiationResponse, error)
	Answer(ctx context.Context, in *NegotiationRequest, opts ...grpc.CallOption) (*NegotiationResponse, error)
	Candidate(ctx context.Context, in *NegotiationRequest, opts ...grpc.CallOption) (*NegotiationResponse, error)
	StartConnect(ctx context.Context, opts ...grpc.CallOption) (Negotiation_StartConnectClient, error)
}

type negotiationClient struct {
	cc grpc.ClientConnInterface
}

func NewNegotiationClient(cc grpc.ClientConnInterface) NegotiationClient {
	return &negotiationClient{cc}
}

func (c *negotiationClient) Offer(ctx context.Context, in *NegotiationRequest, opts ...grpc.CallOption) (*NegotiationResponse, error) {
	out := new(NegotiationResponse)
	err := c.cc.Invoke(ctx, "/protos.Negotiation/Offer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *negotiationClient) Answer(ctx context.Context, in *NegotiationRequest, opts ...grpc.CallOption) (*NegotiationResponse, error) {
	out := new(NegotiationResponse)
	err := c.cc.Invoke(ctx, "/protos.Negotiation/Answer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *negotiationClient) Candidate(ctx context.Context, in *NegotiationRequest, opts ...grpc.CallOption) (*NegotiationResponse, error) {
	out := new(NegotiationResponse)
	err := c.cc.Invoke(ctx, "/protos.Negotiation/Candidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *negotiationClient) StartConnect(ctx context.Context, opts ...grpc.CallOption) (Negotiation_StartConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Negotiation_ServiceDesc.Streams[0], "/protos.Negotiation/StartConnect", opts...)
	if err != nil {
		return nil, err
	}
	x := &negotiationStartConnectClient{stream}
	return x, nil
}

type Negotiation_StartConnectClient interface {
	Send(*NegotiationRequest) error
	Recv() (*NegotiationResponse, error)
	grpc.ClientStream
}

type negotiationStartConnectClient struct {
	grpc.ClientStream
}

func (x *negotiationStartConnectClient) Send(m *NegotiationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *negotiationStartConnectClient) Recv() (*NegotiationResponse, error) {
	m := new(NegotiationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NegotiationServer is the server API for Negotiation service.
// All implementations should embed UnimplementedNegotiationServer
// for forward compatibility
type NegotiationServer interface {
	Offer(context.Context, *NegotiationRequest) (*NegotiationResponse, error)
	Answer(context.Context, *NegotiationRequest) (*NegotiationResponse, error)
	Candidate(context.Context, *NegotiationRequest) (*NegotiationResponse, error)
	StartConnect(Negotiation_StartConnectServer) error
}

// UnimplementedNegotiationServer should be embedded to have forward compatible implementations.
type UnimplementedNegotiationServer struct {
}

func (UnimplementedNegotiationServer) Offer(context.Context, *NegotiationRequest) (*NegotiationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Offer not implemented")
}
func (UnimplementedNegotiationServer) Answer(context.Context, *NegotiationRequest) (*NegotiationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Answer not implemented")
}
func (UnimplementedNegotiationServer) Candidate(context.Context, *NegotiationRequest) (*NegotiationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Candidate not implemented")
}
func (UnimplementedNegotiationServer) StartConnect(Negotiation_StartConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method StartConnect not implemented")
}

// UnsafeNegotiationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NegotiationServer will
// result in compilation errors.
type UnsafeNegotiationServer interface {
	mustEmbedUnimplementedNegotiationServer()
}

func RegisterNegotiationServer(s grpc.ServiceRegistrar, srv NegotiationServer) {
	s.RegisterService(&Negotiation_ServiceDesc, srv)
}

func _Negotiation_Offer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NegotiationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NegotiationServer).Offer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Negotiation/Offer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NegotiationServer).Offer(ctx, req.(*NegotiationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Negotiation_Answer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NegotiationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NegotiationServer).Answer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Negotiation/Answer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NegotiationServer).Answer(ctx, req.(*NegotiationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Negotiation_Candidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NegotiationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NegotiationServer).Candidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Negotiation/Candidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NegotiationServer).Candidate(ctx, req.(*NegotiationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Negotiation_StartConnect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NegotiationServer).StartConnect(&negotiationStartConnectServer{stream})
}

type Negotiation_StartConnectServer interface {
	Send(*NegotiationResponse) error
	Recv() (*NegotiationRequest, error)
	grpc.ServerStream
}

type negotiationStartConnectServer struct {
	grpc.ServerStream
}

func (x *negotiationStartConnectServer) Send(m *NegotiationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *negotiationStartConnectServer) Recv() (*NegotiationRequest, error) {
	m := new(NegotiationRequest)
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
			MethodName: "Offer",
			Handler:    _Negotiation_Offer_Handler,
		},
		{
			MethodName: "Answer",
			Handler:    _Negotiation_Answer_Handler,
		},
		{
			MethodName: "Candidate",
			Handler:    _Negotiation_Candidate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartConnect",
			Handler:       _Negotiation_StartConnect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "notch/dotshake/v1/negotiation.proto",
}
