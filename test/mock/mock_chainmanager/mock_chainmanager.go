// Code generated by MockGen. DO NOT EDIT.
// Source: ./action/protocol/managers.go

// Package mock_chainmanager is a generated GoMock package.
package mock_chainmanager

import (
	gomock "github.com/golang/mock/gomock"
	hash "github.com/iotexproject/go-pkgs/hash"
	protocol "github.com/iotexproject/iotex-core/action/protocol"
	db "github.com/iotexproject/iotex-core/db"
	reflect "reflect"
)

// MockStateReader is a mock of StateReader interface
type MockStateReader struct {
	ctrl     *gomock.Controller
	recorder *MockStateReaderMockRecorder
}

// MockStateReaderMockRecorder is the mock recorder for MockStateReader
type MockStateReaderMockRecorder struct {
	mock *MockStateReader
}

// NewMockStateReader creates a new mock instance
func NewMockStateReader(ctrl *gomock.Controller) *MockStateReader {
	mock := &MockStateReader{ctrl: ctrl}
	mock.recorder = &MockStateReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStateReader) EXPECT() *MockStateReaderMockRecorder {
	return m.recorder
}

// Height mocks base method
func (m *MockStateReader) Height() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Height indicates an expected call of Height
func (mr *MockStateReaderMockRecorder) Height() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockStateReader)(nil).Height))
}

// State mocks base method
func (m *MockStateReader) State(arg0 hash.Hash160, arg1 interface{}, arg2 ...protocol.StateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "State", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// State indicates an expected call of State
func (mr *MockStateReaderMockRecorder) State(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockStateReader)(nil).State), varargs...)
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

// Height mocks base method
func (m *MockStateManager) Height() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Height indicates an expected call of Height
func (mr *MockStateManagerMockRecorder) Height() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockStateManager)(nil).Height))
}

// State mocks base method
func (m *MockStateManager) State(arg0 hash.Hash160, arg1 interface{}, arg2 ...protocol.StateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "State", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// State indicates an expected call of State
func (mr *MockStateManagerMockRecorder) State(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockStateManager)(nil).State), varargs...)
}

// Snapshot mocks base method
func (m *MockStateManager) Snapshot() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot")
	ret0, _ := ret[0].(int)
	return ret0
}

// Snapshot indicates an expected call of Snapshot
func (mr *MockStateManagerMockRecorder) Snapshot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockStateManager)(nil).Snapshot))
}

// Revert mocks base method
func (m *MockStateManager) Revert(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Revert indicates an expected call of Revert
func (mr *MockStateManagerMockRecorder) Revert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revert", reflect.TypeOf((*MockStateManager)(nil).Revert), arg0)
}

// PutState mocks base method
func (m *MockStateManager) PutState(arg0 hash.Hash160, arg1 interface{}, arg2 ...protocol.StateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutState", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutState indicates an expected call of PutState
func (mr *MockStateManagerMockRecorder) PutState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutState", reflect.TypeOf((*MockStateManager)(nil).PutState), varargs...)
}

// DelState mocks base method
func (m *MockStateManager) DelState(arg0 hash.Hash160, arg1 ...protocol.StateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DelState", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelState indicates an expected call of DelState
func (mr *MockStateManagerMockRecorder) DelState(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelState", reflect.TypeOf((*MockStateManager)(nil).DelState), varargs...)
}

// GetDB mocks base method
func (m *MockStateManager) GetDB() db.KVStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDB")
	ret0, _ := ret[0].(db.KVStore)
	return ret0
}

// GetDB indicates an expected call of GetDB
func (mr *MockStateManagerMockRecorder) GetDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDB", reflect.TypeOf((*MockStateManager)(nil).GetDB))
}
