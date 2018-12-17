// Code generated by protoc-gen-go. DO NOT EDIT.
// source: analyticsMessage.proto

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

type CreateAnalyticsJobRequest struct {
	Network    NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Enginename string               `protobuf:"bytes,2,opt,name=enginename,proto3" json:"enginename,omitempty"`
	// Address hash this job is querying.
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// Minimal/threshold value (in satoshis) to query.
	ValueThreshold int32 `protobuf:"varint,4,opt,name=value_threshold,json=valueThreshold,proto3" json:"value_threshold,omitempty"`
	// Limit of results to return.
	Limit int32 `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	// Beginning of time range to query.
	Start string `protobuf:"bytes,6,opt,name=start,proto3" json:"start,omitempty"`
	// End of time range to query.
	End string `protobuf:"bytes,7,opt,name=end,proto3" json:"end,omitempty"`
	// Degree of connectiveness to query.
	Degree int32 `protobuf:"varint,8,opt,name=degree,proto3" json:"degree,omitempty"`
	// IP address and port, of the form “0.0.0.0:80”. Ideally an IP and port combination found from another API lookup (for example, relayed_by from the Transaction Hash Endpoint)
	Source               string   `protobuf:"bytes,9,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAnalyticsJobRequest) Reset()         { *m = CreateAnalyticsJobRequest{} }
func (m *CreateAnalyticsJobRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAnalyticsJobRequest) ProtoMessage()    {}
func (*CreateAnalyticsJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyticsMessage_1e1b57a2d0934132, []int{0}
}
func (m *CreateAnalyticsJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAnalyticsJobRequest.Unmarshal(m, b)
}
func (m *CreateAnalyticsJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAnalyticsJobRequest.Marshal(b, m, deterministic)
}
func (dst *CreateAnalyticsJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAnalyticsJobRequest.Merge(dst, src)
}
func (m *CreateAnalyticsJobRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAnalyticsJobRequest.Size(m)
}
func (m *CreateAnalyticsJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAnalyticsJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAnalyticsJobRequest proto.InternalMessageInfo

func (m *CreateAnalyticsJobRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *CreateAnalyticsJobRequest) GetEnginename() string {
	if m != nil {
		return m.Enginename
	}
	return ""
}

func (m *CreateAnalyticsJobRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CreateAnalyticsJobRequest) GetValueThreshold() int32 {
	if m != nil {
		return m.ValueThreshold
	}
	return 0
}

func (m *CreateAnalyticsJobRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *CreateAnalyticsJobRequest) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *CreateAnalyticsJobRequest) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *CreateAnalyticsJobRequest) GetDegree() int32 {
	if m != nil {
		return m.Degree
	}
	return 0
}

