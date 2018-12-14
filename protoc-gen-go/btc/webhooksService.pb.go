// Code generated by protoc-gen-go. DO NOT EDIT.
// source: webhooksService.proto

package btc // import "github.com/fairway-corp/blockchain-protobuf/protoc-gen-go/btc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WebHooksServiceClient is the client API for WebHooksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WebHooksServiceClient interface {
	CreateWebHookEndpoint(ctx context.Context, in *CreateWebHookEndpointRequest, opts ...grpc.CallOption) (*Event, error)
	ListWebHooksEndpoint(ctx context.Context, in *ListWebHooksEndpointRequest, opts ...grpc.CallOption) (*ArrayEvent, error)
	WebHookIDEndpoint(ctx context.Context, in *WebHookIDEndpointRequest, opts ...grpc.CallOption) (*Event, error)
	DeleteWebHookEndpoint(ctx context.Context, in *DeleteWebHookEndpointRequest, opts ...grpc.CallOption) (*NullValue, error)
}

type webHooksServiceClient struct {
	cc *grpc.ClientConn
}

func NewWebHooksServiceClient(cc *grpc.ClientConn) WebHooksServiceClient {
	return &webHooksServiceClient{cc}
}

func (c *webHooksServiceClient) CreateWebHookEndpoint(ctx context.Context, in *CreateWebHookEndpointRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.WebHooksService/CreateWebHookEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHooksServiceClient) ListWebHooksEndpoint(ctx context.Context, in *ListWebHooksEndpointRequest, opts ...grpc.CallOption) (*ArrayEvent, error) {
	out := new(ArrayEvent)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.WebHooksService/ListWebHooksEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHooksServiceClient) WebHookIDEndpoint(ctx context.Context, in *WebHookIDEndpointRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.WebHooksService/WebHookIDEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webHooksServiceClient) DeleteWebHookEndpoint(ctx context.Context, in *DeleteWebHookEndpointRequest, opts ...grpc.CallOption) (*NullValue, error) {
	out := new(NullValue)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.WebHooksService/DeleteWebHookEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebHooksServiceServer is the server API for WebHooksService service.
type WebHooksServiceServer interface {
	CreateWebHookEndpoint(context.Context, *CreateWebHookEndpointRequest) (*Event, error)
	ListWebHooksEndpoint(context.Context, *ListWebHooksEndpointRequest) (*ArrayEvent, error)
	WebHookIDEndpoint(context.Context, *WebHookIDEndpointRequest) (*Event, error)
	DeleteWebHookEndpoint(context.Context, *DeleteWebHookEndpointRequest) (*NullValue, error)
}

func RegisterWebHooksServiceServer(s *grpc.Server, srv WebHooksServiceServer) {
	s.RegisterService(&_WebHooksService_serviceDesc, srv)
}

func _WebHooksService_CreateWebHookEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWebHookEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHooksServiceServer).CreateWebHookEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.WebHooksService/CreateWebHookEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHooksServiceServer).CreateWebHookEndpoint(ctx, req.(*CreateWebHookEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHooksService_ListWebHooksEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWebHooksEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHooksServiceServer).ListWebHooksEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.WebHooksService/ListWebHooksEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHooksServiceServer).ListWebHooksEndpoint(ctx, req.(*ListWebHooksEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHooksService_WebHookIDEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebHookIDEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHooksServiceServer).WebHookIDEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.WebHooksService/WebHookIDEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHooksServiceServer).WebHookIDEndpoint(ctx, req.(*WebHookIDEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebHooksService_DeleteWebHookEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWebHookEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebHooksServiceServer).DeleteWebHookEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.WebHooksService/DeleteWebHookEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebHooksServiceServer).DeleteWebHookEndpoint(ctx, req.(*DeleteWebHookEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WebHooksService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fairwaycorp.blockchainprotobuf.btc.WebHooksService",
	HandlerType: (*WebHooksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWebHookEndpoint",
			Handler:    _WebHooksService_CreateWebHookEndpoint_Handler,
		},
		{
			MethodName: "ListWebHooksEndpoint",
			Handler:    _WebHooksService_ListWebHooksEndpoint_Handler,
		},
		{
			MethodName: "WebHookIDEndpoint",
			Handler:    _WebHooksService_WebHookIDEndpoint_Handler,
		},
		{
			MethodName: "DeleteWebHookEndpoint",
			Handler:    _WebHooksService_DeleteWebHookEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webhooksService.proto",
}

func init() {
	proto.RegisterFile("webhooksService.proto", fileDescriptor_webhooksService_e78e697da115e7d9)
}

var fileDescriptor_webhooksService_e78e697da115e7d9 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd2, 0x4f, 0x4a, 0xfb, 0x40,
	0x14, 0xc0, 0x71, 0xf2, 0x83, 0x5f, 0x91, 0x08, 0x4a, 0xc7, 0x76, 0x13, 0x5c, 0x75, 0x67, 0x21,
	0x33, 0xa0, 0x3b, 0x51, 0xea, 0x9f, 0x16, 0x14, 0xd4, 0x85, 0x82, 0x82, 0xbb, 0x99, 0xe9, 0x6b,
	0x3a, 0x34, 0x9d, 0x17, 0x27, 0x93, 0x96, 0x22, 0x6e, 0xdc, 0x78, 0x00, 0x17, 0x7a, 0x00, 0x6f,
	0xe4, 0x15, 0x3c, 0x88, 0x34, 0x4d, 0xa2, 0x62, 0xc0, 0xd0, 0x5d, 0x98, 0xc7, 0xfb, 0xf2, 0x49,
	0x32, 0x6e, 0x73, 0x0a, 0x62, 0x88, 0x38, 0x8a, 0xaf, 0xc0, 0x4c, 0x94, 0x04, 0x1a, 0x19, 0xb4,
	0x48, 0x5a, 0x03, 0xae, 0xcc, 0x94, 0xcf, 0x24, 0x9a, 0x88, 0x8a, 0x10, 0xe5, 0x48, 0x0e, 0xb9,
	0xd2, 0xe9, 0x50, 0x24, 0x03, 0x2a, 0xac, 0xf4, 0x36, 0x03, 0xc4, 0x20, 0x04, 0xc6, 0x23, 0xc5,
	0xb8, 0xd6, 0x68, 0xb9, 0x55, 0xa8, 0xe3, 0x45, 0xc1, 0x2b, 0xc2, 0xe7, 0x10, 0xc7, 0x3c, 0xc8,
	0xc2, 0xde, 0x86, 0xc4, 0xf1, 0x18, 0xf5, 0x8f, 0xc3, 0xed, 0xa7, 0xff, 0xee, 0xfa, 0x0d, 0x88,
	0x93, 0x6f, 0x0e, 0xf2, 0xe2, 0xb8, 0xcd, 0x63, 0x03, 0xdc, 0x42, 0x36, 0xe9, 0xe9, 0x7e, 0x84,
	0x4a, 0x5b, 0x72, 0x40, 0xff, 0xc6, 0xd1, 0xd2, 0xd5, 0x4b, 0xb8, 0x4b, 0x20, 0xb6, 0xde, 0x56,
	0x95, 0x42, 0x6f, 0x02, 0xda, 0xb6, 0xea, 0x8f, 0xef, 0x1f, 0xcf, 0xff, 0x56, 0x5b, 0x35, 0x96,
	0xbe, 0xcc, 0xae, 0xd3, 0x9e, 0xcb, 0x1a, 0x67, 0x2a, 0xb6, 0xb9, 0xb8, 0x80, 0x75, 0xaa, 0x64,
	0xcb, 0x36, 0x73, 0x17, 0xad, 0x12, 0x38, 0x34, 0x86, 0xcf, 0x16, 0xb8, 0xb5, 0x14, 0xb7, 0x42,
	0x32, 0x1c, 0x79, 0x75, 0xdc, 0x7a, 0xd6, 0x3e, 0xed, 0x16, 0xac, 0xbd, 0x2a, 0xd5, 0x5f, 0x6b,
	0x4b, 0x7c, 0x2b, 0x2f, 0xe5, 0x34, 0x08, 0x59, 0x70, 0xd8, 0x7d, 0x76, 0x05, 0x54, 0xff, 0x81,
	0xbc, 0x39, 0x6e, 0xb3, 0x0b, 0x21, 0x2c, 0xf9, 0x3b, 0x4b, 0x57, 0x73, 0xa2, 0x5f, 0xa5, 0x70,
	0x91, 0x84, 0xe1, 0x35, 0x0f, 0x13, 0xc8, 0x99, 0xed, 0x12, 0xe6, 0x51, 0xe7, 0x76, 0x3f, 0x50,
	0x76, 0x98, 0x08, 0x2a, 0x71, 0xcc, 0xb2, 0xac, 0x3f, 0xef, 0xb2, 0xaf, 0xae, 0x9f, 0x87, 0x59,
	0xfa, 0x20, 0xfd, 0x00, 0xb4, 0x1f, 0x20, 0x13, 0x56, 0x8a, 0x5a, 0x7a, 0xb4, 0xf3, 0x19, 0x00,
	0x00, 0xff, 0xff, 0x60, 0x46, 0x5f, 0x15, 0x58, 0x03, 0x00, 0x00,
}
