// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	controller "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockDestinationReconciler is a mock of DestinationReconciler interface.
type MockDestinationReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationReconcilerMockRecorder
}

// MockDestinationReconcilerMockRecorder is the mock recorder for MockDestinationReconciler.
type MockDestinationReconcilerMockRecorder struct {
	mock *MockDestinationReconciler
}

// NewMockDestinationReconciler creates a new mock instance.
func NewMockDestinationReconciler(ctrl *gomock.Controller) *MockDestinationReconciler {
	mock := &MockDestinationReconciler{ctrl: ctrl}
	mock.recorder = &MockDestinationReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationReconciler) EXPECT() *MockDestinationReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDestination mocks base method.
func (m *MockDestinationReconciler) ReconcileDestination(obj *v1.Destination) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDestination", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDestination indicates an expected call of ReconcileDestination.
func (mr *MockDestinationReconcilerMockRecorder) ReconcileDestination(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDestination", reflect.TypeOf((*MockDestinationReconciler)(nil).ReconcileDestination), obj)
}

// MockDestinationDeletionReconciler is a mock of DestinationDeletionReconciler interface.
type MockDestinationDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationDeletionReconcilerMockRecorder
}

// MockDestinationDeletionReconcilerMockRecorder is the mock recorder for MockDestinationDeletionReconciler.
type MockDestinationDeletionReconcilerMockRecorder struct {
	mock *MockDestinationDeletionReconciler
}

// NewMockDestinationDeletionReconciler creates a new mock instance.
func NewMockDestinationDeletionReconciler(ctrl *gomock.Controller) *MockDestinationDeletionReconciler {
	mock := &MockDestinationDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockDestinationDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationDeletionReconciler) EXPECT() *MockDestinationDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDestinationDeletion mocks base method.
func (m *MockDestinationDeletionReconciler) ReconcileDestinationDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDestinationDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileDestinationDeletion indicates an expected call of ReconcileDestinationDeletion.
func (mr *MockDestinationDeletionReconcilerMockRecorder) ReconcileDestinationDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDestinationDeletion", reflect.TypeOf((*MockDestinationDeletionReconciler)(nil).ReconcileDestinationDeletion), req)
}

// MockDestinationFinalizer is a mock of DestinationFinalizer interface.
type MockDestinationFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationFinalizerMockRecorder
}

// MockDestinationFinalizerMockRecorder is the mock recorder for MockDestinationFinalizer.
type MockDestinationFinalizerMockRecorder struct {
	mock *MockDestinationFinalizer
}

// NewMockDestinationFinalizer creates a new mock instance.
func NewMockDestinationFinalizer(ctrl *gomock.Controller) *MockDestinationFinalizer {
	mock := &MockDestinationFinalizer{ctrl: ctrl}
	mock.recorder = &MockDestinationFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationFinalizer) EXPECT() *MockDestinationFinalizerMockRecorder {
	return m.recorder
}

// DestinationFinalizerName mocks base method.
func (m *MockDestinationFinalizer) DestinationFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestinationFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DestinationFinalizerName indicates an expected call of DestinationFinalizerName.
func (mr *MockDestinationFinalizerMockRecorder) DestinationFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestinationFinalizerName", reflect.TypeOf((*MockDestinationFinalizer)(nil).DestinationFinalizerName))
}

// FinalizeDestination mocks base method.
func (m *MockDestinationFinalizer) FinalizeDestination(obj *v1.Destination) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeDestination", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeDestination indicates an expected call of FinalizeDestination.
func (mr *MockDestinationFinalizerMockRecorder) FinalizeDestination(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeDestination", reflect.TypeOf((*MockDestinationFinalizer)(nil).FinalizeDestination), obj)
}

// ReconcileDestination mocks base method.
func (m *MockDestinationFinalizer) ReconcileDestination(obj *v1.Destination) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDestination", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDestination indicates an expected call of ReconcileDestination.
func (mr *MockDestinationFinalizerMockRecorder) ReconcileDestination(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDestination", reflect.TypeOf((*MockDestinationFinalizer)(nil).ReconcileDestination), obj)
}

