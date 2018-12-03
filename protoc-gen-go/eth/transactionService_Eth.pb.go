// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transactionService_Eth.proto

package eth // import "github.com/fairway-corp/blockchain-protobuf/protoc-gen-go/eth"

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

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionServiceClient interface {
	TransactionHashEndpoint(ctx context.Context, in *TransactionHashEndpointRequest, opts ...grpc.CallOption) (*TX, error)
	UnconfirmedTransactionsEndpoint(ctx context.Context, in *UnconfirmedTransactionsEndpointRequest, opts ...grpc.CallOption) (*TxArray, error)
	CreateTransactionsEndpoint(ctx context.Context, in *CreateTransactionsEndpointRequest, opts ...grpc.CallOption) (*TXSkeleton, error)
	SendTransactionEndpoint(ctx context.Context, in *SendTransactionEndpointRequest, opts ...grpc.CallOption) (*TXSkeleton, error)
	DecodeRawTransactionEndpoint(ctx context.Context, in *DecodeRawTransactionEndpointRequest, opts ...grpc.CallOption) (*TX, error)
	PushRawTransactionEndpoint(ctx context.Context, in *PushRawTransactionEndpointRequest, opts ...grpc.CallOption) (*TX, error)
}

type transactionServiceClient struct {
	cc *grpc.ClientConn
}

