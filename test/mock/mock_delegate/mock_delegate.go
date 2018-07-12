// Code generated by MockGen. DO NOT EDIT.
// Source: ./delegate/delegate.go

// Package mock_delegate is a generated GoMock package.
package mock_delegate

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPool is a mock of Pool interface
type MockPool struct {
	ctrl     *gomock.Controller
	recorder *MockPoolMockRecorder
}

// MockPoolMockRecorder is the mock recorder for MockPool
type MockPoolMockRecorder struct {
	mock *MockPool
}

// NewMockPool creates a new mock instance
func NewMockPool(ctrl *gomock.Controller) *MockPool {
	mock := &MockPool{ctrl: ctrl}
	mock.recorder = &MockPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPool) EXPECT() *MockPoolMockRecorder {
	return m.recorder
}

// AllDelegates mocks base method
func (m *MockPool) AllDelegates() ([]string, error) {
	ret := m.ctrl.Call(m, "AllDelegates")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllDelegates indicates an expected call of AllDelegates
func (mr *MockPoolMockRecorder) AllDelegates() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllDelegates", reflect.TypeOf((*MockPool)(nil).AllDelegates))
}

// RollDelegates mocks base method
func (m *MockPool) RollDelegates(arg0 uint64) ([]string, error) {
	ret := m.ctrl.Call(m, "RollDelegates", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RollDelegates indicates an expected call of RollDelegates
func (mr *MockPoolMockRecorder) RollDelegates(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollDelegates", reflect.TypeOf((*MockPool)(nil).RollDelegates), arg0)
}

// AnotherDelegate mocks base method
func (m *MockPool) AnotherDelegate(self string) string {
	ret := m.ctrl.Call(m, "AnotherDelegate", self)
	ret0, _ := ret[0].(string)
	return ret0
}

// AnotherDelegate indicates an expected call of AnotherDelegate
func (mr *MockPoolMockRecorder) AnotherDelegate(self interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnotherDelegate", reflect.TypeOf((*MockPool)(nil).AnotherDelegate), self)
}

// NumDelegatesPerEpoch mocks base method
func (m *MockPool) NumDelegatesPerEpoch() (uint, error) {
	ret := m.ctrl.Call(m, "NumDelegatesPerEpoch")
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NumDelegatesPerEpoch indicates an expected call of NumDelegatesPerEpoch
func (mr *MockPoolMockRecorder) NumDelegatesPerEpoch() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumDelegatesPerEpoch", reflect.TypeOf((*MockPool)(nil).NumDelegatesPerEpoch))
}
