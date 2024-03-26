// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: runetale/runetale/v1/acl.proto

package acl

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
	AclService_CreateAcl_FullMethodName = "/protos.AclService/CreateAcl"
	AclService_GetAcl_FullMethodName    = "/protos.AclService/GetAcl"
	AclService_PatchAcl_FullMethodName  = "/protos.AclService/PatchAcl"
)

// AclServiceClient is the client API for AclService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AclServiceClient interface {
	CreateAcl(ctx context.Context, in *CreateAclRequest, opts ...grpc.CallOption) (*AclResponse, error)
	GetAcl(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AclResponse, error)
	PatchAcl(ctx context.Context, in *PatchAclRequest, opts ...grpc.CallOption) (*AclResponse, error)
}

type aclServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAclServiceClient(cc grpc.ClientConnInterface) AclServiceClient {
	return &aclServiceClient{cc}
}

func (c *aclServiceClient) CreateAcl(ctx context.Context, in *CreateAclRequest, opts ...grpc.CallOption) (*AclResponse, error) {
	out := new(AclResponse)
	err := c.cc.Invoke(ctx, AclService_CreateAcl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aclServiceClient) GetAcl(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AclResponse, error) {
	out := new(AclResponse)
	err := c.cc.Invoke(ctx, AclService_GetAcl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aclServiceClient) PatchAcl(ctx context.Context, in *PatchAclRequest, opts ...grpc.CallOption) (*AclResponse, error) {
	out := new(AclResponse)
	err := c.cc.Invoke(ctx, AclService_PatchAcl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AclServiceServer is the server API for AclService service.
// All implementations should embed UnimplementedAclServiceServer
// for forward compatibility
type AclServiceServer interface {
	CreateAcl(context.Context, *CreateAclRequest) (*AclResponse, error)
	GetAcl(context.Context, *emptypb.Empty) (*AclResponse, error)
	PatchAcl(context.Context, *PatchAclRequest) (*AclResponse, error)
}

// UnimplementedAclServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAclServiceServer struct {
}

func (UnimplementedAclServiceServer) CreateAcl(context.Context, *CreateAclRequest) (*AclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAcl not implemented")
}
func (UnimplementedAclServiceServer) GetAcl(context.Context, *emptypb.Empty) (*AclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAcl not implemented")
}
func (UnimplementedAclServiceServer) PatchAcl(context.Context, *PatchAclRequest) (*AclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchAcl not implemented")
}

// UnsafeAclServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AclServiceServer will
// result in compilation errors.
type UnsafeAclServiceServer interface {
	mustEmbedUnimplementedAclServiceServer()
}

func RegisterAclServiceServer(s grpc.ServiceRegistrar, srv AclServiceServer) {
	s.RegisterService(&AclService_ServiceDesc, srv)
}

func _AclService_CreateAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAclRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclServiceServer).CreateAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AclService_CreateAcl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclServiceServer).CreateAcl(ctx, req.(*CreateAclRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AclService_GetAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclServiceServer).GetAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AclService_GetAcl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclServiceServer).GetAcl(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AclService_PatchAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchAclRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclServiceServer).PatchAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AclService_PatchAcl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclServiceServer).PatchAcl(ctx, req.(*PatchAclRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AclService_ServiceDesc is the grpc.ServiceDesc for AclService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AclService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.AclService",
	HandlerType: (*AclServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAcl",
			Handler:    _AclService_CreateAcl_Handler,
		},
		{
			MethodName: "GetAcl",
			Handler:    _AclService_GetAcl_Handler,
		},
		{
			MethodName: "PatchAcl",
			Handler:    _AclService_PatchAcl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runetale/runetale/v1/acl.proto",
}
