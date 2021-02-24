// Code generated by MockGen. DO NOT EDIT.
// Source: ./local_snapshot.go

// Package mock_input is a generated GoMock package.
package mock_input

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1alpha2sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2/sets"
	v1alpha1sets "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1alpha1/sets"
	input "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	v1alpha2sets0 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2/sets"
	v1alpha1sets0 "github.com/solo-io/gloo-mesh/pkg/api/observability.enterprise.mesh.gloo.solo.io/v1alpha1/sets"
	v1alpha2sets1 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2/sets"
	v1alpha1sets1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/sets"
	multicluster "github.com/solo-io/skv2/pkg/multicluster"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockLocalSnapshot is a mock of LocalSnapshot interface
type MockLocalSnapshot struct {
	ctrl     *gomock.Controller
	recorder *MockLocalSnapshotMockRecorder
}

// MockLocalSnapshotMockRecorder is the mock recorder for MockLocalSnapshot
type MockLocalSnapshotMockRecorder struct {
	mock *MockLocalSnapshot
}

// NewMockLocalSnapshot creates a new mock instance
func NewMockLocalSnapshot(ctrl *gomock.Controller) *MockLocalSnapshot {
	mock := &MockLocalSnapshot{ctrl: ctrl}
	mock.recorder = &MockLocalSnapshotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocalSnapshot) EXPECT() *MockLocalSnapshotMockRecorder {
	return m.recorder
}

// Settings mocks base method
func (m *MockLocalSnapshot) Settings() v1alpha2sets1.SettingsSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Settings")
	ret0, _ := ret[0].(v1alpha2sets1.SettingsSet)
	return ret0
}

// Settings indicates an expected call of Settings
func (mr *MockLocalSnapshotMockRecorder) Settings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Settings", reflect.TypeOf((*MockLocalSnapshot)(nil).Settings))
}

// TrafficTargets mocks base method
func (m *MockLocalSnapshot) TrafficTargets() v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficTargets")
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// TrafficTargets indicates an expected call of TrafficTargets
func (mr *MockLocalSnapshotMockRecorder) TrafficTargets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficTargets", reflect.TypeOf((*MockLocalSnapshot)(nil).TrafficTargets))
}

// Workloads mocks base method
func (m *MockLocalSnapshot) Workloads() v1alpha2sets.WorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Workloads")
	ret0, _ := ret[0].(v1alpha2sets.WorkloadSet)
	return ret0
}

// Workloads indicates an expected call of Workloads
func (mr *MockLocalSnapshotMockRecorder) Workloads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Workloads", reflect.TypeOf((*MockLocalSnapshot)(nil).Workloads))
}

// Meshes mocks base method
func (m *MockLocalSnapshot) Meshes() v1alpha2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meshes")
	ret0, _ := ret[0].(v1alpha2sets.MeshSet)
	return ret0
}

// Meshes indicates an expected call of Meshes
func (mr *MockLocalSnapshotMockRecorder) Meshes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meshes", reflect.TypeOf((*MockLocalSnapshot)(nil).Meshes))
}

// TrafficPolicies mocks base method
func (m *MockLocalSnapshot) TrafficPolicies() v1alpha2sets0.TrafficPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficPolicies")
	ret0, _ := ret[0].(v1alpha2sets0.TrafficPolicySet)
	return ret0
}

// TrafficPolicies indicates an expected call of TrafficPolicies
func (mr *MockLocalSnapshotMockRecorder) TrafficPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficPolicies", reflect.TypeOf((*MockLocalSnapshot)(nil).TrafficPolicies))
}

// AccessPolicies mocks base method
func (m *MockLocalSnapshot) AccessPolicies() v1alpha2sets0.AccessPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessPolicies")
	ret0, _ := ret[0].(v1alpha2sets0.AccessPolicySet)
	return ret0
}

// AccessPolicies indicates an expected call of AccessPolicies
func (mr *MockLocalSnapshotMockRecorder) AccessPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessPolicies", reflect.TypeOf((*MockLocalSnapshot)(nil).AccessPolicies))
}

// VirtualMeshes mocks base method
func (m *MockLocalSnapshot) VirtualMeshes() v1alpha2sets0.VirtualMeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualMeshes")
	ret0, _ := ret[0].(v1alpha2sets0.VirtualMeshSet)
	return ret0
}

// VirtualMeshes indicates an expected call of VirtualMeshes
func (mr *MockLocalSnapshotMockRecorder) VirtualMeshes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualMeshes", reflect.TypeOf((*MockLocalSnapshot)(nil).VirtualMeshes))
}

// FailoverServices mocks base method
func (m *MockLocalSnapshot) FailoverServices() v1alpha2sets0.FailoverServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailoverServices")
	ret0, _ := ret[0].(v1alpha2sets0.FailoverServiceSet)
	return ret0
}

// FailoverServices indicates an expected call of FailoverServices
func (mr *MockLocalSnapshotMockRecorder) FailoverServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailoverServices", reflect.TypeOf((*MockLocalSnapshot)(nil).FailoverServices))
}

