// Code generated by protoc-gen-go. DO NOT EDIT.
// source: analyticsMessage.proto

package btc // import "github.com/chainweaver/protobuf/protoc-gen-go/btc"

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
	Network    NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=chainweaver.protobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
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
	return fileDescriptor_analyticsMessage_a65ac115c748173a, []int{0}
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
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=chainweaver.protobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Ticket               string               `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AnalyticsJobRequest) Reset()         { *m = AnalyticsJobRequest{} }
func (m *AnalyticsJobRequest) String() string { return proto.CompactTextString(m) }
func (*AnalyticsJobRequest) ProtoMessage()    {}
func (*AnalyticsJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyticsMessage_a65ac115c748173a, []int{1}
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
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=chainweaver.protobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
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
	return fileDescriptor_analyticsMessage_a65ac115c748173a, []int{2}
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
	proto.RegisterType((*CreateAnalyticsJobRequest)(nil), "chainweaver.protobuf.btc.CreateAnalyticsJobRequest")
	proto.RegisterType((*AnalyticsJobRequest)(nil), "chainweaver.protobuf.btc.AnalyticsJobRequest")
	proto.RegisterType((*AnalyticsJobResultsRequest)(nil), "chainweaver.protobuf.btc.AnalyticsJobResultsRequest")
}

func init() {
	proto.RegisterFile("analyticsMessage.proto", fileDescriptor_analyticsMessage_a65ac115c748173a)
}

var fileDescriptor_analyticsMessage_a65ac115c748173a = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0xbd, 0x6e, 0xdb, 0x30,
	0x10, 0x80, 0x21, 0xdb, 0x92, 0x6b, 0x0e, 0x6e, 0x41, 0x17, 0x06, 0xeb, 0xa1, 0x30, 0xbc, 0xd4,
	0x8b, 0x65, 0xb4, 0x7e, 0x02, 0xb7, 0x4b, 0x51, 0xa0, 0x19, 0x84, 0x4c, 0x59, 0x02, 0x8a, 0xba,
	0x50, 0x84, 0x29, 0xd2, 0x21, 0x4f, 0x16, 0xf2, 0x1a, 0x79, 0x98, 0x3c, 0x5f, 0x60, 0x4a, 0x02,
	0xec, 0x21, 0x63, 0xb2, 0xdd, 0xf7, 0xdd, 0x1f, 0x71, 0x24, 0x73, 0x6e, 0xb8, 0x7e, 0x42, 0x25,
	0xfc, 0x7f, 0xf0, 0x9e, 0x4b, 0x48, 0x8f, 0xce, 0xa2, 0xa5, 0x4c, 0x94, 0x5c, 0x99, 0x06, 0xf8,
	0x09, 0x5c, 0xab, 0xf2, 0xfa, 0x21, 0xcd, 0x51, 0x2c, 0x66, 0xc2, 0x56, 0x95, 0x35, 0x57, 0xe5,
	0xab, 0x97, 0x01, 0xf9, 0xf6, 0xc7, 0x01, 0x47, 0xd8, 0xf7, 0xf3, 0xfe, 0xd9, 0x3c, 0x83, 0xc7,
	0x1a, 0x3c, 0xd2, 0xbf, 0x64, 0x6c, 0x00, 0x1b, 0xeb, 0x0e, 0x2c, 0x5a, 0x46, 0xeb, 0xe9, 0xaf,
	0x34, 0x7d, 0x6b, 0x7c, 0x7a, 0xd3, 0x16, 0xee, 0xb5, 0xb6, 0x8d, 0x32, 0x72, 0xaf, 0x15, 0xf7,
	0x59, 0xdf, 0x4e, 0xbf, 0x13, 0x02, 0x46, 0x2a, 0x03, 0x86, 0x57, 0xc0, 0x06, 0xcb, 0x68, 0x3d,
	0xc9, 0x2e, 0x0c, 0x65, 0x64, 0xcc, 0x8b, 0xc2, 0x81, 0xf7, 0x6c, 0x18, 0x92, 0x3d, 0xd2, 0x1f,
	0xe4, 0xf3, 0x89, 0xeb, 0x1a, 0xee, 0xb1, 0x74, 0xe0, 0x4b, 0xab, 0x0b, 0x36, 0x5a, 0x46, 0xeb,
	0x38, 0x9b, 0x06, 0x7d, 0xdb, 0x5b, 0xfa, 0x95, 0xc4, 0x5a, 0x55, 0x0a, 0x59, 0x1c, 0xd2, 0x2d,
	0x9c, 0xad, 0x47, 0xee, 0x90, 0x25, 0x61, 0x6c, 0x0b, 0xf4, 0x0b, 0x19, 0x82, 0x29, 0xd8, 0x38,
	0xb8, 0x73, 0x48, 0xe7, 0x24, 0x29, 0x40, 0x3a, 0x00, 0xf6, 0x29, 0xb4, 0x77, 0x74, 0xf6, 0xde,
	0xd6, 0x4e, 0x00, 0x9b, 0x84, 0xe2, 0x8e, 0x56, 0x0d, 0x99, 0xbd, 0xef, 0xc5, 0xe6, 0x24, 0x41,
	0x25, 0x0e, 0x80, 0xdd, 0xb5, 0x3a, 0x5a, 0x3d, 0x47, 0x64, 0x71, 0xbd, 0xd9, 0xd7, 0x1a, 0xfd,
	0x87, 0x3d, 0x80, 0x52, 0x32, 0x3a, 0x72, 0x09, 0xe1, 0x9f, 0xe2, 0x2c, 0xc4, 0xbf, 0x77, 0x77,
	0x3f, 0xa5, 0xc2, 0xb2, 0xce, 0x53, 0x61, 0xab, 0xed, 0xc5, 0xc2, 0x6d, 0xbf, 0xb0, 0x0d, 0xc4,
	0x46, 0x82, 0xd9, 0x48, 0xbb, 0xcd, 0x51, 0xe4, 0x49, 0x50, 0xbb, 0xd7, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x55, 0xa3, 0x38, 0xd6, 0xcb, 0x02, 0x00, 0x00,
}
