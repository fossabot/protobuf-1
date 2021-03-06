// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressService.proto

package btc // import "github.com/chainweaver/protobuf/protoc-gen-go/btc"

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

// AddressServiceClient is the client API for AddressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddressServiceClient interface {
	AddressBalanceEndpoint(ctx context.Context, in *AddressBalanceEndpointRequest, opts ...grpc.CallOption) (*Address, error)
	AddressEndpoint(ctx context.Context, in *AddressEndpointRequest, opts ...grpc.CallOption) (*Address, error)
	AddressFullEndpoint(ctx context.Context, in *AddressFullEndpointRequest, opts ...grpc.CallOption) (*Address, error)
	GenerateAddressEndpoint(ctx context.Context, in *GenerateAddressEndpointRequest, opts ...grpc.CallOption) (*AddressKeychain, error)
}

type addressServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddressServiceClient(cc *grpc.ClientConn) AddressServiceClient {
	return &addressServiceClient{cc}
}

func (c *addressServiceClient) AddressBalanceEndpoint(ctx context.Context, in *AddressBalanceEndpointRequest, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/chainweaver.protobuf.btc.AddressService/AddressBalanceEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) AddressEndpoint(ctx context.Context, in *AddressEndpointRequest, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/chainweaver.protobuf.btc.AddressService/AddressEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) AddressFullEndpoint(ctx context.Context, in *AddressFullEndpointRequest, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/chainweaver.protobuf.btc.AddressService/AddressFullEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) GenerateAddressEndpoint(ctx context.Context, in *GenerateAddressEndpointRequest, opts ...grpc.CallOption) (*AddressKeychain, error) {
	out := new(AddressKeychain)
	err := c.cc.Invoke(ctx, "/chainweaver.protobuf.btc.AddressService/GenerateAddressEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServiceServer is the server API for AddressService service.
type AddressServiceServer interface {
	AddressBalanceEndpoint(context.Context, *AddressBalanceEndpointRequest) (*Address, error)
	AddressEndpoint(context.Context, *AddressEndpointRequest) (*Address, error)
	AddressFullEndpoint(context.Context, *AddressFullEndpointRequest) (*Address, error)
	GenerateAddressEndpoint(context.Context, *GenerateAddressEndpointRequest) (*AddressKeychain, error)
}

func RegisterAddressServiceServer(s *grpc.Server, srv AddressServiceServer) {
	s.RegisterService(&_AddressService_serviceDesc, srv)
}

func _AddressService_AddressBalanceEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressBalanceEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).AddressBalanceEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainweaver.protobuf.btc.AddressService/AddressBalanceEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).AddressBalanceEndpoint(ctx, req.(*AddressBalanceEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_AddressEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).AddressEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainweaver.protobuf.btc.AddressService/AddressEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).AddressEndpoint(ctx, req.(*AddressEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_AddressFullEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressFullEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).AddressFullEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainweaver.protobuf.btc.AddressService/AddressFullEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).AddressFullEndpoint(ctx, req.(*AddressFullEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_GenerateAddressEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAddressEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).GenerateAddressEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainweaver.protobuf.btc.AddressService/GenerateAddressEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).GenerateAddressEndpoint(ctx, req.(*GenerateAddressEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddressService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chainweaver.protobuf.btc.AddressService",
	HandlerType: (*AddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddressBalanceEndpoint",
			Handler:    _AddressService_AddressBalanceEndpoint_Handler,
		},
		{
			MethodName: "AddressEndpoint",
			Handler:    _AddressService_AddressEndpoint_Handler,
		},
		{
			MethodName: "AddressFullEndpoint",
			Handler:    _AddressService_AddressFullEndpoint_Handler,
		},
		{
			MethodName: "GenerateAddressEndpoint",
			Handler:    _AddressService_GenerateAddressEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "addressService.proto",
}

func init() {
	proto.RegisterFile("addressService.proto", fileDescriptor_addressService_80b889bff26c371a)
}

var fileDescriptor_addressService_80b889bff26c371a = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xbb, 0x4a, 0x33, 0x41,
	0x14, 0xc7, 0xd9, 0x0f, 0x3e, 0x8b, 0x2d, 0x14, 0x26, 0xc1, 0xcb, 0x22, 0x5e, 0x22, 0x06, 0x2f,
	0x64, 0x47, 0x8d, 0xa0, 0xd8, 0x19, 0x50, 0x0b, 0xb1, 0xd1, 0xce, 0x6e, 0x76, 0x72, 0xb2, 0x19,
	0x9c, 0xcc, 0x89, 0x33, 0xb3, 0x09, 0x12, 0xd2, 0x58, 0x59, 0xd9, 0x58, 0x5a, 0xd8, 0xf8, 0x46,
	0xbe, 0x82, 0x0f, 0x22, 0xd9, 0xdd, 0x84, 0x44, 0xb2, 0x2e, 0xda, 0x0d, 0x7f, 0xfe, 0x97, 0xdf,
	0x1c, 0xb7, 0xc8, 0xea, 0x75, 0x0d, 0xc6, 0xdc, 0x80, 0xee, 0x08, 0x0e, 0x7e, 0x5b, 0xa3, 0x45,
	0xb2, 0xc8, 0x9b, 0x4c, 0xa8, 0x2e, 0xb0, 0x0e, 0xe8, 0x44, 0x0a, 0xa2, 0x86, 0x1f, 0x58, 0xee,
	0x2d, 0x87, 0x88, 0xa1, 0x04, 0xca, 0xda, 0x82, 0x32, 0xa5, 0xd0, 0x32, 0x2b, 0x50, 0x99, 0xc4,
	0xe4, 0x0d, 0xdb, 0xae, 0xc0, 0x18, 0x16, 0xa6, 0x6d, 0x5e, 0x81, 0x63, 0xab, 0x85, 0x6a, 0x42,
	0x3c, 0x78, 0xfa, 0xef, 0xce, 0x9e, 0x4e, 0x6c, 0x93, 0x77, 0xc7, 0x9d, 0x4f, 0xa5, 0x1a, 0x93,
	0x4c, 0x71, 0x38, 0x53, 0xf5, 0x36, 0x0a, 0x65, 0xc9, 0x91, 0x9f, 0x45, 0xe4, 0x4f, 0x4f, 0x5c,
	0xc3, 0x7d, 0x04, 0xc6, 0x7a, 0xeb, 0xb9, 0xc1, 0x92, 0xff, 0xf8, 0xf1, 0xf9, 0xf2, 0x6f, 0x8b,
	0x94, 0x69, 0x60, 0x39, 0xed, 0x29, 0xb0, 0x5d, 0xd4, 0x77, 0x7d, 0x3a, 0xf8, 0x8b, 0xa1, 0xbd,
	0xf4, 0x4b, 0x7d, 0x1a, 0x24, 0x0b, 0xe4, 0xd9, 0x71, 0xe7, 0xd2, 0xec, 0x88, 0x6f, 0x2f, 0x77,
	0xe6, 0x0f, 0x60, 0xe5, 0x18, 0x6c, 0x8d, 0xac, 0xfc, 0x0c, 0x46, 0x5e, 0x1d, 0xb7, 0x90, 0x66,
	0xce, 0x23, 0x29, 0x47, 0x50, 0x87, 0xb9, 0x13, 0xe3, 0xf6, 0x5f, 0x80, 0xed, 0xc6, 0x60, 0x9b,
	0x64, 0x23, 0xe7, 0x62, 0x8d, 0x48, 0x4a, 0xf2, 0xe6, 0xb8, 0x0b, 0x17, 0xa0, 0x40, 0x33, 0x0b,
	0xdf, 0xcf, 0x76, 0x9c, 0xbd, 0x95, 0x11, 0x19, 0x52, 0x6e, 0xe7, 0x52, 0x5e, 0xc2, 0x43, 0x6c,
	0x29, 0xad, 0xc6, 0xb4, 0x4b, 0xa5, 0xe2, 0x34, 0xda, 0x13, 0x67, 0xa7, 0x56, 0xbd, 0xdd, 0x0f,
	0x85, 0x6d, 0x46, 0x81, 0xcf, 0xb1, 0x45, 0xc7, 0x7a, 0xe9, 0xb0, 0x37, 0x79, 0xf0, 0x4a, 0x08,
	0xaa, 0x12, 0xe2, 0xa0, 0x27, 0x98, 0x89, 0xa5, 0xea, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88,
	0x7e, 0x1f, 0xc5, 0x41, 0x03, 0x00, 0x00,
}
