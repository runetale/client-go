// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: notch/dotshake/v1/admin.proto

package admin

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
	AdminService_GetMachines_FullMethodName = "/protos.AdminService/GetMachines"
	AdminService_GetMe_FullMethodName       = "/protos.AdminService/GetMe"
	AdminService_GetUsers_FullMethodName    = "/protos.AdminService/GetUsers"
	AdminService_CreateAcl_FullMethodName   = "/protos.AdminService/CreateAcl"
	AdminService_DeleteAcl_FullMethodName   = "/protos.AdminService/DeleteAcl"
	AdminService_GetAcl_FullMethodName      = "/protos.AdminService/GetAcl"
	AdminService_PatchAcl_FullMethodName    = "/protos.AdminService/PatchAcl"
	AdminService_CreateGroup_FullMethodName = "/protos.AdminService/CreateGroup"
	AdminService_DeleteGroup_FullMethodName = "/protos.AdminService/DeleteGroup"
	AdminService_GetGroup_FullMethodName    = "/protos.AdminService/GetGroup"
	AdminService_PatchGroup_FullMethodName  = "/protos.AdminService/PatchGroup"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	GetMachines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMachinesResponse, error)
	GetMe(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMeResponse, error)
	GetUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetUsersResponse, error)
	CreateAcl(ctx context.Context, in *CreateAclRequest, opts ...grpc.CallOption) (*AclResponse, error)
	DeleteAcl(ctx context.Context, in *DeleteAclRequest, opts ...grpc.CallOption) (*AclResponse, error)
	GetAcl(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AclsResponse, error)
	PatchAcl(ctx context.Context, in *PatchAclRequest, opts ...grpc.CallOption) (*AclResponse, error)
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*GroupResponse, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*GroupResponse, error)
	GetGroup(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GroupsResponse, error)
	PatchGroup(ctx context.Context, in *PatchGroupRequest, opts ...grpc.CallOption) (*GroupResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) GetMachines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMachinesResponse, error) {
	out := new(GetMachinesResponse)
	err := c.cc.Invoke(ctx, AdminService_GetMachines_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetMe(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMeResponse, error) {
	out := new(GetMeResponse)
	err := c.cc.Invoke(ctx, AdminService_GetMe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, AdminService_GetUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreateAcl(ctx context.Context, in *CreateAclRequest, opts ...grpc.CallOption) (*AclResponse, error) {
	out := new(AclResponse)
	err := c.cc.Invoke(ctx, AdminService_CreateAcl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteAcl(ctx context.Context, in *DeleteAclRequest, opts ...grpc.CallOption) (*AclResponse, error) {
	out := new(AclResponse)
	err := c.cc.Invoke(ctx, AdminService_DeleteAcl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetAcl(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AclsResponse, error) {
	out := new(AclsResponse)
	err := c.cc.Invoke(ctx, AdminService_GetAcl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) PatchAcl(ctx context.Context, in *PatchAclRequest, opts ...grpc.CallOption) (*AclResponse, error) {
	out := new(AclResponse)
	err := c.cc.Invoke(ctx, AdminService_PatchAcl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*GroupResponse, error) {
	out := new(GroupResponse)
	err := c.cc.Invoke(ctx, AdminService_CreateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*GroupResponse, error) {
	out := new(GroupResponse)
	err := c.cc.Invoke(ctx, AdminService_DeleteGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetGroup(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GroupsResponse, error) {
	out := new(GroupsResponse)
	err := c.cc.Invoke(ctx, AdminService_GetGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) PatchGroup(ctx context.Context, in *PatchGroupRequest, opts ...grpc.CallOption) (*GroupResponse, error) {
	out := new(GroupResponse)
	err := c.cc.Invoke(ctx, AdminService_PatchGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations should embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	GetMachines(context.Context, *emptypb.Empty) (*GetMachinesResponse, error)
	GetMe(context.Context, *emptypb.Empty) (*GetMeResponse, error)
	GetUsers(context.Context, *emptypb.Empty) (*GetUsersResponse, error)
	CreateAcl(context.Context, *CreateAclRequest) (*AclResponse, error)
	DeleteAcl(context.Context, *DeleteAclRequest) (*AclResponse, error)
	GetAcl(context.Context, *emptypb.Empty) (*AclsResponse, error)
	PatchAcl(context.Context, *PatchAclRequest) (*AclResponse, error)
	CreateGroup(context.Context, *CreateGroupRequest) (*GroupResponse, error)
	DeleteGroup(context.Context, *DeleteGroupRequest) (*GroupResponse, error)
	GetGroup(context.Context, *emptypb.Empty) (*GroupsResponse, error)
	PatchGroup(context.Context, *PatchGroupRequest) (*GroupResponse, error)
}

// UnimplementedAdminServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) GetMachines(context.Context, *emptypb.Empty) (*GetMachinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachines not implemented")
}
func (UnimplementedAdminServiceServer) GetMe(context.Context, *emptypb.Empty) (*GetMeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMe not implemented")
}
func (UnimplementedAdminServiceServer) GetUsers(context.Context, *emptypb.Empty) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedAdminServiceServer) CreateAcl(context.Context, *CreateAclRequest) (*AclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAcl not implemented")
}
func (UnimplementedAdminServiceServer) DeleteAcl(context.Context, *DeleteAclRequest) (*AclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAcl not implemented")
}
func (UnimplementedAdminServiceServer) GetAcl(context.Context, *emptypb.Empty) (*AclsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAcl not implemented")
}
func (UnimplementedAdminServiceServer) PatchAcl(context.Context, *PatchAclRequest) (*AclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchAcl not implemented")
}
func (UnimplementedAdminServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*GroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedAdminServiceServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*GroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedAdminServiceServer) GetGroup(context.Context, *emptypb.Empty) (*GroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedAdminServiceServer) PatchGroup(context.Context, *PatchGroupRequest) (*GroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchGroup not implemented")
}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_GetMachines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetMachines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetMachines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetMachines(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetMe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetMe(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetUsers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreateAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAclRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_CreateAcl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateAcl(ctx, req.(*CreateAclRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAclRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_DeleteAcl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteAcl(ctx, req.(*DeleteAclRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetAcl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetAcl(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_PatchAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchAclRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).PatchAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_PatchAcl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).PatchAcl(ctx, req.(*PatchAclRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_DeleteGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetGroup(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_PatchGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).PatchGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_PatchGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).PatchGroup(ctx, req.(*PatchGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMachines",
			Handler:    _AdminService_GetMachines_Handler,
		},
		{
			MethodName: "GetMe",
			Handler:    _AdminService_GetMe_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _AdminService_GetUsers_Handler,
		},
		{
			MethodName: "CreateAcl",
			Handler:    _AdminService_CreateAcl_Handler,
		},
		{
			MethodName: "DeleteAcl",
			Handler:    _AdminService_DeleteAcl_Handler,
		},
		{
			MethodName: "GetAcl",
			Handler:    _AdminService_GetAcl_Handler,
		},
		{
			MethodName: "PatchAcl",
			Handler:    _AdminService_PatchAcl_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _AdminService_CreateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _AdminService_DeleteGroup_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _AdminService_GetGroup_Handler,
		},
		{
			MethodName: "PatchGroup",
			Handler:    _AdminService_PatchGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notch/dotshake/v1/admin.proto",
}
