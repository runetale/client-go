// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: runetale/runetale/v1/negotiation.proto

package negotiation

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NegotiationService_Offer_FullMethodName     = "/protos.NegotiationService/Offer"
	NegotiationService_Answer_FullMethodName    = "/protos.NegotiationService/Answer"
	NegotiationService_Candidate_FullMethodName = "/protos.NegotiationService/Candidate"
	NegotiationService_Connect_FullMethodName   = "/protos.NegotiationService/Connect"
)

// NegotiationServiceClient is the client API for NegotiationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NegotiationServiceClient interface {
	Offer(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Answer(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Candidate(ctx context.Context, in *CandidateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Connect(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[NegotiationRequest, NegotiationRequest], error)
}

type negotiationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNegotiationServiceClient(cc grpc.ClientConnInterface) NegotiationServiceClient {
	return &negotiationServiceClient{cc}
}

func (c *negotiationServiceClient) Offer(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NegotiationService_Offer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *negotiationServiceClient) Answer(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NegotiationService_Answer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *negotiationServiceClient) Candidate(ctx context.Context, in *CandidateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NegotiationService_Candidate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *negotiationServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[NegotiationRequest, NegotiationRequest], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &NegotiationService_ServiceDesc.Streams[0], NegotiationService_Connect_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[NegotiationRequest, NegotiationRequest]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NegotiationService_ConnectClient = grpc.BidiStreamingClient[NegotiationRequest, NegotiationRequest]

// NegotiationServiceServer is the server API for NegotiationService service.
// All implementations should embed UnimplementedNegotiationServiceServer
// for forward compatibility.
type NegotiationServiceServer interface {
	Offer(context.Context, *HandshakeRequest) (*emptypb.Empty, error)
	Answer(context.Context, *HandshakeRequest) (*emptypb.Empty, error)
	Candidate(context.Context, *CandidateRequest) (*emptypb.Empty, error)
	Connect(grpc.BidiStreamingServer[NegotiationRequest, NegotiationRequest]) error
}

// UnimplementedNegotiationServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNegotiationServiceServer struct{}

func (UnimplementedNegotiationServiceServer) Offer(context.Context, *HandshakeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Offer not implemented")
}
func (UnimplementedNegotiationServiceServer) Answer(context.Context, *HandshakeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Answer not implemented")
}
func (UnimplementedNegotiationServiceServer) Candidate(context.Context, *CandidateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Candidate not implemented")
}
func (UnimplementedNegotiationServiceServer) Connect(grpc.BidiStreamingServer[NegotiationRequest, NegotiationRequest]) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedNegotiationServiceServer) testEmbeddedByValue() {}

// UnsafeNegotiationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NegotiationServiceServer will
// result in compilation errors.
type UnsafeNegotiationServiceServer interface {
	mustEmbedUnimplementedNegotiationServiceServer()
}

func RegisterNegotiationServiceServer(s grpc.ServiceRegistrar, srv NegotiationServiceServer) {
	// If the following call pancis, it indicates UnimplementedNegotiationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
		FullMethod: NegotiationService_Offer_FullMethodName,
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
		FullMethod: NegotiationService_Answer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NegotiationServiceServer).Answer(ctx, req.(*HandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NegotiationService_Candidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NegotiationServiceServer).Candidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NegotiationService_Candidate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NegotiationServiceServer).Candidate(ctx, req.(*CandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NegotiationService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NegotiationServiceServer).Connect(&grpc.GenericServerStream[NegotiationRequest, NegotiationRequest]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NegotiationService_ConnectServer = grpc.BidiStreamingServer[NegotiationRequest, NegotiationRequest]

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
			StreamName:    "Connect",
			Handler:       _NegotiationService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "runetale/runetale/v1/negotiation.proto",
}
