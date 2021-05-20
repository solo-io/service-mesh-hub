// Code generated by MockGen. DO NOT EDIT.
// Source: ./snapshot.go

// Package mock_smi is a generated GoMock package.
package mock_smi

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"
	v1alpha20 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha2"
	v1alpha2sets "github.com/solo-io/external-apis/pkg/api/smi/access.smi-spec.io/v1alpha2/sets"
	v1alpha3sets "github.com/solo-io/external-apis/pkg/api/smi/specs.smi-spec.io/v1alpha3/sets"
	v1alpha2sets0 "github.com/solo-io/external-apis/pkg/api/smi/split.smi-spec.io/v1alpha2/sets"
	smi "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/smi"
	output "github.com/solo-io/skv2/contrib/pkg/output"
	multicluster "github.com/solo-io/skv2/pkg/multicluster"
	resource "github.com/solo-io/skv2/pkg/resource"
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

// HTTPRouteGroups mocks base method.
func (m *MockSnapshot) HTTPRouteGroups() []smi.LabeledHTTPRouteGroupSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPRouteGroups")
	ret0, _ := ret[0].([]smi.LabeledHTTPRouteGroupSet)
	return ret0
}

// HTTPRouteGroups indicates an expected call of HTTPRouteGroups.
func (mr *MockSnapshotMockRecorder) HTTPRouteGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPRouteGroups", reflect.TypeOf((*MockSnapshot)(nil).HTTPRouteGroups))
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

// TrafficSplits mocks base method.
func (m *MockSnapshot) TrafficSplits() []smi.LabeledTrafficSplitSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficSplits")
	ret0, _ := ret[0].([]smi.LabeledTrafficSplitSet)
	return ret0
}

// TrafficSplits indicates an expected call of TrafficSplits.
func (mr *MockSnapshotMockRecorder) TrafficSplits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficSplits", reflect.TypeOf((*MockSnapshot)(nil).TrafficSplits))
}

// TrafficTargets mocks base method.
func (m *MockSnapshot) TrafficTargets() []smi.LabeledTrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficTargets")
	ret0, _ := ret[0].([]smi.LabeledTrafficTargetSet)
	return ret0
}

// TrafficTargets indicates an expected call of TrafficTargets.
func (mr *MockSnapshotMockRecorder) TrafficTargets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficTargets", reflect.TypeOf((*MockSnapshot)(nil).TrafficTargets))
}

// MockLabeledTrafficSplitSet is a mock of LabeledTrafficSplitSet interface.
type MockLabeledTrafficSplitSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledTrafficSplitSetMockRecorder
}

// MockLabeledTrafficSplitSetMockRecorder is the mock recorder for MockLabeledTrafficSplitSet.
type MockLabeledTrafficSplitSetMockRecorder struct {
	mock *MockLabeledTrafficSplitSet
}

// NewMockLabeledTrafficSplitSet creates a new mock instance.
func NewMockLabeledTrafficSplitSet(ctrl *gomock.Controller) *MockLabeledTrafficSplitSet {
	mock := &MockLabeledTrafficSplitSet{ctrl: ctrl}
	mock.recorder = &MockLabeledTrafficSplitSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLabeledTrafficSplitSet) EXPECT() *MockLabeledTrafficSplitSetMockRecorder {
	return m.recorder
}

// Generic mocks base method.
func (m *MockLabeledTrafficSplitSet) Generic() output.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockLabeledTrafficSplitSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledTrafficSplitSet)(nil).Generic))
}

// Labels mocks base method.
func (m *MockLabeledTrafficSplitSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels.
func (mr *MockLabeledTrafficSplitSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledTrafficSplitSet)(nil).Labels))
}

// Set mocks base method.
func (m *MockLabeledTrafficSplitSet) Set() v1alpha2sets0.TrafficSplitSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1alpha2sets0.TrafficSplitSet)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockLabeledTrafficSplitSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledTrafficSplitSet)(nil).Set))
}

