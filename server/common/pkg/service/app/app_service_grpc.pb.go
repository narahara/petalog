// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: service/app/app_service.proto

package app

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AppService_Ping_FullMethodName = "/petalog.service.app.AppService/Ping"
)

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, AppService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations should embed UnimplementedAppServiceServer
// for forward compatibility.
type AppServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// UnimplementedAppServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAppServiceServer struct{}

func (UnimplementedAppServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAppServiceServer) testEmbeddedByValue() {}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	// If the following call pancis, it indicates UnimplementedAppServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "petalog.service.app.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AppService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/app/app_service.proto",
}