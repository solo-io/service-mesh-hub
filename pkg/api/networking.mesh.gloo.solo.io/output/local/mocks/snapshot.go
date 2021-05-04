// Code generated by MockGen. DO NOT EDIT.
// Source: ./snapshot.go

// Package mock_local is a generated GoMock package.
package mock_local

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	local "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/local"
	output "github.com/solo-io/skv2/contrib/pkg/output"
	multicluster "github.com/solo-io/skv2/pkg/multicluster"
	resource "github.com/solo-io/skv2/pkg/resource"
	v1 "k8s.io/api/core/v1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockSnapshot is a mock of Snapshot interface.
type MockSnapshot struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotMockRecorder
}

// MockSnapshotMockRecorder is the mock recorder for MockSnapshot.
type MockSnapshotMockRecorder struct {
	mock *MockSnapshot
}

// NewMockSnapshot creates a new mock instance.
func NewMockSnapshot(ctrl *gomock.Controller) *MockSnapshot {
	mock := &MockSnapshot{ctrl: ctrl}
	mock.recorder = &MockSnapshotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnapshot) EXPECT() *MockSnapshotMockRecorder {
	return m.recorder
}

// ApplyLocalCluster mocks base method.
func (m *MockSnapshot) ApplyLocalCluster(ctx context.Context, clusterClient client.Client, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyLocalCluster", ctx, clusterClient, errHandler)
}

// ApplyLocalCluster indicates an expected call of ApplyLocalCluster.
func (mr *MockSnapshotMockRecorder) ApplyLocalCluster(ctx, clusterClient, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyLocalCluster", reflect.TypeOf((*MockSnapshot)(nil).ApplyLocalCluster), ctx, clusterClient, errHandler)
}

// ApplyMultiCluster mocks base method.
func (m *MockSnapshot) ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyMultiCluster", ctx, multiClusterClient, errHandler)
}

// ApplyMultiCluster indicates an expected call of ApplyMultiCluster.
func (mr *MockSnapshotMockRecorder) ApplyMultiCluster(ctx, multiClusterClient, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyMultiCluster", reflect.TypeOf((*MockSnapshot)(nil).ApplyMultiCluster), ctx, multiClusterClient, errHandler)
}

// MarshalJSON mocks base method.
func (m *MockSnapshot) MarshalJSON() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON.
func (mr *MockSnapshotMockRecorder) MarshalJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockSnapshot)(nil).MarshalJSON))
}

// Secrets mocks base method.
func (m *MockSnapshot) Secrets() []local.LabeledSecretSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Secrets")
	ret0, _ := ret[0].([]local.LabeledSecretSet)
	return ret0
}

// Secrets indicates an expected call of Secrets.
func (mr *MockSnapshotMockRecorder) Secrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Secrets", reflect.TypeOf((*MockSnapshot)(nil).Secrets))
}

// MockLabeledSecretSet is a mock of LabeledSecretSet interface.
type MockLabeledSecretSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledSecretSetMockRecorder
}

// MockLabeledSecretSetMockRecorder is the mock recorder for MockLabeledSecretSet.
type MockLabeledSecretSetMockRecorder struct {
	mock *MockLabeledSecretSet
}

// NewMockLabeledSecretSet creates a new mock instance.
func NewMockLabeledSecretSet(ctrl *gomock.Controller) *MockLabeledSecretSet {
	mock := &MockLabeledSecretSet{ctrl: ctrl}
	mock.recorder = &MockLabeledSecretSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLabeledSecretSet) EXPECT() *MockLabeledSecretSetMockRecorder {
	return m.recorder
}

// Generic mocks base method.
func (m *MockLabeledSecretSet) Generic() output.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockLabeledSecretSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledSecretSet)(nil).Generic))
}

// Labels mocks base method.
func (m *MockLabeledSecretSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels.
func (mr *MockLabeledSecretSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledSecretSet)(nil).Labels))
}

// Set mocks base method.
func (m *MockLabeledSecretSet) Set() v1sets.SecretSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1sets.SecretSet)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockLabeledSecretSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledSecretSet)(nil).Set))
}

// MockBuilder is a mock of Builder interface.
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder.
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance.
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// AddCluster mocks base method.
func (m *MockBuilder) AddCluster(cluster string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddCluster", cluster)
}

// AddCluster indicates an expected call of AddCluster.
func (mr *MockBuilderMockRecorder) AddCluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCluster", reflect.TypeOf((*MockBuilder)(nil).AddCluster), cluster)
}

// AddSecrets mocks base method.
func (m *MockBuilder) AddSecrets(secrets ...*v1.Secret) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range secrets {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddSecrets", varargs...)
}

// AddSecrets indicates an expected call of AddSecrets.
func (mr *MockBuilderMockRecorder) AddSecrets(secrets ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSecrets", reflect.TypeOf((*MockBuilder)(nil).AddSecrets), secrets...)
}

// BuildLabelPartitionedSnapshot mocks base method.
func (m *MockBuilder) BuildLabelPartitionedSnapshot(labelKey string) (local.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildLabelPartitionedSnapshot", labelKey)
	ret0, _ := ret[0].(local.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildLabelPartitionedSnapshot indicates an expected call of BuildLabelPartitionedSnapshot.
func (mr *MockBuilderMockRecorder) BuildLabelPartitionedSnapshot(labelKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildLabelPartitionedSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildLabelPartitionedSnapshot), labelKey)
}

// BuildSinglePartitionedSnapshot mocks base method.
func (m *MockBuilder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (local.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSinglePartitionedSnapshot", snapshotLabels)
	ret0, _ := ret[0].(local.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildSinglePartitionedSnapshot indicates an expected call of BuildSinglePartitionedSnapshot.
func (mr *MockBuilderMockRecorder) BuildSinglePartitionedSnapshot(snapshotLabels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSinglePartitionedSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildSinglePartitionedSnapshot), snapshotLabels)
}

// Clone mocks base method.
func (m *MockBuilder) Clone() local.Builder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(local.Builder)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockBuilderMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockBuilder)(nil).Clone))
}

// Clusters mocks base method.
func (m *MockBuilder) Clusters() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clusters")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Clusters indicates an expected call of Clusters.
func (mr *MockBuilderMockRecorder) Clusters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clusters", reflect.TypeOf((*MockBuilder)(nil).Clusters))
}

// Generic mocks base method.
func (m *MockBuilder) Generic() resource.ClusterSnapshot {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(resource.ClusterSnapshot)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockBuilderMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockBuilder)(nil).Generic))
}

// GetSecrets mocks base method.
func (m *MockBuilder) GetSecrets() v1sets.SecretSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecrets")
	ret0, _ := ret[0].(v1sets.SecretSet)
	return ret0
}

// GetSecrets indicates an expected call of GetSecrets.
func (mr *MockBuilderMockRecorder) GetSecrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecrets", reflect.TypeOf((*MockBuilder)(nil).GetSecrets))
}

// Merge mocks base method.
func (m *MockBuilder) Merge(other local.Builder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Merge", other)
}

// Merge indicates an expected call of Merge.
func (mr *MockBuilderMockRecorder) Merge(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Merge", reflect.TypeOf((*MockBuilder)(nil).Merge), other)
}
