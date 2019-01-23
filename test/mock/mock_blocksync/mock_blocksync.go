// Code generated by MockGen. DO NOT EDIT.
// Source: ./blocksync/blocksync.go

// Package mock_blocksync is a generated GoMock package.
package mock_blocksync

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	block "github.com/iotexproject/iotex-core/blockchain/block"
	proto "github.com/iotexproject/iotex-core/proto"
	go_libp2p_peerstore "github.com/libp2p/go-libp2p-peerstore"
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
func (m *MockBlockSync) Start(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockBlockSyncMockRecorder) Start(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBlockSync)(nil).Start), arg0)
}

// Stop mocks base method
func (m *MockBlockSync) Stop(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockBlockSyncMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBlockSync)(nil).Stop), arg0)
}

// TargetHeight mocks base method
func (m *MockBlockSync) TargetHeight() uint64 {
	ret := m.ctrl.Call(m, "TargetHeight")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// TargetHeight indicates an expected call of TargetHeight
func (mr *MockBlockSyncMockRecorder) TargetHeight() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TargetHeight", reflect.TypeOf((*MockBlockSync)(nil).TargetHeight))
}

// ProcessSyncRequest mocks base method
func (m *MockBlockSync) ProcessSyncRequest(ctx context.Context, peer go_libp2p_peerstore.PeerInfo, sync *proto.BlockSync) error {
	ret := m.ctrl.Call(m, "ProcessSyncRequest", ctx, peer, sync)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessSyncRequest indicates an expected call of ProcessSyncRequest
func (mr *MockBlockSyncMockRecorder) ProcessSyncRequest(ctx, peer, sync interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessSyncRequest", reflect.TypeOf((*MockBlockSync)(nil).ProcessSyncRequest), ctx, peer, sync)
}

// ProcessBlock mocks base method
func (m *MockBlockSync) ProcessBlock(ctx context.Context, blk *block.Block) error {
	ret := m.ctrl.Call(m, "ProcessBlock", ctx, blk)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessBlock indicates an expected call of ProcessBlock
func (mr *MockBlockSyncMockRecorder) ProcessBlock(ctx, blk interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBlock", reflect.TypeOf((*MockBlockSync)(nil).ProcessBlock), ctx, blk)
}

// ProcessBlockSync mocks base method
func (m *MockBlockSync) ProcessBlockSync(ctx context.Context, blk *block.Block) error {
	ret := m.ctrl.Call(m, "ProcessBlockSync", ctx, blk)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessBlockSync indicates an expected call of ProcessBlockSync
func (mr *MockBlockSyncMockRecorder) ProcessBlockSync(ctx, blk interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBlockSync", reflect.TypeOf((*MockBlockSync)(nil).ProcessBlockSync), ctx, blk)
}
