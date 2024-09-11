// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: user_sell_or_buy.proto

package users_service

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
	UserSellOrBuyService_UserSell_FullMethodName               = "/users_service.UserSellOrBuyService/UserSell"
	UserSellOrBuyService_UserBuy_FullMethodName                = "/users_service.UserSellOrBuyService/UserBuy"
	UserSellOrBuyService_GetByIdTransactionSell_FullMethodName = "/users_service.UserSellOrBuyService/GetByIdTransactionSell"
	UserSellOrBuyService_GetByIdTransactionBuy_FullMethodName  = "/users_service.UserSellOrBuyService/GetByIdTransactionBuy"
	UserSellOrBuyService_TransactionUpdate_FullMethodName      = "/users_service.UserSellOrBuyService/TransactionUpdate"
	UserSellOrBuyService_AllUserSell_FullMethodName            = "/users_service.UserSellOrBuyService/AllUserSell"
	UserSellOrBuyService_AllUserBuy_FullMethodName             = "/users_service.UserSellOrBuyService/AllUserBuy"
)

// UserSellOrBuyServiceClient is the client API for UserSellOrBuyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserSellOrBuyServiceClient interface {
	UserSell(ctx context.Context, in *UserSellRequest, opts ...grpc.CallOption) (*Empty, error)
	UserBuy(ctx context.Context, in *UserBuyRequest, opts ...grpc.CallOption) (*Empty, error)
	GetByIdTransactionSell(ctx context.Context, in *TransactioPrimaryKey, opts ...grpc.CallOption) (*UserTransactionSell, error)
	GetByIdTransactionBuy(ctx context.Context, in *TransactioPrimaryKey, opts ...grpc.CallOption) (*UserTransactionBuy, error)
	TransactionUpdate(ctx context.Context, in *UpdateTransaction, opts ...grpc.CallOption) (*Empty, error)
	AllUserSell(ctx context.Context, in *GetListUserTransactionRequest, opts ...grpc.CallOption) (*GetListUserSellTransactionResponse, error)
	AllUserBuy(ctx context.Context, in *GetListUserTransactionRequest, opts ...grpc.CallOption) (*GetListUserBuyTransactionResponse, error)
}

type userSellOrBuyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSellOrBuyServiceClient(cc grpc.ClientConnInterface) UserSellOrBuyServiceClient {
	return &userSellOrBuyServiceClient{cc}
}

