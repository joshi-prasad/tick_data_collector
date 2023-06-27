// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: proto/data_collector.proto

package proto_tick_data_collector

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

const (
	DataCollectionService_SayHello_FullMethodName = "/data_collector.DataCollectionService/SayHello"
	DataCollectionService_Tick_FullMethodName     = "/data_collector.DataCollectionService/Tick"
)

// DataCollectionServiceClient is the client API for DataCollectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataCollectionServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	Tick(ctx context.Context, in *TickData, opts ...grpc.CallOption) (*TickResponse, error)
}

type dataCollectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataCollectionServiceClient(cc grpc.ClientConnInterface) DataCollectionServiceClient {
	return &dataCollectionServiceClient{cc}
}

func (c *dataCollectionServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, DataCollectionService_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataCollectionServiceClient) Tick(ctx context.Context, in *TickData, opts ...grpc.CallOption) (*TickResponse, error) {
	out := new(TickResponse)
	err := c.cc.Invoke(ctx, DataCollectionService_Tick_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataCollectionServiceServer is the server API for DataCollectionService service.
// All implementations must embed UnimplementedDataCollectionServiceServer
// for forward compatibility
type DataCollectionServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	Tick(context.Context, *TickData) (*TickResponse, error)
	mustEmbedUnimplementedDataCollectionServiceServer()
}

// UnimplementedDataCollectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataCollectionServiceServer struct {
}

func (UnimplementedDataCollectionServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedDataCollectionServiceServer) Tick(context.Context, *TickData) (*TickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tick not implemented")
}
func (UnimplementedDataCollectionServiceServer) mustEmbedUnimplementedDataCollectionServiceServer() {}

// UnsafeDataCollectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataCollectionServiceServer will
// result in compilation errors.
type UnsafeDataCollectionServiceServer interface {
	mustEmbedUnimplementedDataCollectionServiceServer()
}

func RegisterDataCollectionServiceServer(s grpc.ServiceRegistrar, srv DataCollectionServiceServer) {
	s.RegisterService(&DataCollectionService_ServiceDesc, srv)
}

func _DataCollectionService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataCollectionServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataCollectionService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataCollectionServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataCollectionService_Tick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TickData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataCollectionServiceServer).Tick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataCollectionService_Tick_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataCollectionServiceServer).Tick(ctx, req.(*TickData))
	}
	return interceptor(ctx, in, info, handler)
}

// DataCollectionService_ServiceDesc is the grpc.ServiceDesc for DataCollectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataCollectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "data_collector.DataCollectionService",
	HandlerType: (*DataCollectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _DataCollectionService_SayHello_Handler,
		},
		{
			MethodName: "Tick",
			Handler:    _DataCollectionService_Tick_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/data_collector.proto",
}