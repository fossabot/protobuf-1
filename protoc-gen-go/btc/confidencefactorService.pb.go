// Code generated by protoc-gen-go. DO NOT EDIT.
// source: confidencefactorService.proto

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

func init() { proto.RegisterFile("confidencefactorService.proto", fileDescriptor_9246c004a4907707) }

var fileDescriptor_9246c004a4907707 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x06, 0x60, 0xce, 0xc2, 0x22, 0x65, 0x2c, 0x84, 0xe0, 0x35, 0x57, 0x88, 0x4d, 0x76, 0x45,
	0x6b, 0x11, 0x14, 0xc5, 0xc6, 0x46, 0x2d, 0xc4, 0x6e, 0x76, 0x6e, 0xb3, 0x59, 0xee, 0x32, 0x13,
	0x77, 0x27, 0xde, 0xc9, 0x71, 0x8d, 0xaf, 0xe0, 0x63, 0x59, 0xfa, 0x0a, 0x82, 0xaf, 0x21, 0x89,
	0xa7, 0x51, 0x04, 0xb5, 0x5b, 0x66, 0x77, 0xff, 0xef, 0x67, 0x92, 0x21, 0x32, 0x15, 0x7e, 0x6c,
	0x09, 0x6d, 0x01, 0x28, 0x1c, 0x2e, 0x6d, 0xb8, 0xf3, 0x68, 0x55, 0x1d, 0x58, 0x38, 0x1d, 0x15,
	0xe0, 0xc3, 0x0c, 0xee, 0x91, 0x43, 0xad, 0xcc, 0x94, 0x71, 0x82, 0x25, 0x78, 0xea, 0x2e, 0x4d,
	0x53, 0x28, 0x23, 0x98, 0x6d, 0x39, 0x66, 0x37, 0xb5, 0x1a, 0x6a, 0xaf, 0x81, 0x88, 0x05, 0xc4,
	0x33, 0xc5, 0xf7, 0x84, 0xec, 0x07, 0x70, 0x6e, 0x63, 0x04, 0xb7, 0x02, 0xb2, 0x0d, 0xe4, 0xaa,
	0x62, 0xfa, 0x36, 0xdc, 0x7b, 0x1d, 0x24, 0x9b, 0xc7, 0x9f, 0xdf, 0x4e, 0xbf, 0xf6, 0x4a, 0x9f,
	0x06, 0xc9, 0xf0, 0x2a, 0x00, 0x45, 0xc0, 0x96, 0xe9, 0x9f, 0x9d, 0xd0, 0xb8, 0x66, 0x4f, 0x92,
	0x9e, 0xa9, 0xbf, 0x4b, 0xab, 0x5f, 0x23, 0x2e, 0xec, 0x6d, 0x63, 0xa3, 0x64, 0xbb, 0xff, 0x4a,
	0xba, 0xee, 0x03, 0x46, 0xea, 0xe1, 0xf9, 0xe5, 0x71, 0x6d, 0x27, 0xdd, 0xd6, 0x46, 0x50, 0x2f,
	0xc8, 0xca, 0x8c, 0xc3, 0x64, 0xa9, 0x65, 0x1e, 0xf5, 0x42, 0xe6, 0x25, 0xc4, 0x72, 0xa9, 0xfb,
	0x8d, 0x1c, 0x1d, 0xde, 0x1c, 0x38, 0x2f, 0x65, 0x63, 0x14, 0x72, 0xa5, 0x57, 0x5a, 0xde, 0x72,
	0xba, 0xe7, 0xf2, 0x0f, 0x4f, 0x77, 0x07, 0xcc, 0x9d, 0xa5, 0xdc, 0x71, 0x4b, 0x98, 0xf5, 0x6e,
	0xb4, 0xff, 0x16, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x71, 0xf6, 0x28, 0xc8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfidenceFactorServiceClient is the client API for ConfidenceFactorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfidenceFactorServiceClient interface {
	TransactionConfidenceEndpoint(ctx context.Context, in *TransactionConfidenceEndpointRequest, opts ...grpc.CallOption) (*TXConfidence, error)
}

type confidenceFactorServiceClient struct {
	cc *grpc.ClientConn
}

func NewConfidenceFactorServiceClient(cc *grpc.ClientConn) ConfidenceFactorServiceClient {
	return &confidenceFactorServiceClient{cc}
}

func (c *confidenceFactorServiceClient) TransactionConfidenceEndpoint(ctx context.Context, in *TransactionConfidenceEndpointRequest, opts ...grpc.CallOption) (*TXConfidence, error) {
	out := new(TXConfidence)
	err := c.cc.Invoke(ctx, "/fairwaycorp.blockchainprotobuf.btc.ConfidenceFactorService/TransactionConfidenceEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfidenceFactorServiceServer is the server API for ConfidenceFactorService service.
type ConfidenceFactorServiceServer interface {
	TransactionConfidenceEndpoint(context.Context, *TransactionConfidenceEndpointRequest) (*TXConfidence, error)
}

func RegisterConfidenceFactorServiceServer(s *grpc.Server, srv ConfidenceFactorServiceServer) {
	s.RegisterService(&_ConfidenceFactorService_serviceDesc, srv)
}

func _ConfidenceFactorService_TransactionConfidenceEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionConfidenceEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfidenceFactorServiceServer).TransactionConfidenceEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairwaycorp.blockchainprotobuf.btc.ConfidenceFactorService/TransactionConfidenceEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfidenceFactorServiceServer).TransactionConfidenceEndpoint(ctx, req.(*TransactionConfidenceEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfidenceFactorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fairwaycorp.blockchainprotobuf.btc.ConfidenceFactorService",
	HandlerType: (*ConfidenceFactorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransactionConfidenceEndpoint",
			Handler:    _ConfidenceFactorService_TransactionConfidenceEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "confidencefactorService.proto",
}