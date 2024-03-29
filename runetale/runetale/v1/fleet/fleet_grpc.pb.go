// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: runetale/runetale/v1/fleet.proto

package fleet

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
	FleetService_CreateFleet_FullMethodName = "/protos.FleetService/CreateFleet"
	FleetService_PatchFleet_FullMethodName  = "/protos.FleetService/PatchFleet"
	FleetService_GetFleet_FullMethodName    = "/protos.FleetService/GetFleet"
	FleetService_GetFleets_FullMethodName   = "/protos.FleetService/GetFleets"
)

// FleetServiceClient is the client API for FleetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FleetServiceClient interface {
	CreateFleet(ctx context.Context, in *CreateFleetRequest, opts ...grpc.CallOption) (*FleetResponse, error)
	PatchFleet(ctx context.Context, in *PatchFleetRequest, opts ...grpc.CallOption) (*FleetResponse, error)
	GetFleet(ctx context.Context, in *GetFleetRequest, opts ...grpc.CallOption) (*FleetResponse, error)
	GetFleets(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetFleetsResponse, error)
}

type fleetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFleetServiceClient(cc grpc.ClientConnInterface) FleetServiceClient {
	return &fleetServiceClient{cc}
}

func (c *fleetServiceClient) CreateFleet(ctx context.Context, in *CreateFleetRequest, opts ...grpc.CallOption) (*FleetResponse, error) {
	out := new(FleetResponse)
	err := c.cc.Invoke(ctx, FleetService_CreateFleet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fleetServiceClient) PatchFleet(ctx context.Context, in *PatchFleetRequest, opts ...grpc.CallOption) (*FleetResponse, error) {
	out := new(FleetResponse)
	err := c.cc.Invoke(ctx, FleetService_PatchFleet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fleetServiceClient) GetFleet(ctx context.Context, in *GetFleetRequest, opts ...grpc.CallOption) (*FleetResponse, error) {
	out := new(FleetResponse)
	err := c.cc.Invoke(ctx, FleetService_GetFleet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fleetServiceClient) GetFleets(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetFleetsResponse, error) {
	out := new(GetFleetsResponse)
	err := c.cc.Invoke(ctx, FleetService_GetFleets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FleetServiceServer is the server API for FleetService service.
// All implementations should embed UnimplementedFleetServiceServer
// for forward compatibility
type FleetServiceServer interface {
	CreateFleet(context.Context, *CreateFleetRequest) (*FleetResponse, error)
	PatchFleet(context.Context, *PatchFleetRequest) (*FleetResponse, error)
	GetFleet(context.Context, *GetFleetRequest) (*FleetResponse, error)
	GetFleets(context.Context, *emptypb.Empty) (*GetFleetsResponse, error)
}

// UnimplementedFleetServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFleetServiceServer struct {
}

func (UnimplementedFleetServiceServer) CreateFleet(context.Context, *CreateFleetRequest) (*FleetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFleet not implemented")
}
func (UnimplementedFleetServiceServer) PatchFleet(context.Context, *PatchFleetRequest) (*FleetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchFleet not implemented")
}
func (UnimplementedFleetServiceServer) GetFleet(context.Context, *GetFleetRequest) (*FleetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFleet not implemented")
}
func (UnimplementedFleetServiceServer) GetFleets(context.Context, *emptypb.Empty) (*GetFleetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFleets not implemented")
}

// UnsafeFleetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FleetServiceServer will
// result in compilation errors.
type UnsafeFleetServiceServer interface {
	mustEmbedUnimplementedFleetServiceServer()
}

func RegisterFleetServiceServer(s grpc.ServiceRegistrar, srv FleetServiceServer) {
	s.RegisterService(&FleetService_ServiceDesc, srv)
}

func _FleetService_CreateFleet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFleetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FleetServiceServer).CreateFleet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FleetService_CreateFleet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FleetServiceServer).CreateFleet(ctx, req.(*CreateFleetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FleetService_PatchFleet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchFleetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FleetServiceServer).PatchFleet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FleetService_PatchFleet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FleetServiceServer).PatchFleet(ctx, req.(*PatchFleetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FleetService_GetFleet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFleetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FleetServiceServer).GetFleet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FleetService_GetFleet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FleetServiceServer).GetFleet(ctx, req.(*GetFleetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FleetService_GetFleets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FleetServiceServer).GetFleets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FleetService_GetFleets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FleetServiceServer).GetFleets(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// FleetService_ServiceDesc is the grpc.ServiceDesc for FleetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FleetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.FleetService",
	HandlerType: (*FleetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFleet",
			Handler:    _FleetService_CreateFleet_Handler,
		},
		{
			MethodName: "PatchFleet",
			Handler:    _FleetService_PatchFleet_Handler,
		},
		{
			MethodName: "GetFleet",
			Handler:    _FleetService_GetFleet_Handler,
		},
		{
			MethodName: "GetFleets",
			Handler:    _FleetService_GetFleets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runetale/runetale/v1/fleet.proto",
}
