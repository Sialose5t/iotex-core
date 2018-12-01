// Code generated by MockGen. DO NOT EDIT.
// Source: ./action/protocol/protocol.go

// Package mock_chainmanager is a generated GoMock package.
package mock_chainmanager

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	action "github.com/iotexproject/iotex-core/action"
	protocol "github.com/iotexproject/iotex-core/action/protocol"
	db "github.com/iotexproject/iotex-core/db"
	hash "github.com/iotexproject/iotex-core/pkg/hash"
	state "github.com/iotexproject/iotex-core/state"
	big "math/big"
	reflect "reflect"
)

// MockProtocol is a mock of Protocol interface
type MockProtocol struct {
	ctrl     *gomock.Controller
	recorder *MockProtocolMockRecorder
}

// MockProtocolMockRecorder is the mock recorder for MockProtocol
type MockProtocolMockRecorder struct {
	mock *MockProtocol
}

// NewMockProtocol creates a new mock instance
func NewMockProtocol(ctrl *gomock.Controller) *MockProtocol {
	mock := &MockProtocol{ctrl: ctrl}
	mock.recorder = &MockProtocolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProtocol) EXPECT() *MockProtocolMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockProtocol) Validate(arg0 context.Context, arg1 action.Action) error {
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockProtocolMockRecorder) Validate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockProtocol)(nil).Validate), arg0, arg1)
}

// Handle mocks base method
func (m *MockProtocol) Handle(arg0 context.Context, arg1 action.Action, arg2 protocol.StateManager) (*action.Receipt, error) {
	ret := m.ctrl.Call(m, "Handle", arg0, arg1, arg2)
	ret0, _ := ret[0].(*action.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle
func (mr *MockProtocolMockRecorder) Handle(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockProtocol)(nil).Handle), arg0, arg1, arg2)
}

// MockActionValidator is a mock of ActionValidator interface
type MockActionValidator struct {
	ctrl     *gomock.Controller
	recorder *MockActionValidatorMockRecorder
}

// MockActionValidatorMockRecorder is the mock recorder for MockActionValidator
type MockActionValidatorMockRecorder struct {
	mock *MockActionValidator
}

// NewMockActionValidator creates a new mock instance
func NewMockActionValidator(ctrl *gomock.Controller) *MockActionValidator {
	mock := &MockActionValidator{ctrl: ctrl}
	mock.recorder = &MockActionValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActionValidator) EXPECT() *MockActionValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockActionValidator) Validate(arg0 context.Context, arg1 action.Action) error {
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockActionValidatorMockRecorder) Validate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockActionValidator)(nil).Validate), arg0, arg1)
}

// MockActionHandler is a mock of ActionHandler interface
type MockActionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockActionHandlerMockRecorder
}

// MockActionHandlerMockRecorder is the mock recorder for MockActionHandler
type MockActionHandlerMockRecorder struct {
	mock *MockActionHandler
}

// NewMockActionHandler creates a new mock instance
func NewMockActionHandler(ctrl *gomock.Controller) *MockActionHandler {
	mock := &MockActionHandler{ctrl: ctrl}
	mock.recorder = &MockActionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActionHandler) EXPECT() *MockActionHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method
func (m *MockActionHandler) Handle(arg0 context.Context, arg1 action.Action, arg2 protocol.StateManager) (*action.Receipt, error) {
	ret := m.ctrl.Call(m, "Handle", arg0, arg1, arg2)
	ret0, _ := ret[0].(*action.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle
func (mr *MockActionHandlerMockRecorder) Handle(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockActionHandler)(nil).Handle), arg0, arg1, arg2)
}

// MockChainManager is a mock of ChainManager interface
type MockChainManager struct {
	ctrl     *gomock.Controller
	recorder *MockChainManagerMockRecorder
}

// MockChainManagerMockRecorder is the mock recorder for MockChainManager
type MockChainManagerMockRecorder struct {
	mock *MockChainManager
}

// NewMockChainManager creates a new mock instance
func NewMockChainManager(ctrl *gomock.Controller) *MockChainManager {
	mock := &MockChainManager{ctrl: ctrl}
	mock.recorder = &MockChainManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainManager) EXPECT() *MockChainManagerMockRecorder {
	return m.recorder
}

// ChainID mocks base method
func (m *MockChainManager) ChainID() uint32 {
	ret := m.ctrl.Call(m, "ChainID")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// ChainID indicates an expected call of ChainID
func (mr *MockChainManagerMockRecorder) ChainID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainID", reflect.TypeOf((*MockChainManager)(nil).ChainID))
}