func (c *userSellOrBuyServiceClient) UserSell(ctx context.Context, in *UserSellRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserSellOrBuyService_UserSell_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSellOrBuyServiceClient) UserBuy(ctx context.Context, in *UserBuyRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserSellOrBuyService_UserBuy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSellOrBuyServiceClient) GetByIdTransactionSell(ctx context.Context, in *TransactioPrimaryKey, opts ...grpc.CallOption) (*UserTransactionSell, error) {
	out := new(UserTransactionSell)
	err := c.cc.Invoke(ctx, UserSellOrBuyService_GetByIdTransactionSell_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSellOrBuyServiceClient) GetByIdTransactionBuy(ctx context.Context, in *TransactioPrimaryKey, opts ...grpc.CallOption) (*UserTransactionBuy, error) {
	out := new(UserTransactionBuy)
	err := c.cc.Invoke(ctx, UserSellOrBuyService_GetByIdTransactionBuy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSellOrBuyServiceClient) TransactionUpdate(ctx context.Context, in *UpdateTransaction, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserSellOrBuyService_TransactionUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSellOrBuyServiceClient) AllUserSell(ctx context.Context, in *GetListUserTransactionRequest, opts ...grpc.CallOption) (*GetListUserSellTransactionResponse, error) {
	out := new(GetListUserSellTransactionResponse)
	err := c.cc.Invoke(ctx, UserSellOrBuyService_AllUserSell_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSellOrBuyServiceClient) AllUserBuy(ctx context.Context, in *GetListUserTransactionRequest, opts ...grpc.CallOption) (*GetListUserBuyTransactionResponse, error) {
	out := new(GetListUserBuyTransactionResponse)
	err := c.cc.Invoke(ctx, UserSellOrBuyService_AllUserBuy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSellOrBuyServiceServer is the server API for UserSellOrBuyService service.
// All implementations must embed UnimplementedUserSellOrBuyServiceServer
// for forward compatibility
type UserSellOrBuyServiceServer interface {
	UserSell(context.Context, *UserSellRequest) (*Empty, error)
	UserBuy(context.Context, *UserBuyRequest) (*Empty, error)
	GetByIdTransactionSell(context.Context, *TransactioPrimaryKey) (*UserTransactionSell, error)
	GetByIdTransactionBuy(context.Context, *TransactioPrimaryKey) (*UserTransactionBuy, error)
	TransactionUpdate(context.Context, *UpdateTransaction) (*Empty, error)
	AllUserSell(context.Context, *GetListUserTransactionRequest) (*GetListUserSellTransactionResponse, error)
	AllUserBuy(context.Context, *GetListUserTransactionRequest) (*GetListUserBuyTransactionResponse, error)
	mustEmbedUnimplementedUserSellOrBuyServiceServer()
}

// UnimplementedUserSellOrBuyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserSellOrBuyServiceServer struct {
}

func (UnimplementedUserSellOrBuyServiceServer) UserSell(context.Context, *UserSellRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSell not implemented")
}
func (UnimplementedUserSellOrBuyServiceServer) UserBuy(context.Context, *UserBuyRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserBuy not implemented")
}
func (UnimplementedUserSellOrBuyServiceServer) GetByIdTransactionSell(context.Context, *TransactioPrimaryKey) (*UserTransactionSell, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdTransactionSell not implemented")
}
func (UnimplementedUserSellOrBuyServiceServer) GetByIdTransactionBuy(context.Context, *TransactioPrimaryKey) (*UserTransactionBuy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdTransactionBuy not implemented")
}
func (UnimplementedUserSellOrBuyServiceServer) TransactionUpdate(context.Context, *UpdateTransaction) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactionUpdate not implemented")
}
func (UnimplementedUserSellOrBuyServiceServer) AllUserSell(context.Context, *GetListUserTransactionRequest) (*GetListUserSellTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllUserSell not implemented")
}
func (UnimplementedUserSellOrBuyServiceServer) AllUserBuy(context.Context, *GetListUserTransactionRequest) (*GetListUserBuyTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllUserBuy not implemented")
}
func (UnimplementedUserSellOrBuyServiceServer) mustEmbedUnimplementedUserSellOrBuyServiceServer() {}

// UnsafeUserSellOrBuyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserSellOrBuyServiceServer will
// result in compilation errors.
type UnsafeUserSellOrBuyServiceServer interface {
	mustEmbedUnimplementedUserSellOrBuyServiceServer()
}

func RegisterUserSellOrBuyServiceServer(s grpc.ServiceRegistrar, srv UserSellOrBuyServiceServer) {
	s.RegisterService(&UserSellOrBuyService_ServiceDesc, srv)
}

func _UserSellOrBuyService_UserSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSellOrBuyServiceServer).UserSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSellOrBuyService_UserSell_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSellOrBuyServiceServer).UserSell(ctx, req.(*UserSellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSellOrBuyService_UserBuy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBuyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSellOrBuyServiceServer).UserBuy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSellOrBuyService_UserBuy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSellOrBuyServiceServer).UserBuy(ctx, req.(*UserBuyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSellOrBuyService_GetByIdTransactionSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactioPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSellOrBuyServiceServer).GetByIdTransactionSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSellOrBuyService_GetByIdTransactionSell_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSellOrBuyServiceServer).GetByIdTransactionSell(ctx, req.(*TransactioPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSellOrBuyService_GetByIdTransactionBuy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactioPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSellOrBuyServiceServer).GetByIdTransactionBuy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSellOrBuyService_GetByIdTransactionBuy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSellOrBuyServiceServer).GetByIdTransactionBuy(ctx, req.(*TransactioPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSellOrBuyService_TransactionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSellOrBuyServiceServer).TransactionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSellOrBuyService_TransactionUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSellOrBuyServiceServer).TransactionUpdate(ctx, req.(*UpdateTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSellOrBuyService_AllUserSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListUserTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSellOrBuyServiceServer).AllUserSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSellOrBuyService_AllUserSell_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSellOrBuyServiceServer).AllUserSell(ctx, req.(*GetListUserTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSellOrBuyService_AllUserBuy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListUserTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSellOrBuyServiceServer).AllUserBuy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSellOrBuyService_AllUserBuy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSellOrBuyServiceServer).AllUserBuy(ctx, req.(*GetListUserTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserSellOrBuyService_ServiceDesc is the grpc.ServiceDesc for UserSellOrBuyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserSellOrBuyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users_service.UserSellOrBuyService",
	HandlerType: (*UserSellOrBuyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserSell",
			Handler:    _UserSellOrBuyService_UserSell_Handler,
		},
		{
			MethodName: "UserBuy",
			Handler:    _UserSellOrBuyService_UserBuy_Handler,
		},
		{
			MethodName: "GetByIdTransactionSell",
			Handler:    _UserSellOrBuyService_GetByIdTransactionSell_Handler,
		},
		{
			MethodName: "GetByIdTransactionBuy",
			Handler:    _UserSellOrBuyService_GetByIdTransactionBuy_Handler,
		},
		{
			MethodName: "TransactionUpdate",
			Handler:    _UserSellOrBuyService_TransactionUpdate_Handler,
		},
		{
			MethodName: "AllUserSell",
			Handler:    _UserSellOrBuyService_AllUserSell_Handler,
		},
		{
			MethodName: "AllUserBuy",
			Handler:    _UserSellOrBuyService_AllUserBuy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_sell_or_buy.proto",
}
