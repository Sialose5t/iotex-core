// Code generated by MockGen. DO NOT EDIT.
// Source: ./blockchain/blockchain.go

// Package mock_blockchain is a generated GoMock package.
package mock_blockchain

import (
	gomock "github.com/golang/mock/gomock"
	blockchain "github.com/iotexproject/iotex-core/blockchain"
	common "github.com/iotexproject/iotex-core/common"
	iotxaddress "github.com/iotexproject/iotex-core/iotxaddress"
	big "math/big"
	reflect "reflect"
)

// MockBlockchain is a mock of Blockchain interface
type MockBlockchain struct {
	ctrl     *gomock.Controller
	recorder *MockBlockchainMockRecorder
}

// MockBlockchainMockRecorder is the mock recorder for MockBlockchain
type MockBlockchainMockRecorder struct {
	mock *MockBlockchain
}

// NewMockBlockchain creates a new mock instance
func NewMockBlockchain(ctrl *gomock.Controller) *MockBlockchain {
	mock := &MockBlockchain{ctrl: ctrl}
	mock.recorder = &MockBlockchainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockchain) EXPECT() *MockBlockchainMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockBlockchain) Init() error {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockBlockchainMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockBlockchain)(nil).Init))
}

// Start mocks base method
func (m *MockBlockchain) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockBlockchainMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBlockchain)(nil).Start))
}

// Stop mocks base method
func (m *MockBlockchain) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockBlockchainMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBlockchain)(nil).Stop))
}

