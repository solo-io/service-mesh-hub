// Code generated by MockGen. DO NOT EDIT.
// Source: ./mesh_service_discovery.go

// Package service_discovery_mocks is a generated GoMock package.
package service_discovery_mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMeshServiceDiscovery is a mock of MeshServiceDiscovery interface.
type MockMeshServiceDiscovery struct {
	ctrl     *gomock.Controller
	recorder *MockMeshServiceDiscoveryMockRecorder
}

// MockMeshServiceDiscoveryMockRecorder is the mock recorder for MockMeshServiceDiscovery.
type MockMeshServiceDiscoveryMockRecorder struct {
	mock *MockMeshServiceDiscovery
}

// NewMockMeshServiceDiscovery creates a new mock instance.
func NewMockMeshServiceDiscovery(ctrl *gomock.Controller) *MockMeshServiceDiscovery {
	mock := &MockMeshServiceDiscovery{ctrl: ctrl}
	mock.recorder = &MockMeshServiceDiscoveryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshServiceDiscovery) EXPECT() *MockMeshServiceDiscoveryMockRecorder {
	return m.recorder
}

// DiscoverMeshServices mocks base method.
func (m *MockMeshServiceDiscovery) DiscoverMeshServices(ctx context.Context, clusterName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoverMeshServices", ctx, clusterName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DiscoverMeshServices indicates an expected call of DiscoverMeshServices.
func (mr *MockMeshServiceDiscoveryMockRecorder) DiscoverMeshServices(ctx, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverMeshServices", reflect.TypeOf((*MockMeshServiceDiscovery)(nil).DiscoverMeshServices), ctx, clusterName)
}
