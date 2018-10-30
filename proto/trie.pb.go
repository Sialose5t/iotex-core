// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trie.proto

package iproto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BranchNodePb struct {
	Index                uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Path                 []byte   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BranchNodePb) Reset()         { *m = BranchNodePb{} }
func (m *BranchNodePb) String() string { return proto.CompactTextString(m) }
func (*BranchNodePb) ProtoMessage()    {}
func (*BranchNodePb) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a69962149106130, []int{0}
}

func (m *BranchNodePb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BranchNodePb.Unmarshal(m, b)
}
func (m *BranchNodePb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BranchNodePb.Marshal(b, m, deterministic)
}
func (m *BranchNodePb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchNodePb.Merge(m, src)
}
func (m *BranchNodePb) XXX_Size() int {
	return xxx_messageInfo_BranchNodePb.Size(m)
}
func (m *BranchNodePb) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchNodePb.DiscardUnknown(m)
}

var xxx_messageInfo_BranchNodePb proto.InternalMessageInfo

func (m *BranchNodePb) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BranchNodePb) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

type BranchPb struct {
	Branches             []*BranchNodePb `protobuf:"bytes,1,rep,name=branches,proto3" json:"branches,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BranchPb) Reset()         { *m = BranchPb{} }
func (m *BranchPb) String() string { return proto.CompactTextString(m) }
func (*BranchPb) ProtoMessage()    {}
func (*BranchPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a69962149106130, []int{1}
}

func (m *BranchPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BranchPb.Unmarshal(m, b)
}
func (m *BranchPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BranchPb.Marshal(b, m, deterministic)
}
func (m *BranchPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchPb.Merge(m, src)
}
func (m *BranchPb) XXX_Size() int {
	return xxx_messageInfo_BranchPb.Size(m)
}
func (m *BranchPb) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchPb.DiscardUnknown(m)
}

var xxx_messageInfo_BranchPb proto.InternalMessageInfo

func (m *BranchPb) GetBranches() []*BranchNodePb {
	if m != nil {
		return m.Branches
	}
	return nil
}

type LeafPb struct {
	Ext                  uint32   `protobuf:"varint,1,opt,name=ext,proto3" json:"ext,omitempty"`
	Path                 []byte   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeafPb) Reset()         { *m = LeafPb{} }
func (m *LeafPb) String() string { return proto.CompactTextString(m) }
func (*LeafPb) ProtoMessage()    {}
func (*LeafPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a69962149106130, []int{2}
}

func (m *LeafPb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeafPb.Unmarshal(m, b)
}
func (m *LeafPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeafPb.Marshal(b, m, deterministic)
}
func (m *LeafPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeafPb.Merge(m, src)
}
func (m *LeafPb) XXX_Size() int {
	return xxx_messageInfo_LeafPb.Size(m)
}
func (m *LeafPb) XXX_DiscardUnknown() {
	xxx_messageInfo_LeafPb.DiscardUnknown(m)
}

var xxx_messageInfo_LeafPb proto.InternalMessageInfo

func (m *LeafPb) GetExt() uint32 {
	if m != nil {
		return m.Ext
	}
	return 0
}

func (m *LeafPb) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *LeafPb) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type NodePb struct {
	// Types that are valid to be assigned to Node:
	//	*NodePb_Branch
	//	*NodePb_Leaf
	Node                 isNodePb_Node `protobuf_oneof:"node"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodePb) Reset()         { *m = NodePb{} }
func (m *NodePb) String() string { return proto.CompactTextString(m) }
func (*NodePb) ProtoMessage()    {}
func (*NodePb) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a69962149106130, []int{3}
}

func (m *NodePb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodePb.Unmarshal(m, b)
}
func (m *NodePb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodePb.Marshal(b, m, deterministic)
}
func (m *NodePb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePb.Merge(m, src)
}
func (m *NodePb) XXX_Size() int {
	return xxx_messageInfo_NodePb.Size(m)
}
func (m *NodePb) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePb.DiscardUnknown(m)
}

var xxx_messageInfo_NodePb proto.InternalMessageInfo

type isNodePb_Node interface {
	isNodePb_Node()
}

type NodePb_Branch struct {
	Branch *BranchPb `protobuf:"bytes,2,opt,name=branch,proto3,oneof"`
}

type NodePb_Leaf struct {
	Leaf *LeafPb `protobuf:"bytes,3,opt,name=leaf,proto3,oneof"`
}

