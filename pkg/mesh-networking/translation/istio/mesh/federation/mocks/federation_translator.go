// Code generated by MockGen. DO NOT EDIT.
// Source: ./federation_translator.go

// Package mock_federation is a generated GoMock package.
package mock_federation

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	input "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/input"
	istio "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/output/istio"
	reporting "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/reporting"
	reflect "reflect"
)

// MockTranslator is a mock of Translator interface
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorMockRecorder
}

// MockTranslatorMockRecorder is the mock recorder for MockTranslator
type MockTranslatorMockRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTranslator) EXPECT() *MockTranslatorMockRecorder {
	return m.recorder
}

// Translate mocks base method
func (m *MockTranslator) Translate(in input.Snapshot, mesh *v1alpha2.Mesh, virtualMesh *v1alpha2.MeshStatus_AppliedVirtualMesh, outputs istio.Builder, reporter reporting.Reporter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Translate", in, mesh, virtualMesh, outputs, reporter)
}

// Translate indicates an expected call of Translate
func (mr *MockTranslatorMockRecorder) Translate(in, mesh, virtualMesh, outputs, reporter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockTranslator)(nil).Translate), in, mesh, virtualMesh, outputs, reporter)
}
