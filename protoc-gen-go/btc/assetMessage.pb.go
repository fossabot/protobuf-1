// Code generated by protoc-gen-go. DO NOT EDIT.
// source: assetMessage.proto

package btc // import "github.com/fairway-corp/blockchain-protobuf/protoc-gen-go/btc"

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

type GenerateAssetAddressEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GenerateAssetAddressEndpointRequest) Reset()         { *m = GenerateAssetAddressEndpointRequest{} }
func (m *GenerateAssetAddressEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateAssetAddressEndpointRequest) ProtoMessage()    {}
func (*GenerateAssetAddressEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_assetMessage_c5355d3599f2f00a, []int{0}
}
func (m *GenerateAssetAddressEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateAssetAddressEndpointRequest.Unmarshal(m, b)
}
func (m *GenerateAssetAddressEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateAssetAddressEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *GenerateAssetAddressEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateAssetAddressEndpointRequest.Merge(dst, src)
}
func (m *GenerateAssetAddressEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateAssetAddressEndpointRequest.Size(m)
}
func (m *GenerateAssetAddressEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateAssetAddressEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateAssetAddressEndpointRequest proto.InternalMessageInfo

func (m *GenerateAssetAddressEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

type IssueAssetEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	// The private key being used to issue or transfer assets.
	FromPrivate string `protobuf:"bytes,2,opt,name=from_private,json=fromPrivate,proto3" json:"from_private,omitempty"`
	// The target OAP address assets for issue or transfer.
	ToAddress string `protobuf:"bytes,3,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// The amount of assets being issued or transfered.
	Amount int32 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	// Optional Hex-encoded metadata that can optionally be encoded into the issue or transfer transaction.
	Metadata             string   `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueAssetEndpointRequest) Reset()         { *m = IssueAssetEndpointRequest{} }
func (m *IssueAssetEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*IssueAssetEndpointRequest) ProtoMessage()    {}
func (*IssueAssetEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_assetMessage_c5355d3599f2f00a, []int{1}
}
func (m *IssueAssetEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueAssetEndpointRequest.Unmarshal(m, b)
}
func (m *IssueAssetEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueAssetEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *IssueAssetEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueAssetEndpointRequest.Merge(dst, src)
}
func (m *IssueAssetEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_IssueAssetEndpointRequest.Size(m)
}
func (m *IssueAssetEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueAssetEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueAssetEndpointRequest proto.InternalMessageInfo

func (m *IssueAssetEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *IssueAssetEndpointRequest) GetFromPrivate() string {
	if m != nil {
		return m.FromPrivate
	}
	return ""
}

func (m *IssueAssetEndpointRequest) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *IssueAssetEndpointRequest) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *IssueAssetEndpointRequest) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

type TransferAssetEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Assetid string               `protobuf:"bytes,2,opt,name=assetid,proto3" json:"assetid,omitempty"`
	// The private key being used to issue or transfer assets.
	FromPrivate string `protobuf:"bytes,3,opt,name=from_private,json=fromPrivate,proto3" json:"from_private,omitempty"`
	// The target OAP address assets for issue or transfer.
	ToAddress string `protobuf:"bytes,4,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// The amount of assets being issued or transfered.
	Amount int32 `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	// Optional Hex-encoded metadata that can optionally be encoded into the issue or transfer transaction.
	Metadata             string   `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferAssetEndpointRequest) Reset()         { *m = TransferAssetEndpointRequest{} }
func (m *TransferAssetEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*TransferAssetEndpointRequest) ProtoMessage()    {}
func (*TransferAssetEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_assetMessage_c5355d3599f2f00a, []int{2}
}
func (m *TransferAssetEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferAssetEndpointRequest.Unmarshal(m, b)
}
func (m *TransferAssetEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferAssetEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *TransferAssetEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferAssetEndpointRequest.Merge(dst, src)
}
func (m *TransferAssetEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_TransferAssetEndpointRequest.Size(m)
}
func (m *TransferAssetEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferAssetEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferAssetEndpointRequest proto.InternalMessageInfo

func (m *TransferAssetEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *TransferAssetEndpointRequest) GetAssetid() string {
	if m != nil {
		return m.Assetid
	}
	return ""
}

func (m *TransferAssetEndpointRequest) GetFromPrivate() string {
	if m != nil {
		return m.FromPrivate
	}
	return ""
}

func (m *TransferAssetEndpointRequest) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *TransferAssetEndpointRequest) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransferAssetEndpointRequest) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

type ListAssetTXsEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Assetid              string               `protobuf:"bytes,2,opt,name=assetid,proto3" json:"assetid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListAssetTXsEndpointRequest) Reset()         { *m = ListAssetTXsEndpointRequest{} }
func (m *ListAssetTXsEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*ListAssetTXsEndpointRequest) ProtoMessage()    {}
func (*ListAssetTXsEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_assetMessage_c5355d3599f2f00a, []int{3}
}
func (m *ListAssetTXsEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAssetTXsEndpointRequest.Unmarshal(m, b)
}
func (m *ListAssetTXsEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAssetTXsEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *ListAssetTXsEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAssetTXsEndpointRequest.Merge(dst, src)
}
func (m *ListAssetTXsEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_ListAssetTXsEndpointRequest.Size(m)
}
func (m *ListAssetTXsEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAssetTXsEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAssetTXsEndpointRequest proto.InternalMessageInfo

func (m *ListAssetTXsEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *ListAssetTXsEndpointRequest) GetAssetid() string {
	if m != nil {
		return m.Assetid
	}
	return ""
}

type AssetTXEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Assetid              string               `protobuf:"bytes,2,opt,name=assetid,proto3" json:"assetid,omitempty"`
	Txhash               string               `protobuf:"bytes,3,opt,name=txhash,proto3" json:"txhash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AssetTXEndpointRequest) Reset()         { *m = AssetTXEndpointRequest{} }
func (m *AssetTXEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*AssetTXEndpointRequest) ProtoMessage()    {}
func (*AssetTXEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_assetMessage_c5355d3599f2f00a, []int{4}
}
func (m *AssetTXEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetTXEndpointRequest.Unmarshal(m, b)
}
func (m *AssetTXEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetTXEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *AssetTXEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetTXEndpointRequest.Merge(dst, src)
}
func (m *AssetTXEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_AssetTXEndpointRequest.Size(m)
}
func (m *AssetTXEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetTXEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssetTXEndpointRequest proto.InternalMessageInfo

func (m *AssetTXEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *AssetTXEndpointRequest) GetAssetid() string {
	if m != nil {
		return m.Assetid
	}
	return ""
}

func (m *AssetTXEndpointRequest) GetTxhash() string {
	if m != nil {
		return m.Txhash
	}
	return ""
}

type AssetAddressEndpointRequest struct {
	Network              string   `protobuf:"bytes,10,opt,name=network,proto3" json:"network,omitempty"`
	Assetid              string   `protobuf:"bytes,1,opt,name=assetid,proto3" json:"assetid,omitempty"`
	Oapaddr              string   `protobuf:"bytes,2,opt,name=oapaddr,proto3" json:"oapaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetAddressEndpointRequest) Reset()         { *m = AssetAddressEndpointRequest{} }
func (m *AssetAddressEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*AssetAddressEndpointRequest) ProtoMessage()    {}
func (*AssetAddressEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_assetMessage_c5355d3599f2f00a, []int{5}
}
func (m *AssetAddressEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetAddressEndpointRequest.Unmarshal(m, b)
}
func (m *AssetAddressEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetAddressEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *AssetAddressEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetAddressEndpointRequest.Merge(dst, src)
}
func (m *AssetAddressEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_AssetAddressEndpointRequest.Size(m)
}
func (m *AssetAddressEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetAddressEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssetAddressEndpointRequest proto.InternalMessageInfo

func (m *AssetAddressEndpointRequest) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *AssetAddressEndpointRequest) GetAssetid() string {
	if m != nil {
		return m.Assetid
	}
	return ""
}

func (m *AssetAddressEndpointRequest) GetOapaddr() string {
	if m != nil {
		return m.Oapaddr
	}
	return ""
}

func init() {
	proto.RegisterType((*GenerateAssetAddressEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.GenerateAssetAddressEndpointRequest")
	proto.RegisterType((*IssueAssetEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.IssueAssetEndpointRequest")
	proto.RegisterType((*TransferAssetEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.TransferAssetEndpointRequest")
	proto.RegisterType((*ListAssetTXsEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.ListAssetTXsEndpointRequest")
	proto.RegisterType((*AssetTXEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.AssetTXEndpointRequest")
	proto.RegisterType((*AssetAddressEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.AssetAddressEndpointRequest")
}

func init() { proto.RegisterFile("assetMessage.proto", fileDescriptor_assetMessage_c5355d3599f2f00a) }

var fileDescriptor_assetMessage_c5355d3599f2f00a = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb5, 0x6d, 0x93, 0xd0, 0x01, 0x71, 0x30, 0x52, 0xb4, 0xb4, 0x20, 0x05, 0x73, 0xc9,
	0x25, 0x8e, 0x04, 0x17, 0x2e, 0x08, 0x05, 0x09, 0x21, 0x24, 0x40, 0xc8, 0xea, 0x01, 0x71, 0xa9,
	0xc6, 0xeb, 0x8d, 0xb3, 0x8a, 0xbd, 0x63, 0x76, 0xc7, 0x84, 0x5e, 0xb9, 0xf2, 0x0e, 0xbc, 0x1d,
	0xef, 0x81, 0x6c, 0xaf, 0xa9, 0x0a, 0x94, 0xde, 0xda, 0x5b, 0xfe, 0x89, 0x67, 0xfe, 0xf9, 0xbf,
	0xb1, 0x21, 0x42, 0xef, 0x35, 0xbf, 0xd3, 0xde, 0x63, 0xa1, 0x93, 0xda, 0x11, 0x53, 0x14, 0xaf,
	0xd1, 0xb8, 0x1d, 0x9e, 0x29, 0x72, 0x75, 0x92, 0x95, 0xa4, 0xb6, 0x6a, 0x83, 0xc6, 0x76, 0x7f,
	0x66, 0xcd, 0x3a, 0xc9, 0x58, 0x1d, 0xdd, 0x53, 0x54, 0x55, 0x64, 0x2f, 0x34, 0xc6, 0x67, 0xf0,
	0xf8, 0xb5, 0xb6, 0xda, 0x21, 0xeb, 0x55, 0x3b, 0x76, 0x95, 0xe7, 0x4e, 0x7b, 0xff, 0xca, 0xe6,
	0x35, 0x19, 0xcb, 0xa9, 0xfe, 0xdc, 0x68, 0xcf, 0x51, 0x0a, 0x13, 0xab, 0x79, 0x47, 0x6e, 0x2b,
	0xc5, 0x4c, 0xcc, 0xef, 0x3e, 0x79, 0x96, 0x5c, 0xed, 0x98, 0xbc, 0xef, 0x5b, 0x56, 0x65, 0x49,
	0x3b, 0x63, 0x8b, 0x55, 0x69, 0xd0, 0xa7, 0xc3, 0xa0, 0xf8, 0xa7, 0x80, 0xfb, 0x6f, 0xbc, 0x6f,
	0x7a, 0xe3, 0x6b, 0x70, 0x8c, 0x1e, 0xc1, 0x9d, 0xb5, 0xa3, 0xea, 0xb4, 0x76, 0xe6, 0x0b, 0xb2,
	0x96, 0x7b, 0x33, 0x31, 0x3f, 0x4c, 0x6f, 0xb7, 0xb5, 0x0f, 0x7d, 0x29, 0x7a, 0x08, 0xc0, 0x74,
	0x8a, 0x3d, 0x05, 0xb9, 0xdf, 0x3d, 0x70, 0xc8, 0x14, 0xb0, 0x44, 0x53, 0x18, 0x63, 0x45, 0x8d,
	0x65, 0x79, 0x30, 0x13, 0xf3, 0x51, 0x1a, 0x54, 0x74, 0x04, 0xb7, 0x2a, 0xcd, 0x98, 0x23, 0xa3,
	0x1c, 0x75, 0x4d, 0xbf, 0x75, 0xfc, 0x6d, 0x0f, 0x1e, 0x9c, 0x38, 0xb4, 0x7e, 0xad, 0xdd, 0xb5,
	0x45, 0x95, 0x30, 0xe9, 0x5e, 0x13, 0x93, 0x87, 0x94, 0x83, 0xfc, 0x0b, 0xc2, 0xfe, 0x55, 0x10,
	0x0e, 0x2e, 0x87, 0x30, 0xba, 0x14, 0xc2, 0xf8, 0x0f, 0x08, 0xdf, 0x05, 0x1c, 0xbf, 0x35, 0x9e,
	0x3b, 0x00, 0x27, 0x1f, 0xfd, 0x8d, 0x32, 0x88, 0x7f, 0x08, 0x98, 0x86, 0x4d, 0x6e, 0xf6, 0x18,
	0x53, 0x18, 0xf3, 0xd7, 0x0d, 0xfa, 0x4d, 0x38, 0x43, 0x50, 0xf1, 0x16, 0x8e, 0xff, 0xf7, 0x39,
	0xca, 0xf3, 0x25, 0xa1, 0x1f, 0xf8, 0x0f, 0x2b, 0x71, 0xd1, 0x4a, 0xc2, 0x84, 0xb0, 0x6e, 0x8f,
	0x3a, 0x2c, 0x11, 0xe4, 0xcb, 0x17, 0x9f, 0x9e, 0x17, 0x86, 0x37, 0x4d, 0x96, 0x28, 0xaa, 0x96,
	0x21, 0xed, 0xa2, 0x8d, 0xbb, 0x3c, 0x8f, 0xbb, 0x18, 0xf2, 0x2e, 0xbb, 0x1f, 0x6a, 0x51, 0x68,
	0xbb, 0x28, 0x68, 0x99, 0xb1, 0xca, 0xc6, 0x5d, 0xe9, 0xe9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7c, 0xa7, 0x80, 0xda, 0x9a, 0x04, 0x00, 0x00,
}