func NewTransactionServiceClient(cc *grpc.ClientConn) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) TransactionHashEndpoint(ctx context.Context, in *TransactionHashEndpointRequest, opts ...grpc.CallOption) (*TX, error) {
	out := new(TX)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.eth.TransactionService/TransactionHashEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) UnconfirmedTransactionsEndpoint(ctx context.Context, in *UnconfirmedTransactionsEndpointRequest, opts ...grpc.CallOption) (*TxArray, error) {
	out := new(TxArray)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.eth.TransactionService/UnconfirmedTransactionsEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreateTransactionsEndpoint(ctx context.Context, in *CreateTransactionsEndpointRequest, opts ...grpc.CallOption) (*TXSkeleton, error) {
	out := new(TXSkeleton)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.eth.TransactionService/CreateTransactionsEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) SendTransactionEndpoint(ctx context.Context, in *SendTransactionEndpointRequest, opts ...grpc.CallOption) (*TXSkeleton, error) {
	out := new(TXSkeleton)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.eth.TransactionService/SendTransactionEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DecodeRawTransactionEndpoint(ctx context.Context, in *DecodeRawTransactionEndpointRequest, opts ...grpc.CallOption) (*TX, error) {
	out := new(TX)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.eth.TransactionService/DecodeRawTransactionEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) PushRawTransactionEndpoint(ctx context.Context, in *PushRawTransactionEndpointRequest, opts ...grpc.CallOption) (*TX, error) {
	out := new(TX)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.eth.TransactionService/PushRawTransactionEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
type TransactionServiceServer interface {
	TransactionHashEndpoint(context.Context, *TransactionHashEndpointRequest) (*TX, error)
	UnconfirmedTransactionsEndpoint(context.Context, *UnconfirmedTransactionsEndpointRequest) (*TxArray, error)
	CreateTransactionsEndpoint(context.Context, *CreateTransactionsEndpointRequest) (*TXSkeleton, error)
	SendTransactionEndpoint(context.Context, *SendTransactionEndpointRequest) (*TXSkeleton, error)
	DecodeRawTransactionEndpoint(context.Context, *DecodeRawTransactionEndpointRequest) (*TX, error)
	PushRawTransactionEndpoint(context.Context, *PushRawTransactionEndpointRequest) (*TX, error)
}

func RegisterTransactionServiceServer(s *grpc.Server, srv TransactionServiceServer) {
	s.RegisterService(&_TransactionService_serviceDesc, srv)
}

func _TransactionService_TransactionHashEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionHashEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).TransactionHashEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.eth.TransactionService/TransactionHashEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).TransactionHashEndpoint(ctx, req.(*TransactionHashEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_UnconfirmedTransactionsEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnconfirmedTransactionsEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).UnconfirmedTransactionsEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.eth.TransactionService/UnconfirmedTransactionsEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).UnconfirmedTransactionsEndpoint(ctx, req.(*UnconfirmedTransactionsEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreateTransactionsEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionsEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateTransactionsEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.eth.TransactionService/CreateTransactionsEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateTransactionsEndpoint(ctx, req.(*CreateTransactionsEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_SendTransactionEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransactionEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).SendTransactionEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.eth.TransactionService/SendTransactionEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).SendTransactionEndpoint(ctx, req.(*SendTransactionEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DecodeRawTransactionEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeRawTransactionEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DecodeRawTransactionEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.eth.TransactionService/DecodeRawTransactionEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DecodeRawTransactionEndpoint(ctx, req.(*DecodeRawTransactionEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_PushRawTransactionEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRawTransactionEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).PushRawTransactionEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.eth.TransactionService/PushRawTransactionEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).PushRawTransactionEndpoint(ctx, req.(*PushRawTransactionEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fairwaycorp.blockchainprotobuf.eth.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransactionHashEndpoint",
			Handler:    _TransactionService_TransactionHashEndpoint_Handler,
		},
		{
			MethodName: "UnconfirmedTransactionsEndpoint",
			Handler:    _TransactionService_UnconfirmedTransactionsEndpoint_Handler,
		},
		{
			MethodName: "CreateTransactionsEndpoint",
			Handler:    _TransactionService_CreateTransactionsEndpoint_Handler,
		},
		{
			MethodName: "SendTransactionEndpoint",
			Handler:    _TransactionService_SendTransactionEndpoint_Handler,
		},
		{
			MethodName: "DecodeRawTransactionEndpoint",
			Handler:    _TransactionService_DecodeRawTransactionEndpoint_Handler,
		},
		{
			MethodName: "PushRawTransactionEndpoint",
			Handler:    _TransactionService_PushRawTransactionEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transactionService_Eth.proto",
}

func init() {
	proto.RegisterFile("transactionService_Eth.proto", fileDescriptor_transactionService_Eth_3ad37b8bac5df2eb)
}

var fileDescriptor_transactionService_Eth_3ad37b8bac5df2eb = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x8a, 0xd4, 0x30,
	0x18, 0xc7, 0xa9, 0x07, 0x91, 0x1e, 0x83, 0x38, 0x98, 0x1d, 0x15, 0xa3, 0x78, 0x58, 0x69, 0x03,
	0x7a, 0x5b, 0x10, 0x71, 0x75, 0x50, 0x04, 0x41, 0x76, 0x14, 0xc4, 0x8b, 0xa4, 0xe9, 0x37, 0x4d,
	0x98, 0x69, 0xbe, 0x9a, 0xa4, 0x76, 0x96, 0x65, 0x2f, 0xbe, 0x82, 0x6f, 0xe1, 0x45, 0x3c, 0xe8,
	0x7b, 0x88, 0xaf, 0xe0, 0x83, 0x48, 0xe3, 0x8e, 0x16, 0xb6, 0x75, 0x3b, 0x7b, 0x0b, 0x09, 0xdf,
	0xef, 0xff, 0xcb, 0x9f, 0x2f, 0x9e, 0x7a, 0x2b, 0x8c, 0x13, 0xd2, 0x6b, 0x34, 0x73, 0xb0, 0x1f,
	0xb4, 0x84, 0x77, 0x33, 0xaf, 0xd2, 0xca, 0xa2, 0x47, 0xc2, 0x16, 0x42, 0xdb, 0x46, 0x1c, 0x4a,
	0xb4, 0x55, 0x9a, 0xad, 0x50, 0x2e, 0xa5, 0x12, 0xda, 0x84, 0xc7, 0xac, 0x5e, 0xa4, 0xe0, 0x15,
	0x9d, 0x16, 0x88, 0xc5, 0x0a, 0xb8, 0xa8, 0x34, 0x17, 0xc6, 0xa0, 0x17, 0x2d, 0xcb, 0xfd, 0x21,
	0xd0, 0x2e, 0xff, 0x05, 0x38, 0x27, 0x8a, 0x0e, 0x9f, 0x4e, 0x24, 0x96, 0x65, 0xcf, 0xc3, 0xbd,
	0x1f, 0x97, 0x62, 0xf2, 0xea, 0x94, 0x19, 0xf9, 0x1c, 0xc5, 0x93, 0xce, 0xf5, 0x33, 0xe1, 0xd4,
	0xcc, 0xe4, 0x15, 0x6a, 0xe3, 0xc9, 0x7e, 0x7a, 0xb6, 0x6c, 0x3a, 0x30, 0x7c, 0x00, 0xef, 0x6b,
	0x70, 0x9e, 0xde, 0x19, 0xc5, 0x78, 0xc3, 0x6e, 0x7d, 0xfc, 0xf9, 0xeb, 0xd3, 0x85, 0x6b, 0x64,
	0x87, 0x83, 0x57, 0xfc, 0xc8, 0x80, 0x6f, 0xd0, 0x2e, 0x8f, 0xb9, 0x5f, 0x3b, 0x7e, 0xe4, 0xd7,
	0x4a, 0x38, 0x75, 0x4c, 0xbe, 0x47, 0xf1, 0x8d, 0xd7, 0x46, 0xa2, 0x59, 0x68, 0x5b, 0x42, 0xde,
	0x89, 0x76, 0x7f, 0xa5, 0x9f, 0x8f, 0x09, 0x3c, 0x03, 0xb2, 0x91, 0xbf, 0x3b, 0x4a, 0x7e, 0xfd,
	0xc8, 0x5a, 0x71, 0xc8, 0x68, 0xf8, 0xc1, 0x65, 0x42, 0x4e, 0xff, 0xa0, 0x15, 0xa7, 0x8f, 0x2d,
	0x08, 0x0f, 0xbd, 0xce, 0xb3, 0x31, 0x39, 0xc3, 0xf3, 0x1b, 0xdd, 0x74, 0x5c, 0xd7, 0xf3, 0x25,
	0xac, 0xc0, 0xa3, 0x61, 0x37, 0x83, 0xf1, 0x0e, 0xbb, 0xd2, 0xd3, 0xb9, 0x81, 0x66, 0x2f, 0xda,
	0x25, 0x5f, 0xa3, 0x78, 0x32, 0x07, 0xd3, 0x6d, 0x69, 0xbb, 0xf5, 0x18, 0x18, 0x3e, 0xaf, 0x32,
	0x0b, 0xca, 0x53, 0x36, 0xe9, 0x51, 0x76, 0x60, 0xf2, 0xd6, 0xf9, 0x5b, 0x14, 0x4f, 0x9f, 0x80,
	0xc4, 0x1c, 0x0e, 0x44, 0xd3, 0x27, 0xfe, 0x74, 0x4c, 0xe8, 0xff, 0x08, 0xdb, 0x2e, 0xf7, 0xed,
	0x60, 0x7d, 0x9d, 0x5d, 0xed, 0xb1, 0xce, 0x43, 0x4e, 0xeb, 0xfd, 0x25, 0x8a, 0xe9, 0xcb, 0xda,
	0xa9, 0x01, 0xeb, 0x51, 0x4b, 0x32, 0x3c, 0xbf, 0xad, 0xf3, 0x49, 0xd3, 0xb4, 0xaf, 0xe9, 0xaa,
	0x76, 0x6a, 0x2f, 0xda, 0xdd, 0x7f, 0xf8, 0xf6, 0x41, 0xa1, 0xbd, 0xaa, 0xb3, 0x54, 0x62, 0xc9,
	0x4f, 0xb8, 0x49, 0x0b, 0xe6, 0xff, 0xc0, 0xc9, 0x86, 0xcc, 0xc3, 0x41, 0x26, 0x05, 0x98, 0xa4,
	0xc0, 0x16, 0x9b, 0x5d, 0x0c, 0x57, 0xf7, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x67, 0xcf,
	0xf2, 0x34, 0x05, 0x00, 0x00,
}
