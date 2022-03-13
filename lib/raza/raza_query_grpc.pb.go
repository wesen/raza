// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: raza_query.proto

package raza

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

// RazaQueryClient is the client API for RazaQuery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RazaQueryClient interface {
	GetSessions(ctx context.Context, in *GetSessionsRequest, opts ...grpc.CallOption) (RazaQuery_GetSessionsClient, error)
	GetCommands(ctx context.Context, in *GetCommandsRequest, opts ...grpc.CallOption) (RazaQuery_GetCommandsClient, error)
}

type razaQueryClient struct {
	cc grpc.ClientConnInterface
}

func NewRazaQueryClient(cc grpc.ClientConnInterface) RazaQueryClient {
	return &razaQueryClient{cc}
}

func (c *razaQueryClient) GetSessions(ctx context.Context, in *GetSessionsRequest, opts ...grpc.CallOption) (RazaQuery_GetSessionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &RazaQuery_ServiceDesc.Streams[0], "/raza.RazaQuery/GetSessions", opts...)
	if err != nil {
		return nil, err
	}
	x := &razaQueryGetSessionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RazaQuery_GetSessionsClient interface {
	Recv() (*Session, error)
	grpc.ClientStream
}

type razaQueryGetSessionsClient struct {
	grpc.ClientStream
}

func (x *razaQueryGetSessionsClient) Recv() (*Session, error) {
	m := new(Session)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *razaQueryClient) GetCommands(ctx context.Context, in *GetCommandsRequest, opts ...grpc.CallOption) (RazaQuery_GetCommandsClient, error) {
	stream, err := c.cc.NewStream(ctx, &RazaQuery_ServiceDesc.Streams[1], "/raza.RazaQuery/GetCommands", opts...)
	if err != nil {
		return nil, err
	}
	x := &razaQueryGetCommandsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RazaQuery_GetCommandsClient interface {
	Recv() (*Command, error)
	grpc.ClientStream
}

type razaQueryGetCommandsClient struct {
	grpc.ClientStream
}

func (x *razaQueryGetCommandsClient) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RazaQueryServer is the server API for RazaQuery service.
// All implementations must embed UnimplementedRazaQueryServer
// for forward compatibility
type RazaQueryServer interface {
	GetSessions(*GetSessionsRequest, RazaQuery_GetSessionsServer) error
	GetCommands(*GetCommandsRequest, RazaQuery_GetCommandsServer) error
	mustEmbedUnimplementedRazaQueryServer()
}

// UnimplementedRazaQueryServer must be embedded to have forward compatible implementations.
type UnimplementedRazaQueryServer struct {
}

func (UnimplementedRazaQueryServer) GetSessions(*GetSessionsRequest, RazaQuery_GetSessionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSessions not implemented")
}
func (UnimplementedRazaQueryServer) GetCommands(*GetCommandsRequest, RazaQuery_GetCommandsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCommands not implemented")
}
func (UnimplementedRazaQueryServer) mustEmbedUnimplementedRazaQueryServer() {}

// UnsafeRazaQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RazaQueryServer will
// result in compilation errors.
type UnsafeRazaQueryServer interface {
	mustEmbedUnimplementedRazaQueryServer()
}

func RegisterRazaQueryServer(s grpc.ServiceRegistrar, srv RazaQueryServer) {
	s.RegisterService(&RazaQuery_ServiceDesc, srv)
}

func _RazaQuery_GetSessions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetSessionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RazaQueryServer).GetSessions(m, &razaQueryGetSessionsServer{stream})
}

type RazaQuery_GetSessionsServer interface {
	Send(*Session) error
	grpc.ServerStream
}

type razaQueryGetSessionsServer struct {
	grpc.ServerStream
}

func (x *razaQueryGetSessionsServer) Send(m *Session) error {
	return x.ServerStream.SendMsg(m)
}

func _RazaQuery_GetCommands_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetCommandsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RazaQueryServer).GetCommands(m, &razaQueryGetCommandsServer{stream})
}

type RazaQuery_GetCommandsServer interface {
	Send(*Command) error
	grpc.ServerStream
}

type razaQueryGetCommandsServer struct {
	grpc.ServerStream
}

func (x *razaQueryGetCommandsServer) Send(m *Command) error {
	return x.ServerStream.SendMsg(m)
}

// RazaQuery_ServiceDesc is the grpc.ServiceDesc for RazaQuery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RazaQuery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raza.RazaQuery",
	HandlerType: (*RazaQueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSessions",
			Handler:       _RazaQuery_GetSessions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetCommands",
			Handler:       _RazaQuery_GetCommands_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "raza_query.proto",
}
