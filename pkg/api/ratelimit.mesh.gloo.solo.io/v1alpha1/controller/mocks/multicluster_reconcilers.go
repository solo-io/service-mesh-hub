// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/ratelimit.mesh.gloo.solo.io/v1alpha1"
	controller "github.com/solo-io/gloo-mesh/pkg/api/ratelimit.mesh.gloo.solo.io/v1alpha1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterRateLimitConfigReconciler is a mock of MulticlusterRateLimitConfigReconciler interface.
type MockMulticlusterRateLimitConfigReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRateLimitConfigReconcilerMockRecorder
}

// MockMulticlusterRateLimitConfigReconcilerMockRecorder is the mock recorder for MockMulticlusterRateLimitConfigReconciler.
type MockMulticlusterRateLimitConfigReconcilerMockRecorder struct {
	mock *MockMulticlusterRateLimitConfigReconciler
}

// NewMockMulticlusterRateLimitConfigReconciler creates a new mock instance.
func NewMockMulticlusterRateLimitConfigReconciler(ctrl *gomock.Controller) *MockMulticlusterRateLimitConfigReconciler {
	mock := &MockMulticlusterRateLimitConfigReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRateLimitConfigReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterRateLimitConfigReconciler) EXPECT() *MockMulticlusterRateLimitConfigReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRateLimitConfig mocks base method.
func (m *MockMulticlusterRateLimitConfigReconciler) ReconcileRateLimitConfig(clusterName string, obj *v1alpha1.RateLimitConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRateLimitConfig", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRateLimitConfig indicates an expected call of ReconcileRateLimitConfig.
func (mr *MockMulticlusterRateLimitConfigReconcilerMockRecorder) ReconcileRateLimitConfig(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRateLimitConfig", reflect.TypeOf((*MockMulticlusterRateLimitConfigReconciler)(nil).ReconcileRateLimitConfig), clusterName, obj)
}

// MockMulticlusterRateLimitConfigDeletionReconciler is a mock of MulticlusterRateLimitConfigDeletionReconciler interface.
type MockMulticlusterRateLimitConfigDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRateLimitConfigDeletionReconcilerMockRecorder
}

// MockMulticlusterRateLimitConfigDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterRateLimitConfigDeletionReconciler.
type MockMulticlusterRateLimitConfigDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterRateLimitConfigDeletionReconciler
}

// NewMockMulticlusterRateLimitConfigDeletionReconciler creates a new mock instance.
func NewMockMulticlusterRateLimitConfigDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterRateLimitConfigDeletionReconciler {
	mock := &MockMulticlusterRateLimitConfigDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRateLimitConfigDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterRateLimitConfigDeletionReconciler) EXPECT() *MockMulticlusterRateLimitConfigDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRateLimitConfigDeletion mocks base method.
func (m *MockMulticlusterRateLimitConfigDeletionReconciler) ReconcileRateLimitConfigDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRateLimitConfigDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileRateLimitConfigDeletion indicates an expected call of ReconcileRateLimitConfigDeletion.
func (mr *MockMulticlusterRateLimitConfigDeletionReconcilerMockRecorder) ReconcileRateLimitConfigDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRateLimitConfigDeletion", reflect.TypeOf((*MockMulticlusterRateLimitConfigDeletionReconciler)(nil).ReconcileRateLimitConfigDeletion), clusterName, req)
}

// MockMulticlusterRateLimitConfigReconcileLoop is a mock of MulticlusterRateLimitConfigReconcileLoop interface.
type MockMulticlusterRateLimitConfigReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRateLimitConfigReconcileLoopMockRecorder
}

// MockMulticlusterRateLimitConfigReconcileLoopMockRecorder is the mock recorder for MockMulticlusterRateLimitConfigReconcileLoop.
type MockMulticlusterRateLimitConfigReconcileLoopMockRecorder struct {
	mock *MockMulticlusterRateLimitConfigReconcileLoop
}

// NewMockMulticlusterRateLimitConfigReconcileLoop creates a new mock instance.
func NewMockMulticlusterRateLimitConfigReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterRateLimitConfigReconcileLoop {
	mock := &MockMulticlusterRateLimitConfigReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRateLimitConfigReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterRateLimitConfigReconcileLoop) EXPECT() *MockMulticlusterRateLimitConfigReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterRateLimitConfigReconciler mocks base method.
func (m *MockMulticlusterRateLimitConfigReconcileLoop) AddMulticlusterRateLimitConfigReconciler(ctx context.Context, rec controller.MulticlusterRateLimitConfigReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterRateLimitConfigReconciler", varargs...)
}

// AddMulticlusterRateLimitConfigReconciler indicates an expected call of AddMulticlusterRateLimitConfigReconciler.
func (mr *MockMulticlusterRateLimitConfigReconcileLoopMockRecorder) AddMulticlusterRateLimitConfigReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterRateLimitConfigReconciler", reflect.TypeOf((*MockMulticlusterRateLimitConfigReconcileLoop)(nil).AddMulticlusterRateLimitConfigReconciler), varargs...)
}