func (*NodePb_Branch) isNodePb_Node() {}

func (*NodePb_Leaf) isNodePb_Node() {}

func (m *NodePb) GetNode() isNodePb_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *NodePb) GetBranch() *BranchPb {
	if x, ok := m.GetNode().(*NodePb_Branch); ok {
		return x.Branch
	}
	return nil
}

func (m *NodePb) GetLeaf() *LeafPb {
	if x, ok := m.GetNode().(*NodePb_Leaf); ok {
		return x.Leaf
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NodePb) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NodePb_OneofMarshaler, _NodePb_OneofUnmarshaler, _NodePb_OneofSizer, []interface{}{
		(*NodePb_Branch)(nil),
		(*NodePb_Leaf)(nil),
	}
}

func _NodePb_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NodePb)
	// node
	switch x := m.Node.(type) {
	case *NodePb_Branch:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Branch); err != nil {
			return err
		}
	case *NodePb_Leaf:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Leaf); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NodePb.Node has unexpected type %T", x)
	}
	return nil
}

func _NodePb_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NodePb)
	switch tag {
	case 2: // node.branch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BranchPb)
		err := b.DecodeMessage(msg)
		m.Node = &NodePb_Branch{msg}
		return true, err
	case 3: // node.leaf
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LeafPb)
		err := b.DecodeMessage(msg)
		m.Node = &NodePb_Leaf{msg}
		return true, err
	default:
		return false, nil
	}
}

func _NodePb_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NodePb)
	// node
	switch x := m.Node.(type) {
	case *NodePb_Branch:
		s := proto.Size(x.Branch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *NodePb_Leaf:
		s := proto.Size(x.Leaf)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*BranchNodePb)(nil), "iproto.branchNodePb")
	proto.RegisterType((*BranchPb)(nil), "iproto.branchPb")
	proto.RegisterType((*LeafPb)(nil), "iproto.leafPb")
	proto.RegisterType((*NodePb)(nil), "iproto.nodePb")
}

func init() { proto.RegisterFile("trie.proto", fileDescriptor_4a69962149106130) }

var fileDescriptor_4a69962149106130 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x29, 0xca, 0x4c,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0x04, 0xd3, 0x4a, 0x16, 0x5c, 0x3c, 0x49,
	0x45, 0x89, 0x79, 0xc9, 0x19, 0x7e, 0xf9, 0x29, 0xa9, 0x01, 0x49, 0x42, 0x22, 0x5c, 0xac, 0x99,
	0x79, 0x29, 0xa9, 0x15, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x10, 0x8e, 0x90, 0x10, 0x17,
	0x4b, 0x41, 0x62, 0x49, 0x86, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98, 0xad, 0x64, 0xc3,
	0xc5, 0x01, 0xd1, 0x19, 0x90, 0x24, 0x64, 0x00, 0x63, 0xa7, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6b,
	0x70, 0x1b, 0x89, 0xe8, 0x41, 0x2c, 0xd0, 0x43, 0x36, 0x3d, 0x08, 0xae, 0x4a, 0xc9, 0x85, 0x8b,
	0x2d, 0x27, 0x35, 0x31, 0x2d, 0x20, 0x49, 0x48, 0x80, 0x8b, 0x39, 0xb5, 0xa2, 0x04, 0x6a, 0x1f,
	0x88, 0x89, 0xcd, 0x36, 0x90, 0xbb, 0xca, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0xc1, 0x82, 0x10,
	0x8e, 0x52, 0x12, 0x17, 0x5b, 0x1e, 0xc4, 0xdd, 0x5a, 0x5c, 0x6c, 0x10, 0xb3, 0xc1, 0xba, 0xb8,
	0x8d, 0x04, 0x50, 0xed, 0x0f, 0x48, 0xf2, 0x60, 0x08, 0x82, 0xaa, 0x10, 0x52, 0xe1, 0x62, 0x01,
	0xd9, 0x0d, 0x36, 0x8a, 0xdb, 0x88, 0x0f, 0xa6, 0x12, 0xe2, 0x1e, 0x0f, 0x86, 0x20, 0xb0, 0xac,
	0x13, 0x1b, 0x17, 0x0b, 0xc8, 0xec, 0x24, 0x36, 0xb0, 0xac, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xf6, 0x80, 0x05, 0xac, 0x3e, 0x01, 0x00, 0x00,
}
