// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simulator.proto

/*
Package simulator is a generated protocol buffer package.

It is generated from these files:
	simulator.proto

It has these top-level messages:
	Request
	Reply
*/
package simulator

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type Request struct {
	PlayerID    int32  `protobuf:"varint,1,opt,name=playerID" json:"playerID,omitempty"`
	SenderID    int32  `protobuf:"varint,2,opt,name=senderID" json:"senderID,omitempty"`
	MessageType int32  `protobuf:"varint,3,opt,name=messageType" json:"messageType,omitempty"`
	Value       string `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *Request) GetSenderID() int32 {
	if m != nil {
		return m.SenderID
	}
	return 0
}

func (m *Request) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *Request) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// The response message containing the greetings
type Reply struct {
	PlayerID    int32  `protobuf:"varint,1,opt,name=playerID" json:"playerID,omitempty"`
	SenderID    int32  `protobuf:"varint,2,opt,name=senderID" json:"senderID,omitempty"`
	MessageType int32  `protobuf:"varint,3,opt,name=messageType" json:"messageType,omitempty"`
	Value       string `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Reply) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *Reply) GetSenderID() int32 {
	if m != nil {
		return m.SenderID
	}
	return 0
}

func (m *Reply) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *Reply) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "simulator.Request")
	proto.RegisterType((*Reply)(nil), "simulator.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Simulator service

type SimulatorClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (Simulator_PingClient, error)
}

type simulatorClient struct {
	cc *grpc.ClientConn
}

func NewSimulatorClient(cc *grpc.ClientConn) SimulatorClient {
	return &simulatorClient{cc}
}

func (c *simulatorClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (Simulator_PingClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Simulator_serviceDesc.Streams[0], c.cc, "/simulator.Simulator/Ping", opts...)
	if err != nil {
		return nil, err
	}
	x := &simulatorPingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Simulator_PingClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type simulatorPingClient struct {
	grpc.ClientStream
}

func (x *simulatorPingClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Simulator service

type SimulatorServer interface {
	Ping(*Request, Simulator_PingServer) error
}

func RegisterSimulatorServer(s *grpc.Server, srv SimulatorServer) {
	s.RegisterService(&_Simulator_serviceDesc, srv)
}

func _Simulator_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimulatorServer).Ping(m, &simulatorPingServer{stream})
}

type Simulator_PingServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type simulatorPingServer struct {
	grpc.ServerStream
}

func (x *simulatorPingServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

var _Simulator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "simulator.Simulator",
	HandlerType: (*SimulatorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Ping",
			Handler:       _Simulator_Ping_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "simulator.proto",
}

func init() { proto.RegisterFile("simulator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xce, 0xcc, 0x2d,
	0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x55, 0x72, 0xb1, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x71, 0x71, 0x14, 0xe4,
	0x24, 0x56, 0xa6, 0x16, 0x79, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0xc1, 0xf9, 0x20,
	0xb9, 0xe2, 0xd4, 0xbc, 0x14, 0xb0, 0x1c, 0x13, 0x44, 0x0e, 0xc6, 0x17, 0x52, 0xe0, 0xe2, 0xce,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x0d, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x06, 0x4b, 0x23, 0x0b,
	0x09, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0xb0, 0x28, 0x30, 0x6a, 0x70, 0x06,
	0x41, 0x38, 0x4a, 0xe5, 0x5c, 0xac, 0x41, 0xa9, 0x05, 0x39, 0x95, 0xf4, 0xb6, 0xd8, 0xc8, 0x96,
	0x8b, 0x33, 0x18, 0x16, 0x00, 0x42, 0x06, 0x5c, 0x2c, 0x01, 0x99, 0x79, 0xe9, 0x42, 0x42, 0x7a,
	0x88, 0x50, 0x82, 0x86, 0x88, 0x94, 0x00, 0x8a, 0x58, 0x41, 0x4e, 0xa5, 0x12, 0x83, 0x01, 0x63,
	0x12, 0x1b, 0x38, 0x10, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x20, 0xdf, 0x1e, 0xc8, 0x57,
	0x01, 0x00, 0x00,
}
