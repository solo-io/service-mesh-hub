// Code generated by MockGen. DO NOT EDIT.
// Source: ./workload_translator.go

// Package mock_workload is a generated GoMock package.
package mock_workload

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/sets"
	sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
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

// TranslateWorkloads mocks base method
func (m *MockTranslator) TranslateWorkloads(deployments v1sets.DeploymentSet, daemonSets v1sets.DaemonSetSet, statefulSets v1sets.StatefulSetSet, meshes sets.MeshSet) sets.WorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateWorkloads", deployments, daemonSets, statefulSets, meshes)
	ret0, _ := ret[0].(sets.WorkloadSet)
	return ret0
}

// TranslateWorkloads indicates an expected call of TranslateWorkloads
func (mr *MockTranslatorMockRecorder) TranslateWorkloads(deployments, daemonSets, statefulSets, meshes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateWorkloads", reflect.TypeOf((*MockTranslator)(nil).TranslateWorkloads), deployments, daemonSets, statefulSets, meshes)
}
