// Code generated by MockGen. DO NOT EDIT.
// Source: ./decorators.go

// Package mock_decorators is a generated GoMock package.
package mock_decorators

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	decorators "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/istio/decorators"
	v1alpha3 "istio.io/api/networking/v1alpha3"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// MakeDecorators mocks base method.
func (m *MockFactory) MakeDecorators(params decorators.Parameters) []decorators.Decorator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeDecorators", params)
	ret0, _ := ret[0].([]decorators.Decorator)
	return ret0
}

// MakeDecorators indicates an expected call of MakeDecorators.
func (mr *MockFactoryMockRecorder) MakeDecorators(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeDecorators", reflect.TypeOf((*MockFactory)(nil).MakeDecorators), params)
}

// MockDecorator is a mock of Decorator interface.
type MockDecorator struct {
	ctrl     *gomock.Controller
	recorder *MockDecoratorMockRecorder
}

// MockDecoratorMockRecorder is the mock recorder for MockDecorator.
type MockDecoratorMockRecorder struct {
	mock *MockDecorator
}

// NewMockDecorator creates a new mock instance.
func NewMockDecorator(ctrl *gomock.Controller) *MockDecorator {
	mock := &MockDecorator{ctrl: ctrl}
	mock.recorder = &MockDecoratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDecorator) EXPECT() *MockDecoratorMockRecorder {
	return m.recorder
}

// DecoratorName mocks base method.
func (m *MockDecorator) DecoratorName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecoratorName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DecoratorName indicates an expected call of DecoratorName.
func (mr *MockDecoratorMockRecorder) DecoratorName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecoratorName", reflect.TypeOf((*MockDecorator)(nil).DecoratorName))
}

// MockVirtualMeshDestinationRuleDecorator is a mock of VirtualMeshDestinationRuleDecorator interface.
type MockVirtualMeshDestinationRuleDecorator struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshDestinationRuleDecoratorMockRecorder
}

// MockVirtualMeshDestinationRuleDecoratorMockRecorder is the mock recorder for MockVirtualMeshDestinationRuleDecorator.
type MockVirtualMeshDestinationRuleDecoratorMockRecorder struct {
	mock *MockVirtualMeshDestinationRuleDecorator
}

// NewMockVirtualMeshDestinationRuleDecorator creates a new mock instance.
func NewMockVirtualMeshDestinationRuleDecorator(ctrl *gomock.Controller) *MockVirtualMeshDestinationRuleDecorator {
	mock := &MockVirtualMeshDestinationRuleDecorator{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshDestinationRuleDecoratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshDestinationRuleDecorator) EXPECT() *MockVirtualMeshDestinationRuleDecoratorMockRecorder {
	return m.recorder
}

// DecoratorName mocks base method.
func (m *MockVirtualMeshDestinationRuleDecorator) DecoratorName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecoratorName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DecoratorName indicates an expected call of DecoratorName.
func (mr *MockVirtualMeshDestinationRuleDecoratorMockRecorder) DecoratorName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecoratorName", reflect.TypeOf((*MockVirtualMeshDestinationRuleDecorator)(nil).DecoratorName))
}

// ApplyVirtualMeshToDestinationRule mocks base method.
func (m *MockVirtualMeshDestinationRuleDecorator) ApplyVirtualMeshToDestinationRule(appliedVirtualMesh *v1alpha2.MeshStatus_AppliedVirtualMesh, service *v1alpha2.MeshService, output *v1alpha3.DestinationRule, registerField decorators.RegisterField) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyVirtualMeshToDestinationRule", appliedVirtualMesh, service, output, registerField)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyVirtualMeshToDestinationRule indicates an expected call of ApplyVirtualMeshToDestinationRule.
func (mr *MockVirtualMeshDestinationRuleDecoratorMockRecorder) ApplyVirtualMeshToDestinationRule(appliedVirtualMesh, service, output, registerField interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyVirtualMeshToDestinationRule", reflect.TypeOf((*MockVirtualMeshDestinationRuleDecorator)(nil).ApplyVirtualMeshToDestinationRule), appliedVirtualMesh, service, output, registerField)
}

// MockVirtualMeshServiceEntryDecorator is a mock of VirtualMeshServiceEntryDecorator interface.
type MockVirtualMeshServiceEntryDecorator struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshServiceEntryDecoratorMockRecorder
}

// MockVirtualMeshServiceEntryDecoratorMockRecorder is the mock recorder for MockVirtualMeshServiceEntryDecorator.
type MockVirtualMeshServiceEntryDecoratorMockRecorder struct {
	mock *MockVirtualMeshServiceEntryDecorator
}

// NewMockVirtualMeshServiceEntryDecorator creates a new mock instance.
func NewMockVirtualMeshServiceEntryDecorator(ctrl *gomock.Controller) *MockVirtualMeshServiceEntryDecorator {
	mock := &MockVirtualMeshServiceEntryDecorator{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshServiceEntryDecoratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshServiceEntryDecorator) EXPECT() *MockVirtualMeshServiceEntryDecoratorMockRecorder {
	return m.recorder
}

// DecoratorName mocks base method.
func (m *MockVirtualMeshServiceEntryDecorator) DecoratorName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecoratorName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DecoratorName indicates an expected call of DecoratorName.
func (mr *MockVirtualMeshServiceEntryDecoratorMockRecorder) DecoratorName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecoratorName", reflect.TypeOf((*MockVirtualMeshServiceEntryDecorator)(nil).DecoratorName))
}

