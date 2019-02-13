// Code generated by protoc-gen-go. DO NOT EDIT.
// source: webhooksMessage.proto

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

type CreateWebHookEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=chainweaver.protobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	// Identifier of the event; generated when a new request is created.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Type of event; can be unconfirmed-tx, new-block, confirmed-tx, tx-confirmation, double-spend-tx, tx-confidence.
	Event string `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	// optional Only objects with a matching hash will be sent. The hash can either be for a block or a transaction.
	Hash string `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	// optional Only transactions associated with the given wallet will be sent; can use a regular or HD wallet name. If used, requires a user token.
	WalletName string `protobuf:"bytes,5,opt,name=wallet_name,json=walletName,proto3" json:"wallet_name,omitempty"`
	// optional Required if wallet_name is used, though generally we advise users to include it (as they can reach API throttling thresholds rapidly).
	Token string `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	// optional Only transactions associated with the given address will be sent. A wallet name can also be used instead of an address, which will then match on any address in the wallet.
	Address string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	// optional Used in concert with the tx-confirmation event type to set the number of confirmations desired for which to receive an update. You’ll receive an updated TX for every confirmation up to this amount. The maximum allowed is 10; if not set, it will default to 6.
	Confirmations int32 `protobuf:"varint,8,opt,name=confirmations,proto3" json:"confirmations,omitempty"`
	// optional Used in concert with the tx-confidence event type to set the minimum confidence for which you’ll receive a notification. You’ll receive a TX once this threshold is met. Will accept any float between 0 and 1, exclusive; if not set, defaults to 0.99.
	Confidence float32 `protobuf:"fixed32,9,opt,name=confidence,proto3" json:"confidence,omitempty"`
	// optional Only transactions with an output script of the provided type will be sent. The recognized types of scripts are: pay-to-pubkey-hash, pay-to-multi-pubkey-hash, pay-to-pubkey, pay-to-script-hash, null-data (sometimes called OP_RETURN), empty or unknown.
	Script string `protobuf:"bytes,10,opt,name=script,proto3" json:"script,omitempty"`
	// optional Callback URL for this Event’s WebHook; not applicable for WebSockets usage.
	Url string `protobuf:"bytes,11,opt,name=url,proto3" json:"url,omitempty"`
	// Number of errors when attempting to POST to callback URL; not applicable for WebSockets usage.
	CallbackErrors       int32    `protobuf:"varint,12,opt,name=callback_errors,json=callbackErrors,proto3" json:"callback_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWebHookEndpointRequest) Reset()         { *m = CreateWebHookEndpointRequest{} }
func (m *CreateWebHookEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWebHookEndpointRequest) ProtoMessage()    {}
func (*CreateWebHookEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_webhooksMessage_24e5fd76531850b1, []int{0}
}
func (m *CreateWebHookEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWebHookEndpointRequest.Unmarshal(m, b)
}
func (m *CreateWebHookEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWebHookEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *CreateWebHookEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWebHookEndpointRequest.Merge(dst, src)
}
func (m *CreateWebHookEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWebHookEndpointRequest.Size(m)
}
func (m *CreateWebHookEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWebHookEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWebHookEndpointRequest proto.InternalMessageInfo

func (m *CreateWebHookEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *CreateWebHookEndpointRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateWebHookEndpointRequest) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *CreateWebHookEndpointRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *CreateWebHookEndpointRequest) GetWalletName() string {
	if m != nil {
		return m.WalletName
	}
	return ""
}

func (m *CreateWebHookEndpointRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CreateWebHookEndpointRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CreateWebHookEndpointRequest) GetConfirmations() int32 {
	if m != nil {
		return m.Confirmations
	}
	return 0
}

func (m *CreateWebHookEndpointRequest) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

func (m *CreateWebHookEndpointRequest) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *CreateWebHookEndpointRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CreateWebHookEndpointRequest) GetCallbackErrors() int32 {
	if m != nil {
		return m.CallbackErrors
	}
	return 0
}

type ListWebHooksEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=chainweaver.protobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Token                string               `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListWebHooksEndpointRequest) Reset()         { *m = ListWebHooksEndpointRequest{} }
func (m *ListWebHooksEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*ListWebHooksEndpointRequest) ProtoMessage()    {}
func (*ListWebHooksEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_webhooksMessage_24e5fd76531850b1, []int{1}
}
func (m *ListWebHooksEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWebHooksEndpointRequest.Unmarshal(m, b)
}
func (m *ListWebHooksEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWebHooksEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *ListWebHooksEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWebHooksEndpointRequest.Merge(dst, src)
}
func (m *ListWebHooksEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_ListWebHooksEndpointRequest.Size(m)
}
func (m *ListWebHooksEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWebHooksEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWebHooksEndpointRequest proto.InternalMessageInfo

func (m *ListWebHooksEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *ListWebHooksEndpointRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Events struct {
	Event                []*Event `protobuf:"bytes,1,rep,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Events) Reset()         { *m = Events{} }
