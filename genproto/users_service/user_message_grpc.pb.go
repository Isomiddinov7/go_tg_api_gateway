// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: user_message.proto

package users_service

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
	UserMessageList_CreateUserMessage_FullMethodName  = "/users_service.UserMessageList/CreateUserMessage"
	UserMessageList_CreateAdminMessage_FullMethodName = "/users_service.UserMessageList/CreateAdminMessage"
	UserMessageList_UpdateMessage_FullMethodName      = "/users_service.UserMessageList/UpdateMessage"
	UserMessageList_GetUserMessage_FullMethodName     = "/users_service.UserMessageList/GetUserMessage"
	UserMessageList_GetAdminAllMessage_FullMethodName = "/users_service.UserMessageList/GetAdminAllMessage"
	UserMessageList_GetMessageAdminID_FullMethodName  = "/users_service.UserMessageList/GetMessageAdminID"
	UserMessageList_SendMessageUser_FullMethodName    = "/users_service.UserMessageList/SendMessageUser"
	UserMessageList_PayMessagePost_FullMethodName     = "/users_service.UserMessageList/PayMessagePost"
	UserMessageList_PayMessageGet_FullMethodName      = "/users_service.UserMessageList/PayMessageGet"
)

// UserMessageListClient is the client API for UserMessageList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserMessageListClient interface {
	CreateUserMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateAdminMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*Empty, error)
	GetUserMessage(ctx context.Context, in *GetMessageUserRequest, opts ...grpc.CallOption) (*GetMessageUserResponse, error)
	GetAdminAllMessage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetMessageAdminResponse, error)
	GetMessageAdminID(ctx context.Context, in *GetMessageUserRequest, opts ...grpc.CallOption) (*GetMessageAdminById, error)
	SendMessageUser(ctx context.Context, in *TelegramMessageUser, opts ...grpc.CallOption) (*TelegramMessageResponse, error)
	PayMessagePost(ctx context.Context, in *PaymsqRequest, opts ...grpc.CallOption) (*Empty, error)
	PayMessageGet(ctx context.Context, in *PaymsqUser, opts ...grpc.CallOption) (*PaymsqResponse, error)
}

type userMessageListClient struct {
	cc grpc.ClientConnInterface
}

func NewUserMessageListClient(cc grpc.ClientConnInterface) UserMessageListClient {
	return &userMessageListClient{cc}
}

