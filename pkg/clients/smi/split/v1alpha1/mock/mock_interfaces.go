// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_v1alpha1 is a generated GoMock package.
package mock_v1alpha1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockTrafficSplitClient is a mock of TrafficSplitClient interface.
type MockTrafficSplitClient struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficSplitClientMockRecorder
}

// MockTrafficSplitClientMockRecorder is the mock recorder for MockTrafficSplitClient.
type MockTrafficSplitClientMockRecorder struct {
	mock *MockTrafficSplitClient
}

// NewMockTrafficSplitClient creates a new mock instance.
func NewMockTrafficSplitClient(ctrl *gomock.Controller) *MockTrafficSplitClient {
	mock := &MockTrafficSplitClient{ctrl: ctrl}
	mock.recorder = &MockTrafficSplitClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficSplitClient) EXPECT() *MockTrafficSplitClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockTrafficSplitClient) Get(ctx context.Context, key client.ObjectKey) (*v1alpha1.TrafficSplit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.TrafficSplit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTrafficSplitClientMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTrafficSplitClient)(nil).Get), ctx, key)
}

// Create mocks base method.
func (m *MockTrafficSplitClient) Create(ctx context.Context, trafficSplitClient *v1alpha1.TrafficSplit, options ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, trafficSplitClient}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTrafficSplitClientMockRecorder) Create(ctx, trafficSplitClient interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, trafficSplitClient}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTrafficSplitClient)(nil).Create), varargs...)
}

// List mocks base method.
func (m *MockTrafficSplitClient) List(ctx context.Context, options ...client.ListOption) (*v1alpha1.TrafficSplitList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v1alpha1.TrafficSplitList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTrafficSplitClientMockRecorder) List(ctx interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTrafficSplitClient)(nil).List), varargs...)
}

// Update mocks base method.
func (m *MockTrafficSplitClient) Update(ctx context.Context, trafficSplitClient *v1alpha1.TrafficSplit, options ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, trafficSplitClient}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTrafficSplitClientMockRecorder) Update(ctx, trafficSplitClient interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, trafficSplitClient}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTrafficSplitClient)(nil).Update), varargs...)
}

// UpsertSpec mocks base method.
func (m *MockTrafficSplitClient) UpsertSpec(ctx context.Context, trafficSplitClient *v1alpha1.TrafficSplit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertSpec", ctx, trafficSplitClient)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertSpec indicates an expected call of UpsertSpec.
func (mr *MockTrafficSplitClientMockRecorder) UpsertSpec(ctx, trafficSplitClient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertSpec", reflect.TypeOf((*MockTrafficSplitClient)(nil).UpsertSpec), ctx, trafficSplitClient)
}

// Delete mocks base method.
func (m *MockTrafficSplitClient) Delete(ctx context.Context, key client.ObjectKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTrafficSplitClientMockRecorder) Delete(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrafficSplitClient)(nil).Delete), ctx, key)
}
