// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package server

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

// SampleClient is the client API for Sample service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleClient interface {
	// rpc function arg is for request, return is response
	Sample(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error)
}

type sampleClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleClient(cc grpc.ClientConnInterface) SampleClient {
	return &sampleClient{cc}
}

func (c *sampleClient) Sample(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/sample.Sample/Sample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleServer is the server API for Sample service.
// All implementations must embed UnimplementedSampleServer
// for forward compatibility
type SampleServer interface {
	// rpc function arg is for request, return is response
	Sample(con context.Context, in *SampleRequest) (*SampleResponse, error)
	mustEmbedUnimplementedSampleServer()
}

// UnimplementedSampleServer must be embedded to have forward compatible implementations.
type UnimplementedSampleServer struct {
}

func (UnimplementedSampleServer) Sample(con context.Context, in *SampleRequest) (*SampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sample not implemented")
}
func (UnimplementedSampleServer) mustEmbedUnimplementedSampleServer() {}

// UnsafeSampleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleServer will
// result in compilation errors.
type UnsafeSampleServer interface {
	mustEmbedUnimplementedSampleServer()
}

func RegisterSampleServer(s grpc.ServiceRegistrar, srv SampleServer) {
	s.RegisterService(&Sample_ServiceDesc, srv)
}

func _Sample_Sample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).Sample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Sample/Sample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).Sample(ctx, req.(*SampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sample_ServiceDesc is the grpc.ServiceDesc for Sample service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sample_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sample.Sample",
	HandlerType: (*SampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sample",
			Handler:    _Sample_Sample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample.proto",
}
