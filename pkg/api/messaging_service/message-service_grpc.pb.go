// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package messaging_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessagingServiceApiClient is the client API for MessagingServiceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingServiceApiClient interface {
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ServiceStatus, error)
	SendInstantEmail(ctx context.Context, in *SendEmailReq, opts ...grpc.CallOption) (*ServiceStatus, error)
	SendMessageToAllUsers(ctx context.Context, in *SendMessageToAllUsersReq, opts ...grpc.CallOption) (*ServiceStatus, error)
	SendMessageToStudyParticipants(ctx context.Context, in *SendMessageToStudyParticipantsReq, opts ...grpc.CallOption) (*ServiceStatus, error)
	GetAutoMessages(ctx context.Context, in *GetAutoMessagesReq, opts ...grpc.CallOption) (*AutoMessages, error)
	SaveAutoMessage(ctx context.Context, in *SaveAutoMessageReq, opts ...grpc.CallOption) (*AutoMessage, error)
	DeleteAutoMessage(ctx context.Context, in *DeleteAutoMessageReq, opts ...grpc.CallOption) (*ServiceStatus, error)
	GetEmailTemplates(ctx context.Context, in *GetEmailTemplatesReq, opts ...grpc.CallOption) (*EmailTemplates, error)
	SaveEmailTemplate(ctx context.Context, in *SaveEmailTemplateReq, opts ...grpc.CallOption) (*EmailTemplate, error)
	DeleteEmailTemplate(ctx context.Context, in *DeleteEmailTemplateReq, opts ...grpc.CallOption) (*ServiceStatus, error)
}

type messagingServiceApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingServiceApiClient(cc grpc.ClientConnInterface) MessagingServiceApiClient {
	return &messagingServiceApiClient{cc}
}

