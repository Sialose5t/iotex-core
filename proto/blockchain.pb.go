// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

/*
Package iproto is a generated protocol buffer package.

It is generated from these files:
	blockchain.proto
	rpc.proto
	utxo.proto

It has these top-level messages:
	TxInputPb
	TxOutputPb
	TxPb
	TransferPb
	VotePb
	ActionPb
	BlockHeaderPb
	BlockPb
	BlockIndex
	PingMsg
	PongMsg
	BlockSync
	BlockContainer
	ViewChangeMsg
	TestPayload
	CreateRawTxRequest
	CreateRawTxReply
	SendTxRequest
	SendTxReply
	UtxoPb
	UtxoEntryPb
	UtxoMapPb
*/
package iproto

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

type ViewChangeMsg_ViewChangeType int32

const (
	ViewChangeMsg_INVALID_VIEW_CHANGE_TYPE ViewChangeMsg_ViewChangeType = 0
	ViewChangeMsg_PROPOSE                  ViewChangeMsg_ViewChangeType = 1
	ViewChangeMsg_PREVOTE                  ViewChangeMsg_ViewChangeType = 2
	ViewChangeMsg_VOTE                     ViewChangeMsg_ViewChangeType = 3
)

var ViewChangeMsg_ViewChangeType_name = map[int32]string{
	0: "INVALID_VIEW_CHANGE_TYPE",
	1: "PROPOSE",
	2: "PREVOTE",
	3: "VOTE",
}
var ViewChangeMsg_ViewChangeType_value = map[string]int32{
	"INVALID_VIEW_CHANGE_TYPE": 0,
	"PROPOSE":                  1,
	"PREVOTE":                  2,
	"VOTE":                     3,
}

func (x ViewChangeMsg_ViewChangeType) String() string {
	return proto.EnumName(ViewChangeMsg_ViewChangeType_name, int32(x))
}
func (ViewChangeMsg_ViewChangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{13, 0}
}

type TxInputPb struct {
	TxHash           []byte `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
	OutIndex         int32  `protobuf:"varint,2,opt,name=outIndex" json:"outIndex,omitempty"`
	UnlockScriptSize uint32 `protobuf:"varint,3,opt,name=unlockScriptSize" json:"unlockScriptSize,omitempty"`
	UnlockScript     []byte `protobuf:"bytes,4,opt,name=unlockScript,proto3" json:"unlockScript,omitempty"`
	Sequence         uint32 `protobuf:"varint,5,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *TxInputPb) Reset()                    { *m = TxInputPb{} }
func (m *TxInputPb) String() string            { return proto.CompactTextString(m) }
func (*TxInputPb) ProtoMessage()               {}
func (*TxInputPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TxInputPb) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *TxInputPb) GetOutIndex() int32 {
	if m != nil {
		return m.OutIndex
	}
	return 0
}

func (m *TxInputPb) GetUnlockScriptSize() uint32 {
	if m != nil {
		return m.UnlockScriptSize
	}
	return 0
}

func (m *TxInputPb) GetUnlockScript() []byte {
	if m != nil {
		return m.UnlockScript
	}
	return nil
}

func (m *TxInputPb) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

// TxOutput stores “coins”. It is indivisible, which means that you cannot reference a part of its value.
// When an output is referenced in a new transaction, it’s spent as a whole. And if its value is greater than required,
// a change is generated and sent back to the sender.
type TxOutputPb struct {
	Value          uint64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	LockScriptSize uint32 `protobuf:"varint,2,opt,name=lockScriptSize" json:"lockScriptSize,omitempty"`
	LockScript     []byte `protobuf:"bytes,3,opt,name=lockScript,proto3" json:"lockScript,omitempty"`
}

func (m *TxOutputPb) Reset()                    { *m = TxOutputPb{} }
func (m *TxOutputPb) String() string            { return proto.CompactTextString(m) }
func (*TxOutputPb) ProtoMessage()               {}
func (*TxOutputPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TxOutputPb) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TxOutputPb) GetLockScriptSize() uint32 {
	if m != nil {
		return m.LockScriptSize
	}
	return 0
}

