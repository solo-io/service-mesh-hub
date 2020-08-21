// Code generated by MockGen. DO NOT EDIT.
// Source: ./mesh_translator.go

// Package mock_mesh is a generated GoMock package.
package mock_mesh

import (
	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/sets"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
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

// TranslateMeshes mocks base method
func (m *MockTranslator) TranslateMeshes(deployments v1sets.DeploymentSet) v1alpha2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateMeshes", deployments)
	ret0, _ := ret[0].(v1alpha2sets.MeshSet)
	return ret0
}

// TranslateMeshes indicates an expected call of TranslateMeshes
func (mr *MockTranslatorMockRecorder) TranslateMeshes(deployments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateMeshes", reflect.TypeOf((*MockTranslator)(nil).TranslateMeshes), deployments)
}
