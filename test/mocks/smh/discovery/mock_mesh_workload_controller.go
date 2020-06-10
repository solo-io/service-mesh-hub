// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/controller (interfaces: MeshWorkloadEventWatcher,MeshServiceEventWatcher,MeshEventWatcher)

// Package mock_smh_discovery is a generated GoMock package.
package mock_smh_discovery

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMeshWorkloadEventWatcher is a mock of MeshWorkloadEventWatcher interface.
type MockMeshWorkloadEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMeshWorkloadEventWatcherMockRecorder
}

// MockMeshWorkloadEventWatcherMockRecorder is the mock recorder for MockMeshWorkloadEventWatcher.
type MockMeshWorkloadEventWatcherMockRecorder struct {
	mock *MockMeshWorkloadEventWatcher
}

// NewMockMeshWorkloadEventWatcher creates a new mock instance.
func NewMockMeshWorkloadEventWatcher(ctrl *gomock.Controller) *MockMeshWorkloadEventWatcher {
	mock := &MockMeshWorkloadEventWatcher{ctrl: ctrl}
	mock.recorder = &MockMeshWorkloadEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshWorkloadEventWatcher) EXPECT() *MockMeshWorkloadEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockMeshWorkloadEventWatcher) AddEventHandler(arg0 context.Context, arg1 controller.MeshWorkloadEventHandler, arg2 ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockMeshWorkloadEventWatcherMockRecorder) AddEventHandler(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockMeshWorkloadEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockMeshServiceEventWatcher is a mock of MeshServiceEventWatcher interface.
type MockMeshServiceEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMeshServiceEventWatcherMockRecorder
}

// MockMeshServiceEventWatcherMockRecorder is the mock recorder for MockMeshServiceEventWatcher.
type MockMeshServiceEventWatcherMockRecorder struct {
	mock *MockMeshServiceEventWatcher
}

// NewMockMeshServiceEventWatcher creates a new mock instance.
func NewMockMeshServiceEventWatcher(ctrl *gomock.Controller) *MockMeshServiceEventWatcher {
	mock := &MockMeshServiceEventWatcher{ctrl: ctrl}
	mock.recorder = &MockMeshServiceEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshServiceEventWatcher) EXPECT() *MockMeshServiceEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockMeshServiceEventWatcher) AddEventHandler(arg0 context.Context, arg1 controller.MeshServiceEventHandler, arg2 ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockMeshServiceEventWatcherMockRecorder) AddEventHandler(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockMeshServiceEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockMeshEventWatcher is a mock of MeshEventWatcher interface.
type MockMeshEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMeshEventWatcherMockRecorder
}

// MockMeshEventWatcherMockRecorder is the mock recorder for MockMeshEventWatcher.
type MockMeshEventWatcherMockRecorder struct {
	mock *MockMeshEventWatcher
}

// NewMockMeshEventWatcher creates a new mock instance.
func NewMockMeshEventWatcher(ctrl *gomock.Controller) *MockMeshEventWatcher {
	mock := &MockMeshEventWatcher{ctrl: ctrl}
	mock.recorder = &MockMeshEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshEventWatcher) EXPECT() *MockMeshEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockMeshEventWatcher) AddEventHandler(arg0 context.Context, arg1 controller.MeshEventHandler, arg2 ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockMeshEventWatcherMockRecorder) AddEventHandler(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockMeshEventWatcher)(nil).AddEventHandler), varargs...)
}
