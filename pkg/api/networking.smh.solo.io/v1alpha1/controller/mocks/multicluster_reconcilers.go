// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	controller "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/controller"
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
func (m *MockMulticlusterTrafficPolicyReconciler) ReconcileTrafficPolicy(clusterName string, obj *v1alpha1.TrafficPolicy) (reconcile.Result, error) {
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

// MockMulticlusterAccessControlPolicyReconciler is a mock of MulticlusterAccessControlPolicyReconciler interface.
type MockMulticlusterAccessControlPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAccessControlPolicyReconcilerMockRecorder
}

// MockMulticlusterAccessControlPolicyReconcilerMockRecorder is the mock recorder for MockMulticlusterAccessControlPolicyReconciler.
type MockMulticlusterAccessControlPolicyReconcilerMockRecorder struct {
	mock *MockMulticlusterAccessControlPolicyReconciler
}

// NewMockMulticlusterAccessControlPolicyReconciler creates a new mock instance.
func NewMockMulticlusterAccessControlPolicyReconciler(ctrl *gomock.Controller) *MockMulticlusterAccessControlPolicyReconciler {
	mock := &MockMulticlusterAccessControlPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAccessControlPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterAccessControlPolicyReconciler) EXPECT() *MockMulticlusterAccessControlPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAccessControlPolicy mocks base method.
func (m *MockMulticlusterAccessControlPolicyReconciler) ReconcileAccessControlPolicy(clusterName string, obj *v1alpha1.AccessControlPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessControlPolicy", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAccessControlPolicy indicates an expected call of ReconcileAccessControlPolicy.
func (mr *MockMulticlusterAccessControlPolicyReconcilerMockRecorder) ReconcileAccessControlPolicy(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessControlPolicy", reflect.TypeOf((*MockMulticlusterAccessControlPolicyReconciler)(nil).ReconcileAccessControlPolicy), clusterName, obj)
}

// MockMulticlusterAccessControlPolicyDeletionReconciler is a mock of MulticlusterAccessControlPolicyDeletionReconciler interface.
type MockMulticlusterAccessControlPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAccessControlPolicyDeletionReconcilerMockRecorder
}

// MockMulticlusterAccessControlPolicyDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterAccessControlPolicyDeletionReconciler.
type MockMulticlusterAccessControlPolicyDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterAccessControlPolicyDeletionReconciler
}

// NewMockMulticlusterAccessControlPolicyDeletionReconciler creates a new mock instance.
func NewMockMulticlusterAccessControlPolicyDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterAccessControlPolicyDeletionReconciler {
	mock := &MockMulticlusterAccessControlPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAccessControlPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterAccessControlPolicyDeletionReconciler) EXPECT() *MockMulticlusterAccessControlPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAccessControlPolicyDeletion mocks base method.
func (m *MockMulticlusterAccessControlPolicyDeletionReconciler) ReconcileAccessControlPolicyDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessControlPolicyDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileAccessControlPolicyDeletion indicates an expected call of ReconcileAccessControlPolicyDeletion.
func (mr *MockMulticlusterAccessControlPolicyDeletionReconcilerMockRecorder) ReconcileAccessControlPolicyDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessControlPolicyDeletion", reflect.TypeOf((*MockMulticlusterAccessControlPolicyDeletionReconciler)(nil).ReconcileAccessControlPolicyDeletion), clusterName, req)
}

// MockMulticlusterAccessControlPolicyReconcileLoop is a mock of MulticlusterAccessControlPolicyReconcileLoop interface.
type MockMulticlusterAccessControlPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAccessControlPolicyReconcileLoopMockRecorder
}

// MockMulticlusterAccessControlPolicyReconcileLoopMockRecorder is the mock recorder for MockMulticlusterAccessControlPolicyReconcileLoop.
type MockMulticlusterAccessControlPolicyReconcileLoopMockRecorder struct {
	mock *MockMulticlusterAccessControlPolicyReconcileLoop
}

// NewMockMulticlusterAccessControlPolicyReconcileLoop creates a new mock instance.
func NewMockMulticlusterAccessControlPolicyReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterAccessControlPolicyReconcileLoop {
	mock := &MockMulticlusterAccessControlPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAccessControlPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterAccessControlPolicyReconcileLoop) EXPECT() *MockMulticlusterAccessControlPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterAccessControlPolicyReconciler mocks base method.
func (m *MockMulticlusterAccessControlPolicyReconcileLoop) AddMulticlusterAccessControlPolicyReconciler(ctx context.Context, rec controller.MulticlusterAccessControlPolicyReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterAccessControlPolicyReconciler", varargs...)
}