// MockDestinationReconcileLoop is a mock of DestinationReconcileLoop interface.
type MockDestinationReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationReconcileLoopMockRecorder
}

// MockDestinationReconcileLoopMockRecorder is the mock recorder for MockDestinationReconcileLoop.
type MockDestinationReconcileLoopMockRecorder struct {
	mock *MockDestinationReconcileLoop
}

// NewMockDestinationReconcileLoop creates a new mock instance.
func NewMockDestinationReconcileLoop(ctrl *gomock.Controller) *MockDestinationReconcileLoop {
	mock := &MockDestinationReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockDestinationReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationReconcileLoop) EXPECT() *MockDestinationReconcileLoopMockRecorder {
	return m.recorder
}

// RunDestinationReconciler mocks base method.
func (m *MockDestinationReconcileLoop) RunDestinationReconciler(ctx context.Context, rec controller.DestinationReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunDestinationReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunDestinationReconciler indicates an expected call of RunDestinationReconciler.
func (mr *MockDestinationReconcileLoopMockRecorder) RunDestinationReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunDestinationReconciler", reflect.TypeOf((*MockDestinationReconcileLoop)(nil).RunDestinationReconciler), varargs...)
}

// MockWorkloadReconciler is a mock of WorkloadReconciler interface.
type MockWorkloadReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockWorkloadReconcilerMockRecorder
}

// MockWorkloadReconcilerMockRecorder is the mock recorder for MockWorkloadReconciler.
type MockWorkloadReconcilerMockRecorder struct {
	mock *MockWorkloadReconciler
}

// NewMockWorkloadReconciler creates a new mock instance.
func NewMockWorkloadReconciler(ctrl *gomock.Controller) *MockWorkloadReconciler {
	mock := &MockWorkloadReconciler{ctrl: ctrl}
	mock.recorder = &MockWorkloadReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkloadReconciler) EXPECT() *MockWorkloadReconcilerMockRecorder {
	return m.recorder
}

// ReconcileWorkload mocks base method.
func (m *MockWorkloadReconciler) ReconcileWorkload(obj *v1.Workload) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileWorkload", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileWorkload indicates an expected call of ReconcileWorkload.
func (mr *MockWorkloadReconcilerMockRecorder) ReconcileWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileWorkload", reflect.TypeOf((*MockWorkloadReconciler)(nil).ReconcileWorkload), obj)
}

// MockWorkloadDeletionReconciler is a mock of WorkloadDeletionReconciler interface.
type MockWorkloadDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockWorkloadDeletionReconcilerMockRecorder
}

// MockWorkloadDeletionReconcilerMockRecorder is the mock recorder for MockWorkloadDeletionReconciler.
type MockWorkloadDeletionReconcilerMockRecorder struct {
	mock *MockWorkloadDeletionReconciler
}

// NewMockWorkloadDeletionReconciler creates a new mock instance.
func NewMockWorkloadDeletionReconciler(ctrl *gomock.Controller) *MockWorkloadDeletionReconciler {
	mock := &MockWorkloadDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockWorkloadDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkloadDeletionReconciler) EXPECT() *MockWorkloadDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileWorkloadDeletion mocks base method.
func (m *MockWorkloadDeletionReconciler) ReconcileWorkloadDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileWorkloadDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileWorkloadDeletion indicates an expected call of ReconcileWorkloadDeletion.
func (mr *MockWorkloadDeletionReconcilerMockRecorder) ReconcileWorkloadDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileWorkloadDeletion", reflect.TypeOf((*MockWorkloadDeletionReconciler)(nil).ReconcileWorkloadDeletion), req)
}

// MockWorkloadFinalizer is a mock of WorkloadFinalizer interface.
type MockWorkloadFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockWorkloadFinalizerMockRecorder
}

// MockWorkloadFinalizerMockRecorder is the mock recorder for MockWorkloadFinalizer.
type MockWorkloadFinalizerMockRecorder struct {
	mock *MockWorkloadFinalizer
}

