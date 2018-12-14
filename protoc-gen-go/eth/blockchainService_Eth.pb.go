// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchainService_Eth.proto

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

// BlockchainServiceClient is the client API for BlockchainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockchainServiceClient interface {
	ChainEndpoint(ctx context.Context, in *ChainEndpointRequest, opts ...grpc.CallOption) (*Blockchain, error)
	BlockHashEndpoint(ctx context.Context, in *BlockHashEndpointRequest, opts ...grpc.CallOption) (*Block, error)
	BlockHeightEndpoint(ctx context.Context, in *BlockHeightEndpointRequest, opts ...grpc.CallOption) (*Block, error)
}

type blockchainServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlockchainServiceClient(cc *grpc.ClientConn) BlockchainServiceClient {
	return &blockchainServiceClient{cc}
}

func (c *blockchainServiceClient) ChainEndpoint(ctx context.Context, in *ChainEndpointRequest, opts ...grpc.CallOption) (*Blockchain, error) {
	out := new(Blockchain)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.eth.BlockchainService/ChainEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) BlockHashEndpoint(ctx context.Context, in *BlockHashEndpointRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.eth.BlockchainService/BlockHashEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) BlockHeightEndpoint(ctx context.Context, in *BlockHeightEndpointRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.eth.BlockchainService/BlockHeightEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServiceServer is the server API for BlockchainService service.
type BlockchainServiceServer interface {
	ChainEndpoint(context.Context, *ChainEndpointRequest) (*Blockchain, error)
	BlockHashEndpoint(context.Context, *BlockHashEndpointRequest) (*Block, error)
	BlockHeightEndpoint(context.Context, *BlockHeightEndpointRequest) (*Block, error)
}

func RegisterBlockchainServiceServer(s *grpc.Server, srv BlockchainServiceServer) {
	s.RegisterService(&_BlockchainService_serviceDesc, srv)
}

func _BlockchainService_ChainEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).ChainEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.eth.BlockchainService/ChainEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).ChainEndpoint(ctx, req.(*ChainEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_BlockHashEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHashEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).BlockHashEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.eth.BlockchainService/BlockHashEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).BlockHashEndpoint(ctx, req.(*BlockHashEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_BlockHeightEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHeightEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).BlockHeightEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.eth.BlockchainService/BlockHeightEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).BlockHeightEndpoint(ctx, req.(*BlockHeightEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockchainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fairwaycorp.blockchainprotobuf.eth.BlockchainService",
	HandlerType: (*BlockchainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChainEndpoint",
			Handler:    _BlockchainService_ChainEndpoint_Handler,
		},
		{
			MethodName: "BlockHashEndpoint",
			Handler:    _BlockchainService_BlockHashEndpoint_Handler,
		},
		{
			MethodName: "BlockHeightEndpoint",
			Handler:    _BlockchainService_BlockHeightEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blockchainService_Eth.proto",
}

func init() {
	proto.RegisterFile("blockchainService_Eth.proto", fileDescriptor_blockchainService_Eth_e5eca99ecb5b6ac9)
}

var fileDescriptor_blockchainService_Eth_e5eca99ecb5b6ac9 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x8e, 0xcd, 0x4a, 0xf4, 0x30,
	0x14, 0x40, 0xe9, 0xc7, 0x87, 0x8b, 0x82, 0xa2, 0x11, 0x14, 0xaa, 0xab, 0x32, 0x08, 0x8a, 0x4d,
	0x40, 0x37, 0x2e, 0xfc, 0x81, 0x91, 0x01, 0x37, 0x6e, 0x74, 0xe7, 0x66, 0x48, 0x63, 0x26, 0x09,
	0x33, 0xcd, 0xad, 0x4d, 0xea, 0x20, 0x43, 0x37, 0xfa, 0x04, 0xe2, 0x4b, 0xb8, 0xf2, 0x65, 0x7c,
	0x05, 0x1f, 0x44, 0x9a, 0xf9, 0x69, 0xab, 0xa0, 0xc5, 0x55, 0xcb, 0xbd, 0x39, 0xf7, 0x1c, 0x7f,
	0x2b, 0x1e, 0x01, 0x1b, 0x32, 0x49, 0x95, 0xbe, 0xe6, 0xd9, 0xbd, 0x62, 0xbc, 0xdf, 0xb3, 0x12,
	0xa7, 0x19, 0x58, 0x40, 0xe1, 0x80, 0xaa, 0x6c, 0x4c, 0x1f, 0x18, 0x64, 0x29, 0xae, 0x1e, 0xba,
	0x65, 0x9c, 0x0f, 0x30, 0xb7, 0x32, 0xd8, 0x16, 0x00, 0x62, 0xc4, 0x09, 0x4d, 0x15, 0xa1, 0x5a,
	0x83, 0xa5, 0x56, 0x81, 0x36, 0xd3, 0x0b, 0xc1, 0x26, 0x83, 0x24, 0x01, 0x7d, 0xc9, 0x8d, 0xa1,
	0xa2, 0x76, 0x3a, 0xa8, 0x79, 0xbf, 0x2d, 0x0f, 0x9e, 0xfe, 0xfb, 0x6b, 0xdd, 0xaf, 0x5d, 0xe8,
	0xd9, 0xf3, 0x97, 0xcf, 0xcb, 0x41, 0x4f, 0xdf, 0xa6, 0xa0, 0xb4, 0x45, 0x47, 0xf8, 0xf7, 0x40,
	0xdc, 0x40, 0xae, 0xf8, 0x5d, 0xce, 0x8d, 0x0d, 0x70, 0x1b, 0xb2, 0x4a, 0x08, 0x37, 0x1e, 0xdf,
	0x3f, 0x5e, 0xfe, 0xad, 0xa2, 0x15, 0xc2, 0xad, 0x24, 0x13, 0xcd, 0xed, 0x18, 0xb2, 0x61, 0x81,
	0x5e, 0xbd, 0x59, 0xe9, 0x05, 0x35, 0x72, 0xd1, 0x75, 0xdc, 0xfa, 0x7a, 0x1d, 0x9b, 0xb7, 0xed,
	0xb6, 0xa6, 0xc3, 0x3d, 0x97, 0xd5, 0x41, 0x61, 0x33, 0x8b, 0x38, 0xc8, 0x90, 0x89, 0xfb, 0xf6,
	0x25, 0x35, 0xb2, 0x40, 0x6f, 0x9e, 0xbf, 0x3e, 0x75, 0x72, 0x25, 0xa4, 0x5d, 0xc4, 0x9e, 0xb6,
	0x8f, 0x6d, 0x80, 0x7f, 0xc8, 0xdd, 0x77, 0xb9, 0x3b, 0xa8, 0xf3, 0x73, 0xae, 0xd3, 0x14, 0xdd,
	0xb3, 0x9b, 0x13, 0xa1, 0xac, 0xcc, 0x63, 0xcc, 0x20, 0x21, 0x33, 0x49, 0x54, 0x5a, 0x48, 0x65,
	0x89, 0xe6, 0x1a, 0xe2, 0x7e, 0x58, 0x24, 0xb8, 0x8e, 0x04, 0x94, 0x82, 0x78, 0xc9, 0x8d, 0x0e,
	0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x6c, 0xec, 0x95, 0xe4, 0x02, 0x00, 0x00,
}
