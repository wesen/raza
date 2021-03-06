// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: raza_user.proto

package raza

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RazaUserClient is the client API for RazaUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RazaUserClient interface {
}

type razaUserClient struct {
	cc grpc.ClientConnInterface
}

func NewRazaUserClient(cc grpc.ClientConnInterface) RazaUserClient {
	return &razaUserClient{cc}
}

// RazaUserServer is the server API for RazaUser service.
// All implementations must embed UnimplementedRazaUserServer
// for forward compatibility
type RazaUserServer interface {
	mustEmbedUnimplementedRazaUserServer()
}

// UnimplementedRazaUserServer must be embedded to have forward compatible implementations.
type UnimplementedRazaUserServer struct {
}

func (UnimplementedRazaUserServer) mustEmbedUnimplementedRazaUserServer() {}

// UnsafeRazaUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RazaUserServer will
// result in compilation errors.
type UnsafeRazaUserServer interface {
	mustEmbedUnimplementedRazaUserServer()
}

func RegisterRazaUserServer(s grpc.ServiceRegistrar, srv RazaUserServer) {
	s.RegisterService(&RazaUser_ServiceDesc, srv)
}

// RazaUser_ServiceDesc is the grpc.ServiceDesc for RazaUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RazaUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raza.RazaUser",
	HandlerType: (*RazaUserServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "raza_user.proto",
}
