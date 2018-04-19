// Code generated by MockGen. DO NOT EDIT.
// Source: ./blocksync/blocksync.go

// Package mock_blocksync is a generated GoMock package.
package mock_blocksync

import (
	gomock "github.com/golang/mock/gomock"
	blockchain "github.com/iotexproject/iotex-core/blockchain"
	network "github.com/iotexproject/iotex-core/network"
	proto "github.com/iotexproject/iotex-core/proto"
	reflect "reflect"
)

// MockBlockSync is a mock of BlockSync interface
type MockBlockSync struct {
	ctrl     *gomock.Controller
	recorder *MockBlockSyncMockRecorder
}

// MockBlockSyncMockRecorder is the mock recorder for MockBlockSync
type MockBlockSyncMockRecorder struct {
	mock *MockBlockSync
}

// NewMockBlockSync creates a new mock instance
func NewMockBlockSync(ctrl *gomock.Controller) *MockBlockSync {
	mock := &MockBlockSync{ctrl: ctrl}
	mock.recorder = &MockBlockSyncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockSync) EXPECT() *MockBlockSyncMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockBlockSync) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockBlockSyncMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBlockSync)(nil).Start))
}

// Stop mocks base method
func (m *MockBlockSync) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockBlockSyncMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBlockSync)(nil).Stop))
}

// P2P mocks base method
func (m *MockBlockSync) P2P() *network.Overlay {
	ret := m.ctrl.Call(m, "P2P")
	ret0, _ := ret[0].(*network.Overlay)
	return ret0
}

// P2P indicates an expected call of P2P
func (mr *MockBlockSyncMockRecorder) P2P() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "P2P", reflect.TypeOf((*MockBlockSync)(nil).P2P))
}

// ProcessSyncRequest mocks base method
func (m *MockBlockSync) ProcessSyncRequest(sender string, sync *proto.BlockSync) error {
	ret := m.ctrl.Call(m, "ProcessSyncRequest", sender, sync)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessSyncRequest indicates an expected call of ProcessSyncRequest
func (mr *MockBlockSyncMockRecorder) ProcessSyncRequest(sender, sync interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessSyncRequest", reflect.TypeOf((*MockBlockSync)(nil).ProcessSyncRequest), sender, sync)
}

// ProcessBlock mocks base method
func (m *MockBlockSync) ProcessBlock(blk *blockchain.Block) error {
	ret := m.ctrl.Call(m, "ProcessBlock", blk)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessBlock indicates an expected call of ProcessBlock
func (mr *MockBlockSyncMockRecorder) ProcessBlock(blk interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBlock", reflect.TypeOf((*MockBlockSync)(nil).ProcessBlock), blk)
}

// ProcessBlockSync mocks base method
func (m *MockBlockSync) ProcessBlockSync(blk *blockchain.Block) error {
	ret := m.ctrl.Call(m, "ProcessBlockSync", blk)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessBlockSync indicates an expected call of ProcessBlockSync
func (mr *MockBlockSyncMockRecorder) ProcessBlockSync(blk interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBlockSync", reflect.TypeOf((*MockBlockSync)(nil).ProcessBlockSync), blk)
}
