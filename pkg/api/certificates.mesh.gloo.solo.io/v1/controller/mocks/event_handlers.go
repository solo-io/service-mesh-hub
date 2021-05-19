// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"
	controller "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockIssuedCertificateEventHandler is a mock of IssuedCertificateEventHandler interface.
type MockIssuedCertificateEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIssuedCertificateEventHandlerMockRecorder
}

// MockIssuedCertificateEventHandlerMockRecorder is the mock recorder for MockIssuedCertificateEventHandler.
type MockIssuedCertificateEventHandlerMockRecorder struct {
	mock *MockIssuedCertificateEventHandler
}

// NewMockIssuedCertificateEventHandler creates a new mock instance.
func NewMockIssuedCertificateEventHandler(ctrl *gomock.Controller) *MockIssuedCertificateEventHandler {
	mock := &MockIssuedCertificateEventHandler{ctrl: ctrl}
	mock.recorder = &MockIssuedCertificateEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssuedCertificateEventHandler) EXPECT() *MockIssuedCertificateEventHandlerMockRecorder {
	return m.recorder
}

// CreateIssuedCertificate mocks base method.
func (m *MockIssuedCertificateEventHandler) CreateIssuedCertificate(obj *v1.IssuedCertificate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIssuedCertificate", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIssuedCertificate indicates an expected call of CreateIssuedCertificate.
func (mr *MockIssuedCertificateEventHandlerMockRecorder) CreateIssuedCertificate(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIssuedCertificate", reflect.TypeOf((*MockIssuedCertificateEventHandler)(nil).CreateIssuedCertificate), obj)
}

// DeleteIssuedCertificate mocks base method.
func (m *MockIssuedCertificateEventHandler) DeleteIssuedCertificate(obj *v1.IssuedCertificate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIssuedCertificate", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIssuedCertificate indicates an expected call of DeleteIssuedCertificate.
func (mr *MockIssuedCertificateEventHandlerMockRecorder) DeleteIssuedCertificate(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIssuedCertificate", reflect.TypeOf((*MockIssuedCertificateEventHandler)(nil).DeleteIssuedCertificate), obj)
}

// GenericIssuedCertificate mocks base method.
func (m *MockIssuedCertificateEventHandler) GenericIssuedCertificate(obj *v1.IssuedCertificate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericIssuedCertificate", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericIssuedCertificate indicates an expected call of GenericIssuedCertificate.
func (mr *MockIssuedCertificateEventHandlerMockRecorder) GenericIssuedCertificate(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericIssuedCertificate", reflect.TypeOf((*MockIssuedCertificateEventHandler)(nil).GenericIssuedCertificate), obj)
}

// UpdateIssuedCertificate mocks base method.
func (m *MockIssuedCertificateEventHandler) UpdateIssuedCertificate(old, new *v1.IssuedCertificate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIssuedCertificate", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIssuedCertificate indicates an expected call of UpdateIssuedCertificate.
func (mr *MockIssuedCertificateEventHandlerMockRecorder) UpdateIssuedCertificate(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIssuedCertificate", reflect.TypeOf((*MockIssuedCertificateEventHandler)(nil).UpdateIssuedCertificate), old, new)
}

// MockIssuedCertificateEventWatcher is a mock of IssuedCertificateEventWatcher interface.
type MockIssuedCertificateEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockIssuedCertificateEventWatcherMockRecorder
}

// MockIssuedCertificateEventWatcherMockRecorder is the mock recorder for MockIssuedCertificateEventWatcher.
type MockIssuedCertificateEventWatcherMockRecorder struct {
	mock *MockIssuedCertificateEventWatcher
}

// NewMockIssuedCertificateEventWatcher creates a new mock instance.
func NewMockIssuedCertificateEventWatcher(ctrl *gomock.Controller) *MockIssuedCertificateEventWatcher {
	mock := &MockIssuedCertificateEventWatcher{ctrl: ctrl}
	mock.recorder = &MockIssuedCertificateEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssuedCertificateEventWatcher) EXPECT() *MockIssuedCertificateEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockIssuedCertificateEventWatcher) AddEventHandler(ctx context.Context, h controller.IssuedCertificateEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockIssuedCertificateEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockIssuedCertificateEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockCertificateRequestEventHandler is a mock of CertificateRequestEventHandler interface.
type MockCertificateRequestEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateRequestEventHandlerMockRecorder
}

// MockCertificateRequestEventHandlerMockRecorder is the mock recorder for MockCertificateRequestEventHandler.
type MockCertificateRequestEventHandlerMockRecorder struct {
	mock *MockCertificateRequestEventHandler
}

// NewMockCertificateRequestEventHandler creates a new mock instance.
func NewMockCertificateRequestEventHandler(ctrl *gomock.Controller) *MockCertificateRequestEventHandler {
	mock := &MockCertificateRequestEventHandler{ctrl: ctrl}
	mock.recorder = &MockCertificateRequestEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCertificateRequestEventHandler) EXPECT() *MockCertificateRequestEventHandlerMockRecorder {
	return m.recorder
}

// CreateCertificateRequest mocks base method.
func (m *MockCertificateRequestEventHandler) CreateCertificateRequest(obj *v1.CertificateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCertificateRequest", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCertificateRequest indicates an expected call of CreateCertificateRequest.
func (mr *MockCertificateRequestEventHandlerMockRecorder) CreateCertificateRequest(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCertificateRequest", reflect.TypeOf((*MockCertificateRequestEventHandler)(nil).CreateCertificateRequest), obj)
}

// DeleteCertificateRequest mocks base method.
func (m *MockCertificateRequestEventHandler) DeleteCertificateRequest(obj *v1.CertificateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCertificateRequest", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCertificateRequest indicates an expected call of DeleteCertificateRequest.
func (mr *MockCertificateRequestEventHandlerMockRecorder) DeleteCertificateRequest(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCertificateRequest", reflect.TypeOf((*MockCertificateRequestEventHandler)(nil).DeleteCertificateRequest), obj)
}

// GenericCertificateRequest mocks base method.
func (m *MockCertificateRequestEventHandler) GenericCertificateRequest(obj *v1.CertificateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericCertificateRequest", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericCertificateRequest indicates an expected call of GenericCertificateRequest.
func (mr *MockCertificateRequestEventHandlerMockRecorder) GenericCertificateRequest(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericCertificateRequest", reflect.TypeOf((*MockCertificateRequestEventHandler)(nil).GenericCertificateRequest), obj)
}

// UpdateCertificateRequest mocks base method.
func (m *MockCertificateRequestEventHandler) UpdateCertificateRequest(old, new *v1.CertificateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCertificateRequest", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCertificateRequest indicates an expected call of UpdateCertificateRequest.
func (mr *MockCertificateRequestEventHandlerMockRecorder) UpdateCertificateRequest(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCertificateRequest", reflect.TypeOf((*MockCertificateRequestEventHandler)(nil).UpdateCertificateRequest), old, new)
}

// MockCertificateRequestEventWatcher is a mock of CertificateRequestEventWatcher interface.
type MockCertificateRequestEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateRequestEventWatcherMockRecorder
}

// MockCertificateRequestEventWatcherMockRecorder is the mock recorder for MockCertificateRequestEventWatcher.
type MockCertificateRequestEventWatcherMockRecorder struct {
	mock *MockCertificateRequestEventWatcher
}

// NewMockCertificateRequestEventWatcher creates a new mock instance.
func NewMockCertificateRequestEventWatcher(ctrl *gomock.Controller) *MockCertificateRequestEventWatcher {
	mock := &MockCertificateRequestEventWatcher{ctrl: ctrl}
	mock.recorder = &MockCertificateRequestEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCertificateRequestEventWatcher) EXPECT() *MockCertificateRequestEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockCertificateRequestEventWatcher) AddEventHandler(ctx context.Context, h controller.CertificateRequestEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockCertificateRequestEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockCertificateRequestEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockPodBounceDirectiveEventHandler is a mock of PodBounceDirectiveEventHandler interface.
type MockPodBounceDirectiveEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPodBounceDirectiveEventHandlerMockRecorder
}

// MockPodBounceDirectiveEventHandlerMockRecorder is the mock recorder for MockPodBounceDirectiveEventHandler.
type MockPodBounceDirectiveEventHandlerMockRecorder struct {
	mock *MockPodBounceDirectiveEventHandler
}

// NewMockPodBounceDirectiveEventHandler creates a new mock instance.
func NewMockPodBounceDirectiveEventHandler(ctrl *gomock.Controller) *MockPodBounceDirectiveEventHandler {
	mock := &MockPodBounceDirectiveEventHandler{ctrl: ctrl}
	mock.recorder = &MockPodBounceDirectiveEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodBounceDirectiveEventHandler) EXPECT() *MockPodBounceDirectiveEventHandlerMockRecorder {
	return m.recorder
}

// CreatePodBounceDirective mocks base method.
func (m *MockPodBounceDirectiveEventHandler) CreatePodBounceDirective(obj *v1.PodBounceDirective) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePodBounceDirective", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePodBounceDirective indicates an expected call of CreatePodBounceDirective.
func (mr *MockPodBounceDirectiveEventHandlerMockRecorder) CreatePodBounceDirective(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePodBounceDirective", reflect.TypeOf((*MockPodBounceDirectiveEventHandler)(nil).CreatePodBounceDirective), obj)
}

// DeletePodBounceDirective mocks base method.
func (m *MockPodBounceDirectiveEventHandler) DeletePodBounceDirective(obj *v1.PodBounceDirective) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePodBounceDirective", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePodBounceDirective indicates an expected call of DeletePodBounceDirective.
func (mr *MockPodBounceDirectiveEventHandlerMockRecorder) DeletePodBounceDirective(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePodBounceDirective", reflect.TypeOf((*MockPodBounceDirectiveEventHandler)(nil).DeletePodBounceDirective), obj)
}

// GenericPodBounceDirective mocks base method.
func (m *MockPodBounceDirectiveEventHandler) GenericPodBounceDirective(obj *v1.PodBounceDirective) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericPodBounceDirective", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericPodBounceDirective indicates an expected call of GenericPodBounceDirective.
func (mr *MockPodBounceDirectiveEventHandlerMockRecorder) GenericPodBounceDirective(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericPodBounceDirective", reflect.TypeOf((*MockPodBounceDirectiveEventHandler)(nil).GenericPodBounceDirective), obj)
}

// UpdatePodBounceDirective mocks base method.
func (m *MockPodBounceDirectiveEventHandler) UpdatePodBounceDirective(old, new *v1.PodBounceDirective) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePodBounceDirective", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePodBounceDirective indicates an expected call of UpdatePodBounceDirective.
func (mr *MockPodBounceDirectiveEventHandlerMockRecorder) UpdatePodBounceDirective(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePodBounceDirective", reflect.TypeOf((*MockPodBounceDirectiveEventHandler)(nil).UpdatePodBounceDirective), old, new)
}

// MockPodBounceDirectiveEventWatcher is a mock of PodBounceDirectiveEventWatcher interface.
type MockPodBounceDirectiveEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockPodBounceDirectiveEventWatcherMockRecorder
}

// MockPodBounceDirectiveEventWatcherMockRecorder is the mock recorder for MockPodBounceDirectiveEventWatcher.
type MockPodBounceDirectiveEventWatcherMockRecorder struct {
	mock *MockPodBounceDirectiveEventWatcher
}

// NewMockPodBounceDirectiveEventWatcher creates a new mock instance.
func NewMockPodBounceDirectiveEventWatcher(ctrl *gomock.Controller) *MockPodBounceDirectiveEventWatcher {
	mock := &MockPodBounceDirectiveEventWatcher{ctrl: ctrl}
	mock.recorder = &MockPodBounceDirectiveEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodBounceDirectiveEventWatcher) EXPECT() *MockPodBounceDirectiveEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockPodBounceDirectiveEventWatcher) AddEventHandler(ctx context.Context, h controller.PodBounceDirectiveEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockPodBounceDirectiveEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockPodBounceDirectiveEventWatcher)(nil).AddEventHandler), varargs...)
}
