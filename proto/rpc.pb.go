// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package iproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateRawTransferRequest struct {
	Sender    string `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Recipient string `protobuf:"bytes,2,opt,name=recipient" json:"recipient,omitempty"`
	Amount    []byte `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Nonce     uint64 `protobuf:"varint,4,opt,name=nonce" json:"nonce,omitempty"`
	Data      []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *CreateRawTransferRequest) Reset()                    { *m = CreateRawTransferRequest{} }
func (m *CreateRawTransferRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRawTransferRequest) ProtoMessage()               {}
func (*CreateRawTransferRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateRawTransferRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *CreateRawTransferRequest) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *CreateRawTransferRequest) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *CreateRawTransferRequest) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *CreateRawTransferRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateRawTransferReply struct {
	SerializedTx []byte `protobuf:"bytes,1,opt,name=serialized_tx,json=serializedTx,proto3" json:"serialized_tx,omitempty"`
}

func (m *CreateRawTransferReply) Reset()                    { *m = CreateRawTransferReply{} }
func (m *CreateRawTransferReply) String() string            { return proto.CompactTextString(m) }
func (*CreateRawTransferReply) ProtoMessage()               {}
func (*CreateRawTransferReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CreateRawTransferReply) GetSerializedTx() []byte {
	if m != nil {
		return m.SerializedTx
	}
	return nil
}

type SendTransferRequest struct {
	SerializedTx []byte `protobuf:"bytes,1,opt,name=serialized_tx,json=serializedTx,proto3" json:"serialized_tx,omitempty"`
}

func (m *SendTransferRequest) Reset()                    { *m = SendTransferRequest{} }
func (m *SendTransferRequest) String() string            { return proto.CompactTextString(m) }
func (*SendTransferRequest) ProtoMessage()               {}
func (*SendTransferRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *SendTransferRequest) GetSerializedTx() []byte {
	if m != nil {
		return m.SerializedTx
	}
	return nil
}

type SendTransferReply struct {
}