func (m *TxOutputPb) GetLockScript() []byte {
	if m != nil {
		return m.LockScript
	}
	return nil
}

type TxPb struct {
	Version  uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	LockTime uint32 `protobuf:"varint,2,opt,name=lockTime" json:"lockTime,omitempty"`
	// used by utxo-based model
	TxIn  []*TxInputPb  `protobuf:"bytes,21,rep,name=txIn" json:"txIn,omitempty"`
	TxOut []*TxOutputPb `protobuf:"bytes,22,rep,name=txOut" json:"txOut,omitempty"`
}

func (m *TxPb) Reset()                    { *m = TxPb{} }
func (m *TxPb) String() string            { return proto.CompactTextString(m) }
func (*TxPb) ProtoMessage()               {}
func (*TxPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TxPb) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TxPb) GetLockTime() uint32 {
	if m != nil {
		return m.LockTime
	}
	return 0
}

func (m *TxPb) GetTxIn() []*TxInputPb {
	if m != nil {
		return m.TxIn
	}
	return nil
}

func (m *TxPb) GetTxOut() []*TxOutputPb {
	if m != nil {
		return m.TxOut
	}
	return nil
}

type TransferPb struct {
	Version  uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	LockTime uint32 `protobuf:"varint,2,opt,name=lockTime" json:"lockTime,omitempty"`
	// used by state-based model
	Nonce        uint64 `protobuf:"varint,31,opt,name=nonce" json:"nonce,omitempty"`
	Amount       []byte `protobuf:"bytes,32,opt,name=amount,proto3" json:"amount,omitempty"`
	Sender       string `protobuf:"bytes,33,opt,name=sender" json:"sender,omitempty"`
	Recipient    string `protobuf:"bytes,34,opt,name=recipient" json:"recipient,omitempty"`
	Payload      []byte `protobuf:"bytes,35,opt,name=payload,proto3" json:"payload,omitempty"`
	SenderPubKey []byte `protobuf:"bytes,36,opt,name=senderPubKey,proto3" json:"senderPubKey,omitempty"`
	Signature    []byte `protobuf:"bytes,37,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *TransferPb) Reset()                    { *m = TransferPb{} }
func (m *TransferPb) String() string            { return proto.CompactTextString(m) }
func (*TransferPb) ProtoMessage()               {}
func (*TransferPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TransferPb) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TransferPb) GetLockTime() uint32 {
	if m != nil {
		return m.LockTime
	}
	return 0
}

func (m *TransferPb) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TransferPb) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *TransferPb) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *TransferPb) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *TransferPb) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TransferPb) GetSenderPubKey() []byte {
	if m != nil {
		return m.SenderPubKey
	}
	return nil
}

func (m *TransferPb) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type VotePb struct {
	Timestamp  uint64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	SelfPubkey []byte `protobuf:"bytes,2,opt,name=selfPubkey,proto3" json:"selfPubkey,omitempty"`
	VotePubkey []byte `protobuf:"bytes,3,opt,name=votePubkey,proto3" json:"votePubkey,omitempty"`
	Signature  []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *VotePb) Reset()                    { *m = VotePb{} }
func (m *VotePb) String() string            { return proto.CompactTextString(m) }
func (*VotePb) ProtoMessage()               {}
func (*VotePb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *VotePb) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *VotePb) GetSelfPubkey() []byte {
	if m != nil {
		return m.SelfPubkey
	}
	return nil
}

func (m *VotePb) GetVotePubkey() []byte {
	if m != nil {
		return m.VotePubkey
	}
	return nil
}

func (m *VotePb) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type ActionPb struct {
	// Types that are valid to be assigned to Action:
	//	*ActionPb_Tx
	//	*ActionPb_Transfer
	//	*ActionPb_Vote
	Action isActionPb_Action `protobuf_oneof:"action"`
}

func (m *ActionPb) Reset()                    { *m = ActionPb{} }
func (m *ActionPb) String() string            { return proto.CompactTextString(m) }
func (*ActionPb) ProtoMessage()               {}
func (*ActionPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type isActionPb_Action interface {
	isActionPb_Action()
}

type ActionPb_Tx struct {
	Tx *TxPb `protobuf:"bytes,1,opt,name=tx,oneof"`
}
type ActionPb_Transfer struct {
	Transfer *TransferPb `protobuf:"bytes,2,opt,name=transfer,oneof"`
}
type ActionPb_Vote struct {
	Vote *VotePb `protobuf:"bytes,3,opt,name=vote,oneof"`
}

func (*ActionPb_Tx) isActionPb_Action()       {}
func (*ActionPb_Transfer) isActionPb_Action() {}
func (*ActionPb_Vote) isActionPb_Action()     {}

func (m *ActionPb) GetAction() isActionPb_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *ActionPb) GetTx() *TxPb {
	if x, ok := m.GetAction().(*ActionPb_Tx); ok {
		return x.Tx
	}
	return nil
}

func (m *ActionPb) GetTransfer() *TransferPb {
	if x, ok := m.GetAction().(*ActionPb_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (m *ActionPb) GetVote() *VotePb {
	if x, ok := m.GetAction().(*ActionPb_Vote); ok {
		return x.Vote
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ActionPb) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ActionPb_OneofMarshaler, _ActionPb_OneofUnmarshaler, _ActionPb_OneofSizer, []interface{}{
		(*ActionPb_Tx)(nil),
		(*ActionPb_Transfer)(nil),
		(*ActionPb_Vote)(nil),
	}
}

func _ActionPb_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ActionPb)
	// action
	switch x := m.Action.(type) {
	case *ActionPb_Tx:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tx); err != nil {
			return err
		}
	case *ActionPb_Transfer:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transfer); err != nil {
			return err
		}
	case *ActionPb_Vote:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Vote); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ActionPb.Action has unexpected type %T", x)
	}
	return nil
}

func _ActionPb_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ActionPb)
	switch tag {
	case 1: // action.tx
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TxPb)
		err := b.DecodeMessage(msg)
		m.Action = &ActionPb_Tx{msg}
		return true, err
	case 2: // action.transfer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferPb)
		err := b.DecodeMessage(msg)
		m.Action = &ActionPb_Transfer{msg}
		return true, err
	case 3: // action.vote
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VotePb)
		err := b.DecodeMessage(msg)
		m.Action = &ActionPb_Vote{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ActionPb_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ActionPb)
	// action
	switch x := m.Action.(type) {
	case *ActionPb_Tx:
		s := proto.Size(x.Tx)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionPb_Transfer:
		s := proto.Size(x.Transfer)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionPb_Vote:
		s := proto.Size(x.Vote)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// header of a block
type BlockHeaderPb struct {
	Version       uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	ChainID       uint32 `protobuf:"varint,2,opt,name=chainID" json:"chainID,omitempty"`
	Height        uint64 `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Timestamp     uint64 `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	PrevBlockHash []byte `protobuf:"bytes,5,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	TxRoot        []byte `protobuf:"bytes,6,opt,name=txRoot,proto3" json:"txRoot,omitempty"`
	StateRoot     []byte `protobuf:"bytes,7,opt,name=stateRoot,proto3" json:"stateRoot,omitempty"`
	TrnxNumber    uint32 `protobuf:"varint,8,opt,name=trnxNumber" json:"trnxNumber,omitempty"`
	TrnxDataSize  uint32 `protobuf:"varint,9,opt,name=trnxDataSize" json:"trnxDataSize,omitempty"`
	Signature     []byte `protobuf:"bytes,10,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *BlockHeaderPb) Reset()                    { *m = BlockHeaderPb{} }
func (m *BlockHeaderPb) String() string            { return proto.CompactTextString(m) }
func (*BlockHeaderPb) ProtoMessage()               {}
func (*BlockHeaderPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BlockHeaderPb) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BlockHeaderPb) GetChainID() uint32 {
	if m != nil {
		return m.ChainID
	}
	return 0
}

func (m *BlockHeaderPb) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeaderPb) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeaderPb) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *BlockHeaderPb) GetTxRoot() []byte {
	if m != nil {
		return m.TxRoot
	}
	return nil
}

func (m *BlockHeaderPb) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *BlockHeaderPb) GetTrnxNumber() uint32 {
	if m != nil {
		return m.TrnxNumber
	}
	return 0
}

func (m *BlockHeaderPb) GetTrnxDataSize() uint32 {
	if m != nil {
		return m.TrnxDataSize
	}
	return 0
}

func (m *BlockHeaderPb) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// block consists of header followed by transactions
// hash of current block can be computed from header hence not stored
type BlockPb struct {
	Header  *BlockHeaderPb `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Actions []*ActionPb    `protobuf:"bytes,2,rep,name=actions" json:"actions,omitempty"`
}

