// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_networking_snapshot is a generated GoMock package.
package mock_networking_snapshot

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	v1alpha10 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	networking_snapshot "github.com/solo-io/service-mesh-hub/pkg/networking-snapshot"
)

// MockMeshNetworkingSnapshotValidator is a mock of MeshNetworkingSnapshotValidator interface.
type MockMeshNetworkingSnapshotValidator struct {
	ctrl     *gomock.Controller
	recorder *MockMeshNetworkingSnapshotValidatorMockRecorder
}

// MockMeshNetworkingSnapshotValidatorMockRecorder is the mock recorder for MockMeshNetworkingSnapshotValidator.
type MockMeshNetworkingSnapshotValidatorMockRecorder struct {
	mock *MockMeshNetworkingSnapshotValidator
}

// NewMockMeshNetworkingSnapshotValidator creates a new mock instance.
func NewMockMeshNetworkingSnapshotValidator(ctrl *gomock.Controller) *MockMeshNetworkingSnapshotValidator {
	mock := &MockMeshNetworkingSnapshotValidator{ctrl: ctrl}
	mock.recorder = &MockMeshNetworkingSnapshotValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshNetworkingSnapshotValidator) EXPECT() *MockMeshNetworkingSnapshotValidatorMockRecorder {
	return m.recorder
}

// ValidateVirtualMeshUpsert mocks base method.
func (m *MockMeshNetworkingSnapshotValidator) ValidateVirtualMeshUpsert(ctx context.Context, obj *v1alpha10.VirtualMesh, snapshot *networking_snapshot.MeshNetworkingSnapshot) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateVirtualMeshUpsert", ctx, obj, snapshot)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidateVirtualMeshUpsert indicates an expected call of ValidateVirtualMeshUpsert.
func (mr *MockMeshNetworkingSnapshotValidatorMockRecorder) ValidateVirtualMeshUpsert(ctx, obj, snapshot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateVirtualMeshUpsert", reflect.TypeOf((*MockMeshNetworkingSnapshotValidator)(nil).ValidateVirtualMeshUpsert), ctx, obj, snapshot)
}

// ValidateVirtualMeshDelete mocks base method.
func (m *MockMeshNetworkingSnapshotValidator) ValidateVirtualMeshDelete(ctx context.Context, obj *v1alpha10.VirtualMesh, snapshot *networking_snapshot.MeshNetworkingSnapshot) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateVirtualMeshDelete", ctx, obj, snapshot)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidateVirtualMeshDelete indicates an expected call of ValidateVirtualMeshDelete.
func (mr *MockMeshNetworkingSnapshotValidatorMockRecorder) ValidateVirtualMeshDelete(ctx, obj, snapshot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateVirtualMeshDelete", reflect.TypeOf((*MockMeshNetworkingSnapshotValidator)(nil).ValidateVirtualMeshDelete), ctx, obj, snapshot)
}

// ValidateMeshServiceUpsert mocks base method.
func (m *MockMeshNetworkingSnapshotValidator) ValidateMeshServiceUpsert(ctx context.Context, obj *v1alpha1.MeshService, snapshot *networking_snapshot.MeshNetworkingSnapshot) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateMeshServiceUpsert", ctx, obj, snapshot)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidateMeshServiceUpsert indicates an expected call of ValidateMeshServiceUpsert.
func (mr *MockMeshNetworkingSnapshotValidatorMockRecorder) ValidateMeshServiceUpsert(ctx, obj, snapshot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateMeshServiceUpsert", reflect.TypeOf((*MockMeshNetworkingSnapshotValidator)(nil).ValidateMeshServiceUpsert), ctx, obj, snapshot)
}

// ValidateMeshServiceDelete mocks base method.
func (m *MockMeshNetworkingSnapshotValidator) ValidateMeshServiceDelete(ctx context.Context, obj *v1alpha1.MeshService, snapshot *networking_snapshot.MeshNetworkingSnapshot) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateMeshServiceDelete", ctx, obj, snapshot)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidateMeshServiceDelete indicates an expected call of ValidateMeshServiceDelete.
func (mr *MockMeshNetworkingSnapshotValidatorMockRecorder) ValidateMeshServiceDelete(ctx, obj, snapshot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateMeshServiceDelete", reflect.TypeOf((*MockMeshNetworkingSnapshotValidator)(nil).ValidateMeshServiceDelete), ctx, obj, snapshot)
}

// ValidateMeshWorkloadUpsert mocks base method.
func (m *MockMeshNetworkingSnapshotValidator) ValidateMeshWorkloadUpsert(ctx context.Context, obj *v1alpha1.MeshWorkload, snapshot *networking_snapshot.MeshNetworkingSnapshot) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateMeshWorkloadUpsert", ctx, obj, snapshot)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidateMeshWorkloadUpsert indicates an expected call of ValidateMeshWorkloadUpsert.
func (mr *MockMeshNetworkingSnapshotValidatorMockRecorder) ValidateMeshWorkloadUpsert(ctx, obj, snapshot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateMeshWorkloadUpsert", reflect.TypeOf((*MockMeshNetworkingSnapshotValidator)(nil).ValidateMeshWorkloadUpsert), ctx, obj, snapshot)
}

