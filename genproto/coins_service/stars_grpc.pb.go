// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: stars.proto

package coins_service

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
	StarsService_Create_FullMethodName                  = "/coins_service.StarsService/Create"
	StarsService_GetById_FullMethodName                 = "/coins_service.StarsService/GetById"
	StarsService_GetList_FullMethodName                 = "/coins_service.StarsService/GetList"
	StarsService_Update_FullMethodName                  = "/coins_service.StarsService/Update"
	StarsService_CreateTransaction_FullMethodName       = "/coins_service.StarsService/CreateTransaction"
	StarsService_GetTransactionById_FullMethodName      = "/coins_service.StarsService/GetTransactionById"
	StarsService_GetListTransactionStars_FullMethodName = "/coins_service.StarsService/GetListTransactionStars"
)

// StarsServiceClient is the client API for StarsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StarsServiceClient interface {
	Create(ctx context.Context, in *CreateStars, opts ...grpc.CallOption) (*Stars, error)
	GetById(ctx context.Context, in *GetStarsPrimeryKey, opts ...grpc.CallOption) (*Stars, error)
	GetList(ctx context.Context, in *GetListStarsRequest, opts ...grpc.CallOption) (*GetListStarsResponse, error)
	Update(ctx context.Context, in *UpdateStars, opts ...grpc.CallOption) (*Empty, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionStars, opts ...grpc.CallOption) (*TransactionStars, error)
	GetTransactionById(ctx context.Context, in *TransactionPrimaryKey, opts ...grpc.CallOption) (*TransactionStars, error)
	GetListTransactionStars(ctx context.Context, in *GetListTransactionRequest, opts ...grpc.CallOption) (*GetListTransactionResponse, error)
}

type starsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStarsServiceClient(cc grpc.ClientConnInterface) StarsServiceClient {
	return &starsServiceClient{cc}
}

func (c *starsServiceClient) Create(ctx context.Context, in *CreateStars, opts ...grpc.CallOption) (*Stars, error) {
	out := new(Stars)
	err := c.cc.Invoke(ctx, StarsService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starsServiceClient) GetById(ctx context.Context, in *GetStarsPrimeryKey, opts ...grpc.CallOption) (*Stars, error) {
	out := new(Stars)
	err := c.cc.Invoke(ctx, StarsService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starsServiceClient) GetList(ctx context.Context, in *GetListStarsRequest, opts ...grpc.CallOption) (*GetListStarsResponse, error) {
	out := new(GetListStarsResponse)
	err := c.cc.Invoke(ctx, StarsService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starsServiceClient) Update(ctx context.Context, in *UpdateStars, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, StarsService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starsServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionStars, opts ...grpc.CallOption) (*TransactionStars, error) {
	out := new(TransactionStars)
	err := c.cc.Invoke(ctx, StarsService_CreateTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starsServiceClient) GetTransactionById(ctx context.Context, in *TransactionPrimaryKey, opts ...grpc.CallOption) (*TransactionStars, error) {
	out := new(TransactionStars)
	err := c.cc.Invoke(ctx, StarsService_GetTransactionById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starsServiceClient) GetListTransactionStars(ctx context.Context, in *GetListTransactionRequest, opts ...grpc.CallOption) (*GetListTransactionResponse, error) {
	out := new(GetListTransactionResponse)
	err := c.cc.Invoke(ctx, StarsService_GetListTransactionStars_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StarsServiceServer is the server API for StarsService service.
// All implementations must embed UnimplementedStarsServiceServer
// for forward compatibility
type StarsServiceServer interface {
	Create(context.Context, *CreateStars) (*Stars, error)
	GetById(context.Context, *GetStarsPrimeryKey) (*Stars, error)
	GetList(context.Context, *GetListStarsRequest) (*GetListStarsResponse, error)
	Update(context.Context, *UpdateStars) (*Empty, error)
	CreateTransaction(context.Context, *CreateTransactionStars) (*TransactionStars, error)
	GetTransactionById(context.Context, *TransactionPrimaryKey) (*TransactionStars, error)
	GetListTransactionStars(context.Context, *GetListTransactionRequest) (*GetListTransactionResponse, error)
	mustEmbedUnimplementedStarsServiceServer()
}

// UnimplementedStarsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStarsServiceServer struct {
}

func (UnimplementedStarsServiceServer) Create(context.Context, *CreateStars) (*Stars, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStarsServiceServer) GetById(context.Context, *GetStarsPrimeryKey) (*Stars, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedStarsServiceServer) GetList(context.Context, *GetListStarsRequest) (*GetListStarsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedStarsServiceServer) Update(context.Context, *UpdateStars) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStarsServiceServer) CreateTransaction(context.Context, *CreateTransactionStars) (*TransactionStars, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedStarsServiceServer) GetTransactionById(context.Context, *TransactionPrimaryKey) (*TransactionStars, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionById not implemented")
}
func (UnimplementedStarsServiceServer) GetListTransactionStars(context.Context, *GetListTransactionRequest) (*GetListTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListTransactionStars not implemented")
}
func (UnimplementedStarsServiceServer) mustEmbedUnimplementedStarsServiceServer() {}

// UnsafeStarsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StarsServiceServer will
// result in compilation errors.
type UnsafeStarsServiceServer interface {
	mustEmbedUnimplementedStarsServiceServer()
}

func RegisterStarsServiceServer(s grpc.ServiceRegistrar, srv StarsServiceServer) {
	s.RegisterService(&StarsService_ServiceDesc, srv)
}

func _StarsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStars)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StarsService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarsServiceServer).Create(ctx, req.(*CreateStars))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarsService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStarsPrimeryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarsServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StarsService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarsServiceServer).GetById(ctx, req.(*GetStarsPrimeryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarsService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListStarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarsServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StarsService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarsServiceServer).GetList(ctx, req.(*GetListStarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStars)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StarsService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarsServiceServer).Update(ctx, req.(*UpdateStars))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarsService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionStars)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarsServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StarsService_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarsServiceServer).CreateTransaction(ctx, req.(*CreateTransactionStars))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarsService_GetTransactionById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarsServiceServer).GetTransactionById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StarsService_GetTransactionById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarsServiceServer).GetTransactionById(ctx, req.(*TransactionPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarsService_GetListTransactionStars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarsServiceServer).GetListTransactionStars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StarsService_GetListTransactionStars_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarsServiceServer).GetListTransactionStars(ctx, req.(*GetListTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StarsService_ServiceDesc is the grpc.ServiceDesc for StarsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StarsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coins_service.StarsService",
	HandlerType: (*StarsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StarsService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _StarsService_GetById_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _StarsService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StarsService_Update_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _StarsService_CreateTransaction_Handler,
		},
		{
			MethodName: "GetTransactionById",
			Handler:    _StarsService_GetTransactionById_Handler,
		},
		{
			MethodName: "GetListTransactionStars",
			Handler:    _StarsService_GetListTransactionStars_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stars.proto",
}
