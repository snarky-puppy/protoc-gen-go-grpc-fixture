// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: use_some_lib.proto

package example

import (
	context "context"
	some_lib "github.com/snarky-puppy/protoc-gen-go-grpc-fixture/example/some_lib"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SomeService_SomeMethod_FullMethodName = "/Some.Lib.SomeService/SomeMethod"
)

// SomeServiceClient is the client API for SomeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SomeServiceClient interface {
	SomeMethod(ctx context.Context, in *some_lib.SomeMessage, opts ...grpc.CallOption) (*some_lib.SomeMessage, error)
}

type someServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSomeServiceClient(cc grpc.ClientConnInterface) SomeServiceClient {
	return &someServiceClient{cc}
}

func (c *someServiceClient) SomeMethod(ctx context.Context, in *some_lib.SomeMessage, opts ...grpc.CallOption) (*some_lib.SomeMessage, error) {
	out := new(some_lib.SomeMessage)
	err := c.cc.Invoke(ctx, SomeService_SomeMethod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SomeServiceServer is the server API for SomeService service.
// All implementations must embed UnimplementedSomeServiceServer
// for forward compatibility
type SomeServiceServer interface {
	SomeMethod(context.Context, *some_lib.SomeMessage) (*some_lib.SomeMessage, error)
	mustEmbedUnimplementedSomeServiceServer()
}

// UnimplementedSomeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSomeServiceServer struct {
}

func (UnimplementedSomeServiceServer) SomeMethod(context.Context, *some_lib.SomeMessage) (*some_lib.SomeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SomeMethod not implemented")
}
func (UnimplementedSomeServiceServer) mustEmbedUnimplementedSomeServiceServer() {}

// UnsafeSomeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SomeServiceServer will
// result in compilation errors.
type UnsafeSomeServiceServer interface {
	mustEmbedUnimplementedSomeServiceServer()
}

func RegisterSomeServiceServer(s grpc.ServiceRegistrar, srv SomeServiceServer) {
	s.RegisterService(&SomeService_ServiceDesc, srv)
}

func _SomeService_SomeMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(some_lib.SomeMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SomeServiceServer).SomeMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SomeService_SomeMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SomeServiceServer).SomeMethod(ctx, req.(*some_lib.SomeMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// SomeService_ServiceDesc is the grpc.ServiceDesc for SomeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Some.Lib.SomeService",
	HandlerType: (*SomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SomeMethod",
			Handler:    _SomeService_SomeMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "use_some_lib.proto",
}
