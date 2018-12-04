// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contractMessage_Eth.proto

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

type CreateContractEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias" json:"network,omitempty"`
	// Solidity code of this contract; required when creating a contract. In responses, only returned with contracts initially compiled by BlockCypher.
	Solidity string `protobuf:"bytes,2,opt,name=solidity,proto3" json:"solidity,omitempty"`
	// Parameters for either contract creation or method execution.
	Params []string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
	// Named contract(s) to publish; necessary to specify when first creating a contract.
	Publish []string `protobuf:"bytes,4,rep,name=publish,proto3" json:"publish,omitempty"`
	// Private key associated with a funded Ethereum external account used to publish a contract or execute a method.
	Private string `protobuf:"bytes,5,opt,name=private,proto3" json:"private,omitempty"`
	// Maximum amount of gas to use in contract creation or method execution.
	GasLimit int32 `protobuf:"varint,6,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// Amount (in wei) to transfer; can be used when creating a contract or calling a method.
	Value int32 `protobuf:"varint,7,opt,name=value,proto3" json:"value,omitempty"`
	// Name of contract as published.
	Name string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	// Hex-encoded binary compilation of this contract.
	Bin string `protobuf:"bytes,9,opt,name=bin,proto3" json:"bin,omitempty"`
	// JSON-encoded ABI. Only returned with contracts initially compiled by BlockCypher.
	Abi string `protobuf:"bytes,10,opt,name=abi,proto3" json:"abi,omitempty"`
	// Hex-encoded address of this contract.
	Address string `protobuf:"bytes,11,opt,name=address,proto3" json:"address,omitempty"`
	// Timestamp when this contract was confirmed in the Ethereum blockchain.
	Created string `protobuf:"bytes,12,opt,name=created,proto3" json:"created,omitempty"`
	// Hex-encoded transaction hash that created this contract.
	CreationTxHash string `protobuf:"bytes,13,opt,name=creation_tx_hash,json=creationTxHash,proto3" json:"creation_tx_hash,omitempty"`
	// If this is a response from a calling a contract method, this array of results may appear if the method returns any.
	Results              []string `protobuf:"bytes,14,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateContractEndpointRequest) Reset()         { *m = CreateContractEndpointRequest{} }
func (m *CreateContractEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*CreateContractEndpointRequest) ProtoMessage()    {}
func (*CreateContractEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_contractMessage_Eth_c10f399ebc726d4e, []int{0}
}
func (m *CreateContractEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateContractEndpointRequest.Unmarshal(m, b)
}
func (m *CreateContractEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateContractEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *CreateContractEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateContractEndpointRequest.Merge(dst, src)
}
func (m *CreateContractEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_CreateContractEndpointRequest.Size(m)
}
func (m *CreateContractEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateContractEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateContractEndpointRequest proto.InternalMessageInfo

func (m *CreateContractEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAIN
}

func (m *CreateContractEndpointRequest) GetSolidity() string {
	if m != nil {
		return m.Solidity
	}
	return ""
}

func (m *CreateContractEndpointRequest) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *CreateContractEndpointRequest) GetPublish() []string {
	if m != nil {
		return m.Publish
	}
	return nil
}

func (m *CreateContractEndpointRequest) GetPrivate() string {
	if m != nil {
		return m.Private
	}
	return ""
}

func (m *CreateContractEndpointRequest) GetGasLimit() int32 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *CreateContractEndpointRequest) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *CreateContractEndpointRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateContractEndpointRequest) GetBin() string {
	if m != nil {
		return m.Bin
	}
	return ""
}

func (m *CreateContractEndpointRequest) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *CreateContractEndpointRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CreateContractEndpointRequest) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *CreateContractEndpointRequest) GetCreationTxHash() string {
	if m != nil {
		return m.CreationTxHash
	}
	return ""
}

func (m *CreateContractEndpointRequest) GetResults() []string {
	if m != nil {
		return m.Results
	}
	return nil
}

type GetContractAddressEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias" json:"network,omitempty"`
	QueryAddress         string               `protobuf:"bytes,2,opt,name=query_address,json=queryAddress,proto3" json:"query_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetContractAddressEndpointRequest) Reset()         { *m = GetContractAddressEndpointRequest{} }
func (m *GetContractAddressEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*GetContractAddressEndpointRequest) ProtoMessage()    {}
func (*GetContractAddressEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_contractMessage_Eth_c10f399ebc726d4e, []int{1}
}
func (m *GetContractAddressEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContractAddressEndpointRequest.Unmarshal(m, b)
}
func (m *GetContractAddressEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContractAddressEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *GetContractAddressEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContractAddressEndpointRequest.Merge(dst, src)
}
func (m *GetContractAddressEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_GetContractAddressEndpointRequest.Size(m)
}
func (m *GetContractAddressEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContractAddressEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetContractAddressEndpointRequest proto.InternalMessageInfo

func (m *GetContractAddressEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAIN
}

func (m *GetContractAddressEndpointRequest) GetQueryAddress() string {
	if m != nil {
		return m.QueryAddress
	}
	return ""
}

type CallContractMethodEndpointRequest struct {
	Network      NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias" json:"network,omitempty"`
	QueryAddress string               `protobuf:"bytes,2,opt,name=query_address,json=queryAddress,proto3" json:"query_address,omitempty"`
	Method       string               `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	// Solidity code of this contract; required when creating a contract. In responses, only returned with contracts initially compiled by BlockCypher.
	Solidity string `protobuf:"bytes,4,opt,name=solidity,proto3" json:"solidity,omitempty"`
	// Parameters for either contract creation or method execution.
	Params []string `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty"`
	// Named contract(s) to publish; necessary to specify when first creating a contract.
	Publish []string `protobuf:"bytes,6,rep,name=publish,proto3" json:"publish,omitempty"`
	// Private key associated with a funded Ethereum external account used to publish a contract or execute a method.
	Private string `protobuf:"bytes,7,opt,name=private,proto3" json:"private,omitempty"`
	// Maximum amount of gas to use in contract creation or method execution.
	GasLimit int32 `protobuf:"varint,8,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// Amount (in wei) to transfer; can be used when creating a contract or calling a method.
	Value int32 `protobuf:"varint,9,opt,name=value,proto3" json:"value,omitempty"`
	// Name of contract as published.
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// Hex-encoded binary compilation of this contract.
	Bin string `protobuf:"bytes,11,opt,name=bin,proto3" json:"bin,omitempty"`
	// JSON-encoded ABI. Only returned with contracts initially compiled by BlockCypher.
	Abi string `protobuf:"bytes,12,opt,name=abi,proto3" json:"abi,omitempty"`
	// Hex-encoded address of this contract.
	Address string `protobuf:"bytes,13,opt,name=address,proto3" json:"address,omitempty"`
	// Timestamp when this contract was confirmed in the Ethereum blockchain.
	Created string `protobuf:"bytes,14,opt,name=created,proto3" json:"created,omitempty"`
	// Hex-encoded transaction hash that created this contract.
	CreationTxHash string `protobuf:"bytes,15,opt,name=creation_tx_hash,json=creationTxHash,proto3" json:"creation_tx_hash,omitempty"`
	// If this is a response from a calling a contract method, this array of results may appear if the method returns any.
	Results              []string `protobuf:"bytes,16,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallContractMethodEndpointRequest) Reset()         { *m = CallContractMethodEndpointRequest{} }
func (m *CallContractMethodEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*CallContractMethodEndpointRequest) ProtoMessage()    {}
func (*CallContractMethodEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_contractMessage_Eth_c10f399ebc726d4e, []int{2}
}
func (m *CallContractMethodEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallContractMethodEndpointRequest.Unmarshal(m, b)
}
func (m *CallContractMethodEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallContractMethodEndpointRequest.Marshal(b, m, deterministic)
}
func (dst *CallContractMethodEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallContractMethodEndpointRequest.Merge(dst, src)
}
func (m *CallContractMethodEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_CallContractMethodEndpointRequest.Size(m)
}
func (m *CallContractMethodEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallContractMethodEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallContractMethodEndpointRequest proto.InternalMessageInfo

func (m *CallContractMethodEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAIN
}

func (m *CallContractMethodEndpointRequest) GetQueryAddress() string {
	if m != nil {
		return m.QueryAddress
	}
	return ""
}

func (m *CallContractMethodEndpointRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *CallContractMethodEndpointRequest) GetSolidity() string {
	if m != nil {
		return m.Solidity
	}
	return ""
}

func (m *CallContractMethodEndpointRequest) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *CallContractMethodEndpointRequest) GetPublish() []string {
	if m != nil {
		return m.Publish
	}
	return nil
}

func (m *CallContractMethodEndpointRequest) GetPrivate() string {
	if m != nil {
		return m.Private
	}
	return ""
}

func (m *CallContractMethodEndpointRequest) GetGasLimit() int32 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *CallContractMethodEndpointRequest) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *CallContractMethodEndpointRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CallContractMethodEndpointRequest) GetBin() string {
	if m != nil {
		return m.Bin
	}
	return ""
}