// ValidateMeshWorkloadDelete mocks base method.
func (m *MockMeshNetworkingSnapshotValidator) ValidateMeshWorkloadDelete(ctx context.Context, obj *v1alpha1.MeshWorkload, snapshot *networking_snapshot.MeshNetworkingSnapshot) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateMeshWorkloadDelete", ctx, obj, snapshot)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidateMeshWorkloadDelete indicates an expected call of ValidateMeshWorkloadDelete.
func (mr *MockMeshNetworkingSnapshotValidatorMockRecorder) ValidateMeshWorkloadDelete(ctx, obj, snapshot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateMeshWorkloadDelete", reflect.TypeOf((*MockMeshNetworkingSnapshotValidator)(nil).ValidateMeshWorkloadDelete), ctx, obj, snapshot)
}

// MockMeshNetworkingSnapshotGenerator is a mock of MeshNetworkingSnapshotGenerator interface.
type MockMeshNetworkingSnapshotGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockMeshNetworkingSnapshotGeneratorMockRecorder
}

// MockMeshNetworkingSnapshotGeneratorMockRecorder is the mock recorder for MockMeshNetworkingSnapshotGenerator.
type MockMeshNetworkingSnapshotGeneratorMockRecorder struct {
	mock *MockMeshNetworkingSnapshotGenerator
}

// NewMockMeshNetworkingSnapshotGenerator creates a new mock instance.
func NewMockMeshNetworkingSnapshotGenerator(ctrl *gomock.Controller) *MockMeshNetworkingSnapshotGenerator {
	mock := &MockMeshNetworkingSnapshotGenerator{ctrl: ctrl}
	mock.recorder = &MockMeshNetworkingSnapshotGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshNetworkingSnapshotGenerator) EXPECT() *MockMeshNetworkingSnapshotGeneratorMockRecorder {
	return m.recorder
}

// RegisterListener mocks base method.
func (m *MockMeshNetworkingSnapshotGenerator) RegisterListener(arg0 networking_snapshot.MeshNetworkingSnapshotListener) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterListener", arg0)
}

// RegisterListener indicates an expected call of RegisterListener.
func (mr *MockMeshNetworkingSnapshotGeneratorMockRecorder) RegisterListener(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterListener", reflect.TypeOf((*MockMeshNetworkingSnapshotGenerator)(nil).RegisterListener), arg0)
}

// StartPushingSnapshots mocks base method.
func (m *MockMeshNetworkingSnapshotGenerator) StartPushingSnapshots(ctx context.Context, snapshotFrequency time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartPushingSnapshots", ctx, snapshotFrequency)
}

// StartPushingSnapshots indicates an expected call of StartPushingSnapshots.
func (mr *MockMeshNetworkingSnapshotGeneratorMockRecorder) StartPushingSnapshots(ctx, snapshotFrequency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPushingSnapshots", reflect.TypeOf((*MockMeshNetworkingSnapshotGenerator)(nil).StartPushingSnapshots), ctx, snapshotFrequency)
}

// MockMeshNetworkingSnapshotListener is a mock of MeshNetworkingSnapshotListener interface.
type MockMeshNetworkingSnapshotListener struct {
	ctrl     *gomock.Controller
	recorder *MockMeshNetworkingSnapshotListenerMockRecorder
}

// MockMeshNetworkingSnapshotListenerMockRecorder is the mock recorder for MockMeshNetworkingSnapshotListener.
type MockMeshNetworkingSnapshotListenerMockRecorder struct {
	mock *MockMeshNetworkingSnapshotListener
}

// NewMockMeshNetworkingSnapshotListener creates a new mock instance.
func NewMockMeshNetworkingSnapshotListener(ctrl *gomock.Controller) *MockMeshNetworkingSnapshotListener {
	mock := &MockMeshNetworkingSnapshotListener{ctrl: ctrl}
	mock.recorder = &MockMeshNetworkingSnapshotListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshNetworkingSnapshotListener) EXPECT() *MockMeshNetworkingSnapshotListenerMockRecorder {
	return m.recorder
}

// Sync mocks base method.
func (m *MockMeshNetworkingSnapshotListener) Sync(arg0 context.Context, arg1 *networking_snapshot.MeshNetworkingSnapshot) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Sync", arg0, arg1)
}

// Sync indicates an expected call of Sync.
func (mr *MockMeshNetworkingSnapshotListenerMockRecorder) Sync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockMeshNetworkingSnapshotListener)(nil).Sync), arg0, arg1)
}
