// Code generated by MockGen. DO NOT EDIT.
// Source: ./ioctl/client.go

// Package mock_ioctlclient is a generated GoMock package.
package mock_ioctlclient

import (
	context "context"
	keystore "github.com/ethereum/go-ethereum/accounts/keystore"
	gomock "github.com/golang/mock/gomock"
	ioctl "github.com/iotexproject/iotex-core/ioctl"
	config "github.com/iotexproject/iotex-core/ioctl/config"
	iotexapi "github.com/iotexproject/iotex-proto/golang/iotexapi"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockClient) Start(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockClientMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockClient)(nil).Start), arg0)
}

// Stop mocks base method
func (m *MockClient) Stop(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockClientMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockClient)(nil).Stop), arg0)
}

// Config mocks base method
func (m *MockClient) Config() config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(config.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockClientMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockClient)(nil).Config))
}

// APIServiceClient mocks base method
func (m *MockClient) APIServiceClient(arg0 ioctl.APIServiceConfig) (iotexapi.APIServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIServiceClient", arg0)
	ret0, _ := ret[0].(iotexapi.APIServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIServiceClient indicates an expected call of APIServiceClient
func (mr *MockClientMockRecorder) APIServiceClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIServiceClient", reflect.TypeOf((*MockClient)(nil).APIServiceClient), arg0)
}

// SelectTranslation mocks base method
func (m *MockClient) SelectTranslation(arg0 map[config.Language]string) (string, config.Language) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectTranslation", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(config.Language)
	return ret0, ret1
}

// SelectTranslation indicates an expected call of SelectTranslation
func (mr *MockClientMockRecorder) SelectTranslation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectTranslation", reflect.TypeOf((*MockClient)(nil).SelectTranslation), arg0)
}

// AskToConfirm mocks base method
func (m *MockClient) AskToConfirm() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskToConfirm")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AskToConfirm indicates an expected call of AskToConfirm
func (mr *MockClientMockRecorder) AskToConfirm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskToConfirm", reflect.TypeOf((*MockClient)(nil).AskToConfirm))
}

// ReadSecret mocks base method
func (m *MockClient) ReadSecret() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadSecret")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadSecret indicates an expected call of ReadSecret
func (mr *MockClientMockRecorder) ReadSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadSecret", reflect.TypeOf((*MockClient)(nil).ReadSecret))
}

// Execute mocks base method
func (m *MockClient) Execute(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockClientMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockClient)(nil).Execute), arg0)
}

// GetAddress mocks base method
func (m *MockClient) GetAddress(in string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress", in)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddress indicates an expected call of GetAddress
func (mr *MockClientMockRecorder) GetAddress(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockClient)(nil).GetAddress), in)
}

// Address mocks base method
func (m *MockClient) Address(in string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address", in)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Address indicates an expected call of Address
func (mr *MockClientMockRecorder) Address(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockClient)(nil).Address), in)
}

// NewKeyStore mocks base method
func (m *MockClient) NewKeyStore(arg0 string, arg1, arg2 int) *keystore.KeyStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewKeyStore", arg0, arg1, arg2)
	ret0, _ := ret[0].(*keystore.KeyStore)
	return ret0
}

// NewKeyStore indicates an expected call of NewKeyStore
func (mr *MockClientMockRecorder) NewKeyStore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewKeyStore", reflect.TypeOf((*MockClient)(nil).NewKeyStore), arg0, arg1, arg2)
}

// GetAliasMap mocks base method
func (m *MockClient) GetAliasMap() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAliasMap")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAliasMap indicates an expected call of GetAliasMap
func (mr *MockClientMockRecorder) GetAliasMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAliasMap", reflect.TypeOf((*MockClient)(nil).GetAliasMap))
}