// WasmDeployments mocks base method
func (m *MockLocalSnapshot) WasmDeployments() v1alpha1sets.WasmDeploymentSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WasmDeployments")
	ret0, _ := ret[0].(v1alpha1sets.WasmDeploymentSet)
	return ret0
}

// WasmDeployments indicates an expected call of WasmDeployments
func (mr *MockLocalSnapshotMockRecorder) WasmDeployments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WasmDeployments", reflect.TypeOf((*MockLocalSnapshot)(nil).WasmDeployments))
}

// VirtualDestinations mocks base method
func (m *MockLocalSnapshot) VirtualDestinations() v1alpha1sets.VirtualDestinationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualDestinations")
	ret0, _ := ret[0].(v1alpha1sets.VirtualDestinationSet)
	return ret0
}

// VirtualDestinations indicates an expected call of VirtualDestinations
func (mr *MockLocalSnapshotMockRecorder) VirtualDestinations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualDestinations", reflect.TypeOf((*MockLocalSnapshot)(nil).VirtualDestinations))
}

// AccessLogRecords mocks base method
func (m *MockLocalSnapshot) AccessLogRecords() v1alpha1sets0.AccessLogRecordSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessLogRecords")
	ret0, _ := ret[0].(v1alpha1sets0.AccessLogRecordSet)
	return ret0
}

// AccessLogRecords indicates an expected call of AccessLogRecords
func (mr *MockLocalSnapshotMockRecorder) AccessLogRecords() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessLogRecords", reflect.TypeOf((*MockLocalSnapshot)(nil).AccessLogRecords))
}

// Secrets mocks base method
func (m *MockLocalSnapshot) Secrets() v1sets.SecretSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Secrets")
	ret0, _ := ret[0].(v1sets.SecretSet)
	return ret0
}

// Secrets indicates an expected call of Secrets
func (mr *MockLocalSnapshotMockRecorder) Secrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Secrets", reflect.TypeOf((*MockLocalSnapshot)(nil).Secrets))
}

// KubernetesClusters mocks base method
func (m *MockLocalSnapshot) KubernetesClusters() v1alpha1sets1.KubernetesClusterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubernetesClusters")
	ret0, _ := ret[0].(v1alpha1sets1.KubernetesClusterSet)
	return ret0
}

// KubernetesClusters indicates an expected call of KubernetesClusters
func (mr *MockLocalSnapshotMockRecorder) KubernetesClusters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubernetesClusters", reflect.TypeOf((*MockLocalSnapshot)(nil).KubernetesClusters))
}

// SyncStatusesMultiCluster mocks base method
func (m *MockLocalSnapshot) SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client, opts input.LocalSyncStatusOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatusesMultiCluster", ctx, mcClient, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStatusesMultiCluster indicates an expected call of SyncStatusesMultiCluster
func (mr *MockLocalSnapshotMockRecorder) SyncStatusesMultiCluster(ctx, mcClient, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatusesMultiCluster", reflect.TypeOf((*MockLocalSnapshot)(nil).SyncStatusesMultiCluster), ctx, mcClient, opts)
}

// SyncStatuses mocks base method
func (m *MockLocalSnapshot) SyncStatuses(ctx context.Context, c client.Client, opts input.LocalSyncStatusOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatuses", ctx, c, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStatuses indicates an expected call of SyncStatuses
func (mr *MockLocalSnapshotMockRecorder) SyncStatuses(ctx, c, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatuses", reflect.TypeOf((*MockLocalSnapshot)(nil).SyncStatuses), ctx, c, opts)
}

// MarshalJSON mocks base method
func (m *MockLocalSnapshot) MarshalJSON() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON
func (mr *MockLocalSnapshotMockRecorder) MarshalJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockLocalSnapshot)(nil).MarshalJSON))
}

// MockLocalBuilder is a mock of LocalBuilder interface
type MockLocalBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockLocalBuilderMockRecorder
}

// MockLocalBuilderMockRecorder is the mock recorder for MockLocalBuilder
type MockLocalBuilderMockRecorder struct {
	mock *MockLocalBuilder
}

// NewMockLocalBuilder creates a new mock instance
func NewMockLocalBuilder(ctrl *gomock.Controller) *MockLocalBuilder {
	mock := &MockLocalBuilder{ctrl: ctrl}
	mock.recorder = &MockLocalBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocalBuilder) EXPECT() *MockLocalBuilderMockRecorder {
	return m.recorder
}

// BuildSnapshot mocks base method
func (m *MockLocalBuilder) BuildSnapshot(ctx context.Context, name string, opts input.LocalBuildOptions) (input.LocalSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSnapshot", ctx, name, opts)
	ret0, _ := ret[0].(input.LocalSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildSnapshot indicates an expected call of BuildSnapshot
func (mr *MockLocalBuilderMockRecorder) BuildSnapshot(ctx, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSnapshot", reflect.TypeOf((*MockLocalBuilder)(nil).BuildSnapshot), ctx, name, opts)
}