// MockLabeledTrafficTargetSet is a mock of LabeledTrafficTargetSet interface.
type MockLabeledTrafficTargetSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledTrafficTargetSetMockRecorder
}

// MockLabeledTrafficTargetSetMockRecorder is the mock recorder for MockLabeledTrafficTargetSet.
type MockLabeledTrafficTargetSetMockRecorder struct {
	mock *MockLabeledTrafficTargetSet
}

// NewMockLabeledTrafficTargetSet creates a new mock instance.
func NewMockLabeledTrafficTargetSet(ctrl *gomock.Controller) *MockLabeledTrafficTargetSet {
	mock := &MockLabeledTrafficTargetSet{ctrl: ctrl}
	mock.recorder = &MockLabeledTrafficTargetSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLabeledTrafficTargetSet) EXPECT() *MockLabeledTrafficTargetSetMockRecorder {
	return m.recorder
}

// Generic mocks base method.
func (m *MockLabeledTrafficTargetSet) Generic() output.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockLabeledTrafficTargetSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledTrafficTargetSet)(nil).Generic))
}

// Labels mocks base method.
func (m *MockLabeledTrafficTargetSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels.
func (mr *MockLabeledTrafficTargetSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledTrafficTargetSet)(nil).Labels))
}

// Set mocks base method.
func (m *MockLabeledTrafficTargetSet) Set() v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockLabeledTrafficTargetSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledTrafficTargetSet)(nil).Set))
}

// MockLabeledHTTPRouteGroupSet is a mock of LabeledHTTPRouteGroupSet interface.
type MockLabeledHTTPRouteGroupSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledHTTPRouteGroupSetMockRecorder
}

// MockLabeledHTTPRouteGroupSetMockRecorder is the mock recorder for MockLabeledHTTPRouteGroupSet.
type MockLabeledHTTPRouteGroupSetMockRecorder struct {
	mock *MockLabeledHTTPRouteGroupSet
}

// NewMockLabeledHTTPRouteGroupSet creates a new mock instance.
func NewMockLabeledHTTPRouteGroupSet(ctrl *gomock.Controller) *MockLabeledHTTPRouteGroupSet {
	mock := &MockLabeledHTTPRouteGroupSet{ctrl: ctrl}
	mock.recorder = &MockLabeledHTTPRouteGroupSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLabeledHTTPRouteGroupSet) EXPECT() *MockLabeledHTTPRouteGroupSetMockRecorder {
	return m.recorder
}

// Generic mocks base method.
func (m *MockLabeledHTTPRouteGroupSet) Generic() output.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockLabeledHTTPRouteGroupSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledHTTPRouteGroupSet)(nil).Generic))
}

// Labels mocks base method.
func (m *MockLabeledHTTPRouteGroupSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels.
func (mr *MockLabeledHTTPRouteGroupSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledHTTPRouteGroupSet)(nil).Labels))
}

// Set mocks base method.
func (m *MockLabeledHTTPRouteGroupSet) Set() v1alpha3sets.HTTPRouteGroupSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1alpha3sets.HTTPRouteGroupSet)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockLabeledHTTPRouteGroupSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledHTTPRouteGroupSet)(nil).Set))
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

// AddHTTPRouteGroups mocks base method.
func (m *MockBuilder) AddHTTPRouteGroups(hTTPRouteGroups ...*v1alpha3.HTTPRouteGroup) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range hTTPRouteGroups {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddHTTPRouteGroups", varargs...)
}

// AddHTTPRouteGroups indicates an expected call of AddHTTPRouteGroups.
func (mr *MockBuilderMockRecorder) AddHTTPRouteGroups(hTTPRouteGroups ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHTTPRouteGroups", reflect.TypeOf((*MockBuilder)(nil).AddHTTPRouteGroups), hTTPRouteGroups...)
}

