// Code generated by MockGen. DO NOT EDIT.
// Source: ./failover_service_validation.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	failoverservice "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/validation/failoverservice"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"
)

// MockFailoverServiceValidator is a mock of FailoverServiceValidator interface
type MockFailoverServiceValidator struct {
	ctrl     *gomock.Controller
	recorder *MockFailoverServiceValidatorMockRecorder
}

// MockFailoverServiceValidatorMockRecorder is the mock recorder for MockFailoverServiceValidator
type MockFailoverServiceValidatorMockRecorder struct {
	mock *MockFailoverServiceValidator
}

// NewMockFailoverServiceValidator creates a new mock instance
func NewMockFailoverServiceValidator(ctrl *gomock.Controller) *MockFailoverServiceValidator {
	mock := &MockFailoverServiceValidator{ctrl: ctrl}
	mock.recorder = &MockFailoverServiceValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFailoverServiceValidator) EXPECT() *MockFailoverServiceValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockFailoverServiceValidator) Validate(inputs failoverservice.Inputs, failoverService *v1alpha2.FailoverServiceSpec) []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", inputs, failoverService)
	ret0, _ := ret[0].([]error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockFailoverServiceValidatorMockRecorder) Validate(inputs, failoverService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockFailoverServiceValidator)(nil).Validate), inputs, failoverService)
}