// NewMockWorkloadFinalizer creates a new mock instance.
func NewMockWorkloadFinalizer(ctrl *gomock.Controller) *MockWorkloadFinalizer {
	mock := &MockWorkloadFinalizer{ctrl: ctrl}
	mock.recorder = &MockWorkloadFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkloadFinalizer) EXPECT() *MockWorkloadFinalizerMockRecorder {
	return m.recorder
}

// FinalizeWorkload mocks base method.
func (m *MockWorkloadFinalizer) FinalizeWorkload(obj *v1.Workload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeWorkload", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeWorkload indicates an expected call of FinalizeWorkload.
func (mr *MockWorkloadFinalizerMockRecorder) FinalizeWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeWorkload", reflect.TypeOf((*MockWorkloadFinalizer)(nil).FinalizeWorkload), obj)
}

// ReconcileWorkload mocks base method.
func (m *MockWorkloadFinalizer) ReconcileWorkload(obj *v1.Workload) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileWorkload", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileWorkload indicates an expected call of ReconcileWorkload.
func (mr *MockWorkloadFinalizerMockRecorder) ReconcileWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileWorkload", reflect.TypeOf((*MockWorkloadFinalizer)(nil).ReconcileWorkload), obj)
}

// WorkloadFinalizerName mocks base method.
func (m *MockWorkloadFinalizer) WorkloadFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkloadFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// WorkloadFinalizerName indicates an expected call of WorkloadFinalizerName.
func (mr *MockWorkloadFinalizerMockRecorder) WorkloadFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkloadFinalizerName", reflect.TypeOf((*MockWorkloadFinalizer)(nil).WorkloadFinalizerName))
}

// MockWorkloadReconcileLoop is a mock of WorkloadReconcileLoop interface.
type MockWorkloadReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockWorkloadReconcileLoopMockRecorder
}

// MockWorkloadReconcileLoopMockRecorder is the mock recorder for MockWorkloadReconcileLoop.
type MockWorkloadReconcileLoopMockRecorder struct {
	mock *MockWorkloadReconcileLoop
}

// NewMockWorkloadReconcileLoop creates a new mock instance.
func NewMockWorkloadReconcileLoop(ctrl *gomock.Controller) *MockWorkloadReconcileLoop {
	mock := &MockWorkloadReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockWorkloadReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkloadReconcileLoop) EXPECT() *MockWorkloadReconcileLoopMockRecorder {
	return m.recorder
}

// RunWorkloadReconciler mocks base method.
func (m *MockWorkloadReconcileLoop) RunWorkloadReconciler(ctx context.Context, rec controller.WorkloadReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunWorkloadReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunWorkloadReconciler indicates an expected call of RunWorkloadReconciler.
func (mr *MockWorkloadReconcileLoopMockRecorder) RunWorkloadReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunWorkloadReconciler", reflect.TypeOf((*MockWorkloadReconcileLoop)(nil).RunWorkloadReconciler), varargs...)
}

// MockMeshReconciler is a mock of MeshReconciler interface.
type MockMeshReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMeshReconcilerMockRecorder
}

// MockMeshReconcilerMockRecorder is the mock recorder for MockMeshReconciler.
type MockMeshReconcilerMockRecorder struct {
	mock *MockMeshReconciler
}

// NewMockMeshReconciler creates a new mock instance.
func NewMockMeshReconciler(ctrl *gomock.Controller) *MockMeshReconciler {
	mock := &MockMeshReconciler{ctrl: ctrl}
	mock.recorder = &MockMeshReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshReconciler) EXPECT() *MockMeshReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMesh mocks base method.
func (m *MockMeshReconciler) ReconcileMesh(obj *v1.Mesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMesh", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMesh indicates an expected call of ReconcileMesh.
func (mr *MockMeshReconcilerMockRecorder) ReconcileMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMesh", reflect.TypeOf((*MockMeshReconciler)(nil).ReconcileMesh), obj)
}

