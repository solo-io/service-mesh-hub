// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	controller "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterTrafficPolicyReconciler is a mock of MulticlusterTrafficPolicyReconciler interface.
type MockMulticlusterTrafficPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterTrafficPolicyReconcilerMockRecorder
}

// MockMulticlusterTrafficPolicyReconcilerMockRecorder is the mock recorder for MockMulticlusterTrafficPolicyReconciler.
type MockMulticlusterTrafficPolicyReconcilerMockRecorder struct {
	mock *MockMulticlusterTrafficPolicyReconciler
}

// NewMockMulticlusterTrafficPolicyReconciler creates a new mock instance.
func NewMockMulticlusterTrafficPolicyReconciler(ctrl *gomock.Controller) *MockMulticlusterTrafficPolicyReconciler {
	mock := &MockMulticlusterTrafficPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterTrafficPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterTrafficPolicyReconciler) EXPECT() *MockMulticlusterTrafficPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTrafficPolicy mocks base method.
func (m *MockMulticlusterTrafficPolicyReconciler) ReconcileTrafficPolicy(clusterName string, obj *v1.TrafficPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficPolicy", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTrafficPolicy indicates an expected call of ReconcileTrafficPolicy.
func (mr *MockMulticlusterTrafficPolicyReconcilerMockRecorder) ReconcileTrafficPolicy(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficPolicy", reflect.TypeOf((*MockMulticlusterTrafficPolicyReconciler)(nil).ReconcileTrafficPolicy), clusterName, obj)
}

// MockMulticlusterTrafficPolicyDeletionReconciler is a mock of MulticlusterTrafficPolicyDeletionReconciler interface.
type MockMulticlusterTrafficPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterTrafficPolicyDeletionReconcilerMockRecorder
}

// MockMulticlusterTrafficPolicyDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterTrafficPolicyDeletionReconciler.
type MockMulticlusterTrafficPolicyDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterTrafficPolicyDeletionReconciler
}

// NewMockMulticlusterTrafficPolicyDeletionReconciler creates a new mock instance.
func NewMockMulticlusterTrafficPolicyDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterTrafficPolicyDeletionReconciler {
	mock := &MockMulticlusterTrafficPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterTrafficPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterTrafficPolicyDeletionReconciler) EXPECT() *MockMulticlusterTrafficPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTrafficPolicyDeletion mocks base method.
func (m *MockMulticlusterTrafficPolicyDeletionReconciler) ReconcileTrafficPolicyDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficPolicyDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileTrafficPolicyDeletion indicates an expected call of ReconcileTrafficPolicyDeletion.
func (mr *MockMulticlusterTrafficPolicyDeletionReconcilerMockRecorder) ReconcileTrafficPolicyDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficPolicyDeletion", reflect.TypeOf((*MockMulticlusterTrafficPolicyDeletionReconciler)(nil).ReconcileTrafficPolicyDeletion), clusterName, req)
}

// MockMulticlusterTrafficPolicyReconcileLoop is a mock of MulticlusterTrafficPolicyReconcileLoop interface.
type MockMulticlusterTrafficPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterTrafficPolicyReconcileLoopMockRecorder
}

// MockMulticlusterTrafficPolicyReconcileLoopMockRecorder is the mock recorder for MockMulticlusterTrafficPolicyReconcileLoop.
type MockMulticlusterTrafficPolicyReconcileLoopMockRecorder struct {
	mock *MockMulticlusterTrafficPolicyReconcileLoop
}

// NewMockMulticlusterTrafficPolicyReconcileLoop creates a new mock instance.
func NewMockMulticlusterTrafficPolicyReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterTrafficPolicyReconcileLoop {
	mock := &MockMulticlusterTrafficPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterTrafficPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterTrafficPolicyReconcileLoop) EXPECT() *MockMulticlusterTrafficPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterTrafficPolicyReconciler mocks base method.
func (m *MockMulticlusterTrafficPolicyReconcileLoop) AddMulticlusterTrafficPolicyReconciler(ctx context.Context, rec controller.MulticlusterTrafficPolicyReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterTrafficPolicyReconciler", varargs...)
}

// AddMulticlusterTrafficPolicyReconciler indicates an expected call of AddMulticlusterTrafficPolicyReconciler.
func (mr *MockMulticlusterTrafficPolicyReconcileLoopMockRecorder) AddMulticlusterTrafficPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterTrafficPolicyReconciler", reflect.TypeOf((*MockMulticlusterTrafficPolicyReconcileLoop)(nil).AddMulticlusterTrafficPolicyReconciler), varargs...)
}

