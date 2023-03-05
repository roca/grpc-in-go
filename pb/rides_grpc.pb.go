// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: rides.proto

package pb

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

// RidesClient is the client API for Rides service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RidesClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
}

type ridesClient struct {
	cc grpc.ClientConnInterface
}

func NewRidesClient(cc grpc.ClientConnInterface) RidesClient {
	return &ridesClient{cc}
}

func (c *ridesClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/Rides/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RidesServer is the server API for Rides service.
// All implementations must embed UnimplementedRidesServer
// for forward compatibility
type RidesServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	mustEmbedUnimplementedRidesServer()
}

// UnimplementedRidesServer must be embedded to have forward compatible implementations.
type UnimplementedRidesServer struct {
}

func (UnimplementedRidesServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedRidesServer) mustEmbedUnimplementedRidesServer() {}

// UnsafeRidesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RidesServer will
// result in compilation errors.
type UnsafeRidesServer interface {
	mustEmbedUnimplementedRidesServer()
}

func RegisterRidesServer(s grpc.ServiceRegistrar, srv RidesServer) {
	s.RegisterService(&Rides_ServiceDesc, srv)
}

func _Rides_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RidesServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Rides/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RidesServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rides_ServiceDesc is the grpc.ServiceDesc for Rides service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rides_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Rides",
	HandlerType: (*RidesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Rides_Start_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rides.proto",
}
