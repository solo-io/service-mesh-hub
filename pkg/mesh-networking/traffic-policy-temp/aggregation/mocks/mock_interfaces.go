// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_traffic_policy_aggregation is a generated GoMock package.
package mock_traffic_policy_aggregation

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	types "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/types"
	v1alpha10 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	types0 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/types"
	traffic_policy_aggregation "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/traffic-policy-temp/aggregation"
	mesh_translation "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/traffic-policy-temp/translation/translators"
)

// MockPolicyCollector is a mock of PolicyCollector interface.
type MockPolicyCollector struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyCollectorMockRecorder
}

// MockPolicyCollectorMockRecorder is the mock recorder for MockPolicyCollector.
type MockPolicyCollectorMockRecorder struct {
	mock *MockPolicyCollector
}

// NewMockPolicyCollector creates a new mock instance.
func NewMockPolicyCollector(ctrl *gomock.Controller) *MockPolicyCollector {
	mock := &MockPolicyCollector{ctrl: ctrl}
	mock.recorder = &MockPolicyCollectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicyCollector) EXPECT() *MockPolicyCollectorMockRecorder {
	return m.recorder
}

// CollectForService mocks base method.
func (m *MockPolicyCollector) CollectForService(ctx context.Context, meshService *v1alpha1.MeshService, allMeshServices []*v1alpha1.MeshService, mesh *v1alpha1.Mesh, translationValidator mesh_translation.TranslationValidator, allTrafficPolicies []*v1alpha10.TrafficPolicy) (*traffic_policy_aggregation.CollectionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollectForService", ctx, meshService, allMeshServices, mesh, translationValidator, allTrafficPolicies)
	ret0, _ := ret[0].(*traffic_policy_aggregation.CollectionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollectForService indicates an expected call of CollectForService.
func (mr *MockPolicyCollectorMockRecorder) CollectForService(ctx, meshService, allMeshServices, mesh, translationValidator, allTrafficPolicies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollectForService", reflect.TypeOf((*MockPolicyCollector)(nil).CollectForService), ctx, meshService, allMeshServices, mesh, translationValidator, allTrafficPolicies)
}

// MockInMemoryStatusMutator is a mock of InMemoryStatusMutator interface.
type MockInMemoryStatusMutator struct {
	ctrl     *gomock.Controller
	recorder *MockInMemoryStatusMutatorMockRecorder
}

// MockInMemoryStatusMutatorMockRecorder is the mock recorder for MockInMemoryStatusMutator.
type MockInMemoryStatusMutatorMockRecorder struct {
	mock *MockInMemoryStatusMutator
}

// NewMockInMemoryStatusMutator creates a new mock instance.
func NewMockInMemoryStatusMutator(ctrl *gomock.Controller) *MockInMemoryStatusMutator {
	mock := &MockInMemoryStatusMutator{ctrl: ctrl}
	mock.recorder = &MockInMemoryStatusMutatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInMemoryStatusMutator) EXPECT() *MockInMemoryStatusMutatorMockRecorder {
	return m.recorder
}

// MutateServicePolicies mocks base method.
func (m *MockInMemoryStatusMutator) MutateServicePolicies(meshService *v1alpha1.MeshService, newlyComputedMergeablePolicies []*types.MeshServiceStatus_ValidatedTrafficPolicy) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MutateServicePolicies", meshService, newlyComputedMergeablePolicies)
	ret0, _ := ret[0].(bool)
	return ret0
}

// MutateServicePolicies indicates an expected call of MutateServicePolicies.
func (mr *MockInMemoryStatusMutatorMockRecorder) MutateServicePolicies(meshService, newlyComputedMergeablePolicies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MutateServicePolicies", reflect.TypeOf((*MockInMemoryStatusMutator)(nil).MutateServicePolicies), meshService, newlyComputedMergeablePolicies)
}

// MutateTrafficPolicyTranslationStatus mocks base method.
func (m *MockInMemoryStatusMutator) MutateTrafficPolicyTranslationStatus(policy *v1alpha10.TrafficPolicy, newConflictErrors []*types0.TrafficPolicyStatus_ConflictError, newTranslationErrors []*types0.TrafficPolicyStatus_TranslatorError) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MutateTrafficPolicyTranslationStatus", policy, newConflictErrors, newTranslationErrors)
}

// MutateTrafficPolicyTranslationStatus indicates an expected call of MutateTrafficPolicyTranslationStatus.
func (mr *MockInMemoryStatusMutatorMockRecorder) MutateTrafficPolicyTranslationStatus(policy, newConflictErrors, newTranslationErrors interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MutateTrafficPolicyTranslationStatus", reflect.TypeOf((*MockInMemoryStatusMutator)(nil).MutateTrafficPolicyTranslationStatus), policy, newConflictErrors, newTranslationErrors)
}

// MockAggregator is a mock of Aggregator interface.
type MockAggregator struct {
	ctrl     *gomock.Controller
	recorder *MockAggregatorMockRecorder
}

// MockAggregatorMockRecorder is the mock recorder for MockAggregator.
type MockAggregatorMockRecorder struct {
	mock *MockAggregator
}

// NewMockAggregator creates a new mock instance.
func NewMockAggregator(ctrl *gomock.Controller) *MockAggregator {
	mock := &MockAggregator{ctrl: ctrl}
	mock.recorder = &MockAggregatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAggregator) EXPECT() *MockAggregatorMockRecorder {
	return m.recorder
}

// FindMergeConflict mocks base method.
func (m *MockAggregator) FindMergeConflict(trafficPolicyToMerge *types0.TrafficPolicySpec, policiesToMergeWith []*types0.TrafficPolicySpec, meshService *v1alpha1.MeshService) *types0.TrafficPolicyStatus_ConflictError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMergeConflict", trafficPolicyToMerge, policiesToMergeWith, meshService)
	ret0, _ := ret[0].(*types0.TrafficPolicyStatus_ConflictError)
	return ret0
}

// FindMergeConflict indicates an expected call of FindMergeConflict.
func (mr *MockAggregatorMockRecorder) FindMergeConflict(trafficPolicyToMerge, policiesToMergeWith, meshService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMergeConflict", reflect.TypeOf((*MockAggregator)(nil).FindMergeConflict), trafficPolicyToMerge, policiesToMergeWith, meshService)
}

// PoliciesForService mocks base method.
func (m *MockAggregator) PoliciesForService(trafficPolicies []*v1alpha10.TrafficPolicy, meshService *v1alpha1.MeshService) ([]*v1alpha10.TrafficPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PoliciesForService", trafficPolicies, meshService)
	ret0, _ := ret[0].([]*v1alpha10.TrafficPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PoliciesForService indicates an expected call of PoliciesForService.
func (mr *MockAggregatorMockRecorder) PoliciesForService(trafficPolicies, meshService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PoliciesForService", reflect.TypeOf((*MockAggregator)(nil).PoliciesForService), trafficPolicies, meshService)
}