// MockMulticlusterAccessPolicyReconciler is a mock of MulticlusterAccessPolicyReconciler interface.
type MockMulticlusterAccessPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAccessPolicyReconcilerMockRecorder
}

// MockMulticlusterAccessPolicyReconcilerMockRecorder is the mock recorder for MockMulticlusterAccessPolicyReconciler.
type MockMulticlusterAccessPolicyReconcilerMockRecorder struct {
	mock *MockMulticlusterAccessPolicyReconciler
}

// NewMockMulticlusterAccessPolicyReconciler creates a new mock instance.
func NewMockMulticlusterAccessPolicyReconciler(ctrl *gomock.Controller) *MockMulticlusterAccessPolicyReconciler {
	mock := &MockMulticlusterAccessPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAccessPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterAccessPolicyReconciler) EXPECT() *MockMulticlusterAccessPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAccessPolicy mocks base method.
func (m *MockMulticlusterAccessPolicyReconciler) ReconcileAccessPolicy(clusterName string, obj *v1.AccessPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessPolicy", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAccessPolicy indicates an expected call of ReconcileAccessPolicy.
func (mr *MockMulticlusterAccessPolicyReconcilerMockRecorder) ReconcileAccessPolicy(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessPolicy", reflect.TypeOf((*MockMulticlusterAccessPolicyReconciler)(nil).ReconcileAccessPolicy), clusterName, obj)
}

// MockMulticlusterAccessPolicyDeletionReconciler is a mock of MulticlusterAccessPolicyDeletionReconciler interface.
type MockMulticlusterAccessPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAccessPolicyDeletionReconcilerMockRecorder
}

// MockMulticlusterAccessPolicyDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterAccessPolicyDeletionReconciler.
type MockMulticlusterAccessPolicyDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterAccessPolicyDeletionReconciler
}

// NewMockMulticlusterAccessPolicyDeletionReconciler creates a new mock instance.
func NewMockMulticlusterAccessPolicyDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterAccessPolicyDeletionReconciler {
	mock := &MockMulticlusterAccessPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAccessPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterAccessPolicyDeletionReconciler) EXPECT() *MockMulticlusterAccessPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAccessPolicyDeletion mocks base method.
func (m *MockMulticlusterAccessPolicyDeletionReconciler) ReconcileAccessPolicyDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessPolicyDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileAccessPolicyDeletion indicates an expected call of ReconcileAccessPolicyDeletion.
func (mr *MockMulticlusterAccessPolicyDeletionReconcilerMockRecorder) ReconcileAccessPolicyDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessPolicyDeletion", reflect.TypeOf((*MockMulticlusterAccessPolicyDeletionReconciler)(nil).ReconcileAccessPolicyDeletion), clusterName, req)
}

// MockMulticlusterAccessPolicyReconcileLoop is a mock of MulticlusterAccessPolicyReconcileLoop interface.
type MockMulticlusterAccessPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAccessPolicyReconcileLoopMockRecorder
}

// MockMulticlusterAccessPolicyReconcileLoopMockRecorder is the mock recorder for MockMulticlusterAccessPolicyReconcileLoop.
type MockMulticlusterAccessPolicyReconcileLoopMockRecorder struct {
	mock *MockMulticlusterAccessPolicyReconcileLoop
}

// NewMockMulticlusterAccessPolicyReconcileLoop creates a new mock instance.
func NewMockMulticlusterAccessPolicyReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterAccessPolicyReconcileLoop {
	mock := &MockMulticlusterAccessPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAccessPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterAccessPolicyReconcileLoop) EXPECT() *MockMulticlusterAccessPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterAccessPolicyReconciler mocks base method.
func (m *MockMulticlusterAccessPolicyReconcileLoop) AddMulticlusterAccessPolicyReconciler(ctx context.Context, rec controller.MulticlusterAccessPolicyReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterAccessPolicyReconciler", varargs...)
}

// AddMulticlusterAccessPolicyReconciler indicates an expected call of AddMulticlusterAccessPolicyReconciler.
func (mr *MockMulticlusterAccessPolicyReconcileLoopMockRecorder) AddMulticlusterAccessPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterAccessPolicyReconciler", reflect.TypeOf((*MockMulticlusterAccessPolicyReconcileLoop)(nil).AddMulticlusterAccessPolicyReconciler), varargs...)
}