// GetHashByHeight mocks base method
func (m *MockChainManager) GetHashByHeight(height uint64) (hash.Hash32B, error) {
	ret := m.ctrl.Call(m, "GetHashByHeight", height)
	ret0, _ := ret[0].(hash.Hash32B)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHashByHeight indicates an expected call of GetHashByHeight
func (mr *MockChainManagerMockRecorder) GetHashByHeight(height interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashByHeight", reflect.TypeOf((*MockChainManager)(nil).GetHashByHeight), height)
}

// StateByAddr mocks base method
func (m *MockChainManager) StateByAddr(address string) (*state.Account, error) {
	ret := m.ctrl.Call(m, "StateByAddr", address)
	ret0, _ := ret[0].(*state.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateByAddr indicates an expected call of StateByAddr
func (mr *MockChainManagerMockRecorder) StateByAddr(address interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateByAddr", reflect.TypeOf((*MockChainManager)(nil).StateByAddr), address)
}

// Nonce mocks base method
func (m *MockChainManager) Nonce(addr string) (uint64, error) {
	ret := m.ctrl.Call(m, "Nonce", addr)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Nonce indicates an expected call of Nonce
func (mr *MockChainManagerMockRecorder) Nonce(addr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nonce", reflect.TypeOf((*MockChainManager)(nil).Nonce), addr)
}

// MockStateManager is a mock of StateManager interface
type MockStateManager struct {
	ctrl     *gomock.Controller
	recorder *MockStateManagerMockRecorder
}

// MockStateManagerMockRecorder is the mock recorder for MockStateManager
type MockStateManagerMockRecorder struct {
	mock *MockStateManager
}

// NewMockStateManager creates a new mock instance
func NewMockStateManager(ctrl *gomock.Controller) *MockStateManager {
	mock := &MockStateManager{ctrl: ctrl}
	mock.recorder = &MockStateManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStateManager) EXPECT() *MockStateManagerMockRecorder {
	return m.recorder
}

// LoadOrCreateAccountState mocks base method
func (m *MockStateManager) LoadOrCreateAccountState(arg0 string, arg1 *big.Int) (*state.Account, error) {
	ret := m.ctrl.Call(m, "LoadOrCreateAccountState", arg0, arg1)
	ret0, _ := ret[0].(*state.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadOrCreateAccountState indicates an expected call of LoadOrCreateAccountState
func (mr *MockStateManagerMockRecorder) LoadOrCreateAccountState(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadOrCreateAccountState", reflect.TypeOf((*MockStateManager)(nil).LoadOrCreateAccountState), arg0, arg1)
}

// CachedAccountState mocks base method
func (m *MockStateManager) CachedAccountState(arg0 string) (*state.Account, error) {
	ret := m.ctrl.Call(m, "CachedAccountState", arg0)
	ret0, _ := ret[0].(*state.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CachedAccountState indicates an expected call of CachedAccountState
func (mr *MockStateManagerMockRecorder) CachedAccountState(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CachedAccountState", reflect.TypeOf((*MockStateManager)(nil).CachedAccountState), arg0)
}

// Height mocks base method
func (m *MockStateManager) Height() uint64 {
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Height indicates an expected call of Height
func (mr *MockStateManagerMockRecorder) Height() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockStateManager)(nil).Height))
}

// State mocks base method
func (m *MockStateManager) State(arg0 hash.PKHash, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "State", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// State indicates an expected call of State
func (mr *MockStateManagerMockRecorder) State(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockStateManager)(nil).State), arg0, arg1)
}

// CachedState mocks base method
func (m *MockStateManager) CachedState(arg0 hash.PKHash, arg1 state.State) (state.State, error) {
	ret := m.ctrl.Call(m, "CachedState", arg0, arg1)
	ret0, _ := ret[0].(state.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CachedState indicates an expected call of CachedState
func (mr *MockStateManagerMockRecorder) CachedState(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CachedState", reflect.TypeOf((*MockStateManager)(nil).CachedState), arg0, arg1)
}

// PutState mocks base method
func (m *MockStateManager) PutState(arg0 hash.PKHash, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "PutState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutState indicates an expected call of PutState
func (mr *MockStateManagerMockRecorder) PutState(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutState", reflect.TypeOf((*MockStateManager)(nil).PutState), arg0, arg1)
}

// GetDB mocks base method
func (m *MockStateManager) GetDB() db.KVStore {
	ret := m.ctrl.Call(m, "GetDB")
	ret0, _ := ret[0].(db.KVStore)
	return ret0
}

// GetDB indicates an expected call of GetDB
func (mr *MockStateManagerMockRecorder) GetDB() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDB", reflect.TypeOf((*MockStateManager)(nil).GetDB))
}

// GetCachedBatch mocks base method
func (m *MockStateManager) GetCachedBatch() db.CachedBatch {
	ret := m.ctrl.Call(m, "GetCachedBatch")
	ret0, _ := ret[0].(db.CachedBatch)
	return ret0
}

// GetCachedBatch indicates an expected call of GetCachedBatch
func (mr *MockStateManagerMockRecorder) GetCachedBatch() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCachedBatch", reflect.TypeOf((*MockStateManager)(nil).GetCachedBatch))
}
