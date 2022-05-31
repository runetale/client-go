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

// NegotiationServiceClient is the client API for NegotiationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NegotiationServiceClient interface {
	Offer(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*NegotiationResponse, error)
	Answer(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*NegotiationResponse, error)
	Candidate(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*NegotiationResponse, error)
	StartConnect(ctx context.Context, opts ...grpc.CallOption) (NegotiationService_StartConnectClient, error)
}

type negotiationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNegotiationServiceClient(cc grpc.ClientConnInterface) NegotiationServiceClient {
	return &negotiationServiceClient{cc}
}

func (c *negotiationServiceClient) Offer(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*NegotiationResponse, error) {
	out := new(NegotiationResponse)
	err := c.cc.Invoke(ctx, "/protos.NegotiationService/Offer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *negotiationServiceClient) Answer(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*NegotiationResponse, error) {
	out := new(NegotiationResponse)
	err := c.cc.Invoke(ctx, "/protos.NegotiationService/Answer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *negotiationServiceClient) Candidate(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*NegotiationResponse, error) {
	out := new(NegotiationResponse)
	err := c.cc.Invoke(ctx, "/protos.NegotiationService/Candidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *negotiationServiceClient) StartConnect(ctx context.Context, opts ...grpc.CallOption) (NegotiationService_StartConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &NegotiationService_ServiceDesc.Streams[0], "/protos.NegotiationService/StartConnect", opts...)
	if err != nil {
		return nil, err
	}
	x := &negotiationServiceStartConnectClient{stream}
	return x, nil
}

type NegotiationService_StartConnectClient interface {
	Send(*NegotiationRequest) error
	Recv() (*NegotiationResponse, error)
	grpc.ClientStream
}

type negotiationServiceStartConnectClient struct {
	grpc.ClientStream
}

func (x *negotiationServiceStartConnectClient) Send(m *NegotiationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *negotiationServiceStartConnectClient) Recv() (*NegotiationResponse, error) {
	m := new(NegotiationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NegotiationServiceServer is the server API for NegotiationService service.
// All implementations should embed UnimplementedNegotiationServiceServer
// for forward compatibility
type NegotiationServiceServer interface {
	Offer(context.Context, *HandshakeRequest) (*NegotiationResponse, error)
	Answer(context.Context, *HandshakeRequest) (*NegotiationResponse, error)
	Candidate(context.Context, *HandshakeRequest) (*NegotiationResponse, error)
	StartConnect(NegotiationService_StartConnectServer) error
}

// UnimplementedNegotiationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNegotiationServiceServer struct {
}

func (UnimplementedNegotiationServiceServer) Offer(context.Context, *HandshakeRequest) (*NegotiationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Offer not implemented")
}
func (UnimplementedNegotiationServiceServer) Answer(context.Context, *HandshakeRequest) (*NegotiationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Answer not implemented")
}
func (UnimplementedNegotiationServiceServer) Candidate(context.Context, *HandshakeRequest) (*NegotiationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Candidate not implemented")
}
func (UnimplementedNegotiationServiceServer) StartConnect(NegotiationService_StartConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method StartConnect not implemented")
}

// UnsafeNegotiationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NegotiationServiceServer will
// result in compilation errors.
type UnsafeNegotiationServiceServer interface {
	mustEmbedUnimplementedNegotiationServiceServer()
}

func RegisterNegotiationServiceServer(s grpc.ServiceRegistrar, srv NegotiationServiceServer) {
	s.RegisterService(&NegotiationService_ServiceDesc, srv)
}

func _NegotiationService_Offer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NegotiationServiceServer).Offer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.NegotiationService/Offer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NegotiationServiceServer).Offer(ctx, req.(*HandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NegotiationService_Answer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NegotiationServiceServer).Answer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.NegotiationService/Answer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NegotiationServiceServer).Answer(ctx, req.(*HandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NegotiationService_Candidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NegotiationServiceServer).Candidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.NegotiationService/Candidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NegotiationServiceServer).Candidate(ctx, req.(*HandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NegotiationService_StartConnect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NegotiationServiceServer).StartConnect(&negotiationServiceStartConnectServer{stream})
}

type NegotiationService_StartConnectServer interface {
	Send(*NegotiationResponse) error
	Recv() (*NegotiationRequest, error)
	grpc.ServerStream
}

type negotiationServiceStartConnectServer struct {
	grpc.ServerStream
}

func (x *negotiationServiceStartConnectServer) Send(m *NegotiationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *negotiationServiceStartConnectServer) Recv() (*NegotiationRequest, error) {
	m := new(NegotiationRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NegotiationService_ServiceDesc is the grpc.ServiceDesc for NegotiationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NegotiationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.NegotiationService",
	HandlerType: (*NegotiationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Offer",
			Handler:    _NegotiationService_Offer_Handler,
		},
		{
			MethodName: "Answer",
			Handler:    _NegotiationService_Answer_Handler,
		},
		{
			MethodName: "Candidate",
			Handler:    _NegotiationService_Candidate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartConnect",
			Handler:       _NegotiationService_StartConnect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "notch/dotshake/v1/negotiation.proto",
}