// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: buy_or_sell.proto

package coins_service

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
	BuyOrSell_GetSell_FullMethodName     = "/coins_service.BuyOrSell/GetSell"
	BuyOrSell_GetBuy_FullMethodName      = "/coins_service.BuyOrSell/GetBuy"
	BuyOrSell_GetUserBuy_FullMethodName  = "/coins_service.BuyOrSell/GetUserBuy"
	BuyOrSell_GetUserSell_FullMethodName = "/coins_service.BuyOrSell/GetUserSell"
)

// BuyOrSellClient is the client API for BuyOrSell service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuyOrSellClient interface {
	GetSell(ctx context.Context, in *BuyOrSellRequest, opts ...grpc.CallOption) (*BuyOrSellResponse, error)
	GetBuy(ctx context.Context, in *BuyOrSellRequest, opts ...grpc.CallOption) (*BuyOrSellResponse, error)
	GetUserBuy(ctx context.Context, in *GetListUserBuyOrSellRequest, opts ...grpc.CallOption) (*UserBuy, error)
	GetUserSell(ctx context.Context, in *GetListUserBuyOrSellRequest, opts ...grpc.CallOption) (*UserSell, error)
}

type buyOrSellClient struct {
	cc grpc.ClientConnInterface
}

func NewBuyOrSellClient(cc grpc.ClientConnInterface) BuyOrSellClient {
	return &buyOrSellClient{cc}
}

func (c *buyOrSellClient) GetSell(ctx context.Context, in *BuyOrSellRequest, opts ...grpc.CallOption) (*BuyOrSellResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BuyOrSellResponse)
	err := c.cc.Invoke(ctx, BuyOrSell_GetSell_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buyOrSellClient) GetBuy(ctx context.Context, in *BuyOrSellRequest, opts ...grpc.CallOption) (*BuyOrSellResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BuyOrSellResponse)
	err := c.cc.Invoke(ctx, BuyOrSell_GetBuy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buyOrSellClient) GetUserBuy(ctx context.Context, in *GetListUserBuyOrSellRequest, opts ...grpc.CallOption) (*UserBuy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserBuy)
	err := c.cc.Invoke(ctx, BuyOrSell_GetUserBuy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buyOrSellClient) GetUserSell(ctx context.Context, in *GetListUserBuyOrSellRequest, opts ...grpc.CallOption) (*UserSell, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserSell)
	err := c.cc.Invoke(ctx, BuyOrSell_GetUserSell_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuyOrSellServer is the server API for BuyOrSell service.
// All implementations must embed UnimplementedBuyOrSellServer
// for forward compatibility.
type BuyOrSellServer interface {
	GetSell(context.Context, *BuyOrSellRequest) (*BuyOrSellResponse, error)
	GetBuy(context.Context, *BuyOrSellRequest) (*BuyOrSellResponse, error)
	GetUserBuy(context.Context, *GetListUserBuyOrSellRequest) (*UserBuy, error)
	GetUserSell(context.Context, *GetListUserBuyOrSellRequest) (*UserSell, error)
	mustEmbedUnimplementedBuyOrSellServer()
}

// UnimplementedBuyOrSellServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBuyOrSellServer struct{}

func (UnimplementedBuyOrSellServer) GetSell(context.Context, *BuyOrSellRequest) (*BuyOrSellResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSell not implemented")
}
func (UnimplementedBuyOrSellServer) GetBuy(context.Context, *BuyOrSellRequest) (*BuyOrSellResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuy not implemented")
}
func (UnimplementedBuyOrSellServer) GetUserBuy(context.Context, *GetListUserBuyOrSellRequest) (*UserBuy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBuy not implemented")
}
func (UnimplementedBuyOrSellServer) GetUserSell(context.Context, *GetListUserBuyOrSellRequest) (*UserSell, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSell not implemented")
}
func (UnimplementedBuyOrSellServer) mustEmbedUnimplementedBuyOrSellServer() {}
func (UnimplementedBuyOrSellServer) testEmbeddedByValue()                   {}

// UnsafeBuyOrSellServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuyOrSellServer will
// result in compilation errors.
type UnsafeBuyOrSellServer interface {
	mustEmbedUnimplementedBuyOrSellServer()
}

func RegisterBuyOrSellServer(s grpc.ServiceRegistrar, srv BuyOrSellServer) {
	// If the following call pancis, it indicates UnimplementedBuyOrSellServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BuyOrSell_ServiceDesc, srv)
}

func _BuyOrSell_GetSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyOrSellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuyOrSellServer).GetSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BuyOrSell_GetSell_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuyOrSellServer).GetSell(ctx, req.(*BuyOrSellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuyOrSell_GetBuy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyOrSellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuyOrSellServer).GetBuy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BuyOrSell_GetBuy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuyOrSellServer).GetBuy(ctx, req.(*BuyOrSellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuyOrSell_GetUserBuy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListUserBuyOrSellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuyOrSellServer).GetUserBuy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BuyOrSell_GetUserBuy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuyOrSellServer).GetUserBuy(ctx, req.(*GetListUserBuyOrSellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuyOrSell_GetUserSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListUserBuyOrSellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuyOrSellServer).GetUserSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BuyOrSell_GetUserSell_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuyOrSellServer).GetUserSell(ctx, req.(*GetListUserBuyOrSellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BuyOrSell_ServiceDesc is the grpc.ServiceDesc for BuyOrSell service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BuyOrSell_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coins_service.BuyOrSell",
	HandlerType: (*BuyOrSellServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSell",
			Handler:    _BuyOrSell_GetSell_Handler,
		},
		{
			MethodName: "GetBuy",
			Handler:    _BuyOrSell_GetBuy_Handler,
		},
		{
			MethodName: "GetUserBuy",
			Handler:    _BuyOrSell_GetUserBuy_Handler,
		},
		{
			MethodName: "GetUserSell",
			Handler:    _BuyOrSell_GetUserSell_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buy_or_sell.proto",
}
