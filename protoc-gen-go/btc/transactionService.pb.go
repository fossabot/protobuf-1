// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transactionService.proto

package btc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() { proto.RegisterFile("transactionService.proto", fileDescriptor_cd57cf5433ebed65) }

var fileDescriptor_cd57cf5433ebed65 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd4, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0x06, 0x70, 0xe2, 0x85, 0x17, 0xa3, 0x57, 0xa3, 0x6c, 0x35, 0x5b, 0x57, 0x8c, 0xe2, 0xa2,
	0xd8, 0x0c, 0xac, 0x17, 0x82, 0x20, 0x62, 0xd9, 0xc5, 0x3f, 0xa0, 0x2c, 0x5b, 0x05, 0xf1, 0xee,
	0x64, 0x72, 0x9a, 0x84, 0xa6, 0x73, 0x62, 0x66, 0x62, 0x56, 0x96, 0xbd, 0xf1, 0x15, 0xbc, 0xf2,
	0x15, 0xbc, 0x50, 0x14, 0x7d, 0x0f, 0xc1, 0x57, 0xf0, 0x41, 0x24, 0x71, 0xd3, 0x8d, 0x6e, 0xda,
	0x4e, 0xbd, 0x2b, 0x53, 0xce, 0xf7, 0xfd, 0x7a, 0x3a, 0x0c, 0xbb, 0x60, 0x72, 0x50, 0x1a, 0xa4,
	0x49, 0x48, 0x8d, 0x30, 0x7f, 0x93, 0x48, 0xf4, 0xb3, 0x9c, 0x0c, 0x71, 0x6f, 0x0c, 0x49, 0x5e,
	0xc2, 0x5b, 0x49, 0x79, 0xe6, 0x07, 0x29, 0xc9, 0x89, 0x8c, 0x21, 0x51, 0xf5, 0x97, 0x41, 0x31,
	0xf6, 0x03, 0x23, 0xdd, 0x7e, 0x44, 0x14, 0xa5, 0x28, 0x20, 0x4b, 0x04, 0x28, 0x45, 0x06, 0xaa,
	0x1c, 0xfd, 0x27, 0xc1, 0x6d, 0x67, 0x3f, 0x45, 0xad, 0x21, 0x3a, 0xca, 0x76, 0xcf, 0x49, 0x9a,
	0x4e, 0xff, 0x39, 0xdc, 0xfa, 0x76, 0x86, 0xf1, 0xe7, 0x27, 0x34, 0xfc, 0xa3, 0xc3, 0x7a, 0xad,
	0xe3, 0x47, 0xa0, 0xe3, 0x1d, 0x15, 0x66, 0x94, 0x28, 0xc3, 0x87, 0xfe, 0x72, 0xa4, 0x3f, 0x67,
	0x78, 0x0f, 0x5f, 0x17, 0xa8, 0x8d, 0x7b, 0xdd, 0x2a, 0xe3, 0xa5, 0x77, 0xf5, 0xdd, 0xcf, 0x5f,
	0xef, 0x4f, 0x5d, 0xe2, 0xeb, 0x22, 0x30, 0x52, 0x1c, 0x28, 0x34, 0x25, 0xe5, 0x93, 0x43, 0x61,
	0xf6, 0xb5, 0x38, 0x30, 0xfb, 0x31, 0xe8, 0xf8, 0x90, 0x7f, 0x75, 0xd8, 0xe5, 0x17, 0x4a, 0x92,
	0x1a, 0x27, 0xf9, 0x14, 0xc3, 0x56, 0xb5, 0x9e, 0xa1, 0x9f, 0xd8, 0x14, 0x2e, 0x09, 0x69, 0xf0,
	0x9b, 0x76, 0xf8, 0x91, 0xe7, 0xd6, 0xfa, 0xf3, 0x9c, 0x9f, 0xd4, 0xf3, 0x4f, 0x0e, 0x5b, 0x7b,
	0x86, 0x65, 0xab, 0x67, 0x66, 0x7d, 0x60, 0x93, 0xdf, 0x3d, 0xdb, 0x10, 0x7d, 0x4b, 0xe2, 0x04,
	0x53, 0x34, 0xa4, 0xbc, 0x2b, 0xb5, 0x74, 0xdd, 0x5b, 0xeb, 0xd8, 0xb3, 0xc2, 0xf2, 0xae, 0x73,
	0x93, 0x7f, 0x71, 0x58, 0x6f, 0x84, 0x2a, 0xec, 0x12, 0x5b, 0x5d, 0x89, 0x39, 0xc3, 0xff, 0x4b,
	0xf6, 0x6a, 0x72, 0xdf, 0xeb, 0x75, 0x90, 0x35, 0xaa, 0xb0, 0x32, 0x7f, 0x76, 0x98, 0xbb, 0x5b,
	0xe8, 0x78, 0x0f, 0x3a, 0x17, 0xbd, 0x63, 0x53, 0x39, 0x7f, 0x7e, 0xd5, 0xcb, 0xbc, 0x48, 0x9c,
	0x15, 0x3a, 0xae, 0xc4, 0xdf, 0x1d, 0xd6, 0xdf, 0x46, 0x49, 0x21, 0xce, 0x31, 0x3f, 0xb4, 0x29,
	0x5b, 0x94, 0xb0, 0xaa, 0xfa, 0x5a, 0xad, 0xde, 0xf0, 0x2e, 0x76, 0xa8, 0xc3, 0xba, 0xa7, 0x72,
	0x7f, 0x70, 0xd8, 0xd9, 0x6d, 0x30, 0x30, 0x73, 0xde, 0xb1, 0x72, 0xb6, 0x26, 0x1a, 0xd7, 0x2d,
	0xab, 0xdb, 0x5f, 0xa4, 0x69, 0x35, 0xbc, 0x70, 0xa7, 0x21, 0x18, 0xa8, 0x6c, 0x3f, 0x1c, 0xb6,
	0xd1, 0x5a, 0xc4, 0x6e, 0x4e, 0x19, 0x44, 0xf0, 0xd7, 0x56, 0x1f, 0xaf, 0xf8, 0xa6, 0x75, 0x64,
	0x34, 0xfe, 0x2d, 0xbb, 0x3f, 0x48, 0xcb, 0x3c, 0x09, 0x30, 0x1c, 0x62, 0x4a, 0xa5, 0x27, 0xea,
	0x5f, 0x71, 0x83, 0x6f, 0x2e, 0x78, 0xe6, 0x44, 0x76, 0xdc, 0x39, 0xbc, 0xff, 0xea, 0x5e, 0x94,
	0x98, 0xb8, 0x08, 0x7c, 0x49, 0x53, 0x71, 0x54, 0x38, 0xa8, 0x1a, 0xc5, 0x71, 0xe3, 0xa0, 0xa9,
	0x14, 0xf5, 0x07, 0x39, 0x88, 0x50, 0x0d, 0x22, 0xaa, 0x3a, 0x82, 0xd3, 0xf5, 0xd1, 0xed, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x97, 0x5f, 0xcb, 0x8b, 0x06, 0x00, 0x00,
}

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
	UnconfirmedTransactionsEndpoint(ctx context.Context, in *UnconfirmedTransactionsEndpointRequest, opts ...grpc.CallOption) (*TXS, error)
	NewTransactionEndpoint(ctx context.Context, in *NewTransactionEndpointRequest, opts ...grpc.CallOption) (*TXSkeleton, error)
	SendTransactionEndpoint(ctx context.Context, in *SendTransactionEndpointRequest, opts ...grpc.CallOption) (*TXSkeleton, error)
	PushRawTransactionEndpoint(ctx context.Context, in *PushRawTransactionEndpointRequest, opts ...grpc.CallOption) (*TX, error)
	DecodeRawTransactionEndpoint(ctx context.Context, in *DecodeRawTransactionEndpointRequest, opts ...grpc.CallOption) (*TX, error)
	DataEndpoint(ctx context.Context, in *DataEndpointRequest, opts ...grpc.CallOption) (*NullData, error)
	TransactionPropagationEndpoint(ctx context.Context, in *TransactionPropagationEndpointRequest, opts ...grpc.CallOption) (*DescribedBelow, error)
}