// MockMeshDeletionReconciler is a mock of MeshDeletionReconciler interface.
type MockMeshDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMeshDeletionReconcilerMockRecorder
}

// MockMeshDeletionReconcilerMockRecorder is the mock recorder for MockMeshDeletionReconciler.
type MockMeshDeletionReconcilerMockRecorder struct {
	mock *MockMeshDeletionReconciler
}

// NewMockMeshDeletionReconciler creates a new mock instance.
func NewMockMeshDeletionReconciler(ctrl *gomock.Controller) *MockMeshDeletionReconciler {
	mock := &MockMeshDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMeshDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshDeletionReconciler) EXPECT() *MockMeshDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMeshDeletion mocks base method.
func (m *MockMeshDeletionReconciler) ReconcileMeshDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMeshDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileMeshDeletion indicates an expected call of ReconcileMeshDeletion.
func (mr *MockMeshDeletionReconcilerMockRecorder) ReconcileMeshDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMeshDeletion", reflect.TypeOf((*MockMeshDeletionReconciler)(nil).ReconcileMeshDeletion), req)
}

// MockMeshFinalizer is a mock of MeshFinalizer interface.
type MockMeshFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockMeshFinalizerMockRecorder
}

// MockMeshFinalizerMockRecorder is the mock recorder for MockMeshFinalizer.
type MockMeshFinalizerMockRecorder struct {
	mock *MockMeshFinalizer
}

// NewMockMeshFinalizer creates a new mock instance.
func NewMockMeshFinalizer(ctrl *gomock.Controller) *MockMeshFinalizer {
	mock := &MockMeshFinalizer{ctrl: ctrl}
	mock.recorder = &MockMeshFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshFinalizer) EXPECT() *MockMeshFinalizerMockRecorder {
	return m.recorder
}

// FinalizeMesh mocks base method.
func (m *MockMeshFinalizer) FinalizeMesh(obj *v1.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeMesh", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeMesh indicates an expected call of FinalizeMesh.
func (mr *MockMeshFinalizerMockRecorder) FinalizeMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeMesh", reflect.TypeOf((*MockMeshFinalizer)(nil).FinalizeMesh), obj)
}

// MeshFinalizerName mocks base method.
func (m *MockMeshFinalizer) MeshFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MeshFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// MeshFinalizerName indicates an expected call of MeshFinalizerName.
func (mr *MockMeshFinalizerMockRecorder) MeshFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MeshFinalizerName", reflect.TypeOf((*MockMeshFinalizer)(nil).MeshFinalizerName))
}

// ReconcileMesh mocks base method.
func (m *MockMeshFinalizer) ReconcileMesh(obj *v1.Mesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMesh", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMesh indicates an expected call of ReconcileMesh.
func (mr *MockMeshFinalizerMockRecorder) ReconcileMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMesh", reflect.TypeOf((*MockMeshFinalizer)(nil).ReconcileMesh), obj)
}

// MockMeshReconcileLoop is a mock of MeshReconcileLoop interface.
type MockMeshReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMeshReconcileLoopMockRecorder
}

// MockMeshReconcileLoopMockRecorder is the mock recorder for MockMeshReconcileLoop.
type MockMeshReconcileLoopMockRecorder struct {
	mock *MockMeshReconcileLoop
}

// NewMockMeshReconcileLoop creates a new mock instance.
func NewMockMeshReconcileLoop(ctrl *gomock.Controller) *MockMeshReconcileLoop {
	mock := &MockMeshReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMeshReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshReconcileLoop) EXPECT() *MockMeshReconcileLoopMockRecorder {
	return m.recorder
}

// RunMeshReconciler mocks base method.
func (m *MockMeshReconcileLoop) RunMeshReconciler(ctx context.Context, rec controller.MeshReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunMeshReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunMeshReconciler indicates an expected call of RunMeshReconciler.
func (mr *MockMeshReconcileLoopMockRecorder) RunMeshReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunMeshReconciler", reflect.TypeOf((*MockMeshReconcileLoop)(nil).RunMeshReconciler), varargs...)
}