func (m *CreateAnalyticsJobRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type AnalyticsJobRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Ticket               string               `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AnalyticsJobRequest) Reset()         { *m = AnalyticsJobRequest{} }
func (m *AnalyticsJobRequest) String() string { return proto.CompactTextString(m) }
func (*AnalyticsJobRequest) ProtoMessage()    {}
func (*AnalyticsJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyticsMessage_1e1b57a2d0934132, []int{1}
}
func (m *AnalyticsJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyticsJobRequest.Unmarshal(m, b)
}
func (m *AnalyticsJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyticsJobRequest.Marshal(b, m, deterministic)
}
func (dst *AnalyticsJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyticsJobRequest.Merge(dst, src)
}
func (m *AnalyticsJobRequest) XXX_Size() int {
	return xxx_messageInfo_AnalyticsJobRequest.Size(m)
}
func (m *AnalyticsJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyticsJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyticsJobRequest proto.InternalMessageInfo

func (m *AnalyticsJobRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *AnalyticsJobRequest) GetTicket() string {
	if m != nil {
		return m.Ticket
	}
	return ""
}

type AnalyticsJobResultsRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Ticket               string               `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Page                 int32                `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AnalyticsJobResultsRequest) Reset()         { *m = AnalyticsJobResultsRequest{} }
func (m *AnalyticsJobResultsRequest) String() string { return proto.CompactTextString(m) }
func (*AnalyticsJobResultsRequest) ProtoMessage()    {}
func (*AnalyticsJobResultsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyticsMessage_1e1b57a2d0934132, []int{2}
}
func (m *AnalyticsJobResultsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyticsJobResultsRequest.Unmarshal(m, b)
}
func (m *AnalyticsJobResultsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyticsJobResultsRequest.Marshal(b, m, deterministic)
}
func (dst *AnalyticsJobResultsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyticsJobResultsRequest.Merge(dst, src)
}
func (m *AnalyticsJobResultsRequest) XXX_Size() int {
	return xxx_messageInfo_AnalyticsJobResultsRequest.Size(m)
}
func (m *AnalyticsJobResultsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyticsJobResultsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyticsJobResultsRequest proto.InternalMessageInfo

func (m *AnalyticsJobResultsRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *AnalyticsJobResultsRequest) GetTicket() string {
	if m != nil {
		return m.Ticket
	}
	return ""
}

func (m *AnalyticsJobResultsRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateAnalyticsJobRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.CreateAnalyticsJobRequest")
	proto.RegisterType((*AnalyticsJobRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.AnalyticsJobRequest")
	proto.RegisterType((*AnalyticsJobResultsRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.AnalyticsJobResultsRequest")
}

func init() {
	proto.RegisterFile("analyticsMessage.proto", fileDescriptor_analyticsMessage_1e1b57a2d0934132)
}

var fileDescriptor_analyticsMessage_1e1b57a2d0934132 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x91, 0xb1, 0x6e, 0xdb, 0x30,
	0x10, 0x86, 0x21, 0xdb, 0x92, 0x6b, 0x0e, 0x6e, 0x41, 0x17, 0x06, 0xeb, 0xa1, 0x30, 0xb4, 0xd4,
	0x8b, 0x24, 0xa0, 0x5d, 0xba, 0x04, 0x81, 0x93, 0x2d, 0x40, 0x32, 0x08, 0x99, 0xb2, 0x04, 0x14,
	0x75, 0xa6, 0x08, 0x53, 0xa4, 0x43, 0x52, 0x31, 0x3c, 0xe6, 0x3d, 0xf2, 0x30, 0x79, 0xb4, 0x40,
	0xb4, 0x84, 0xc4, 0x53, 0xb6, 0x64, 0xbb, 0xff, 0x3b, 0xde, 0xff, 0x1f, 0x8e, 0x68, 0x4e, 0x15,
	0x95, 0x07, 0x27, 0x98, 0xbd, 0x06, 0x6b, 0x29, 0x87, 0x74, 0x67, 0xb4, 0xd3, 0x38, 0xde, 0x50,
	0x61, 0xf6, 0xf4, 0xc0, 0xb4, 0xd9, 0xa5, 0x85, 0xd4, 0x6c, 0xcb, 0x2a, 0x2a, 0x94, 0x6f, 0x16,
	0xcd, 0x26, 0x2d, 0x1c, 0x5b, 0xcc, 0x98, 0xae, 0x6b, 0xad, 0x4e, 0x06, 0xe3, 0x97, 0x01, 0xfa,
	0x75, 0x69, 0x80, 0x3a, 0x58, 0xf7, 0xce, 0x57, 0xba, 0xc8, 0xe1, 0xa1, 0x01, 0xeb, 0x70, 0x8e,
	0xc6, 0x0a, 0xdc, 0x5e, 0x9b, 0x2d, 0x09, 0x96, 0xc1, 0x6a, 0xfa, 0xf7, 0x7f, 0xfa, 0x71, 0x50,
	0x7a, 0x73, 0x1c, 0x59, 0x4b, 0xa9, 0xf7, 0x42, 0xf1, 0xb5, 0x14, 0xd4, 0xe6, 0xbd, 0x11, 0xfe,
	0x8d, 0x10, 0x28, 0x2e, 0x14, 0x28, 0x5a, 0x03, 0x19, 0x2c, 0x83, 0xd5, 0x24, 0x7f, 0x47, 0x30,
	0x41, 0x63, 0x5a, 0x96, 0x06, 0xac, 0x25, 0x43, 0xdf, 0xec, 0x25, 0xfe, 0x83, 0xbe, 0x3f, 0x52,
	0xd9, 0xc0, 0xbd, 0xab, 0x0c, 0xd8, 0x4a, 0xcb, 0x92, 0x8c, 0x96, 0xc1, 0x2a, 0xcc, 0xa7, 0x1e,
	0xdf, 0xf6, 0x14, 0xff, 0x44, 0xa1, 0x14, 0xb5, 0x70, 0x24, 0xf4, 0xed, 0xa3, 0x68, 0xa9, 0x75,
	0xd4, 0x38, 0x12, 0x79, 0xdb, 0xa3, 0xc0, 0x3f, 0xd0, 0x10, 0x54, 0x49, 0xc6, 0x9e, 0xb5, 0x25,
	0x9e, 0xa3, 0xa8, 0x04, 0x6e, 0x00, 0xc8, 0x37, 0x3f, 0xde, 0xa9, 0x96, 0x5b, 0xdd, 0x18, 0x06,
	0x64, 0xe2, 0x1f, 0x77, 0x2a, 0x7e, 0x0a, 0xd0, 0xec, 0xb3, 0x8e, 0x37, 0x47, 0x91, 0x13, 0x6c,
	0x0b, 0xae, 0x3b, 0x5c, 0xa7, 0xe2, 0xe7, 0x00, 0x2d, 0x4e, 0x77, 0xb0, 0x8d, 0x74, 0xf6, 0x0b,
	0x56, 0xc1, 0x18, 0x8d, 0x76, 0x94, 0x83, 0xff, 0xbc, 0x30, 0xf7, 0xf5, 0xc5, 0xf9, 0xdd, 0x19,
	0x17, 0xae, 0x6a, 0x8a, 0x94, 0xe9, 0x3a, 0xeb, 0xa2, 0x93, 0x36, 0x3b, 0x7b, 0xcb, 0x4e, 0xfa,
	0xf0, 0xcc, 0x17, 0x2c, 0xe1, 0xa0, 0x12, 0xae, 0xb3, 0xc2, 0xb1, 0x22, 0xf2, 0xe8, 0xdf, 0x6b,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0xfb, 0xa4, 0xe7, 0x00, 0x03, 0x00, 0x00,
}
