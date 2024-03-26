// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: runetale/runetale/v1/oidc.proto

package oidc

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

const (
	OIDCService_Login_FullMethodName        = "/protos.OIDCService/Login"
	OIDCService_Authenticate_FullMethodName = "/protos.OIDCService/Authenticate"
	OIDCService_RefreshToken_FullMethodName = "/protos.OIDCService/RefreshToken"
)

// OIDCServiceClient is the client API for OIDCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OIDCServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Authenticate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	RefreshToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type oIDCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOIDCServiceClient(cc grpc.ClientConnInterface) OIDCServiceClient {
	return &oIDCServiceClient{cc}
}

func (c *oIDCServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, OIDCService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oIDCServiceClient) Authenticate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, OIDCService_Authenticate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oIDCServiceClient) RefreshToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, OIDCService_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OIDCServiceServer is the server API for OIDCService service.
// All implementations should embed UnimplementedOIDCServiceServer
// for forward compatibility
type OIDCServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Authenticate(context.Context, *emptypb.Empty) (*AuthenticateResponse, error)
	RefreshToken(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

// UnimplementedOIDCServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOIDCServiceServer struct {
}

func (UnimplementedOIDCServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedOIDCServiceServer) Authenticate(context.Context, *emptypb.Empty) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedOIDCServiceServer) RefreshToken(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}

// UnsafeOIDCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OIDCServiceServer will
// result in compilation errors.
type UnsafeOIDCServiceServer interface {
	mustEmbedUnimplementedOIDCServiceServer()
}

func RegisterOIDCServiceServer(s grpc.ServiceRegistrar, srv OIDCServiceServer) {
	s.RegisterService(&OIDCService_ServiceDesc, srv)
}

func _OIDCService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OIDCService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OIDCService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OIDCService_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCServiceServer).Authenticate(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OIDCService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OIDCService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCServiceServer).RefreshToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// OIDCService_ServiceDesc is the grpc.ServiceDesc for OIDCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OIDCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.OIDCService",
	HandlerType: (*OIDCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _OIDCService_Login_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _OIDCService_Authenticate_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _OIDCService_RefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runetale/runetale/v1/oidc.proto",
}