func (m *Events) String() string { return proto.CompactTextString(m) }
func (*Events) ProtoMessage()    {}
func (*Events) Descriptor() ([]byte, []int) {
	return fileDescriptor_webhooksMessage_24e5fd76531850b1, []int{2}
}
func (m *Events) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Events.Unmarshal(m, b)
}
func (m *Events) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Events.Marshal(b, m, deterministic)
}
func (dst *Events) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Events.Merge(dst, src)
}
func (m *Events) XXX_Size() int {
	return xxx_messageInfo_Events.Size(m)
}
func (m *Events) XXX_DiscardUnknown() {
	xxx_messageInfo_Events.DiscardUnknown(m)
}

var xxx_messageInfo_Events proto.InternalMessageInfo

func (m *Events) GetEvent() []*Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type WebHookIDEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=chainweaver.protobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Webhookid            string               `protobuf:"bytes,2,opt,name=webhookid,proto3" json:"webhookid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WebHookIDEndpointRequest) Reset()         { *m = WebHookIDEndpointRequest{} }
func (m *WebHookIDEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*WebHookIDEndpointRequest) ProtoMessage()    {}
func (*WebHookIDEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_webhooksMessage_24e5fd76531850b1, []int{3}
}
func (m *WebHookIDEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebHookIDEndpointRequest.Unmarshal(m, b)
}
func (m *WebHookIDEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebHookIDEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *WebHookIDEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebHookIDEndpointRequest.Merge(dst, src)
}
func (m *WebHookIDEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_WebHookIDEndpointRequest.Size(m)
}
func (m *WebHookIDEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WebHookIDEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WebHookIDEndpointRequest proto.InternalMessageInfo

func (m *WebHookIDEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *WebHookIDEndpointRequest) GetWebhookid() string {
	if m != nil {
		return m.Webhookid
	}
	return ""
}

type DeleteWebHookEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=chainweaver.protobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Webhookid            string               `protobuf:"bytes,2,opt,name=webhookid,proto3" json:"webhookid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeleteWebHookEndpointRequest) Reset()         { *m = DeleteWebHookEndpointRequest{} }
func (m *DeleteWebHookEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteWebHookEndpointRequest) ProtoMessage()    {}
func (*DeleteWebHookEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_webhooksMessage_24e5fd76531850b1, []int{4}
}
func (m *DeleteWebHookEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteWebHookEndpointRequest.Unmarshal(m, b)
}
func (m *DeleteWebHookEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteWebHookEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteWebHookEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteWebHookEndpointRequest.Merge(dst, src)
}
func (m *DeleteWebHookEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteWebHookEndpointRequest.Size(m)
}
func (m *DeleteWebHookEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteWebHookEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteWebHookEndpointRequest proto.InternalMessageInfo

func (m *DeleteWebHookEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *DeleteWebHookEndpointRequest) GetWebhookid() string {
	if m != nil {
		return m.Webhookid
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateWebHookEndpointRequest)(nil), "chainweaver.protobuf.btc.CreateWebHookEndpointRequest")
	proto.RegisterType((*ListWebHooksEndpointRequest)(nil), "chainweaver.protobuf.btc.ListWebHooksEndpointRequest")
	proto.RegisterType((*Events)(nil), "chainweaver.protobuf.btc.Events")
	proto.RegisterType((*WebHookIDEndpointRequest)(nil), "chainweaver.protobuf.btc.WebHookIDEndpointRequest")
	proto.RegisterType((*DeleteWebHookEndpointRequest)(nil), "chainweaver.protobuf.btc.DeleteWebHookEndpointRequest")
}

func init() {
	proto.RegisterFile("webhooksMessage.proto", fileDescriptor_webhooksMessage_24e5fd76531850b1)
}