// ApplyVirtualMeshToServiceEntry mocks base method.
func (m *MockVirtualMeshServiceEntryDecorator) ApplyVirtualMeshToServiceEntry(appliedVirtualMesh *v1alpha2.MeshStatus_AppliedVirtualMesh, service *v1alpha2.MeshService, output *v1alpha3.ServiceEntry, registerField decorators.RegisterField) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyVirtualMeshToServiceEntry", appliedVirtualMesh, service, output, registerField)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyVirtualMeshToServiceEntry indicates an expected call of ApplyVirtualMeshToServiceEntry.
func (mr *MockVirtualMeshServiceEntryDecoratorMockRecorder) ApplyVirtualMeshToServiceEntry(appliedVirtualMesh, service, output, registerField interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyVirtualMeshToServiceEntry", reflect.TypeOf((*MockVirtualMeshServiceEntryDecorator)(nil).ApplyVirtualMeshToServiceEntry), appliedVirtualMesh, service, output, registerField)
}

// MockTrafficPolicyDestinationRuleDecorator is a mock of TrafficPolicyDestinationRuleDecorator interface.
type MockTrafficPolicyDestinationRuleDecorator struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyDestinationRuleDecoratorMockRecorder
}

// MockTrafficPolicyDestinationRuleDecoratorMockRecorder is the mock recorder for MockTrafficPolicyDestinationRuleDecorator.
type MockTrafficPolicyDestinationRuleDecoratorMockRecorder struct {
	mock *MockTrafficPolicyDestinationRuleDecorator
}

// NewMockTrafficPolicyDestinationRuleDecorator creates a new mock instance.
func NewMockTrafficPolicyDestinationRuleDecorator(ctrl *gomock.Controller) *MockTrafficPolicyDestinationRuleDecorator {
	mock := &MockTrafficPolicyDestinationRuleDecorator{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyDestinationRuleDecoratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyDestinationRuleDecorator) EXPECT() *MockTrafficPolicyDestinationRuleDecoratorMockRecorder {
	return m.recorder
}

// DecoratorName mocks base method.
func (m *MockTrafficPolicyDestinationRuleDecorator) DecoratorName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecoratorName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DecoratorName indicates an expected call of DecoratorName.
func (mr *MockTrafficPolicyDestinationRuleDecoratorMockRecorder) DecoratorName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecoratorName", reflect.TypeOf((*MockTrafficPolicyDestinationRuleDecorator)(nil).DecoratorName))
}

// ApplyTrafficPolicyToDestinationRule mocks base method.
func (m *MockTrafficPolicyDestinationRuleDecorator) ApplyTrafficPolicyToDestinationRule(appliedPolicy *v1alpha2.MeshServiceStatus_AppliedTrafficPolicy, service *v1alpha2.MeshService, output *v1alpha3.DestinationRule, registerField decorators.RegisterField) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyTrafficPolicyToDestinationRule", appliedPolicy, service, output, registerField)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyTrafficPolicyToDestinationRule indicates an expected call of ApplyTrafficPolicyToDestinationRule.
func (mr *MockTrafficPolicyDestinationRuleDecoratorMockRecorder) ApplyTrafficPolicyToDestinationRule(appliedPolicy, service, output, registerField interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyTrafficPolicyToDestinationRule", reflect.TypeOf((*MockTrafficPolicyDestinationRuleDecorator)(nil).ApplyTrafficPolicyToDestinationRule), appliedPolicy, service, output, registerField)
}

// MockTrafficPolicyVirtualServiceDecorator is a mock of TrafficPolicyVirtualServiceDecorator interface.
type MockTrafficPolicyVirtualServiceDecorator struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyVirtualServiceDecoratorMockRecorder
}

// MockTrafficPolicyVirtualServiceDecoratorMockRecorder is the mock recorder for MockTrafficPolicyVirtualServiceDecorator.
type MockTrafficPolicyVirtualServiceDecoratorMockRecorder struct {
	mock *MockTrafficPolicyVirtualServiceDecorator
}

// NewMockTrafficPolicyVirtualServiceDecorator creates a new mock instance.
func NewMockTrafficPolicyVirtualServiceDecorator(ctrl *gomock.Controller) *MockTrafficPolicyVirtualServiceDecorator {
	mock := &MockTrafficPolicyVirtualServiceDecorator{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyVirtualServiceDecoratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyVirtualServiceDecorator) EXPECT() *MockTrafficPolicyVirtualServiceDecoratorMockRecorder {
	return m.recorder
}

// DecoratorName mocks base method.
func (m *MockTrafficPolicyVirtualServiceDecorator) DecoratorName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecoratorName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DecoratorName indicates an expected call of DecoratorName.
func (mr *MockTrafficPolicyVirtualServiceDecoratorMockRecorder) DecoratorName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecoratorName", reflect.TypeOf((*MockTrafficPolicyVirtualServiceDecorator)(nil).DecoratorName))
}

// ApplyTrafficPolicyToVirtualService mocks base method.
func (m *MockTrafficPolicyVirtualServiceDecorator) ApplyTrafficPolicyToVirtualService(appliedPolicy *v1alpha2.MeshServiceStatus_AppliedTrafficPolicy, service *v1alpha2.MeshService, output *v1alpha3.HTTPRoute, registerField decorators.RegisterField) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyTrafficPolicyToVirtualService", appliedPolicy, service, output, registerField)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyTrafficPolicyToVirtualService indicates an expected call of ApplyTrafficPolicyToVirtualService.
func (mr *MockTrafficPolicyVirtualServiceDecoratorMockRecorder) ApplyTrafficPolicyToVirtualService(appliedPolicy, service, output, registerField interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyTrafficPolicyToVirtualService", reflect.TypeOf((*MockTrafficPolicyVirtualServiceDecorator)(nil).ApplyTrafficPolicyToVirtualService), appliedPolicy, service, output, registerField)
}