func (m *SendTransferReply) Reset()                    { *m = SendTransferReply{} }
func (m *SendTransferReply) String() string            { return proto.CompactTextString(m) }
func (*SendTransferReply) ProtoMessage()               {}
func (*SendTransferReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func init() {
	proto.RegisterType((*CreateRawTransferRequest)(nil), "iproto.CreateRawTransferRequest")
	proto.RegisterType((*CreateRawTransferReply)(nil), "iproto.CreateRawTransferReply")
	proto.RegisterType((*SendTransferRequest)(nil), "iproto.SendTransferRequest")
	proto.RegisterType((*SendTransferReply)(nil), "iproto.SendTransferReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChainService service

type ChainServiceClient interface {
	CreateRawTx(ctx context.Context, in *CreateRawTransferRequest, opts ...grpc.CallOption) (*CreateRawTransferReply, error)
	SendTx(ctx context.Context, in *SendTransferRequest, opts ...grpc.CallOption) (*SendTransferReply, error)
}

type chainServiceClient struct {
	cc *grpc.ClientConn
}

func NewChainServiceClient(cc *grpc.ClientConn) ChainServiceClient {
	return &chainServiceClient{cc}
}

func (c *chainServiceClient) CreateRawTx(ctx context.Context, in *CreateRawTransferRequest, opts ...grpc.CallOption) (*CreateRawTransferReply, error) {
	out := new(CreateRawTransferReply)
	err := grpc.Invoke(ctx, "/iproto.ChainService/CreateRawTx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) SendTx(ctx context.Context, in *SendTransferRequest, opts ...grpc.CallOption) (*SendTransferReply, error) {
	out := new(SendTransferReply)
	err := grpc.Invoke(ctx, "/iproto.ChainService/SendTx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChainService service

type ChainServiceServer interface {
	CreateRawTx(context.Context, *CreateRawTransferRequest) (*CreateRawTransferReply, error)
	SendTx(context.Context, *SendTransferRequest) (*SendTransferReply, error)
}

func RegisterChainServiceServer(s *grpc.Server, srv ChainServiceServer) {
	s.RegisterService(&_ChainService_serviceDesc, srv)
}

func _ChainService_CreateRawTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRawTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).CreateRawTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iproto.ChainService/CreateRawTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).CreateRawTx(ctx, req.(*CreateRawTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_SendTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).SendTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iproto.ChainService/SendTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).SendTx(ctx, req.(*SendTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iproto.ChainService",
	HandlerType: (*ChainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRawTx",
			Handler:    _ChainService_CreateRawTx_Handler,
		},
		{
			MethodName: "SendTx",
			Handler:    _ChainService_SendTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x50, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0xc4, 0x90, 0x46, 0xca, 0x12, 0x0e, 0x6c, 0x51, 0x65, 0x0a, 0x42, 0x51, 0xb8, 0xe4, 0x94,
	0x03, 0xdc, 0x90, 0xb8, 0xd0, 0x17, 0xe0, 0xf6, 0x8e, 0x4c, 0xb2, 0x08, 0x4b, 0xc1, 0x31, 0x8e,
	0x0b, 0x09, 0xcf, 0xe0, 0x0f, 0xfc, 0x13, 0xe1, 0x06, 0x45, 0xa0, 0x80, 0x7a, 0xb2, 0x67, 0x67,
	0xc6, 0x9e, 0x1d, 0x88, 0xac, 0x29, 0x72, 0x63, 0x6b, 0x57, 0x63, 0xa8, 0xfc, 0x99, 0xbe, 0x33,
	0xe0, 0x0b, 0x4b, 0xd2, 0x91, 0x90, 0xaf, 0x2b, 0x2b, 0x75, 0xf3, 0x40, 0x56, 0xd0, 0xf3, 0x9a,
	0x1a, 0x87, 0x33, 0x08, 0x1b, 0xd2, 0x25, 0x59, 0xce, 0x12, 0x96, 0x45, 0xa2, 0x47, 0x78, 0x0a,
	0x91, 0xa5, 0x42, 0x19, 0x45, 0xda, 0xf1, 0x5d, 0x4f, 0x0d, 0x83, 0x2f, 0x97, 0x7c, 0xaa, 0xd7,
	0xda, 0xf1, 0xbd, 0x84, 0x65, 0xb1, 0xe8, 0x11, 0x1e, 0xc1, 0x44, 0xd7, 0xba, 0x20, 0x1e, 0x24,
	0x2c, 0x0b, 0xc4, 0x06, 0x20, 0x42, 0x50, 0x4a, 0x27, 0xf9, 0xc4, 0x6b, 0xfd, 0x3d, 0xbd, 0x86,
	0xd9, 0x48, 0x26, 0x53, 0x75, 0x78, 0x0e, 0x07, 0x0d, 0x59, 0x25, 0x2b, 0xf5, 0x46, 0xe5, 0x9d,
	0x6b, 0x7d, 0xb0, 0x58, 0xc4, 0xc3, 0x70, 0xd5, 0xa6, 0x57, 0x30, 0x5d, 0x92, 0x2e, 0x7f, 0x6f,
	0xb3, 0x95, 0x77, 0x0a, 0x87, 0x3f, 0xbd, 0xa6, 0xea, 0x2e, 0x3e, 0x18, 0xc4, 0x8b, 0x47, 0xa9,
	0xf4, 0x92, 0xec, 0x8b, 0x2a, 0x08, 0x6f, 0x61, 0x7f, 0x08, 0xd8, 0x62, 0x92, 0x6f, 0xda, 0xcc,
	0xff, 0x6a, 0x72, 0x7e, 0xf6, 0x8f, 0xc2, 0x54, 0x5d, 0xba, 0x83, 0x37, 0x10, 0xfa, 0x8f, 0x5b,
	0x3c, 0xf9, 0xd6, 0x8e, 0x2c, 0x31, 0x3f, 0x1e, 0x27, 0xfd, 0x1b, 0xf7, 0xa1, 0xa7, 0x2e, 0x3f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x02, 0x82, 0x35, 0xe8, 0x01, 0x00, 0x00,
}
