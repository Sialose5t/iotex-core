// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rewarding.proto

package rewardingpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Admin struct {
	Admin                      []byte   `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	BlockReward                string   `protobuf:"bytes,2,opt,name=blockReward,proto3" json:"blockReward,omitempty"`
	EpochReward                string   `protobuf:"bytes,3,opt,name=epochReward,proto3" json:"epochReward,omitempty"`
	NumDelegatesForEpochReward uint64   `protobuf:"varint,4,opt,name=numDelegatesForEpochReward,proto3" json:"numDelegatesForEpochReward,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Admin) Reset()         { *m = Admin{} }
func (m *Admin) String() string { return proto.CompactTextString(m) }
func (*Admin) ProtoMessage()    {}
func (*Admin) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a8d72c965c1359, []int{0}
}

func (m *Admin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Admin.Unmarshal(m, b)
}
func (m *Admin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Admin.Marshal(b, m, deterministic)
}
func (m *Admin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Admin.Merge(m, src)
}
func (m *Admin) XXX_Size() int {
	return xxx_messageInfo_Admin.Size(m)
}
func (m *Admin) XXX_DiscardUnknown() {
	xxx_messageInfo_Admin.DiscardUnknown(m)
}

var xxx_messageInfo_Admin proto.InternalMessageInfo

func (m *Admin) GetAdmin() []byte {
	if m != nil {
		return m.Admin
	}
	return nil
}

func (m *Admin) GetBlockReward() string {
	if m != nil {
		return m.BlockReward
	}
	return ""
}

func (m *Admin) GetEpochReward() string {
	if m != nil {
		return m.EpochReward
	}
	return ""
}

func (m *Admin) GetNumDelegatesForEpochReward() uint64 {
	if m != nil {
		return m.NumDelegatesForEpochReward
	}
	return 0
}

