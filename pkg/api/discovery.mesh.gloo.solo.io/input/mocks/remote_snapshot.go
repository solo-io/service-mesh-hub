// Code generated by MockGen. DO NOT EDIT.
// Source: ./remote_snapshot.go

// Package mock_input is a generated GoMock package.
package mock_input

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta2sets "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.k8s.aws/v1beta2/sets"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/sets"
	v1sets0 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	input "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/input"
)

// MockDiscoveryInputSnapshot is a mock of DiscoveryInputSnapshot interface
type MockDiscoveryInputSnapshot struct {
	ctrl     *gomock.Controller
	recorder *MockDiscoveryInputSnapshotMockRecorder
}

// MockDiscoveryInputSnapshotMockRecorder is the mock recorder for MockDiscoveryInputSnapshot
type MockDiscoveryInputSnapshotMockRecorder struct {
	mock *MockDiscoveryInputSnapshot
}

// NewMockDiscoveryInputSnapshot creates a new mock instance
func NewMockDiscoveryInputSnapshot(ctrl *gomock.Controller) *MockDiscoveryInputSnapshot {
	mock := &MockDiscoveryInputSnapshot{ctrl: ctrl}
	mock.recorder = &MockDiscoveryInputSnapshotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiscoveryInputSnapshot) EXPECT() *MockDiscoveryInputSnapshotMockRecorder {
	return m.recorder
}

// Meshes mocks base method
func (m *MockDiscoveryInputSnapshot) Meshes() v1beta2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meshes")
	ret0, _ := ret[0].(v1beta2sets.MeshSet)
	return ret0
}

// Meshes indicates an expected call of Meshes
func (mr *MockDiscoveryInputSnapshotMockRecorder) Meshes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meshes", reflect.TypeOf((*MockDiscoveryInputSnapshot)(nil).Meshes))
}

// ConfigMaps mocks base method
func (m *MockDiscoveryInputSnapshot) ConfigMaps() v1sets0.ConfigMapSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigMaps")
	ret0, _ := ret[0].(v1sets0.ConfigMapSet)
	return ret0
}

// ConfigMaps indicates an expected call of ConfigMaps
func (mr *MockDiscoveryInputSnapshotMockRecorder) ConfigMaps() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigMaps", reflect.TypeOf((*MockDiscoveryInputSnapshot)(nil).ConfigMaps))
}

// Services mocks base method
func (m *MockDiscoveryInputSnapshot) Services() v1sets0.ServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Services")
	ret0, _ := ret[0].(v1sets0.ServiceSet)
	return ret0
}

// Services indicates an expected call of Services
func (mr *MockDiscoveryInputSnapshotMockRecorder) Services() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Services", reflect.TypeOf((*MockDiscoveryInputSnapshot)(nil).Services))
}

// Pods mocks base method
func (m *MockDiscoveryInputSnapshot) Pods() v1sets0.PodSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pods")
	ret0, _ := ret[0].(v1sets0.PodSet)
	return ret0
}

// Pods indicates an expected call of Pods
func (mr *MockDiscoveryInputSnapshotMockRecorder) Pods() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pods", reflect.TypeOf((*MockDiscoveryInputSnapshot)(nil).Pods))
}

// Endpoints mocks base method
func (m *MockDiscoveryInputSnapshot) Endpoints() v1sets0.EndpointsSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoints")
	ret0, _ := ret[0].(v1sets0.EndpointsSet)
	return ret0
}

// Endpoints indicates an expected call of Endpoints
func (mr *MockDiscoveryInputSnapshotMockRecorder) Endpoints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoints", reflect.TypeOf((*MockDiscoveryInputSnapshot)(nil).Endpoints))
}

// Nodes mocks base method
func (m *MockDiscoveryInputSnapshot) Nodes() v1sets0.NodeSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nodes")
	ret0, _ := ret[0].(v1sets0.NodeSet)
	return ret0
}