// MockMulticlusterVirtualMeshReconciler is a mock of MulticlusterVirtualMeshReconciler interface.
type MockMulticlusterVirtualMeshReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualMeshReconcilerMockRecorder
}

// MockMulticlusterVirtualMeshReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualMeshReconciler.
type MockMulticlusterVirtualMeshReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualMeshReconciler
}

// NewMockMulticlusterVirtualMeshReconciler creates a new mock instance.
func NewMockMulticlusterVirtualMeshReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualMeshReconciler {
	mock := &MockMulticlusterVirtualMeshReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualMeshReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualMeshReconciler) EXPECT() *MockMulticlusterVirtualMeshReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualMesh mocks base method.
func (m *MockMulticlusterVirtualMeshReconciler) ReconcileVirtualMesh(clusterName string, obj *v1.VirtualMesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualMesh", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualMesh indicates an expected call of ReconcileVirtualMesh.
func (mr *MockMulticlusterVirtualMeshReconcilerMockRecorder) ReconcileVirtualMesh(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualMesh", reflect.TypeOf((*MockMulticlusterVirtualMeshReconciler)(nil).ReconcileVirtualMesh), clusterName, obj)
}

// MockMulticlusterVirtualMeshDeletionReconciler is a mock of MulticlusterVirtualMeshDeletionReconciler interface.
type MockMulticlusterVirtualMeshDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualMeshDeletionReconcilerMockRecorder
}

// MockMulticlusterVirtualMeshDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualMeshDeletionReconciler.
type MockMulticlusterVirtualMeshDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualMeshDeletionReconciler
}

// NewMockMulticlusterVirtualMeshDeletionReconciler creates a new mock instance.
func NewMockMulticlusterVirtualMeshDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualMeshDeletionReconciler {
	mock := &MockMulticlusterVirtualMeshDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualMeshDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualMeshDeletionReconciler) EXPECT() *MockMulticlusterVirtualMeshDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualMeshDeletion mocks base method.
func (m *MockMulticlusterVirtualMeshDeletionReconciler) ReconcileVirtualMeshDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualMeshDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualMeshDeletion indicates an expected call of ReconcileVirtualMeshDeletion.
func (mr *MockMulticlusterVirtualMeshDeletionReconcilerMockRecorder) ReconcileVirtualMeshDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualMeshDeletion", reflect.TypeOf((*MockMulticlusterVirtualMeshDeletionReconciler)(nil).ReconcileVirtualMeshDeletion), clusterName, req)
}

// MockMulticlusterVirtualMeshReconcileLoop is a mock of MulticlusterVirtualMeshReconcileLoop interface.
type MockMulticlusterVirtualMeshReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualMeshReconcileLoopMockRecorder
}

// MockMulticlusterVirtualMeshReconcileLoopMockRecorder is the mock recorder for MockMulticlusterVirtualMeshReconcileLoop.
type MockMulticlusterVirtualMeshReconcileLoopMockRecorder struct {
	mock *MockMulticlusterVirtualMeshReconcileLoop
}

// NewMockMulticlusterVirtualMeshReconcileLoop creates a new mock instance.
func NewMockMulticlusterVirtualMeshReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterVirtualMeshReconcileLoop {
	mock := &MockMulticlusterVirtualMeshReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualMeshReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualMeshReconcileLoop) EXPECT() *MockMulticlusterVirtualMeshReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterVirtualMeshReconciler mocks base method.
func (m *MockMulticlusterVirtualMeshReconcileLoop) AddMulticlusterVirtualMeshReconciler(ctx context.Context, rec controller.MulticlusterVirtualMeshReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterVirtualMeshReconciler", varargs...)
}

// AddMulticlusterVirtualMeshReconciler indicates an expected call of AddMulticlusterVirtualMeshReconciler.
func (mr *MockMulticlusterVirtualMeshReconcileLoopMockRecorder) AddMulticlusterVirtualMeshReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterVirtualMeshReconciler", reflect.TypeOf((*MockMulticlusterVirtualMeshReconcileLoop)(nil).AddMulticlusterVirtualMeshReconciler), varargs...)
}