func (m *BlockPb) Reset()                    { *m = BlockPb{} }
func (m *BlockPb) String() string            { return proto.CompactTextString(m) }
func (*BlockPb) ProtoMessage()               {}
func (*BlockPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BlockPb) GetHeader() *BlockHeaderPb {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BlockPb) GetActions() []*ActionPb {
	if m != nil {
		return m.Actions
	}
	return nil
}

// index of block raw data file
type BlockIndex struct {
	Start  uint64   `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End    uint64   `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
	Offset []uint32 `protobuf:"varint,3,rep,packed,name=offset" json:"offset,omitempty"`
}

func (m *BlockIndex) Reset()                    { *m = BlockIndex{} }
func (m *BlockIndex) String() string            { return proto.CompactTextString(m) }
func (*BlockIndex) ProtoMessage()               {}
func (*BlockIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BlockIndex) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *BlockIndex) GetEnd() uint64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *BlockIndex) GetOffset() []uint32 {
	if m != nil {
		return m.Offset
	}
	return nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////
// BELOW ARE DEFINITIONS FOR ON-WIRE MESSAGES!
// //////////////////////////////////////////////////////////////////////////////////////////////////
type PingMsg struct {
	Nonce uint64 `protobuf:"varint,1,opt,name=nonce" json:"nonce,omitempty"`
}

func (m *PingMsg) Reset()                    { *m = PingMsg{} }
func (m *PingMsg) String() string            { return proto.CompactTextString(m) }
func (*PingMsg) ProtoMessage()               {}
func (*PingMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PingMsg) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type PongMsg struct {
	AckNonce uint64 `protobuf:"varint,1,opt,name=ack_nonce,json=ackNonce" json:"ack_nonce,omitempty"`
}

func (m *PongMsg) Reset()                    { *m = PongMsg{} }
func (m *PongMsg) String() string            { return proto.CompactTextString(m) }
func (*PongMsg) ProtoMessage()               {}
func (*PongMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PongMsg) GetAckNonce() uint64 {
	if m != nil {
		return m.AckNonce
	}
	return 0
}

type BlockSync struct {
	Start uint64 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	End   uint64 `protobuf:"varint,3,opt,name=end" json:"end,omitempty"`
}

func (m *BlockSync) Reset()                    { *m = BlockSync{} }
func (m *BlockSync) String() string            { return proto.CompactTextString(m) }
func (*BlockSync) ProtoMessage()               {}
func (*BlockSync) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *BlockSync) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *BlockSync) GetEnd() uint64 {
	if m != nil {
		return m.End
	}
	return 0
}