type transactionServiceClient struct {
	cc *grpc.ClientConn
}

func NewTransactionServiceClient(cc *grpc.ClientConn) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) TransactionHashEndpoint(ctx context.Context, in *TransactionHashEndpointRequest, opts ...grpc.CallOption) (*TX, error) {
	out := new(TX)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.TransactionService/TransactionHashEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) UnconfirmedTransactionsEndpoint(ctx context.Context, in *UnconfirmedTransactionsEndpointRequest, opts ...grpc.CallOption) (*TXS, error) {
	out := new(TXS)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.TransactionService/UnconfirmedTransactionsEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) NewTransactionEndpoint(ctx context.Context, in *NewTransactionEndpointRequest, opts ...grpc.CallOption) (*TXSkeleton, error) {
	out := new(TXSkeleton)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.TransactionService/NewTransactionEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) SendTransactionEndpoint(ctx context.Context, in *SendTransactionEndpointRequest, opts ...grpc.CallOption) (*TXSkeleton, error) {
	out := new(TXSkeleton)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.TransactionService/SendTransactionEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) PushRawTransactionEndpoint(ctx context.Context, in *PushRawTransactionEndpointRequest, opts ...grpc.CallOption) (*TX, error) {
	out := new(TX)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.TransactionService/PushRawTransactionEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DecodeRawTransactionEndpoint(ctx context.Context, in *DecodeRawTransactionEndpointRequest, opts ...grpc.CallOption) (*TX, error) {
	out := new(TX)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.TransactionService/DecodeRawTransactionEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DataEndpoint(ctx context.Context, in *DataEndpointRequest, opts ...grpc.CallOption) (*NullData, error) {
	out := new(NullData)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.TransactionService/DataEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) TransactionPropagationEndpoint(ctx context.Context, in *TransactionPropagationEndpointRequest, opts ...grpc.CallOption) (*DescribedBelow, error) {
	out := new(DescribedBelow)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.TransactionService/TransactionPropagationEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
type TransactionServiceServer interface {
	TransactionHashEndpoint(context.Context, *TransactionHashEndpointRequest) (*TX, error)
	UnconfirmedTransactionsEndpoint(context.Context, *UnconfirmedTransactionsEndpointRequest) (*TXS, error)
	NewTransactionEndpoint(context.Context, *NewTransactionEndpointRequest) (*TXSkeleton, error)
	SendTransactionEndpoint(context.Context, *SendTransactionEndpointRequest) (*TXSkeleton, error)
	PushRawTransactionEndpoint(context.Context, *PushRawTransactionEndpointRequest) (*TX, error)
	DecodeRawTransactionEndpoint(context.Context, *DecodeRawTransactionEndpointRequest) (*TX, error)
	DataEndpoint(context.Context, *DataEndpointRequest) (*NullData, error)
	TransactionPropagationEndpoint(context.Context, *TransactionPropagationEndpointRequest) (*DescribedBelow, error)
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
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.TransactionService/TransactionHashEndpoint",
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
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.TransactionService/UnconfirmedTransactionsEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).UnconfirmedTransactionsEndpoint(ctx, req.(*UnconfirmedTransactionsEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_NewTransactionEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTransactionEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).NewTransactionEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.TransactionService/NewTransactionEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).NewTransactionEndpoint(ctx, req.(*NewTransactionEndpointRequest))
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
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.TransactionService/SendTransactionEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).SendTransactionEndpoint(ctx, req.(*SendTransactionEndpointRequest))
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
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.TransactionService/PushRawTransactionEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).PushRawTransactionEndpoint(ctx, req.(*PushRawTransactionEndpointRequest))
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
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.TransactionService/DecodeRawTransactionEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DecodeRawTransactionEndpoint(ctx, req.(*DecodeRawTransactionEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DataEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DataEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.TransactionService/DataEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DataEndpoint(ctx, req.(*DataEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_TransactionPropagationEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionPropagationEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).TransactionPropagationEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.TransactionService/TransactionPropagationEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).TransactionPropagationEndpoint(ctx, req.(*TransactionPropagationEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fairwaycorp.blockchainprotobuf.btc.TransactionService",
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
			MethodName: "NewTransactionEndpoint",
			Handler:    _TransactionService_NewTransactionEndpoint_Handler,
		},
		{
			MethodName: "SendTransactionEndpoint",
			Handler:    _TransactionService_SendTransactionEndpoint_Handler,
		},
		{
			MethodName: "PushRawTransactionEndpoint",
			Handler:    _TransactionService_PushRawTransactionEndpoint_Handler,
		},
		{
			MethodName: "DecodeRawTransactionEndpoint",
			Handler:    _TransactionService_DecodeRawTransactionEndpoint_Handler,
		},
		{
			MethodName: "DataEndpoint",
			Handler:    _TransactionService_DataEndpoint_Handler,
		},
		{
			MethodName: "TransactionPropagationEndpoint",
			Handler:    _TransactionService_TransactionPropagationEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transactionService.proto",
}