// AddTrafficSplits mocks base method.
func (m *MockBuilder) AddTrafficSplits(trafficSplits ...*v1alpha20.TrafficSplit) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range trafficSplits {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddTrafficSplits", varargs...)
}

// AddTrafficSplits indicates an expected call of AddTrafficSplits.
func (mr *MockBuilderMockRecorder) AddTrafficSplits(trafficSplits ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrafficSplits", reflect.TypeOf((*MockBuilder)(nil).AddTrafficSplits), trafficSplits...)
}

// AddTrafficTargets mocks base method.
func (m *MockBuilder) AddTrafficTargets(trafficTargets ...*v1alpha2.TrafficTarget) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range trafficTargets {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddTrafficTargets", varargs...)
}

// AddTrafficTargets indicates an expected call of AddTrafficTargets.
func (mr *MockBuilderMockRecorder) AddTrafficTargets(trafficTargets ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrafficTargets", reflect.TypeOf((*MockBuilder)(nil).AddTrafficTargets), trafficTargets...)
}

// BuildLabelPartitionedSnapshot mocks base method.
func (m *MockBuilder) BuildLabelPartitionedSnapshot(labelKey string) (smi.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildLabelPartitionedSnapshot", labelKey)
	ret0, _ := ret[0].(smi.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildLabelPartitionedSnapshot indicates an expected call of BuildLabelPartitionedSnapshot.
func (mr *MockBuilderMockRecorder) BuildLabelPartitionedSnapshot(labelKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildLabelPartitionedSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildLabelPartitionedSnapshot), labelKey)
}

// BuildSinglePartitionedSnapshot mocks base method.
func (m *MockBuilder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (smi.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSinglePartitionedSnapshot", snapshotLabels)
	ret0, _ := ret[0].(smi.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildSinglePartitionedSnapshot indicates an expected call of BuildSinglePartitionedSnapshot.
func (mr *MockBuilderMockRecorder) BuildSinglePartitionedSnapshot(snapshotLabels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSinglePartitionedSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildSinglePartitionedSnapshot), snapshotLabels)
}

// Clone mocks base method.
func (m *MockBuilder) Clone() smi.Builder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(smi.Builder)
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

// GetHTTPRouteGroups mocks base method.
func (m *MockBuilder) GetHTTPRouteGroups() v1alpha3sets.HTTPRouteGroupSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPRouteGroups")
	ret0, _ := ret[0].(v1alpha3sets.HTTPRouteGroupSet)
	return ret0
}

// GetHTTPRouteGroups indicates an expected call of GetHTTPRouteGroups.
func (mr *MockBuilderMockRecorder) GetHTTPRouteGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPRouteGroups", reflect.TypeOf((*MockBuilder)(nil).GetHTTPRouteGroups))
}

// GetTrafficSplits mocks base method.
func (m *MockBuilder) GetTrafficSplits() v1alpha2sets0.TrafficSplitSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrafficSplits")
	ret0, _ := ret[0].(v1alpha2sets0.TrafficSplitSet)
	return ret0
}

// GetTrafficSplits indicates an expected call of GetTrafficSplits.
func (mr *MockBuilderMockRecorder) GetTrafficSplits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrafficSplits", reflect.TypeOf((*MockBuilder)(nil).GetTrafficSplits))
}

// GetTrafficTargets mocks base method.
func (m *MockBuilder) GetTrafficTargets() v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrafficTargets")
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// GetTrafficTargets indicates an expected call of GetTrafficTargets.
func (mr *MockBuilderMockRecorder) GetTrafficTargets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrafficTargets", reflect.TypeOf((*MockBuilder)(nil).GetTrafficTargets))
}

// Merge mocks base method.
func (m *MockBuilder) Merge(other smi.Builder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Merge", other)
}

// Merge indicates an expected call of Merge.
func (mr *MockBuilderMockRecorder) Merge(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Merge", reflect.TypeOf((*MockBuilder)(nil).Merge), other)
}
