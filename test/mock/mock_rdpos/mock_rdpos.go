// Code generated by MockGen. DO NOT EDIT.
// Source: ./consensus/scheme/rdpos/rdpos.go

// Package mock_rdpos is a generated GoMock package.
package mock_rdpos

import (
	gomock "github.com/golang/mock/gomock"
	proto "github.com/golang/protobuf/proto"
	net "net"
	reflect "reflect"
)

// MockDNet is a mock of DNet interface
type MockDNet struct {
	ctrl     *gomock.Controller
	recorder *MockDNetMockRecorder
}

// MockDNetMockRecorder is the mock recorder for MockDNet
type MockDNetMockRecorder struct {
	mock *MockDNet
}

// NewMockDNet creates a new mock instance
func NewMockDNet(ctrl *gomock.Controller) *MockDNet {
	mock := &MockDNet{ctrl: ctrl}
	mock.recorder = &MockDNetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDNet) EXPECT() *MockDNetMockRecorder {
	return m.recorder
}

// Tell mocks base method
func (m *MockDNet) Tell(node net.Addr, msg proto.Message) error {
	ret := m.ctrl.Call(m, "Tell", node, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tell indicates an expected call of Tell
func (mr *MockDNetMockRecorder) Tell(node, msg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tell", reflect.TypeOf((*MockDNet)(nil).Tell), node, msg)
}

// Self mocks base method
func (m *MockDNet) Self() net.Addr {
	ret := m.ctrl.Call(m, "Self")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// Self indicates an expected call of Self
func (mr *MockDNetMockRecorder) Self() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Self", reflect.TypeOf((*MockDNet)(nil).Self))
}

// Broadcast mocks base method
func (m *MockDNet) Broadcast(msg proto.Message) error {
	ret := m.ctrl.Call(m, "Broadcast", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockDNetMockRecorder) Broadcast(msg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockDNet)(nil).Broadcast), msg)
}
