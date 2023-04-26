// Code generated by protoc-gen-go-grpc-fixture. DO NOT EDIT.
// source: use_some_lib.proto

package example

import (
	context "context"
	protoc_gen_go_grpc_fixture "github.com/snarky-puppy/protoc-gen-go-grpc-fixture"
	some_lib "github.com/snarky-puppy/protoc-gen-go-grpc-fixture/example/some_lib"
	grpc "google.golang.org/grpc"
	path "path"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SomeServiceClient is the fixture stub API for SomeService service.
type someServiceFixtures struct {
	baseDir string
}

func NewSomeServiceFixtures(baseDir string) SomeServiceClient {
	return &someServiceFixtures{baseDir: baseDir}
}

func (c *someServiceFixtures) SomeMethod(ctx context.Context, in *some_lib.SomeMessage, opts ...grpc.CallOption) (*some_lib.SomeMessage, error) {
	out := new(some_lib.SomeMessage)
	return protoc_gen_go_grpc_fixture.Fixture(path.Join(c.baseDir, SomeService_SomeMethod_FullMethodName+".json"), out)
}