// AddMulticlusterAccessControlPolicyReconciler indicates an expected call of AddMulticlusterAccessControlPolicyReconciler.
func (mr *MockMulticlusterAccessControlPolicyReconcileLoopMockRecorder) AddMulticlusterAccessControlPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterAccessControlPolicyReconciler", reflect.TypeOf((*MockMulticlusterAccessControlPolicyReconcileLoop)(nil).AddMulticlusterAccessControlPolicyReconciler), varargs...)
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
func (m *MockMulticlusterVirtualMeshReconciler) ReconcileVirtualMesh(clusterName string, obj *v1alpha1.VirtualMesh) (reconcile.Result, error) {
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

// MockMulticlusterFailoverServiceReconciler is a mock of MulticlusterFailoverServiceReconciler interface.
type MockMulticlusterFailoverServiceReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFailoverServiceReconcilerMockRecorder
}

// MockMulticlusterFailoverServiceReconcilerMockRecorder is the mock recorder for MockMulticlusterFailoverServiceReconciler.
type MockMulticlusterFailoverServiceReconcilerMockRecorder struct {
	mock *MockMulticlusterFailoverServiceReconciler
}

// NewMockMulticlusterFailoverServiceReconciler creates a new mock instance.
func NewMockMulticlusterFailoverServiceReconciler(ctrl *gomock.Controller) *MockMulticlusterFailoverServiceReconciler {
	mock := &MockMulticlusterFailoverServiceReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFailoverServiceReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFailoverServiceReconciler) EXPECT() *MockMulticlusterFailoverServiceReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFailoverService mocks base method.
func (m *MockMulticlusterFailoverServiceReconciler) ReconcileFailoverService(clusterName string, obj *v1alpha1.FailoverService) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFailoverService", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFailoverService indicates an expected call of ReconcileFailoverService.
func (mr *MockMulticlusterFailoverServiceReconcilerMockRecorder) ReconcileFailoverService(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFailoverService", reflect.TypeOf((*MockMulticlusterFailoverServiceReconciler)(nil).ReconcileFailoverService), clusterName, obj)
}

// MockMulticlusterFailoverServiceDeletionReconciler is a mock of MulticlusterFailoverServiceDeletionReconciler interface.
type MockMulticlusterFailoverServiceDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFailoverServiceDeletionReconcilerMockRecorder
}

// MockMulticlusterFailoverServiceDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterFailoverServiceDeletionReconciler.
type MockMulticlusterFailoverServiceDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterFailoverServiceDeletionReconciler
}

// NewMockMulticlusterFailoverServiceDeletionReconciler creates a new mock instance.
func NewMockMulticlusterFailoverServiceDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterFailoverServiceDeletionReconciler {
	mock := &MockMulticlusterFailoverServiceDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFailoverServiceDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFailoverServiceDeletionReconciler) EXPECT() *MockMulticlusterFailoverServiceDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFailoverServiceDeletion mocks base method.
func (m *MockMulticlusterFailoverServiceDeletionReconciler) ReconcileFailoverServiceDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFailoverServiceDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFailoverServiceDeletion indicates an expected call of ReconcileFailoverServiceDeletion.
func (mr *MockMulticlusterFailoverServiceDeletionReconcilerMockRecorder) ReconcileFailoverServiceDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFailoverServiceDeletion", reflect.TypeOf((*MockMulticlusterFailoverServiceDeletionReconciler)(nil).ReconcileFailoverServiceDeletion), clusterName, req)
}

// MockMulticlusterFailoverServiceReconcileLoop is a mock of MulticlusterFailoverServiceReconcileLoop interface.
type MockMulticlusterFailoverServiceReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFailoverServiceReconcileLoopMockRecorder
}

// MockMulticlusterFailoverServiceReconcileLoopMockRecorder is the mock recorder for MockMulticlusterFailoverServiceReconcileLoop.
type MockMulticlusterFailoverServiceReconcileLoopMockRecorder struct {
	mock *MockMulticlusterFailoverServiceReconcileLoop
}

// NewMockMulticlusterFailoverServiceReconcileLoop creates a new mock instance.
func NewMockMulticlusterFailoverServiceReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterFailoverServiceReconcileLoop {
	mock := &MockMulticlusterFailoverServiceReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFailoverServiceReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterFailoverServiceReconcileLoop) EXPECT() *MockMulticlusterFailoverServiceReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterFailoverServiceReconciler mocks base method.
func (m *MockMulticlusterFailoverServiceReconcileLoop) AddMulticlusterFailoverServiceReconciler(ctx context.Context, rec controller.MulticlusterFailoverServiceReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterFailoverServiceReconciler", varargs...)
}

// AddMulticlusterFailoverServiceReconciler indicates an expected call of AddMulticlusterFailoverServiceReconciler.
func (mr *MockMulticlusterFailoverServiceReconcileLoopMockRecorder) AddMulticlusterFailoverServiceReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterFailoverServiceReconciler", reflect.TypeOf((*MockMulticlusterFailoverServiceReconcileLoop)(nil).AddMulticlusterFailoverServiceReconciler), varargs...)
}