// block container
// used to send old/existing blocks in block sync
type BlockContainer struct {
	Block *BlockPb `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
}

func (m *BlockContainer) Reset()                    { *m = BlockContainer{} }
func (m *BlockContainer) String() string            { return proto.CompactTextString(m) }
func (*BlockContainer) ProtoMessage()               {}
func (*BlockContainer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *BlockContainer) GetBlock() *BlockPb {
	if m != nil {
		return m.Block
	}
	return nil
}

type ViewChangeMsg struct {
	Vctype     ViewChangeMsg_ViewChangeType `protobuf:"varint,1,opt,name=vctype,enum=iproto.ViewChangeMsg_ViewChangeType" json:"vctype,omitempty"`
	Block      *BlockPb                     `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
	BlockHash  []byte                       `protobuf:"bytes,3,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	SenderAddr string                       `protobuf:"bytes,4,opt,name=senderAddr" json:"senderAddr,omitempty"`
}

func (m *ViewChangeMsg) Reset()                    { *m = ViewChangeMsg{} }
func (m *ViewChangeMsg) String() string            { return proto.CompactTextString(m) }
func (*ViewChangeMsg) ProtoMessage()               {}
func (*ViewChangeMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ViewChangeMsg) GetVctype() ViewChangeMsg_ViewChangeType {
	if m != nil {
		return m.Vctype
	}
	return ViewChangeMsg_INVALID_VIEW_CHANGE_TYPE
}

func (m *ViewChangeMsg) GetBlock() *BlockPb {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *ViewChangeMsg) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *ViewChangeMsg) GetSenderAddr() string {
	if m != nil {
		return m.SenderAddr
	}
	return ""
}

// //////////////////////////////////////////////////////////////////////////////////////////////////
// BELOW ARE DEFINITIONS FOR TEST-ONLY MESSAGES!
// //////////////////////////////////////////////////////////////////////////////////////////////////
type TestPayload struct {
	MsgBody []byte `protobuf:"bytes,1,opt,name=msg_body,json=msgBody,proto3" json:"msg_body,omitempty"`
}

func (m *TestPayload) Reset()                    { *m = TestPayload{} }
func (m *TestPayload) String() string            { return proto.CompactTextString(m) }
func (*TestPayload) ProtoMessage()               {}
func (*TestPayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *TestPayload) GetMsgBody() []byte {
	if m != nil {
		return m.MsgBody
	}
	return nil
}

func init() {
	proto.RegisterType((*TxInputPb)(nil), "iproto.TxInputPb")
	proto.RegisterType((*TxOutputPb)(nil), "iproto.TxOutputPb")
	proto.RegisterType((*TxPb)(nil), "iproto.TxPb")
	proto.RegisterType((*TransferPb)(nil), "iproto.TransferPb")
	proto.RegisterType((*VotePb)(nil), "iproto.VotePb")
	proto.RegisterType((*ActionPb)(nil), "iproto.ActionPb")
	proto.RegisterType((*BlockHeaderPb)(nil), "iproto.BlockHeaderPb")
	proto.RegisterType((*BlockPb)(nil), "iproto.BlockPb")
	proto.RegisterType((*BlockIndex)(nil), "iproto.BlockIndex")
	proto.RegisterType((*PingMsg)(nil), "iproto.PingMsg")
	proto.RegisterType((*PongMsg)(nil), "iproto.PongMsg")
	proto.RegisterType((*BlockSync)(nil), "iproto.BlockSync")
	proto.RegisterType((*BlockContainer)(nil), "iproto.BlockContainer")
	proto.RegisterType((*ViewChangeMsg)(nil), "iproto.ViewChangeMsg")
	proto.RegisterType((*TestPayload)(nil), "iproto.TestPayload")
	proto.RegisterEnum("iproto.ViewChangeMsg_ViewChangeType", ViewChangeMsg_ViewChangeType_name, ViewChangeMsg_ViewChangeType_value)
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 933 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5d, 0x6f, 0xe3, 0x44,
	0x14, 0xad, 0x1d, 0xe7, 0xeb, 0xb6, 0x29, 0x61, 0xd4, 0x5d, 0x0d, 0xb0, 0xda, 0x0d, 0xa6, 0x5d,
	0x45, 0x2b, 0x51, 0xa1, 0xee, 0x03, 0x2f, 0xbc, 0xf4, 0x23, 0x22, 0x11, 0x4b, 0x6b, 0x4d, 0xa3,
	0x20, 0x9e, 0xa2, 0xb1, 0x3d, 0x4d, 0x4c, 0x9b, 0x71, 0xb0, 0xc7, 0x21, 0xe1, 0x11, 0x21, 0x9e,
	0xe0, 0x97, 0xf0, 0x03, 0xf8, 0x7b, 0x68, 0xee, 0xd8, 0xb1, 0x13, 0x10, 0x48, 0xfb, 0xd4, 0x9e,
	0x33, 0xd7, 0x33, 0x67, 0xce, 0xb9, 0x73, 0x03, 0x5d, 0xff, 0x29, 0x0e, 0x1e, 0x83, 0x39, 0x8f,
	0xe4, 0xf9, 0x32, 0x89, 0x55, 0x4c, 0x1a, 0x11, 0xfe, 0x75, 0xff, 0xb4, 0xa0, 0x3d, 0x5e, 0x8f,
	0xe4, 0x32, 0x53, 0x9e, 0x4f, 0x9e, 0x43, 0x43, 0xad, 0x87, 0x3c, 0x9d, 0x53, 0xab, 0x67, 0xf5,
	0x8f, 0x58, 0x8e, 0xc8, 0xc7, 0xd0, 0x8a, 0x33, 0x35, 0x92, 0xa1, 0x58, 0x53, 0xbb, 0x67, 0xf5,
	0xeb, 0x6c, 0x8b, 0xc9, 0x1b, 0xe8, 0x66, 0x52, 0x6f, 0x7f, 0x1f, 0x24, 0xd1, 0x52, 0xdd, 0x47,
	0x3f, 0x0b, 0x5a, 0xeb, 0x59, 0xfd, 0x0e, 0xfb, 0x07, 0x4f, 0x5c, 0x38, 0xaa, 0x72, 0xd4, 0xc1,
	0x53, 0x76, 0x38, 0x7d, 0x56, 0x2a, 0x7e, 0xcc, 0x84, 0x0c, 0x04, 0xad, 0xe3, 0x3e, 0x5b, 0xec,
	0xfe, 0x00, 0x30, 0x5e, 0xdf, 0x65, 0xca, 0xa8, 0x3d, 0x81, 0xfa, 0x8a, 0x3f, 0x65, 0x02, 0xc5,
	0x3a, 0xcc, 0x00, 0xf2, 0x1a, 0x8e, 0xf7, 0xd4, 0xd8, 0xb8, 0xcb, 0x1e, 0x4b, 0x5e, 0x02, 0x54,
	0x94, 0xd4, 0x50, 0x49, 0x85, 0x71, 0xff, 0xb0, 0xc0, 0x19, 0xaf, 0x3d, 0x9f, 0x50, 0x68, 0xae,
	0x44, 0x92, 0x46, 0xb1, 0xc4, 0x83, 0x3a, 0xac, 0x80, 0x5a, 0xaa, 0xfe, 0x60, 0x1c, 0x2d, 0x8a,
	0x43, 0xb6, 0x98, 0x9c, 0x81, 0xa3, 0xd6, 0x23, 0x49, 0x9f, 0xf5, 0x6a, 0xfd, 0xc3, 0x8b, 0x0f,
	0xcf, 0x8d, 0xdf, 0xe7, 0x5b, 0xaf, 0x19, 0x2e, 0x93, 0x3e, 0xd4, 0x95, 0xbe, 0x11, 0x7d, 0x8e,
	0x75, 0xa4, 0xac, 0x2b, 0xae, 0xc9, 0x4c, 0x81, 0xfb, 0x8b, 0x0d, 0x30, 0x4e, 0xb8, 0x4c, 0x1f,
	0x44, 0xf2, 0xde, 0xaa, 0x4e, 0xa0, 0x2e, 0x63, 0xed, 0xec, 0x2b, 0x63, 0x19, 0x02, 0x1d, 0x3b,
	0x5f, 0xc4, 0x99, 0x54, 0xb4, 0x67, 0x62, 0x37, 0x48, 0xf3, 0xa9, 0x90, 0xa1, 0x48, 0xe8, 0xa7,
	0x3d, 0xab, 0xdf, 0x66, 0x39, 0x22, 0x2f, 0xa0, 0x9d, 0x88, 0x20, 0x5a, 0x46, 0x42, 0x2a, 0xea,
	0xe2, 0x52, 0x49, 0x68, 0x65, 0x4b, 0xbe, 0x79, 0x8a, 0x79, 0x48, 0x3f, 0xc3, 0xed, 0x0a, 0xa8,
	0xe3, 0x37, 0x3b, 0x78, 0x99, 0xff, 0x8d, 0xd8, 0xd0, 0x53, 0x13, 0x7f, 0x95, 0xd3, 0x7b, 0xa7,
	0xd1, 0x4c, 0x72, 0x95, 0x25, 0x82, 0x9e, 0x61, 0x41, 0x49, 0xb8, 0xbf, 0x5a, 0xd0, 0x98, 0xc4,
	0x4a, 0x78, 0xbe, 0x2e, 0x54, 0xd1, 0x42, 0xa4, 0x8a, 0x2f, 0x96, 0x79, 0x07, 0x94, 0x84, 0x4e,
	0x37, 0x15, 0x4f, 0x0f, 0x5e, 0xe6, 0x3f, 0x8a, 0x0d, 0xda, 0x70, 0xc4, 0x2a, 0x8c, 0x5e, 0x5f,
	0xe9, 0x7d, 0xcc, 0x7a, 0x9e, 0x7e, 0xc9, 0xec, 0xca, 0x70, 0xf6, 0x65, 0xfc, 0x6e, 0x41, 0xeb,
	0x32, 0x50, 0x51, 0x2c, 0x3d, 0x9f, 0xbc, 0x04, 0x5b, 0xad, 0x51, 0xc1, 0xe1, 0xc5, 0x51, 0x99,
	0x9f, 0xe7, 0x0f, 0x0f, 0x98, 0xad, 0xd6, 0xe4, 0x0b, 0x68, 0xa9, 0x3c, 0x37, 0x14, 0x52, 0x4d,
	0x79, 0x9b, 0xe7, 0xf0, 0x80, 0x6d, 0xab, 0xc8, 0x29, 0x38, 0x5a, 0x0a, 0xca, 0x3a, 0xbc, 0x38,
	0x2e, 0xaa, 0xcd, 0xc5, 0x87, 0x07, 0x0c, 0x57, 0xaf, 0x5a, 0xd0, 0xe0, 0xa8, 0xc1, 0xfd, 0xcb,
	0x86, 0xce, 0x95, 0xce, 0x78, 0x28, 0x78, 0xf8, 0x3f, 0xdd, 0x41, 0xa1, 0x89, 0x73, 0x60, 0x74,
	0x93, 0x37, 0x47, 0x01, 0x75, 0xda, 0x73, 0x11, 0xcd, 0xe6, 0xe6, 0x31, 0x38, 0x2c, 0x47, 0xbb,
	0x46, 0x3b, 0xfb, 0x46, 0x9f, 0x42, 0x67, 0x99, 0x88, 0x95, 0x39, 0x5e, 0x4f, 0x8e, 0x3a, 0x9a,
	0xb5, 0x4b, 0x9a, 0xc1, 0xc2, 0xe2, 0x58, 0xd1, 0x46, 0x31, 0x58, 0x34, 0x42, 0x9b, 0x15, 0x57,
	0x02, 0x97, 0x9a, 0xb9, 0xcd, 0x05, 0xa1, 0x43, 0x52, 0x89, 0x5c, 0xdf, 0x66, 0x0b, 0x5f, 0x24,
	0xb4, 0x85, 0x72, 0x2b, 0x8c, 0xee, 0x27, 0x8d, 0x6e, 0xb8, 0xe2, 0xf8, 0xd0, 0xdb, 0x58, 0xb1,
	0xc3, 0xed, 0x06, 0x09, 0xfb, 0x41, 0x86, 0xd0, 0x44, 0x91, 0x9e, 0x4f, 0x3e, 0xd7, 0xd7, 0xd7,
	0xf6, 0xe5, 0x51, 0x3e, 0x2b, 0x6c, 0xdf, 0x71, 0x96, 0xe5, 0x45, 0xe4, 0x0d, 0x34, 0x8d, 0xfb,
	0x29, 0xb5, 0xf1, 0xe9, 0x76, 0x8b, 0xfa, 0xa2, 0x31, 0x58, 0x51, 0xe0, 0xbe, 0x03, 0xc0, 0x4d,
	0xcc, 0xc0, 0x3c, 0x81, 0x7a, 0xaa, 0x78, 0xa2, 0x8a, 0xb1, 0x85, 0x80, 0x74, 0xa1, 0x26, 0x64,
	0x88, 0x99, 0x38, 0x4c, 0xff, 0xab, 0x3d, 0x8b, 0x1f, 0x1e, 0x52, 0xa1, 0xf3, 0xa8, 0xf5, 0x3b,
	0x2c, 0x47, 0xee, 0x2b, 0x68, 0x7a, 0x91, 0x9c, 0x7d, 0x9b, 0xce, 0xca, 0xe7, 0x6c, 0x55, 0x9e,
	0xb3, 0xfb, 0x1a, 0x9a, 0x5e, 0x6c, 0x0a, 0x3e, 0x81, 0x36, 0x0f, 0x1e, 0xa7, 0xd5, 0xa2, 0x16,
	0x0f, 0x1e, 0x6f, 0xb1, 0xee, 0x2d, 0xb4, 0x51, 0xd6, 0xfd, 0x46, 0x06, 0xa5, 0x2a, 0xfb, 0x5f,
	0x54, 0xd5, 0xb6, 0xaa, 0xdc, 0x2f, 0xe1, 0x18, 0x3f, 0xba, 0x8e, 0xa5, 0xe2, 0x91, 0x14, 0x09,
	0x39, 0x83, 0x3a, 0xfe, 0xbc, 0xe4, 0xbe, 0x7d, 0xb0, 0xe3, 0x9b, 0x9e, 0x5f, 0xb8, 0xea, 0xfe,
	0x66, 0x43, 0x67, 0x12, 0x89, 0x9f, 0xae, 0xe7, 0x5c, 0xce, 0x84, 0x16, 0xf7, 0x15, 0x34, 0x56,
	0x81, 0xda, 0x2c, 0x8d, 0xb2, 0xe3, 0x8b, 0xd3, 0x6d, 0xa3, 0x57, 0xcb, 0x2a, 0x68, 0xbc, 0x59,
	0x0a, 0x96, 0x7f, 0x53, 0x1e, 0x6b, 0xff, 0xd7, 0xb1, 0x3a, 0x7f, 0x7f, 0xdb, 0x9b, 0xe6, 0x9d,
	0x97, 0x84, 0x19, 0x13, 0x7a, 0xfa, 0x5c, 0x86, 0x61, 0x82, 0xcd, 0xdd, 0x66, 0x15, 0xc6, 0x65,
	0x70, 0xbc, 0x7b, 0x3c, 0x79, 0x01, 0x74, 0x74, 0x3b, 0xb9, 0x7c, 0x37, 0xba, 0x99, 0x4e, 0x46,
	0x83, 0xef, 0xa6, 0xd7, 0xc3, 0xcb, 0xdb, 0xaf, 0x07, 0xd3, 0xf1, 0xf7, 0xde, 0xa0, 0x7b, 0x40,
	0x0e, 0xa1, 0xe9, 0xb1, 0x3b, 0xef, 0xee, 0x7e, 0xd0, 0xb5, 0x0c, 0x18, 0x4c, 0xee, 0xc6, 0x83,
	0xae, 0x4d, 0x5a, 0xe0, 0xe0, 0x7f, 0x35, 0xb7, 0x0f, 0x87, 0x63, 0x91, 0x2a, 0x2f, 0x1f, 0x8a,
	0x1f, 0x41, 0x6b, 0x91, 0xce, 0xa6, 0x7e, 0x1c, 0x6e, 0xf2, 0x5f, 0xdd, 0xe6, 0x22, 0x9d, 0x5d,
	0xc5, 0xe1, 0xc6, 0x6f, 0xe0, 0x8d, 0xde, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x07, 0x73, 0xb0,
	0x12, 0xbf, 0x07, 0x00, 0x00,
}