var fileDescriptor_webhooksMessage_24e5fd76531850b1 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0x6f, 0xd4, 0x3e,
	0x10, 0xc5, 0x95, 0x6c, 0x77, 0xf7, 0xbb, 0xb3, 0x5f, 0x16, 0x64, 0x7e, 0xc8, 0x82, 0x15, 0x8d,
	0x22, 0x24, 0x72, 0x69, 0x56, 0xb4, 0xe2, 0x8c, 0x0a, 0x5d, 0xa9, 0x48, 0xd0, 0x43, 0x2e, 0x48,
	0x5c, 0x2a, 0xc7, 0x99, 0x26, 0x56, 0x1c, 0x7b, 0xb1, 0xbd, 0xcd, 0x89, 0x0b, 0x07, 0x4e, 0xfc,
	0xd1, 0x68, 0x9d, 0xa4, 0x6c, 0x0f, 0xe5, 0x56, 0x6e, 0xf3, 0x3e, 0xb6, 0xe7, 0x8d, 0xfc, 0x06,
	0x9e, 0xb6, 0x98, 0x57, 0x5a, 0xd7, 0xf6, 0x33, 0x5a, 0xcb, 0x4a, 0x4c, 0x37, 0x46, 0x3b, 0x4d,
	0x28, 0xaf, 0x98, 0x50, 0x2d, 0xb2, 0x6b, 0x34, 0x1d, 0xca, 0xb7, 0x57, 0x69, 0xee, 0xf8, 0xf3,
	0xc7, 0x5c, 0x37, 0x8d, 0x56, 0xb7, 0xae, 0xc7, 0xbf, 0x46, 0xb0, 0xfc, 0x60, 0x90, 0x39, 0xfc,
	0x82, 0xf9, 0xb9, 0xd6, 0xf5, 0x5a, 0x15, 0x1b, 0x2d, 0x94, 0xcb, 0xf0, 0xdb, 0x16, 0xad, 0x23,
	0xe7, 0x30, 0x55, 0xe8, 0x5a, 0x6d, 0x6a, 0x1a, 0x44, 0x41, 0xb2, 0x38, 0x4e, 0xd3, 0xbb, 0x1c,
	0xd2, 0x8b, 0xee, 0xe2, 0xa9, 0x94, 0xba, 0x15, 0xaa, 0x3c, 0x95, 0x82, 0xd9, 0x6c, 0x78, 0x4e,
	0x16, 0x10, 0x8a, 0x82, 0x86, 0x51, 0x90, 0xcc, 0xb2, 0x50, 0x14, 0xe4, 0x09, 0x8c, 0xf1, 0x1a,
	0x95, 0xa3, 0x23, 0x8f, 0x3a, 0x41, 0x08, 0x1c, 0x54, 0xcc, 0x56, 0xf4, 0xc0, 0x43, 0x5f, 0x93,
	0x43, 0x98, 0xb7, 0x4c, 0x4a, 0x74, 0x97, 0x8a, 0x35, 0x48, 0xc7, 0xfe, 0x08, 0x3a, 0x74, 0xc1,
	0x1a, 0xdc, 0xb5, 0x72, 0xba, 0x46, 0x45, 0x27, 0x5d, 0x2b, 0x2f, 0x08, 0x85, 0x29, 0x2b, 0x0a,
	0x83, 0xd6, 0xd2, 0xa9, 0xe7, 0x83, 0x24, 0xaf, 0xe0, 0x01, 0xd7, 0xea, 0x4a, 0x98, 0x86, 0x39,
	0xa1, 0x95, 0xa5, 0xff, 0x45, 0x41, 0x32, 0xce, 0x6e, 0x43, 0xf2, 0x12, 0xc0, 0x83, 0x02, 0x15,
	0x47, 0x3a, 0x8b, 0x82, 0x24, 0xcc, 0xf6, 0x08, 0x79, 0x06, 0x13, 0xcb, 0x8d, 0xd8, 0x38, 0x0a,
	0xbe, 0x7d, 0xaf, 0xc8, 0x23, 0x18, 0x6d, 0x8d, 0xa4, 0x73, 0x0f, 0x77, 0x25, 0x79, 0x0d, 0x0f,
	0x39, 0x93, 0x32, 0x67, 0xbc, 0xbe, 0x44, 0x63, 0xb4, 0xb1, 0xf4, 0x7f, 0xef, 0xb8, 0x18, 0xf0,
	0xda, 0xd3, 0xf8, 0x3b, 0xbc, 0xf8, 0x24, 0xac, 0xeb, 0xb3, 0xb0, 0xf7, 0x17, 0xc6, 0xcd, 0x8f,
	0x85, 0x7b, 0x3f, 0x16, 0xbf, 0x83, 0xc9, 0x7a, 0x97, 0x82, 0x25, 0x6f, 0x87, 0x70, 0x82, 0x68,
	0x94, 0xcc, 0x8f, 0x0f, 0xef, 0xf6, 0xf1, 0x0f, 0xfa, 0xf4, 0xe2, 0x1f, 0x01, 0xd0, 0x7e, 0xf8,
	0x8f, 0x67, 0xf7, 0x37, 0xfd, 0x12, 0x66, 0xfd, 0xf6, 0xdf, 0x6c, 0xd4, 0x1f, 0x10, 0xff, 0x0c,
	0x60, 0x79, 0x86, 0x12, 0xff, 0xc1, 0x4e, 0xff, 0x75, 0x90, 0xf7, 0x27, 0x5f, 0xdf, 0x94, 0xc2,
	0x55, 0xdb, 0x3c, 0xe5, 0xba, 0x59, 0xed, 0x59, 0xac, 0x06, 0x8b, 0xae, 0xe0, 0x47, 0x25, 0xaa,
	0xa3, 0x52, 0xaf, 0x72, 0xc7, 0xf3, 0x89, 0x47, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb5,
	0xc0, 0xd9, 0x69, 0xe0, 0x03, 0x00, 0x00,
}
