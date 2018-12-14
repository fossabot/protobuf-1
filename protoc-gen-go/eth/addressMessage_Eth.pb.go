// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressMessage_Eth.proto

package eth // import "github.com/fairway-corp/blockchain-protobuf/protoc-gen-go/eth"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddressBalanceEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias" json:"network,omitempty"`
	Address              string               `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AddressBalanceEndpointRequest) Reset()         { *m = AddressBalanceEndpointRequest{} }
func (m *AddressBalanceEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*AddressBalanceEndpointRequest) ProtoMessage()    {}
func (*AddressBalanceEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_addressMessage_Eth_9ec4b757126248d7, []int{0}
}
func (m *AddressBalanceEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBalanceEndpointRequest.Unmarshal(m, b)
}
func (m *AddressBalanceEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBalanceEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *AddressBalanceEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBalanceEndpointRequest.Merge(dst, src)
}
func (m *AddressBalanceEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_AddressBalanceEndpointRequest.Size(m)
}
func (m *AddressBalanceEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBalanceEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBalanceEndpointRequest proto.InternalMessageInfo

func (m *AddressBalanceEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAIN
}

func (m *AddressBalanceEndpointRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type AddressEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias" json:"network,omitempty"`
	Address string               `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Filters response to only include transactions below before height in the blockchain.
	Before int32 `protobuf:"varint,3,opt,name=before,proto3" json:"before,omitempty"`
	// 	Filters response to only include transactions above after height in the blockchain.
	After int32 `protobuf:"varint,4,opt,name=after,proto3" json:"after,omitempty"`
	// limit sets the minimum number of returned TXRefs; there can be less if there are less than limit TXRefs associated with this address, but there can be more in the rare case of more TXRefs in the block at the bottom of your call. This ensures paging by block height never misses TXRefs. Defaults to 200, maximum is 2000.
	Limit int32 `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	// 	If set, only returns the balance and TXRefs that have at least this number of confirmations.
	Confirmations        int32    `protobuf:"varint,6,opt,name=confirmations,proto3" json:"confirmations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressEndpointRequest) Reset()         { *m = AddressEndpointRequest{} }
func (m *AddressEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*AddressEndpointRequest) ProtoMessage()    {}
func (*AddressEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_addressMessage_Eth_9ec4b757126248d7, []int{1}
}
func (m *AddressEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressEndpointRequest.Unmarshal(m, b)
}
func (m *AddressEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *AddressEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressEndpointRequest.Merge(dst, src)
}
func (m *AddressEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_AddressEndpointRequest.Size(m)
}
func (m *AddressEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressEndpointRequest proto.InternalMessageInfo

func (m *AddressEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAIN
}

func (m *AddressEndpointRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddressEndpointRequest) GetBefore() int32 {
	if m != nil {
		return m.Before
	}
	return 0
}

func (m *AddressEndpointRequest) GetAfter() int32 {
	if m != nil {
		return m.After
	}
	return 0
}

func (m *AddressEndpointRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *AddressEndpointRequest) GetConfirmations() int32 {
	if m != nil {
		return m.Confirmations
	}
	return 0
}

type GenerateAddressEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GenerateAddressEndpointRequest) Reset()         { *m = GenerateAddressEndpointRequest{} }
func (m *GenerateAddressEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateAddressEndpointRequest) ProtoMessage()    {}
func (*GenerateAddressEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_addressMessage_Eth_9ec4b757126248d7, []int{2}
}
func (m *GenerateAddressEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateAddressEndpointRequest.Unmarshal(m, b)
}
func (m *GenerateAddressEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateAddressEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *GenerateAddressEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateAddressEndpointRequest.Merge(dst, src)
}
func (m *GenerateAddressEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateAddressEndpointRequest.Size(m)
}
func (m *GenerateAddressEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateAddressEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateAddressEndpointRequest proto.InternalMessageInfo

func (m *GenerateAddressEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAIN
}

func init() {
	proto.RegisterType((*AddressBalanceEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.eth.AddressBalanceEndpointRequest")
	proto.RegisterType((*AddressEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.eth.AddressEndpointRequest")
	proto.RegisterType((*GenerateAddressEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.eth.GenerateAddressEndpointRequest")
}

func init() {
	proto.RegisterFile("addressMessage_Eth.proto", fileDescriptor_addressMessage_Eth_9ec4b757126248d7)
}

var fileDescriptor_addressMessage_Eth_9ec4b757126248d7 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x8f, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0xd9, 0xef, 0xb3, 0x2d, 0x06, 0xf4, 0x10, 0xa4, 0x06, 0x41, 0x29, 0x8b, 0x87, 0x5e,
	0x36, 0x0b, 0x7a, 0xf1, 0x22, 0xd2, 0x42, 0xf1, 0xa4, 0x87, 0x3d, 0x7a, 0x91, 0x6c, 0x3a, 0xbb,
	0x1b, 0x9a, 0xcd, 0xd4, 0x64, 0x4a, 0xf1, 0x47, 0xf8, 0x5f, 0xfd, 0x09, 0xd2, 0xec, 0x16, 0x11,
	0x0f, 0xde, 0xf4, 0x96, 0xf7, 0x19, 0xe6, 0xc9, 0xbc, 0x4c, 0xa8, 0xe5, 0xd2, 0x43, 0x08, 0x0f,
	0x10, 0x82, 0xaa, 0xe1, 0x79, 0x41, 0x8d, 0x5c, 0x7b, 0x24, 0xe4, 0x69, 0xa5, 0x8c, 0xdf, 0xaa,
	0x57, 0x8d, 0x7e, 0x2d, 0x4b, 0x8b, 0x7a, 0xa5, 0x1b, 0x65, 0x5c, 0x1c, 0x96, 0x9b, 0x4a, 0x02,
	0x35, 0x67, 0xa7, 0x1a, 0xdb, 0x16, 0xdd, 0xb7, 0xe5, 0xf4, 0x2d, 0x61, 0xe7, 0xb3, 0xce, 0x3c,
	0x57, 0x56, 0x39, 0x0d, 0x0b, 0xb7, 0x5c, 0xa3, 0x71, 0x54, 0xc0, 0xcb, 0x06, 0x02, 0xf1, 0x82,
	0x8d, 0x1c, 0xd0, 0x16, 0xfd, 0x4a, 0x24, 0x93, 0x64, 0x7a, 0x7c, 0x75, 0x23, 0x7f, 0xfe, 0x50,
	0x3e, 0x76, 0x2b, 0x33, 0x6b, 0x71, 0x6b, 0x5c, 0x3d, 0xb3, 0x46, 0x85, 0x62, 0x2f, 0xe2, 0x82,
	0x8d, 0xfa, 0x3a, 0xe2, 0xdf, 0x24, 0x99, 0x1e, 0x16, 0xfb, 0x98, 0xbe, 0x27, 0x6c, 0xdc, 0xdf,
	0xf3, 0xa7, 0x87, 0xf0, 0x31, 0x1b, 0x96, 0x50, 0xa1, 0x07, 0xf1, 0x7f, 0x92, 0x4c, 0x07, 0x45,
	0x9f, 0xf8, 0x09, 0x1b, 0xa8, 0x8a, 0xc0, 0x8b, 0x83, 0x88, 0xbb, 0xb0, 0xa3, 0xd6, 0xb4, 0x86,
	0xc4, 0xa0, 0xa3, 0x31, 0xf0, 0x4b, 0x76, 0xa4, 0xd1, 0x55, 0xc6, 0xb7, 0x8a, 0x0c, 0xba, 0x20,
	0x86, 0x71, 0xfa, 0x15, 0xa6, 0xc4, 0x2e, 0xee, 0xc1, 0x81, 0x57, 0x04, 0xbf, 0xd7, 0x7c, 0x7e,
	0xf7, 0x74, 0x5b, 0x1b, 0x6a, 0x36, 0xa5, 0xd4, 0xd8, 0xe6, 0xbd, 0x2e, 0xdb, 0xf9, 0xf2, 0x4f,
	0x5f, 0xb6, 0x17, 0xe6, 0xf1, 0xa1, 0xb3, 0x1a, 0x5c, 0x56, 0x63, 0x0e, 0xd4, 0x94, 0xc3, 0x88,
	0xae, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xc2, 0x77, 0x22, 0x99, 0x02, 0x00, 0x00,
}
