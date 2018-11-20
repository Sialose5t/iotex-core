// Code generated by MockGen. DO NOT EDIT.
// Source: ./state/factory/factory.go

// Package mock_factory is a generated GoMock package.
package mock_factory

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	protocol "github.com/iotexproject/iotex-core/action/protocol"
	hash "github.com/iotexproject/iotex-core/pkg/hash"
	state "github.com/iotexproject/iotex-core/state"
	factory "github.com/iotexproject/iotex-core/state/factory"
	big "math/big"
	reflect "reflect"
)

// MockFactory is a mock of Factory interface
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockFactory) Start(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockFactoryMockRecorder) Start(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockFactory)(nil).Start), arg0)
}

// Stop mocks base method
func (m *MockFactory) Stop(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockFactoryMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockFactory)(nil).Stop), arg0)
}

// Balance mocks base method
func (m *MockFactory) Balance(arg0 string) (*big.Int, error) {
	ret := m.ctrl.Call(m, "Balance", arg0)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Balance indicates an expected call of Balance
func (mr *MockFactoryMockRecorder) Balance(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockFactory)(nil).Balance), arg0)
}

// Nonce mocks base method
func (m *MockFactory) Nonce(arg0 string) (uint64, error) {
	ret := m.ctrl.Call(m, "Nonce", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Nonce indicates an expected call of Nonce
func (mr *MockFactoryMockRecorder) Nonce(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nonce", reflect.TypeOf((*MockFactory)(nil).Nonce), arg0)
}

// AccountState mocks base method
func (m *MockFactory) AccountState(arg0 string) (*state.Account, error) {
	ret := m.ctrl.Call(m, "AccountState", arg0)
	ret0, _ := ret[0].(*state.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountState indicates an expected call of AccountState
func (mr *MockFactoryMockRecorder) AccountState(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountState", reflect.TypeOf((*MockFactory)(nil).AccountState), arg0)
}

// RootHash mocks base method
func (m *MockFactory) RootHash() hash.Hash32B {
	ret := m.ctrl.Call(m, "RootHash")
	ret0, _ := ret[0].(hash.Hash32B)
	return ret0
}

// RootHash indicates an expected call of RootHash
func (mr *MockFactoryMockRecorder) RootHash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootHash", reflect.TypeOf((*MockFactory)(nil).RootHash))
}

// Height mocks base method
func (m *MockFactory) Height() (uint64, error) {
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Height indicates an expected call of Height
func (mr *MockFactoryMockRecorder) Height() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockFactory)(nil).Height))
}

// NewWorkingSet mocks base method
func (m *MockFactory) NewWorkingSet() (factory.WorkingSet, error) {
	ret := m.ctrl.Call(m, "NewWorkingSet")
	ret0, _ := ret[0].(factory.WorkingSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewWorkingSet indicates an expected call of NewWorkingSet
func (mr *MockFactoryMockRecorder) NewWorkingSet() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWorkingSet", reflect.TypeOf((*MockFactory)(nil).NewWorkingSet))
}

// Commit mocks base method
func (m *MockFactory) Commit(arg0 factory.WorkingSet) error {
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockFactoryMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockFactory)(nil).Commit), arg0)
}

// CandidatesByHeight mocks base method
func (m *MockFactory) CandidatesByHeight(arg0 uint64) ([]*state.Candidate, error) {
	ret := m.ctrl.Call(m, "CandidatesByHeight", arg0)
	ret0, _ := ret[0].([]*state.Candidate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CandidatesByHeight indicates an expected call of CandidatesByHeight
func (mr *MockFactoryMockRecorder) CandidatesByHeight(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CandidatesByHeight", reflect.TypeOf((*MockFactory)(nil).CandidatesByHeight), arg0)
}

// State mocks base method
func (m *MockFactory) State(arg0 hash.PKHash, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "State", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// State indicates an expected call of State
func (mr *MockFactoryMockRecorder) State(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockFactory)(nil).State), arg0, arg1)
}

// AddActionHandlers mocks base method
func (m *MockFactory) AddActionHandlers(arg0 ...protocol.ActionHandler) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddActionHandlers", varargs...)
}

// AddActionHandlers indicates an expected call of AddActionHandlers
func (mr *MockFactoryMockRecorder) AddActionHandlers(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddActionHandlers", reflect.TypeOf((*MockFactory)(nil).AddActionHandlers), arg0...)
}
