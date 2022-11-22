// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/rchat.proto

package protos

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

// RChatClient is the client API for RChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RChatClient interface {
	// Basic RPC for chat
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
	ChatSession(ctx context.Context, opts ...grpc.CallOption) (RChat_ChatSessionClient, error)
}

type rChatClient struct {
	cc grpc.ClientConnInterface
}

func NewRChatClient(cc grpc.ClientConnInterface) RChatClient {
	return &rChatClient{cc}
}

func (c *rChatClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.RChat/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rChatClient) ChatSession(ctx context.Context, opts ...grpc.CallOption) (RChat_ChatSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &RChat_ServiceDesc.Streams[0], "/protos.RChat/ChatSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &rChatChatSessionClient{stream}
	return x, nil
}

type RChat_ChatSessionClient interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ClientStream
}

type rChatChatSessionClient struct {
	grpc.ClientStream
}

func (x *rChatChatSessionClient) Send(m *RouteNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rChatChatSessionClient) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RChatServer is the server API for RChat service.
// All implementations must embed UnimplementedRChatServer
// for forward compatibility
type RChatServer interface {
	// Basic RPC for chat
	SendMessage(context.Context, *Message) (*Response, error)
	ChatSession(RChat_ChatSessionServer) error
	mustEmbedUnimplementedRChatServer()
}

// UnimplementedRChatServer must be embedded to have forward compatible implementations.
type UnimplementedRChatServer struct {
}

func (UnimplementedRChatServer) SendMessage(context.Context, *Message) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedRChatServer) ChatSession(RChat_ChatSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatSession not implemented")
}
func (UnimplementedRChatServer) mustEmbedUnimplementedRChatServer() {}

// UnsafeRChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RChatServer will
// result in compilation errors.
type UnsafeRChatServer interface {
	mustEmbedUnimplementedRChatServer()
}

func RegisterRChatServer(s grpc.ServiceRegistrar, srv RChatServer) {
	s.RegisterService(&RChat_ServiceDesc, srv)
}

func _RChat_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RChatServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.RChat/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RChatServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _RChat_ChatSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RChatServer).ChatSession(&rChatChatSessionServer{stream})
}

type RChat_ChatSessionServer interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ServerStream
}

type rChatChatSessionServer struct {
	grpc.ServerStream
}

func (x *rChatChatSessionServer) Send(m *RouteNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rChatChatSessionServer) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RChat_ServiceDesc is the grpc.ServiceDesc for RChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.RChat",
	HandlerType: (*RChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _RChat_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatSession",
			Handler:       _RChat_ChatSession_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/rchat.proto",
}