func (c *messagingServiceApiClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ServiceStatus, error) {
	out := new(ServiceStatus)
	err := c.cc.Invoke(ctx, "/influenzanet.message_service.MessagingServiceApi/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceApiClient) SendInstantEmail(ctx context.Context, in *SendEmailReq, opts ...grpc.CallOption) (*ServiceStatus, error) {
	out := new(ServiceStatus)
	err := c.cc.Invoke(ctx, "/influenzanet.message_service.MessagingServiceApi/SendInstantEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceApiClient) SendMessageToAllUsers(ctx context.Context, in *SendMessageToAllUsersReq, opts ...grpc.CallOption) (*ServiceStatus, error) {
	out := new(ServiceStatus)
	err := c.cc.Invoke(ctx, "/influenzanet.message_service.MessagingServiceApi/SendMessageToAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceApiClient) SendMessageToStudyParticipants(ctx context.Context, in *SendMessageToStudyParticipantsReq, opts ...grpc.CallOption) (*ServiceStatus, error) {
	out := new(ServiceStatus)
	err := c.cc.Invoke(ctx, "/influenzanet.message_service.MessagingServiceApi/SendMessageToStudyParticipants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceApiClient) GetAutoMessages(ctx context.Context, in *GetAutoMessagesReq, opts ...grpc.CallOption) (*AutoMessages, error) {
	out := new(AutoMessages)
	err := c.cc.Invoke(ctx, "/influenzanet.message_service.MessagingServiceApi/GetAutoMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceApiClient) SaveAutoMessage(ctx context.Context, in *SaveAutoMessageReq, opts ...grpc.CallOption) (*AutoMessage, error) {
	out := new(AutoMessage)
	err := c.cc.Invoke(ctx, "/influenzanet.message_service.MessagingServiceApi/SaveAutoMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceApiClient) DeleteAutoMessage(ctx context.Context, in *DeleteAutoMessageReq, opts ...grpc.CallOption) (*ServiceStatus, error) {
	out := new(ServiceStatus)
	err := c.cc.Invoke(ctx, "/influenzanet.message_service.MessagingServiceApi/DeleteAutoMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceApiClient) GetEmailTemplates(ctx context.Context, in *GetEmailTemplatesReq, opts ...grpc.CallOption) (*EmailTemplates, error) {
	out := new(EmailTemplates)
	err := c.cc.Invoke(ctx, "/influenzanet.message_service.MessagingServiceApi/GetEmailTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceApiClient) SaveEmailTemplate(ctx context.Context, in *SaveEmailTemplateReq, opts ...grpc.CallOption) (*EmailTemplate, error) {
	out := new(EmailTemplate)
	err := c.cc.Invoke(ctx, "/influenzanet.message_service.MessagingServiceApi/SaveEmailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceApiClient) DeleteEmailTemplate(ctx context.Context, in *DeleteEmailTemplateReq, opts ...grpc.CallOption) (*ServiceStatus, error) {
	out := new(ServiceStatus)
	err := c.cc.Invoke(ctx, "/influenzanet.message_service.MessagingServiceApi/DeleteEmailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingServiceApiServer is the server API for MessagingServiceApi service.
// All implementations must embed UnimplementedMessagingServiceApiServer
// for forward compatibility
type MessagingServiceApiServer interface {
	Status(context.Context, *emptypb.Empty) (*ServiceStatus, error)
	SendInstantEmail(context.Context, *SendEmailReq) (*ServiceStatus, error)
	SendMessageToAllUsers(context.Context, *SendMessageToAllUsersReq) (*ServiceStatus, error)
	SendMessageToStudyParticipants(context.Context, *SendMessageToStudyParticipantsReq) (*ServiceStatus, error)
	GetAutoMessages(context.Context, *GetAutoMessagesReq) (*AutoMessages, error)
	SaveAutoMessage(context.Context, *SaveAutoMessageReq) (*AutoMessage, error)
	DeleteAutoMessage(context.Context, *DeleteAutoMessageReq) (*ServiceStatus, error)
	GetEmailTemplates(context.Context, *GetEmailTemplatesReq) (*EmailTemplates, error)
	SaveEmailTemplate(context.Context, *SaveEmailTemplateReq) (*EmailTemplate, error)
	DeleteEmailTemplate(context.Context, *DeleteEmailTemplateReq) (*ServiceStatus, error)
	mustEmbedUnimplementedMessagingServiceApiServer()
}

// UnimplementedMessagingServiceApiServer must be embedded to have forward compatible implementations.
type UnimplementedMessagingServiceApiServer struct {
}

func (UnimplementedMessagingServiceApiServer) Status(context.Context, *emptypb.Empty) (*ServiceStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedMessagingServiceApiServer) SendInstantEmail(context.Context, *SendEmailReq) (*ServiceStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendInstantEmail not implemented")
}
func (UnimplementedMessagingServiceApiServer) SendMessageToAllUsers(context.Context, *SendMessageToAllUsersReq) (*ServiceStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToAllUsers not implemented")
}
func (UnimplementedMessagingServiceApiServer) SendMessageToStudyParticipants(context.Context, *SendMessageToStudyParticipantsReq) (*ServiceStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToStudyParticipants not implemented")
}
func (UnimplementedMessagingServiceApiServer) GetAutoMessages(context.Context, *GetAutoMessagesReq) (*AutoMessages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAutoMessages not implemented")
}
func (UnimplementedMessagingServiceApiServer) SaveAutoMessage(context.Context, *SaveAutoMessageReq) (*AutoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveAutoMessage not implemented")
}
func (UnimplementedMessagingServiceApiServer) DeleteAutoMessage(context.Context, *DeleteAutoMessageReq) (*ServiceStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAutoMessage not implemented")
}
func (UnimplementedMessagingServiceApiServer) GetEmailTemplates(context.Context, *GetEmailTemplatesReq) (*EmailTemplates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmailTemplates not implemented")
}
func (UnimplementedMessagingServiceApiServer) SaveEmailTemplate(context.Context, *SaveEmailTemplateReq) (*EmailTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveEmailTemplate not implemented")
}
func (UnimplementedMessagingServiceApiServer) DeleteEmailTemplate(context.Context, *DeleteEmailTemplateReq) (*ServiceStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmailTemplate not implemented")
}
func (UnimplementedMessagingServiceApiServer) mustEmbedUnimplementedMessagingServiceApiServer() {}

// UnsafeMessagingServiceApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServiceApiServer will
// result in compilation errors.
type UnsafeMessagingServiceApiServer interface {
	mustEmbedUnimplementedMessagingServiceApiServer()
}

func RegisterMessagingServiceApiServer(s grpc.ServiceRegistrar, srv MessagingServiceApiServer) {
	s.RegisterService(&MessagingServiceApi_ServiceDesc, srv)
}

func _MessagingServiceApi_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceApiServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.message_service.MessagingServiceApi/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceApiServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingServiceApi_SendInstantEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceApiServer).SendInstantEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.message_service.MessagingServiceApi/SendInstantEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceApiServer).SendInstantEmail(ctx, req.(*SendEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingServiceApi_SendMessageToAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageToAllUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceApiServer).SendMessageToAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.message_service.MessagingServiceApi/SendMessageToAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceApiServer).SendMessageToAllUsers(ctx, req.(*SendMessageToAllUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingServiceApi_SendMessageToStudyParticipants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageToStudyParticipantsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceApiServer).SendMessageToStudyParticipants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.message_service.MessagingServiceApi/SendMessageToStudyParticipants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceApiServer).SendMessageToStudyParticipants(ctx, req.(*SendMessageToStudyParticipantsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingServiceApi_GetAutoMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAutoMessagesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceApiServer).GetAutoMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.message_service.MessagingServiceApi/GetAutoMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceApiServer).GetAutoMessages(ctx, req.(*GetAutoMessagesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingServiceApi_SaveAutoMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveAutoMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceApiServer).SaveAutoMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.message_service.MessagingServiceApi/SaveAutoMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceApiServer).SaveAutoMessage(ctx, req.(*SaveAutoMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingServiceApi_DeleteAutoMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAutoMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceApiServer).DeleteAutoMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.message_service.MessagingServiceApi/DeleteAutoMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceApiServer).DeleteAutoMessage(ctx, req.(*DeleteAutoMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingServiceApi_GetEmailTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmailTemplatesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceApiServer).GetEmailTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.message_service.MessagingServiceApi/GetEmailTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceApiServer).GetEmailTemplates(ctx, req.(*GetEmailTemplatesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingServiceApi_SaveEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveEmailTemplateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceApiServer).SaveEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.message_service.MessagingServiceApi/SaveEmailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceApiServer).SaveEmailTemplate(ctx, req.(*SaveEmailTemplateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingServiceApi_DeleteEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmailTemplateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceApiServer).DeleteEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.message_service.MessagingServiceApi/DeleteEmailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceApiServer).DeleteEmailTemplate(ctx, req.(*DeleteEmailTemplateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MessagingServiceApi_ServiceDesc is the grpc.ServiceDesc for MessagingServiceApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagingServiceApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "influenzanet.message_service.MessagingServiceApi",
	HandlerType: (*MessagingServiceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _MessagingServiceApi_Status_Handler,
		},
		{
			MethodName: "SendInstantEmail",
			Handler:    _MessagingServiceApi_SendInstantEmail_Handler,
		},
		{
			MethodName: "SendMessageToAllUsers",
			Handler:    _MessagingServiceApi_SendMessageToAllUsers_Handler,
		},
		{
			MethodName: "SendMessageToStudyParticipants",
			Handler:    _MessagingServiceApi_SendMessageToStudyParticipants_Handler,
		},
		{
			MethodName: "GetAutoMessages",
			Handler:    _MessagingServiceApi_GetAutoMessages_Handler,
		},
		{
			MethodName: "SaveAutoMessage",
			Handler:    _MessagingServiceApi_SaveAutoMessage_Handler,
		},
		{
			MethodName: "DeleteAutoMessage",
			Handler:    _MessagingServiceApi_DeleteAutoMessage_Handler,
		},
		{
			MethodName: "GetEmailTemplates",
			Handler:    _MessagingServiceApi_GetEmailTemplates_Handler,
		},
		{
			MethodName: "SaveEmailTemplate",
			Handler:    _MessagingServiceApi_SaveEmailTemplate_Handler,
		},
		{
			MethodName: "DeleteEmailTemplate",
			Handler:    _MessagingServiceApi_DeleteEmailTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messaging_service/message-service.proto",
}