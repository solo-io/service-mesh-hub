// Code generated by MockGen. DO NOT EDIT.
// Source: ./snapshot.go

// Package mock_certagent is a generated GoMock package.
package mock_certagent

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	certagent "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/agent/output/certagent"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2/sets"
	output "github.com/solo-io/skv2/contrib/pkg/output"
	multicluster "github.com/solo-io/skv2/pkg/multicluster"
	v1 "k8s.io/api/core/v1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockSnapshot is a mock of Snapshot interface
type MockSnapshot struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotMockRecorder
}

// MockSnapshotMockRecorder is the mock recorder for MockSnapshot
type MockSnapshotMockRecorder struct {
	mock *MockSnapshot
}

// NewMockSnapshot creates a new mock instance
func NewMockSnapshot(ctrl *gomock.Controller) *MockSnapshot {
	mock := &MockSnapshot{ctrl: ctrl}
	mock.recorder = &MockSnapshotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSnapshot) EXPECT() *MockSnapshotMockRecorder {
	return m.recorder
}

// CertificateRequests mocks base method
func (m *MockSnapshot) CertificateRequests() []certagent.LabeledCertificateRequestSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CertificateRequests")
	ret0, _ := ret[0].([]certagent.LabeledCertificateRequestSet)
	return ret0
}

// CertificateRequests indicates an expected call of CertificateRequests
func (mr *MockSnapshotMockRecorder) CertificateRequests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CertificateRequests", reflect.TypeOf((*MockSnapshot)(nil).CertificateRequests))
}

// Secrets mocks base method
func (m *MockSnapshot) Secrets() []certagent.LabeledSecretSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Secrets")
	ret0, _ := ret[0].([]certagent.LabeledSecretSet)
	return ret0
}

// Secrets indicates an expected call of Secrets
func (mr *MockSnapshotMockRecorder) Secrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Secrets", reflect.TypeOf((*MockSnapshot)(nil).Secrets))
}

// ApplyLocalCluster mocks base method
func (m *MockSnapshot) ApplyLocalCluster(ctx context.Context, clusterClient client.Client, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyLocalCluster", ctx, clusterClient, errHandler)
}

// ApplyLocalCluster indicates an expected call of ApplyLocalCluster
func (mr *MockSnapshotMockRecorder) ApplyLocalCluster(ctx, clusterClient, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyLocalCluster", reflect.TypeOf((*MockSnapshot)(nil).ApplyLocalCluster), ctx, clusterClient, errHandler)
}

// ApplyMultiCluster mocks base method
func (m *MockSnapshot) ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyMultiCluster", ctx, multiClusterClient, errHandler)
}

// ApplyMultiCluster indicates an expected call of ApplyMultiCluster
func (mr *MockSnapshotMockRecorder) ApplyMultiCluster(ctx, multiClusterClient, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyMultiCluster", reflect.TypeOf((*MockSnapshot)(nil).ApplyMultiCluster), ctx, multiClusterClient, errHandler)
}

// MarshalJSON mocks base method
func (m *MockSnapshot) MarshalJSON() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON
func (mr *MockSnapshotMockRecorder) MarshalJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockSnapshot)(nil).MarshalJSON))
}

// MockLabeledCertificateRequestSet is a mock of LabeledCertificateRequestSet interface
type MockLabeledCertificateRequestSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledCertificateRequestSetMockRecorder
}

// MockLabeledCertificateRequestSetMockRecorder is the mock recorder for MockLabeledCertificateRequestSet
type MockLabeledCertificateRequestSetMockRecorder struct {
	mock *MockLabeledCertificateRequestSet
}

// NewMockLabeledCertificateRequestSet creates a new mock instance
func NewMockLabeledCertificateRequestSet(ctrl *gomock.Controller) *MockLabeledCertificateRequestSet {
	mock := &MockLabeledCertificateRequestSet{ctrl: ctrl}
	mock.recorder = &MockLabeledCertificateRequestSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLabeledCertificateRequestSet) EXPECT() *MockLabeledCertificateRequestSetMockRecorder {
	return m.recorder
}

// Labels mocks base method
func (m *MockLabeledCertificateRequestSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels
func (mr *MockLabeledCertificateRequestSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledCertificateRequestSet)(nil).Labels))
}

// Set mocks base method
func (m *MockLabeledCertificateRequestSet) Set() v1alpha2sets.CertificateRequestSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1alpha2sets.CertificateRequestSet)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockLabeledCertificateRequestSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledCertificateRequestSet)(nil).Set))
}

// Generic mocks base method
func (m *MockLabeledCertificateRequestSet) Generic() output.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockLabeledCertificateRequestSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledCertificateRequestSet)(nil).Generic))
}

// MockLabeledSecretSet is a mock of LabeledSecretSet interface
type MockLabeledSecretSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledSecretSetMockRecorder
}