type Fund struct {
	TotalBalance         string   `protobuf:"bytes,1,opt,name=totalBalance,proto3" json:"totalBalance,omitempty"`
	UnclaimedBalance     string   `protobuf:"bytes,2,opt,name=unclaimedBalance,proto3" json:"unclaimedBalance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fund) Reset()         { *m = Fund{} }
func (m *Fund) String() string { return proto.CompactTextString(m) }
func (*Fund) ProtoMessage()    {}
func (*Fund) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a8d72c965c1359, []int{1}
}

func (m *Fund) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fund.Unmarshal(m, b)
}
func (m *Fund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fund.Marshal(b, m, deterministic)
}
func (m *Fund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fund.Merge(m, src)
}
func (m *Fund) XXX_Size() int {
	return xxx_messageInfo_Fund.Size(m)
}
func (m *Fund) XXX_DiscardUnknown() {
	xxx_messageInfo_Fund.DiscardUnknown(m)
}

var xxx_messageInfo_Fund proto.InternalMessageInfo

func (m *Fund) GetTotalBalance() string {
	if m != nil {
		return m.TotalBalance
	}
	return ""
}

func (m *Fund) GetUnclaimedBalance() string {
	if m != nil {
		return m.UnclaimedBalance
	}
	return ""
}

type RewardHistory struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RewardHistory) Reset()         { *m = RewardHistory{} }
func (m *RewardHistory) String() string { return proto.CompactTextString(m) }
func (*RewardHistory) ProtoMessage()    {}
func (*RewardHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a8d72c965c1359, []int{2}
}

func (m *RewardHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RewardHistory.Unmarshal(m, b)
}
func (m *RewardHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RewardHistory.Marshal(b, m, deterministic)
}
func (m *RewardHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardHistory.Merge(m, src)
}
func (m *RewardHistory) XXX_Size() int {
	return xxx_messageInfo_RewardHistory.Size(m)
}
func (m *RewardHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardHistory.DiscardUnknown(m)
}

var xxx_messageInfo_RewardHistory proto.InternalMessageInfo

type Account struct {
	Balance              string   `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a8d72c965c1359, []int{3}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

type Exempt struct {
	Addrs                [][]byte `protobuf:"bytes,1,rep,name=addrs,proto3" json:"addrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Exempt) Reset()         { *m = Exempt{} }
func (m *Exempt) String() string { return proto.CompactTextString(m) }
func (*Exempt) ProtoMessage()    {}
func (*Exempt) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a8d72c965c1359, []int{4}
}

func (m *Exempt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Exempt.Unmarshal(m, b)
}
func (m *Exempt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Exempt.Marshal(b, m, deterministic)
}
func (m *Exempt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Exempt.Merge(m, src)
}
func (m *Exempt) XXX_Size() int {
	return xxx_messageInfo_Exempt.Size(m)
}
func (m *Exempt) XXX_DiscardUnknown() {
	xxx_messageInfo_Exempt.DiscardUnknown(m)
}

var xxx_messageInfo_Exempt proto.InternalMessageInfo

func (m *Exempt) GetAddrs() [][]byte {
	if m != nil {
		return m.Addrs
	}
	return nil
}

func init() {
	proto.RegisterType((*Admin)(nil), "rewardingpb.Admin")
	proto.RegisterType((*Fund)(nil), "rewardingpb.Fund")
	proto.RegisterType((*RewardHistory)(nil), "rewardingpb.RewardHistory")
	proto.RegisterType((*Account)(nil), "rewardingpb.Account")
	proto.RegisterType((*Exempt)(nil), "rewardingpb.Exempt")
}

func init() { proto.RegisterFile("rewarding.proto", fileDescriptor_a5a8d72c965c1359) }

var fileDescriptor_a5a8d72c965c1359 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x89, 0xdd, 0xb6, 0x74, 0x5a, 0xa9, 0x04, 0x0f, 0x8b, 0x07, 0x09, 0xf1, 0xb2, 0x78,
	0xf0, 0xe2, 0x5d, 0xa8, 0xd8, 0xe2, 0x39, 0x07, 0xef, 0xd9, 0x64, 0xa8, 0x8b, 0xd9, 0xcc, 0x92,
	0x4d, 0x50, 0x7f, 0x8e, 0xff, 0x54, 0xba, 0xb1, 0xb2, 0x45, 0xf0, 0x96, 0xf9, 0xf8, 0x1e, 0x79,
	0x3c, 0x58, 0x07, 0x7c, 0xd7, 0xc1, 0x36, 0x7e, 0x7f, 0xd7, 0x05, 0x8a, 0xc4, 0x97, 0xbf, 0xa0,
	0xab, 0xe5, 0x17, 0x83, 0xe9, 0xc6, 0xb6, 0x8d, 0xe7, 0x97, 0x30, 0xd5, 0x87, 0x47, 0xc9, 0x04,
	0xab, 0x56, 0x2a, 0x1f, 0x5c, 0xc0, 0xb2, 0x76, 0x64, 0xde, 0xd4, 0x90, 0x29, 0xcf, 0x04, 0xab,
	0x16, 0x6a, 0x8c, 0x0e, 0x06, 0x76, 0x64, 0x5e, 0x7f, 0x8c, 0x49, 0x36, 0x46, 0x88, 0x3f, 0xc0,
	0x95, 0x4f, 0xed, 0x13, 0x3a, 0xdc, 0xeb, 0x88, 0xfd, 0x8e, 0xc2, 0x76, 0x14, 0x28, 0x04, 0xab,
	0x0a, 0xf5, 0x8f, 0x21, 0x5f, 0xa0, 0xd8, 0x25, 0x6f, 0xb9, 0x84, 0x55, 0xa4, 0xa8, 0xdd, 0xa3,
	0x76, 0xda, 0x1b, 0x1c, 0x8a, 0x2e, 0xd4, 0x09, 0xe3, 0xb7, 0x70, 0x91, 0xbc, 0x71, 0xba, 0x69,
	0xd1, 0x1e, 0xbd, 0x5c, 0xfa, 0x0f, 0x97, 0x6b, 0x38, 0xcf, 0x3f, 0x3c, 0x37, 0x7d, 0xa4, 0xf0,
	0x29, 0x6f, 0x60, 0xbe, 0x31, 0x86, 0x92, 0x8f, 0xbc, 0x84, 0x79, 0x7d, 0x12, 0x3f, 0x9e, 0xf2,
	0x1a, 0x66, 0xdb, 0x0f, 0x6c, 0xbb, 0x98, 0x17, 0xb3, 0xa1, 0x2f, 0x99, 0x98, 0xe4, 0xc5, 0x6c,
	0xe8, 0xeb, 0xd9, 0xb0, 0xf2, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x97, 0xe8, 0x8b,
	0x78, 0x01, 0x00, 0x00,
}