// Nodes indicates an expected call of Nodes
func (mr *MockDiscoveryInputSnapshotMockRecorder) Nodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nodes", reflect.TypeOf((*MockDiscoveryInputSnapshot)(nil).Nodes))
}

// Deployments mocks base method
func (m *MockDiscoveryInputSnapshot) Deployments() v1sets.DeploymentSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deployments")
	ret0, _ := ret[0].(v1sets.DeploymentSet)
	return ret0
}

// Deployments indicates an expected call of Deployments
func (mr *MockDiscoveryInputSnapshotMockRecorder) Deployments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployments", reflect.TypeOf((*MockDiscoveryInputSnapshot)(nil).Deployments))
}

// ReplicaSets mocks base method
func (m *MockDiscoveryInputSnapshot) ReplicaSets() v1sets.ReplicaSetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicaSets")
	ret0, _ := ret[0].(v1sets.ReplicaSetSet)
	return ret0
}

// ReplicaSets indicates an expected call of ReplicaSets
func (mr *MockDiscoveryInputSnapshotMockRecorder) ReplicaSets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicaSets", reflect.TypeOf((*MockDiscoveryInputSnapshot)(nil).ReplicaSets))
}

// DaemonSets mocks base method
func (m *MockDiscoveryInputSnapshot) DaemonSets() v1sets.DaemonSetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DaemonSets")
	ret0, _ := ret[0].(v1sets.DaemonSetSet)
	return ret0
}

// DaemonSets indicates an expected call of DaemonSets
func (mr *MockDiscoveryInputSnapshotMockRecorder) DaemonSets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DaemonSets", reflect.TypeOf((*MockDiscoveryInputSnapshot)(nil).DaemonSets))
}

// StatefulSets mocks base method
func (m *MockDiscoveryInputSnapshot) StatefulSets() v1sets.StatefulSetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatefulSets")
	ret0, _ := ret[0].(v1sets.StatefulSetSet)
	return ret0
}

// StatefulSets indicates an expected call of StatefulSets
func (mr *MockDiscoveryInputSnapshotMockRecorder) StatefulSets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatefulSets", reflect.TypeOf((*MockDiscoveryInputSnapshot)(nil).StatefulSets))
}

// MarshalJSON mocks base method
func (m *MockDiscoveryInputSnapshot) MarshalJSON() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON
func (mr *MockDiscoveryInputSnapshotMockRecorder) MarshalJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockDiscoveryInputSnapshot)(nil).MarshalJSON))
}

// MockDiscoveryInputBuilder is a mock of DiscoveryInputBuilder interface
type MockDiscoveryInputBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockDiscoveryInputBuilderMockRecorder
}

// MockDiscoveryInputBuilderMockRecorder is the mock recorder for MockDiscoveryInputBuilder
type MockDiscoveryInputBuilderMockRecorder struct {
	mock *MockDiscoveryInputBuilder
}

// NewMockDiscoveryInputBuilder creates a new mock instance
func NewMockDiscoveryInputBuilder(ctrl *gomock.Controller) *MockDiscoveryInputBuilder {
	mock := &MockDiscoveryInputBuilder{ctrl: ctrl}
	mock.recorder = &MockDiscoveryInputBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiscoveryInputBuilder) EXPECT() *MockDiscoveryInputBuilderMockRecorder {
	return m.recorder
}

// BuildSnapshot mocks base method
func (m *MockDiscoveryInputBuilder) BuildSnapshot(ctx context.Context, name string, opts input.DiscoveryInputBuildOptions) (input.DiscoveryInputSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSnapshot", ctx, name, opts)
	ret0, _ := ret[0].(input.DiscoveryInputSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildSnapshot indicates an expected call of BuildSnapshot
func (mr *MockDiscoveryInputBuilderMockRecorder) BuildSnapshot(ctx, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSnapshot", reflect.TypeOf((*MockDiscoveryInputBuilder)(nil).BuildSnapshot), ctx, name, opts)
}