func (c *userMessageListClient) CreateUserMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserMessageList_CreateUserMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMessageListClient) CreateAdminMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserMessageList_CreateAdminMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMessageListClient) UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserMessageList_UpdateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMessageListClient) GetUserMessage(ctx context.Context, in *GetMessageUserRequest, opts ...grpc.CallOption) (*GetMessageUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessageUserResponse)
	err := c.cc.Invoke(ctx, UserMessageList_GetUserMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMessageListClient) GetAdminAllMessage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetMessageAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessageAdminResponse)
	err := c.cc.Invoke(ctx, UserMessageList_GetAdminAllMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMessageListClient) GetMessageAdminID(ctx context.Context, in *GetMessageUserRequest, opts ...grpc.CallOption) (*GetMessageAdminById, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessageAdminById)
	err := c.cc.Invoke(ctx, UserMessageList_GetMessageAdminID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMessageListClient) SendMessageUser(ctx context.Context, in *TelegramMessageUser, opts ...grpc.CallOption) (*TelegramMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TelegramMessageResponse)
	err := c.cc.Invoke(ctx, UserMessageList_SendMessageUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMessageListClient) PayMessagePost(ctx context.Context, in *PaymsqRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserMessageList_PayMessagePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMessageListClient) PayMessageGet(ctx context.Context, in *PaymsqUser, opts ...grpc.CallOption) (*PaymsqResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymsqResponse)
	err := c.cc.Invoke(ctx, UserMessageList_PayMessageGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserMessageListServer is the server API for UserMessageList service.
// All implementations must embed UnimplementedUserMessageListServer
// for forward compatibility.
type UserMessageListServer interface {
	CreateUserMessage(context.Context, *MessageRequest) (*Empty, error)
	CreateAdminMessage(context.Context, *MessageRequest) (*Empty, error)
	UpdateMessage(context.Context, *UpdateMessageRequest) (*Empty, error)
	GetUserMessage(context.Context, *GetMessageUserRequest) (*GetMessageUserResponse, error)
	GetAdminAllMessage(context.Context, *Empty) (*GetMessageAdminResponse, error)
	GetMessageAdminID(context.Context, *GetMessageUserRequest) (*GetMessageAdminById, error)
	SendMessageUser(context.Context, *TelegramMessageUser) (*TelegramMessageResponse, error)
	PayMessagePost(context.Context, *PaymsqRequest) (*Empty, error)
	PayMessageGet(context.Context, *PaymsqUser) (*PaymsqResponse, error)
	mustEmbedUnimplementedUserMessageListServer()
}

// UnimplementedUserMessageListServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserMessageListServer struct{}

func (UnimplementedUserMessageListServer) CreateUserMessage(context.Context, *MessageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserMessage not implemented")
}
func (UnimplementedUserMessageListServer) CreateAdminMessage(context.Context, *MessageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdminMessage not implemented")
}
func (UnimplementedUserMessageListServer) UpdateMessage(context.Context, *UpdateMessageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedUserMessageListServer) GetUserMessage(context.Context, *GetMessageUserRequest) (*GetMessageUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMessage not implemented")
}
func (UnimplementedUserMessageListServer) GetAdminAllMessage(context.Context, *Empty) (*GetMessageAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminAllMessage not implemented")
}
func (UnimplementedUserMessageListServer) GetMessageAdminID(context.Context, *GetMessageUserRequest) (*GetMessageAdminById, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageAdminID not implemented")
}
func (UnimplementedUserMessageListServer) SendMessageUser(context.Context, *TelegramMessageUser) (*TelegramMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageUser not implemented")
}
func (UnimplementedUserMessageListServer) PayMessagePost(context.Context, *PaymsqRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayMessagePost not implemented")
}
func (UnimplementedUserMessageListServer) PayMessageGet(context.Context, *PaymsqUser) (*PaymsqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayMessageGet not implemented")
}
func (UnimplementedUserMessageListServer) mustEmbedUnimplementedUserMessageListServer() {}
func (UnimplementedUserMessageListServer) testEmbeddedByValue()                         {}

// UnsafeUserMessageListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserMessageListServer will
// result in compilation errors.
type UnsafeUserMessageListServer interface {
	mustEmbedUnimplementedUserMessageListServer()
}

func RegisterUserMessageListServer(s grpc.ServiceRegistrar, srv UserMessageListServer) {
	// If the following call pancis, it indicates UnimplementedUserMessageListServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserMessageList_ServiceDesc, srv)
}

func _UserMessageList_CreateUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageListServer).CreateUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMessageList_CreateUserMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageListServer).CreateUserMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMessageList_CreateAdminMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageListServer).CreateAdminMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMessageList_CreateAdminMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageListServer).CreateAdminMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMessageList_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageListServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMessageList_UpdateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageListServer).UpdateMessage(ctx, req.(*UpdateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMessageList_GetUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageListServer).GetUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMessageList_GetUserMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageListServer).GetUserMessage(ctx, req.(*GetMessageUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMessageList_GetAdminAllMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageListServer).GetAdminAllMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMessageList_GetAdminAllMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageListServer).GetAdminAllMessage(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMessageList_GetMessageAdminID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageListServer).GetMessageAdminID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMessageList_GetMessageAdminID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageListServer).GetMessageAdminID(ctx, req.(*GetMessageUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMessageList_SendMessageUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelegramMessageUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageListServer).SendMessageUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMessageList_SendMessageUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageListServer).SendMessageUser(ctx, req.(*TelegramMessageUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMessageList_PayMessagePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymsqRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageListServer).PayMessagePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMessageList_PayMessagePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageListServer).PayMessagePost(ctx, req.(*PaymsqRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMessageList_PayMessageGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymsqUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageListServer).PayMessageGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMessageList_PayMessageGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageListServer).PayMessageGet(ctx, req.(*PaymsqUser))
	}
	return interceptor(ctx, in, info, handler)
}

// UserMessageList_ServiceDesc is the grpc.ServiceDesc for UserMessageList service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserMessageList_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users_service.UserMessageList",
	HandlerType: (*UserMessageListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserMessage",
			Handler:    _UserMessageList_CreateUserMessage_Handler,
		},
		{
			MethodName: "CreateAdminMessage",
			Handler:    _UserMessageList_CreateAdminMessage_Handler,
		},
		{
			MethodName: "UpdateMessage",
			Handler:    _UserMessageList_UpdateMessage_Handler,
		},
		{
			MethodName: "GetUserMessage",
			Handler:    _UserMessageList_GetUserMessage_Handler,
		},
		{
			MethodName: "GetAdminAllMessage",
			Handler:    _UserMessageList_GetAdminAllMessage_Handler,
		},
		{
			MethodName: "GetMessageAdminID",
			Handler:    _UserMessageList_GetMessageAdminID_Handler,
		},
		{
			MethodName: "SendMessageUser",
			Handler:    _UserMessageList_SendMessageUser_Handler,
		},
		{
			MethodName: "PayMessagePost",
			Handler:    _UserMessageList_PayMessagePost_Handler,
		},
		{
			MethodName: "PayMessageGet",
			Handler:    _UserMessageList_PayMessageGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_message.proto",
}
