// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/service-mesh-hub/cli/pkg/common/config (interfaces: KubeLoader)

// Package cli_mocks is a generated GoMock package.
package cli_mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	resource "k8s.io/cli-runtime/pkg/resource"
	rest "k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	api "k8s.io/client-go/tools/clientcmd/api"
)

// MockKubeLoader is a mock of KubeLoader interface.
type MockKubeLoader struct {
	ctrl     *gomock.Controller
	recorder *MockKubeLoaderMockRecorder
}

// MockKubeLoaderMockRecorder is the mock recorder for MockKubeLoader.
type MockKubeLoaderMockRecorder struct {
	mock *MockKubeLoader
}

// NewMockKubeLoader creates a new mock instance.
func NewMockKubeLoader(ctrl *gomock.Controller) *MockKubeLoader {
	mock := &MockKubeLoader{ctrl: ctrl}
	mock.recorder = &MockKubeLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeLoader) EXPECT() *MockKubeLoaderMockRecorder {
	return m.recorder
}

// GetConfigWithContext mocks base method.
func (m *MockKubeLoader) GetConfigWithContext(arg0, arg1, arg2 string) (clientcmd.ClientConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigWithContext", arg0, arg1, arg2)
	ret0, _ := ret[0].(clientcmd.ClientConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigWithContext indicates an expected call of GetConfigWithContext.
func (mr *MockKubeLoaderMockRecorder) GetConfigWithContext(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigWithContext", reflect.TypeOf((*MockKubeLoader)(nil).GetConfigWithContext), arg0, arg1, arg2)
}

// GetRawConfigForContext mocks base method.
func (m *MockKubeLoader) GetRawConfigForContext(arg0, arg1 string) (api.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawConfigForContext", arg0, arg1)
	ret0, _ := ret[0].(api.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawConfigForContext indicates an expected call of GetRawConfigForContext.
func (mr *MockKubeLoaderMockRecorder) GetRawConfigForContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawConfigForContext", reflect.TypeOf((*MockKubeLoader)(nil).GetRawConfigForContext), arg0, arg1)
}

// GetRestConfigForContext mocks base method.
func (m *MockKubeLoader) GetRestConfigForContext(arg0, arg1 string) (*rest.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestConfigForContext", arg0, arg1)
	ret0, _ := ret[0].(*rest.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestConfigForContext indicates an expected call of GetRestConfigForContext.
func (mr *MockKubeLoaderMockRecorder) GetRestConfigForContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestConfigForContext", reflect.TypeOf((*MockKubeLoader)(nil).GetRestConfigForContext), arg0, arg1)
}

// GetRestConfigFromBytes mocks base method.
func (m *MockKubeLoader) GetRestConfigFromBytes(arg0 []byte) (*rest.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestConfigFromBytes", arg0)
	ret0, _ := ret[0].(*rest.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestConfigFromBytes indicates an expected call of GetRestConfigFromBytes.
func (mr *MockKubeLoaderMockRecorder) GetRestConfigFromBytes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestConfigFromBytes", reflect.TypeOf((*MockKubeLoader)(nil).GetRestConfigFromBytes), arg0)
}

// RESTClientGetter mocks base method.
func (m *MockKubeLoader) RESTClientGetter(arg0, arg1 string) resource.RESTClientGetter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClientGetter", arg0, arg1)
	ret0, _ := ret[0].(resource.RESTClientGetter)
	return ret0
}

// RESTClientGetter indicates an expected call of RESTClientGetter.
func (mr *MockKubeLoaderMockRecorder) RESTClientGetter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClientGetter", reflect.TypeOf((*MockKubeLoader)(nil).RESTClientGetter), arg0, arg1)
}
