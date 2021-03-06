// Code generated by MockGen. DO NOT EDIT.
// Source: ./istio_networking_extender.go

// Package mock_extensions is a generated GoMock package.
package mock_extensions

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	input "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	istio "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/istio"
)

// MockIstioExtender is a mock of IstioExtender interface.
type MockIstioExtender struct {
	ctrl     *gomock.Controller
	recorder *MockIstioExtenderMockRecorder
}

// MockIstioExtenderMockRecorder is the mock recorder for MockIstioExtender.
type MockIstioExtenderMockRecorder struct {
	mock *MockIstioExtender
}

// NewMockIstioExtender creates a new mock instance.
func NewMockIstioExtender(ctrl *gomock.Controller) *MockIstioExtender {
	mock := &MockIstioExtender{ctrl: ctrl}
	mock.recorder = &MockIstioExtenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIstioExtender) EXPECT() *MockIstioExtenderMockRecorder {
	return m.recorder
}

// PatchOutputs mocks base method.
func (m *MockIstioExtender) PatchOutputs(ctx context.Context, inputs input.LocalSnapshot, outputs istio.Builder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchOutputs", ctx, inputs, outputs)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchOutputs indicates an expected call of PatchOutputs.
func (mr *MockIstioExtenderMockRecorder) PatchOutputs(ctx, inputs, outputs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchOutputs", reflect.TypeOf((*MockIstioExtender)(nil).PatchOutputs), ctx, inputs, outputs)
}