// MockLabeledSecretSetMockRecorder is the mock recorder for MockLabeledSecretSet
type MockLabeledSecretSetMockRecorder struct {
	mock *MockLabeledSecretSet
}

// NewMockLabeledSecretSet creates a new mock instance
func NewMockLabeledSecretSet(ctrl *gomock.Controller) *MockLabeledSecretSet {
	mock := &MockLabeledSecretSet{ctrl: ctrl}
	mock.recorder = &MockLabeledSecretSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLabeledSecretSet) EXPECT() *MockLabeledSecretSetMockRecorder {
	return m.recorder
}

// Labels mocks base method
func (m *MockLabeledSecretSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels
func (mr *MockLabeledSecretSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledSecretSet)(nil).Labels))
}

// Set mocks base method
func (m *MockLabeledSecretSet) Set() v1sets.SecretSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1sets.SecretSet)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockLabeledSecretSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledSecretSet)(nil).Set))
}

// Generic mocks base method
func (m *MockLabeledSecretSet) Generic() output.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockLabeledSecretSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledSecretSet)(nil).Generic))
}

// MockBuilder is a mock of Builder interface
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// AddCertificateRequests mocks base method
func (m *MockBuilder) AddCertificateRequests(certificateRequests ...*v1alpha2.CertificateRequest) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range certificateRequests {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddCertificateRequests", varargs...)
}

// AddCertificateRequests indicates an expected call of AddCertificateRequests
func (mr *MockBuilderMockRecorder) AddCertificateRequests(certificateRequests ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCertificateRequests", reflect.TypeOf((*MockBuilder)(nil).AddCertificateRequests), certificateRequests...)
}

// GetCertificateRequests mocks base method
func (m *MockBuilder) GetCertificateRequests() v1alpha2sets.CertificateRequestSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificateRequests")
	ret0, _ := ret[0].(v1alpha2sets.CertificateRequestSet)
	return ret0
}

// GetCertificateRequests indicates an expected call of GetCertificateRequests
func (mr *MockBuilderMockRecorder) GetCertificateRequests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateRequests", reflect.TypeOf((*MockBuilder)(nil).GetCertificateRequests))
}

// AddSecrets mocks base method
func (m *MockBuilder) AddSecrets(secrets ...*v1.Secret) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range secrets {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddSecrets", varargs...)
}

// AddSecrets indicates an expected call of AddSecrets
func (mr *MockBuilderMockRecorder) AddSecrets(secrets ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSecrets", reflect.TypeOf((*MockBuilder)(nil).AddSecrets), secrets...)
}

// GetSecrets mocks base method
func (m *MockBuilder) GetSecrets() v1sets.SecretSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecrets")
	ret0, _ := ret[0].(v1sets.SecretSet)
	return ret0
}

// GetSecrets indicates an expected call of GetSecrets
func (mr *MockBuilderMockRecorder) GetSecrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecrets", reflect.TypeOf((*MockBuilder)(nil).GetSecrets))
}

// BuildLabelPartitionedSnapshot mocks base method
func (m *MockBuilder) BuildLabelPartitionedSnapshot(labelKey string) (certagent.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildLabelPartitionedSnapshot", labelKey)
	ret0, _ := ret[0].(certagent.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildLabelPartitionedSnapshot indicates an expected call of BuildLabelPartitionedSnapshot
func (mr *MockBuilderMockRecorder) BuildLabelPartitionedSnapshot(labelKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildLabelPartitionedSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildLabelPartitionedSnapshot), labelKey)
}

// BuildSinglePartitionedSnapshot mocks base method
func (m *MockBuilder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (certagent.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSinglePartitionedSnapshot", snapshotLabels)
	ret0, _ := ret[0].(certagent.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildSinglePartitionedSnapshot indicates an expected call of BuildSinglePartitionedSnapshot
func (mr *MockBuilderMockRecorder) BuildSinglePartitionedSnapshot(snapshotLabels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSinglePartitionedSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildSinglePartitionedSnapshot), snapshotLabels)
}

// AddCluster mocks base method
func (m *MockBuilder) AddCluster(cluster string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddCluster", cluster)
}

// AddCluster indicates an expected call of AddCluster
func (mr *MockBuilderMockRecorder) AddCluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCluster", reflect.TypeOf((*MockBuilder)(nil).AddCluster), cluster)
}

// Merge mocks base method
func (m *MockBuilder) Merge(other certagent.Builder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Merge", other)
}

// Merge indicates an expected call of Merge
func (mr *MockBuilderMockRecorder) Merge(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Merge", reflect.TypeOf((*MockBuilder)(nil).Merge), other)
}

// Clone mocks base method
func (m *MockBuilder) Clone() certagent.Builder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(certagent.Builder)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockBuilderMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockBuilder)(nil).Clone))
}
