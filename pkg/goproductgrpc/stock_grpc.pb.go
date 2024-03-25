// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/goproductgrpc/stock.proto

package goproductgrpc

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

// StockClient is the client API for Stock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockClient interface {
	IncrementStock(ctx context.Context, in *ReqIncrementStock, opts ...grpc.CallOption) (*StockEmpty, error)
	DecrementStock(ctx context.Context, in *ReqDecrementStock, opts ...grpc.CallOption) (*StockEmpty, error)
}

type stockClient struct {
	cc grpc.ClientConnInterface
}

func NewStockClient(cc grpc.ClientConnInterface) StockClient {
	return &stockClient{cc}
}

func (c *stockClient) IncrementStock(ctx context.Context, in *ReqIncrementStock, opts ...grpc.CallOption) (*StockEmpty, error) {
	out := new(StockEmpty)
	err := c.cc.Invoke(ctx, "/goproductgrpc.Stock/IncrementStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) DecrementStock(ctx context.Context, in *ReqDecrementStock, opts ...grpc.CallOption) (*StockEmpty, error) {
	out := new(StockEmpty)
	err := c.cc.Invoke(ctx, "/goproductgrpc.Stock/DecrementStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServer is the server API for Stock service.
// All implementations must embed UnimplementedStockServer
// for forward compatibility
type StockServer interface {
	IncrementStock(context.Context, *ReqIncrementStock) (*StockEmpty, error)
	DecrementStock(context.Context, *ReqDecrementStock) (*StockEmpty, error)
	mustEmbedUnimplementedStockServer()
}

// UnimplementedStockServer must be embedded to have forward compatible implementations.
type UnimplementedStockServer struct {
}

func (UnimplementedStockServer) IncrementStock(context.Context, *ReqIncrementStock) (*StockEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementStock not implemented")
}
func (UnimplementedStockServer) DecrementStock(context.Context, *ReqDecrementStock) (*StockEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrementStock not implemented")
}
func (UnimplementedStockServer) mustEmbedUnimplementedStockServer() {}

// UnsafeStockServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServer will
// result in compilation errors.
type UnsafeStockServer interface {
	mustEmbedUnimplementedStockServer()
}

func RegisterStockServer(s grpc.ServiceRegistrar, srv StockServer) {
	s.RegisterService(&Stock_ServiceDesc, srv)
}

func _Stock_IncrementStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqIncrementStock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).IncrementStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goproductgrpc.Stock/IncrementStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).IncrementStock(ctx, req.(*ReqIncrementStock))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_DecrementStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDecrementStock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).DecrementStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goproductgrpc.Stock/DecrementStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).DecrementStock(ctx, req.(*ReqDecrementStock))
	}
	return interceptor(ctx, in, info, handler)
}

// Stock_ServiceDesc is the grpc.ServiceDesc for Stock service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stock_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goproductgrpc.Stock",
	HandlerType: (*StockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncrementStock",
			Handler:    _Stock_IncrementStock_Handler,
		},
		{
			MethodName: "DecrementStock",
			Handler:    _Stock_DecrementStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/goproductgrpc/stock.proto",
}
