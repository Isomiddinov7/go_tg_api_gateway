// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: coin.proto

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
	CoinsService_Create_FullMethodName  = "/coins_service.CoinsService/Create"
	CoinsService_GetById_FullMethodName = "/coins_service.CoinsService/GetById"
	CoinsService_GetList_FullMethodName = "/coins_service.CoinsService/GetList"
	CoinsService_Update_FullMethodName  = "/coins_service.CoinsService/Update"
	CoinsService_Delete_FullMethodName  = "/coins_service.CoinsService/Delete"
)

// CoinsServiceClient is the client API for CoinsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoinsServiceClient interface {
	Create(ctx context.Context, in *CreateCoin, opts ...grpc.CallOption) (*CoinPrimaryKey, error)
	GetById(ctx context.Context, in *CoinPrimaryKey, opts ...grpc.CallOption) (*Coin, error)
	GetList(ctx context.Context, in *GetListCoinRequest, opts ...grpc.CallOption) (*GetListCoinResponse, error)
	Update(ctx context.Context, in *UpdateCoin, opts ...grpc.CallOption) (*Coin, error)
	Delete(ctx context.Context, in *CoinPrimaryKey, opts ...grpc.CallOption) (*Empty, error)
}

type coinsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoinsServiceClient(cc grpc.ClientConnInterface) CoinsServiceClient {
	return &coinsServiceClient{cc}
}

func (c *coinsServiceClient) Create(ctx context.Context, in *CreateCoin, opts ...grpc.CallOption) (*CoinPrimaryKey, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CoinPrimaryKey)
	err := c.cc.Invoke(ctx, CoinsService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinsServiceClient) GetById(ctx context.Context, in *CoinPrimaryKey, opts ...grpc.CallOption) (*Coin, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Coin)
	err := c.cc.Invoke(ctx, CoinsService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinsServiceClient) GetList(ctx context.Context, in *GetListCoinRequest, opts ...grpc.CallOption) (*GetListCoinResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetListCoinResponse)
	err := c.cc.Invoke(ctx, CoinsService_GetList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinsServiceClient) Update(ctx context.Context, in *UpdateCoin, opts ...grpc.CallOption) (*Coin, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Coin)
	err := c.cc.Invoke(ctx, CoinsService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinsServiceClient) Delete(ctx context.Context, in *CoinPrimaryKey, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, CoinsService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoinsServiceServer is the server API for CoinsService service.
// All implementations must embed UnimplementedCoinsServiceServer
// for forward compatibility.
type CoinsServiceServer interface {
	Create(context.Context, *CreateCoin) (*CoinPrimaryKey, error)
	GetById(context.Context, *CoinPrimaryKey) (*Coin, error)
	GetList(context.Context, *GetListCoinRequest) (*GetListCoinResponse, error)
	Update(context.Context, *UpdateCoin) (*Coin, error)
	Delete(context.Context, *CoinPrimaryKey) (*Empty, error)
	mustEmbedUnimplementedCoinsServiceServer()
}

// UnimplementedCoinsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCoinsServiceServer struct{}

func (UnimplementedCoinsServiceServer) Create(context.Context, *CreateCoin) (*CoinPrimaryKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCoinsServiceServer) GetById(context.Context, *CoinPrimaryKey) (*Coin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedCoinsServiceServer) GetList(context.Context, *GetListCoinRequest) (*GetListCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedCoinsServiceServer) Update(context.Context, *UpdateCoin) (*Coin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCoinsServiceServer) Delete(context.Context, *CoinPrimaryKey) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCoinsServiceServer) mustEmbedUnimplementedCoinsServiceServer() {}
func (UnimplementedCoinsServiceServer) testEmbeddedByValue()                      {}

// UnsafeCoinsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoinsServiceServer will
// result in compilation errors.
type UnsafeCoinsServiceServer interface {
	mustEmbedUnimplementedCoinsServiceServer()
}

func RegisterCoinsServiceServer(s grpc.ServiceRegistrar, srv CoinsServiceServer) {
	// If the following call pancis, it indicates UnimplementedCoinsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CoinsService_ServiceDesc, srv)
}

func _CoinsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoinsService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinsServiceServer).Create(ctx, req.(*CreateCoin))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoinsService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoinPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinsServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoinsService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinsServiceServer).GetById(ctx, req.(*CoinPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoinsService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinsServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoinsService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinsServiceServer).GetList(ctx, req.(*GetListCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoinsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoinsService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinsServiceServer).Update(ctx, req.(*UpdateCoin))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoinsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoinPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoinsService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinsServiceServer).Delete(ctx, req.(*CoinPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// CoinsService_ServiceDesc is the grpc.ServiceDesc for CoinsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoinsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coins_service.CoinsService",
	HandlerType: (*CoinsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CoinsService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _CoinsService_GetById_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _CoinsService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CoinsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CoinsService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coin.proto",
}