func (m *CallContractMethodEndpointRequest) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *CallContractMethodEndpointRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CallContractMethodEndpointRequest) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *CallContractMethodEndpointRequest) GetCreationTxHash() string {
	if m != nil {
		return m.CreationTxHash
	}
	return ""
}

func (m *CallContractMethodEndpointRequest) GetResults() []string {
	if m != nil {
		return m.Results
	}
	return nil
}

type ContractArray struct {
	Contract             []*Contract `protobuf:"bytes,1,rep,name=contract,proto3" json:"contract,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ContractArray) Reset()         { *m = ContractArray{} }
func (m *ContractArray) String() string { return proto.CompactTextString(m) }
func (*ContractArray) ProtoMessage()    {}
func (*ContractArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_contractMessage_Eth_c10f399ebc726d4e, []int{3}
}
func (m *ContractArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractArray.Unmarshal(m, b)
}
func (m *ContractArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractArray.Marshal(b, m, deterministic)
}
func (dst *ContractArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractArray.Merge(dst, src)
}
func (m *ContractArray) XXX_Size() int {
	return xxx_messageInfo_ContractArray.Size(m)
}
func (m *ContractArray) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractArray.DiscardUnknown(m)
}

var xxx_messageInfo_ContractArray proto.InternalMessageInfo

func (m *ContractArray) GetContract() []*Contract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateContractEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.eth.CreateContractEndpointRequest")
	proto.RegisterType((*GetContractAddressEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.eth.GetContractAddressEndpointRequest")
	proto.RegisterType((*CallContractMethodEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.eth.CallContractMethodEndpointRequest")
	proto.RegisterType((*ContractArray)(nil), "fairwaycorp.blockchainprotobuf.eth.ContractArray")
}

func init() {
	proto.RegisterFile("contractMessage_Eth.proto", fileDescriptor_contractMessage_Eth_c10f399ebc726d4e)
}

var fileDescriptor_contractMessage_Eth_c10f399ebc726d4e = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x15, 0xd2, 0x5f, 0xf7, 0x87, 0xca, 0x42, 0x60, 0x06, 0x21, 0x75, 0xca, 0x26, 0x0b,
	0x9a, 0x4a, 0xc3, 0x86, 0x0d, 0x42, 0xa5, 0x1a, 0x31, 0x0b, 0x60, 0x11, 0xb1, 0x81, 0x4d, 0x75,
	0x93, 0x78, 0x12, 0x6b, 0x1c, 0xbb, 0x63, 0x3b, 0xd3, 0xe9, 0xbb, 0xf0, 0x2c, 0xbc, 0x1a, 0x28,
	0x4e, 0x52, 0x34, 0xea, 0x84, 0xb2, 0x42, 0xb3, 0xbb, 0xe7, 0x9e, 0xfa, 0xf8, 0xba, 0xf7, 0x53,
	0xd0, 0xf3, 0x48, 0x0a, 0xa3, 0x20, 0x32, 0x9f, 0xa9, 0xd6, 0x90, 0xd0, 0xf5, 0xb9, 0x49, 0xfd,
	0x8d, 0x92, 0x46, 0xe2, 0xd9, 0x25, 0x30, 0xb5, 0x85, 0x5d, 0x24, 0xd5, 0xc6, 0x0f, 0xb9, 0x8c,
	0xae, 0xa2, 0x14, 0x98, 0xb0, 0x66, 0x98, 0x5f, 0xfa, 0xd4, 0xa4, 0x27, 0xcf, 0x22, 0x99, 0x65,
	0x52, 0x1c, 0x1c, 0x9e, 0xfd, 0x74, 0xd1, 0xcb, 0x95, 0xa2, 0x60, 0xe8, 0xaa, 0xba, 0xe0, 0x5c,
	0xc4, 0x1b, 0xc9, 0x84, 0x09, 0xe8, 0x75, 0x4e, 0xb5, 0xc1, 0x01, 0xea, 0x0a, 0x6a, 0xb6, 0x52,
	0x5d, 0x11, 0x67, 0xea, 0x78, 0xe3, 0xb3, 0xb7, 0xfe, 0xf1, 0x0b, 0xfd, 0x2f, 0xe5, 0x91, 0x25,
	0xe7, 0x72, 0xcb, 0x44, 0xb2, 0xe4, 0x0c, 0x74, 0x50, 0x07, 0xe1, 0x13, 0xd4, 0xd3, 0x92, 0xb3,
	0x98, 0x99, 0x1d, 0x79, 0x34, 0x75, 0xbc, 0x7e, 0xb0, 0xd7, 0xf8, 0x29, 0xea, 0x6c, 0x40, 0x41,
	0xa6, 0x89, 0x3b, 0x75, 0xbd, 0x7e, 0x50, 0x29, 0x4c, 0x50, 0x77, 0x93, 0x87, 0x9c, 0xe9, 0x94,
	0xb4, 0xac, 0x51, 0x4b, 0xeb, 0x28, 0x76, 0x03, 0x86, 0x92, 0xb6, 0x0d, 0xab, 0x25, 0x7e, 0x81,
	0xfa, 0x09, 0xe8, 0x35, 0x67, 0x19, 0x33, 0xa4, 0x33, 0x75, 0xbc, 0x76, 0xd0, 0x4b, 0x40, 0x7f,
	0x2a, 0x34, 0x7e, 0x82, 0xda, 0x37, 0xc0, 0x73, 0x4a, 0xba, 0xd6, 0x28, 0x05, 0xc6, 0xa8, 0x25,
	0x20, 0xa3, 0xa4, 0x67, 0x93, 0x6c, 0x8d, 0x27, 0xc8, 0x0d, 0x99, 0x20, 0x7d, 0xdb, 0x2a, 0xca,
	0xa2, 0x03, 0x21, 0x23, 0xa8, 0xec, 0x40, 0xc8, 0x8a, 0x21, 0x20, 0x8e, 0x15, 0xd5, 0x9a, 0x0c,
	0xca, 0x21, 0x2a, 0x59, 0x38, 0x91, 0xfd, 0x87, 0x63, 0x32, 0x2c, 0x9d, 0x4a, 0x62, 0x0f, 0x4d,
	0x6c, 0xc9, 0xa4, 0x58, 0x9b, 0xdb, 0x75, 0x0a, 0x3a, 0x25, 0x23, 0xfb, 0x93, 0x71, 0xdd, 0xff,
	0x7a, 0x7b, 0x01, 0xe5, 0x13, 0x15, 0xd5, 0x39, 0x37, 0x9a, 0x8c, 0xcb, 0xc7, 0x57, 0x72, 0xf6,
	0xc3, 0x41, 0xa7, 0x1f, 0xa9, 0xa9, 0xb7, 0xb7, 0x2c, 0x2f, 0xfd, 0x1f, 0x4b, 0x7c, 0x85, 0x46,
	0xd7, 0x39, 0x55, 0xbb, 0x75, 0xfd, 0xee, 0x72, 0x93, 0x43, 0xdb, 0xac, 0xe6, 0x98, 0xfd, 0x72,
	0xd1, 0xe9, 0x0a, 0x38, 0x5f, 0xed, 0xf1, 0x35, 0xa9, 0x8c, 0x1f, 0xca, 0x78, 0x05, 0x6c, 0x99,
	0x9d, 0x88, 0xb8, 0xd6, 0xad, 0xd4, 0x1d, 0x40, 0x5b, 0x8d, 0x80, 0xb6, 0x9b, 0x00, 0xed, 0x34,
	0x02, 0xda, 0xfd, 0x0b, 0xa0, 0xbd, 0x26, 0x40, 0xfb, 0xf7, 0x01, 0x8a, 0x0e, 0x01, 0x1d, 0x1c,
	0x00, 0x3a, 0xbc, 0x17, 0xd0, 0x51, 0x23, 0xa0, 0xe3, 0xe3, 0x80, 0x3e, 0x3e, 0x06, 0xe8, 0xe4,
	0x2e, 0xa0, 0xdf, 0xd0, 0x68, 0x0f, 0xa7, 0x52, 0xb0, 0xc3, 0x17, 0xa8, 0x57, 0x7f, 0xcc, 0x88,
	0x33, 0x75, 0xbd, 0xc1, 0xd9, 0xeb, 0x7f, 0xd9, 0x76, 0x1d, 0x12, 0xec, 0x4f, 0x7f, 0x78, 0xff,
	0xfd, 0x5d, 0xc2, 0x4c, 0x9a, 0x87, 0x7e, 0x24, 0xb3, 0x45, 0x95, 0x31, 0x2f, 0x42, 0x16, 0x7f,
	0x42, 0xe6, 0x75, 0xca, 0xc2, 0x16, 0xd1, 0x3c, 0xa1, 0x62, 0x9e, 0xc8, 0x05, 0x35, 0x69, 0xd8,
	0xb1, 0xad, 0x37, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x75, 0xdf, 0xed, 0x8a, 0x5e, 0x05, 0x00,
	0x00,
}