// GetHeightByHash mocks base method
func (m *MockBlockchain) GetHeightByHash(hash common.Hash32B) (uint64, error) {
	ret := m.ctrl.Call(m, "GetHeightByHash", hash)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeightByHash indicates an expected call of GetHeightByHash
func (mr *MockBlockchainMockRecorder) GetHeightByHash(hash interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeightByHash", reflect.TypeOf((*MockBlockchain)(nil).GetHeightByHash), hash)
}

// GetHashByHeight mocks base method
func (m *MockBlockchain) GetHashByHeight(height uint64) (common.Hash32B, error) {
	ret := m.ctrl.Call(m, "GetHashByHeight", height)
	ret0, _ := ret[0].(common.Hash32B)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHashByHeight indicates an expected call of GetHashByHeight
func (mr *MockBlockchainMockRecorder) GetHashByHeight(height interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashByHeight", reflect.TypeOf((*MockBlockchain)(nil).GetHashByHeight), height)
}

// GetBlockByHeight mocks base method
func (m *MockBlockchain) GetBlockByHeight(height uint64) (*blockchain.Block, error) {
	ret := m.ctrl.Call(m, "GetBlockByHeight", height)
	ret0, _ := ret[0].(*blockchain.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHeight indicates an expected call of GetBlockByHeight
func (mr *MockBlockchainMockRecorder) GetBlockByHeight(height interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHeight", reflect.TypeOf((*MockBlockchain)(nil).GetBlockByHeight), height)
}

// GetBlockByHash mocks base method
func (m *MockBlockchain) GetBlockByHash(hash common.Hash32B) (*blockchain.Block, error) {
	ret := m.ctrl.Call(m, "GetBlockByHash", hash)
	ret0, _ := ret[0].(*blockchain.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHash indicates an expected call of GetBlockByHash
func (mr *MockBlockchainMockRecorder) GetBlockByHash(hash interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHash", reflect.TypeOf((*MockBlockchain)(nil).GetBlockByHash), hash)
}

// TipHash mocks base method
func (m *MockBlockchain) TipHash() (common.Hash32B, error) {
	ret := m.ctrl.Call(m, "TipHash")
	ret0, _ := ret[0].(common.Hash32B)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TipHash indicates an expected call of TipHash
func (mr *MockBlockchainMockRecorder) TipHash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TipHash", reflect.TypeOf((*MockBlockchain)(nil).TipHash))
}

// TipHeight mocks base method
func (m *MockBlockchain) TipHeight() (uint64, error) {
	ret := m.ctrl.Call(m, "TipHeight")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TipHeight indicates an expected call of TipHeight
func (mr *MockBlockchainMockRecorder) TipHeight() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TipHeight", reflect.TypeOf((*MockBlockchain)(nil).TipHeight))
}

// Reset mocks base method
func (m *MockBlockchain) Reset() {
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset
func (mr *MockBlockchainMockRecorder) Reset() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockBlockchain)(nil).Reset))
}

// ValidateBlock mocks base method
func (m *MockBlockchain) ValidateBlock(blk *blockchain.Block) error {
	ret := m.ctrl.Call(m, "ValidateBlock", blk)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateBlock indicates an expected call of ValidateBlock
func (mr *MockBlockchainMockRecorder) ValidateBlock(blk interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateBlock", reflect.TypeOf((*MockBlockchain)(nil).ValidateBlock), blk)
}

// MintNewBlock mocks base method
func (m *MockBlockchain) MintNewBlock(arg0 []*blockchain.Tx, arg1 *iotxaddress.Address, arg2 string) (*blockchain.Block, error) {
	ret := m.ctrl.Call(m, "MintNewBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(*blockchain.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MintNewBlock indicates an expected call of MintNewBlock
func (mr *MockBlockchainMockRecorder) MintNewBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintNewBlock", reflect.TypeOf((*MockBlockchain)(nil).MintNewBlock), arg0, arg1, arg2)
}

// AddBlockCommit mocks base method
func (m *MockBlockchain) AddBlockCommit(blk *blockchain.Block) error {
	ret := m.ctrl.Call(m, "AddBlockCommit", blk)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBlockCommit indicates an expected call of AddBlockCommit
func (mr *MockBlockchainMockRecorder) AddBlockCommit(blk interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlockCommit", reflect.TypeOf((*MockBlockchain)(nil).AddBlockCommit), blk)
}

// AddBlockSync mocks base method
func (m *MockBlockchain) AddBlockSync(blk *blockchain.Block) error {
	ret := m.ctrl.Call(m, "AddBlockSync", blk)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBlockSync indicates an expected call of AddBlockSync
func (mr *MockBlockchainMockRecorder) AddBlockSync(blk interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlockSync", reflect.TypeOf((*MockBlockchain)(nil).AddBlockSync), blk)
}

// BalanceOf mocks base method
func (m *MockBlockchain) BalanceOf(arg0 string) *big.Int {
	ret := m.ctrl.Call(m, "BalanceOf", arg0)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// BalanceOf indicates an expected call of BalanceOf
func (mr *MockBlockchainMockRecorder) BalanceOf(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceOf", reflect.TypeOf((*MockBlockchain)(nil).BalanceOf), arg0)
}

// UtxoPool mocks base method
func (m *MockBlockchain) UtxoPool() map[common.Hash32B][]*blockchain.TxOutput {
	ret := m.ctrl.Call(m, "UtxoPool")
	ret0, _ := ret[0].(map[common.Hash32B][]*blockchain.TxOutput)
	return ret0
}

// UtxoPool indicates an expected call of UtxoPool
func (mr *MockBlockchainMockRecorder) UtxoPool() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UtxoPool", reflect.TypeOf((*MockBlockchain)(nil).UtxoPool))
}

// CreateTransaction mocks base method
func (m *MockBlockchain) CreateTransaction(from *iotxaddress.Address, amount uint64, to []*blockchain.Payee) *blockchain.Tx {
	ret := m.ctrl.Call(m, "CreateTransaction", from, amount, to)
	ret0, _ := ret[0].(*blockchain.Tx)
	return ret0
}

// CreateTransaction indicates an expected call of CreateTransaction
func (mr *MockBlockchainMockRecorder) CreateTransaction(from, amount, to interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockBlockchain)(nil).CreateTransaction), from, amount, to)
}

// CreateRawTransaction mocks base method
func (m *MockBlockchain) CreateRawTransaction(from *iotxaddress.Address, amount uint64, to []*blockchain.Payee) *blockchain.Tx {
	ret := m.ctrl.Call(m, "CreateRawTransaction", from, amount, to)
	ret0, _ := ret[0].(*blockchain.Tx)
	return ret0
}

// CreateRawTransaction indicates an expected call of CreateRawTransaction
func (mr *MockBlockchainMockRecorder) CreateRawTransaction(from, amount, to interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRawTransaction", reflect.TypeOf((*MockBlockchain)(nil).CreateRawTransaction), from, amount, to)
}
