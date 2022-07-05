// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package storagepb

import (
	context "context"
	ceresprompb "github.com/CeresDB/ceresdbproto/pkg/ceresprompb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	Route(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	StreamWrite(ctx context.Context, opts ...grpc.CallOption) (StorageService_StreamWriteClient, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	StreamQuery(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (StorageService_StreamQueryClient, error)
	PromQuery(ctx context.Context, in *ceresprompb.PrometheusQueryRequest, opts ...grpc.CallOption) (*ceresprompb.PrometheusQueryResponse, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) Route(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	out := new(RouteResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) StreamWrite(ctx context.Context, opts ...grpc.CallOption) (StorageService_StreamWriteClient, error) {
	stream, err := c.cc.NewStream(ctx, &StorageService_ServiceDesc.Streams[0], "/storage.StorageService/StreamWrite", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageServiceStreamWriteClient{stream}
	return x, nil
}

type StorageService_StreamWriteClient interface {
	Send(*WriteRequest) error
	CloseAndRecv() (*WriteResponse, error)
	grpc.ClientStream
}

type storageServiceStreamWriteClient struct {
	grpc.ClientStream
}

func (x *storageServiceStreamWriteClient) Send(m *WriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageServiceStreamWriteClient) CloseAndRecv() (*WriteResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) StreamQuery(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (StorageService_StreamQueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &StorageService_ServiceDesc.Streams[1], "/storage.StorageService/StreamQuery", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageServiceStreamQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StorageService_StreamQueryClient interface {
	Recv() (*QueryResponse, error)
	grpc.ClientStream
}

type storageServiceStreamQueryClient struct {
	grpc.ClientStream
}

func (x *storageServiceStreamQueryClient) Recv() (*QueryResponse, error) {
	m := new(QueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageServiceClient) PromQuery(ctx context.Context, in *ceresprompb.PrometheusQueryRequest, opts ...grpc.CallOption) (*ceresprompb.PrometheusQueryResponse, error) {
	out := new(ceresprompb.PrometheusQueryResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/PromQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
// All implementations must embed UnimplementedStorageServiceServer
// for forward compatibility
type StorageServiceServer interface {
	Route(context.Context, *RouteRequest) (*RouteResponse, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	StreamWrite(StorageService_StreamWriteServer) error
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	StreamQuery(*QueryRequest, StorageService_StreamQueryServer) error
	PromQuery(context.Context, *ceresprompb.PrometheusQueryRequest) (*ceresprompb.PrometheusQueryResponse, error)
	mustEmbedUnimplementedStorageServiceServer()
}

// UnimplementedStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServiceServer struct {
}

func (UnimplementedStorageServiceServer) Route(context.Context, *RouteRequest) (*RouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (UnimplementedStorageServiceServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedStorageServiceServer) StreamWrite(StorageService_StreamWriteServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamWrite not implemented")
}
func (UnimplementedStorageServiceServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedStorageServiceServer) StreamQuery(*QueryRequest, StorageService_StreamQueryServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamQuery not implemented")
}
func (UnimplementedStorageServiceServer) PromQuery(context.Context, *ceresprompb.PrometheusQueryRequest) (*ceresprompb.PrometheusQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromQuery not implemented")
}
func (UnimplementedStorageServiceServer) mustEmbedUnimplementedStorageServiceServer() {}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).Route(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_StreamWrite_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServiceServer).StreamWrite(&storageServiceStreamWriteServer{stream})
}

type StorageService_StreamWriteServer interface {
	SendAndClose(*WriteResponse) error
	Recv() (*WriteRequest, error)
	grpc.ServerStream
}

type storageServiceStreamWriteServer struct {
	grpc.ServerStream
}

func (x *storageServiceStreamWriteServer) SendAndClose(m *WriteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageServiceStreamWriteServer) Recv() (*WriteRequest, error) {
	m := new(WriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StorageService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_StreamQuery_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServiceServer).StreamQuery(m, &storageServiceStreamQueryServer{stream})
}

type StorageService_StreamQueryServer interface {
	Send(*QueryResponse) error
	grpc.ServerStream
}

type storageServiceStreamQueryServer struct {
	grpc.ServerStream
}

func (x *storageServiceStreamQueryServer) Send(m *QueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StorageService_PromQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ceresprompb.PrometheusQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).PromQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/PromQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).PromQuery(ctx, req.(*ceresprompb.PrometheusQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Route",
			Handler:    _StorageService_Route_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _StorageService_Write_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _StorageService_Query_Handler,
		},
		{
			MethodName: "PromQuery",
			Handler:    _StorageService_PromQuery_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamWrite",
			Handler:       _StorageService_StreamWrite_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamQuery",
			Handler:       _StorageService_StreamQuery_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "storage.proto",
}