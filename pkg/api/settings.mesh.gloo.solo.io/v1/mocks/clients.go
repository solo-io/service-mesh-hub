// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockMulticlusterClientset is a mock of MulticlusterClientset interface
type MockMulticlusterClientset struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClientsetMockRecorder
}

// MockMulticlusterClientsetMockRecorder is the mock recorder for MockMulticlusterClientset
type MockMulticlusterClientsetMockRecorder struct {
	mock *MockMulticlusterClientset
}

// NewMockMulticlusterClientset creates a new mock instance
func NewMockMulticlusterClientset(ctrl *gomock.Controller) *MockMulticlusterClientset {
	mock := &MockMulticlusterClientset{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClientset) EXPECT() *MockMulticlusterClientsetMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterClientset) Cluster(cluster string) (v1.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1.Clientset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterClientsetMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterClientset)(nil).Cluster), cluster)
}

// MockClientset is a mock of Clientset interface
type MockClientset struct {
	ctrl     *gomock.Controller
	recorder *MockClientsetMockRecorder
}

// MockClientsetMockRecorder is the mock recorder for MockClientset
type MockClientsetMockRecorder struct {
	mock *MockClientset
}

// NewMockClientset creates a new mock instance
func NewMockClientset(ctrl *gomock.Controller) *MockClientset {
	mock := &MockClientset{ctrl: ctrl}
	mock.recorder = &MockClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientset) EXPECT() *MockClientsetMockRecorder {
	return m.recorder
}

// Settings mocks base method
func (m *MockClientset) Settings() v1.SettingsClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Settings")
	ret0, _ := ret[0].(v1.SettingsClient)
	return ret0
}

// Settings indicates an expected call of Settings
func (mr *MockClientsetMockRecorder) Settings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Settings", reflect.TypeOf((*MockClientset)(nil).Settings))
}

// MockSettingsReader is a mock of SettingsReader interface
type MockSettingsReader struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsReaderMockRecorder
}

// MockSettingsReaderMockRecorder is the mock recorder for MockSettingsReader
type MockSettingsReaderMockRecorder struct {
	mock *MockSettingsReader
}

// NewMockSettingsReader creates a new mock instance
func NewMockSettingsReader(ctrl *gomock.Controller) *MockSettingsReader {
	mock := &MockSettingsReader{ctrl: ctrl}
	mock.recorder = &MockSettingsReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsReader) EXPECT() *MockSettingsReaderMockRecorder {
	return m.recorder
}

// GetSettings mocks base method
func (m *MockSettingsReader) GetSettings(ctx context.Context, key client.ObjectKey) (*v1.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettings", ctx, key)
	ret0, _ := ret[0].(*v1.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettings indicates an expected call of GetSettings
func (mr *MockSettingsReaderMockRecorder) GetSettings(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockSettingsReader)(nil).GetSettings), ctx, key)
}

// ListSettings mocks base method
func (m *MockSettingsReader) ListSettings(ctx context.Context, opts ...client.ListOption) (*v1.SettingsList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSettings", varargs...)
	ret0, _ := ret[0].(*v1.SettingsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSettings indicates an expected call of ListSettings
func (mr *MockSettingsReaderMockRecorder) ListSettings(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSettings", reflect.TypeOf((*MockSettingsReader)(nil).ListSettings), varargs...)
}

// MockSettingsWriter is a mock of SettingsWriter interface
type MockSettingsWriter struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsWriterMockRecorder
}

// MockSettingsWriterMockRecorder is the mock recorder for MockSettingsWriter
type MockSettingsWriterMockRecorder struct {
	mock *MockSettingsWriter
}

// NewMockSettingsWriter creates a new mock instance
func NewMockSettingsWriter(ctrl *gomock.Controller) *MockSettingsWriter {
	mock := &MockSettingsWriter{ctrl: ctrl}
	mock.recorder = &MockSettingsWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsWriter) EXPECT() *MockSettingsWriterMockRecorder {
	return m.recorder
}

// CreateSettings mocks base method
func (m *MockSettingsWriter) CreateSettings(ctx context.Context, obj *v1.Settings, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSettings indicates an expected call of CreateSettings
func (mr *MockSettingsWriterMockRecorder) CreateSettings(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSettings", reflect.TypeOf((*MockSettingsWriter)(nil).CreateSettings), varargs...)
}

// DeleteSettings mocks base method
func (m *MockSettingsWriter) DeleteSettings(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSettings indicates an expected call of DeleteSettings
func (mr *MockSettingsWriterMockRecorder) DeleteSettings(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSettings", reflect.TypeOf((*MockSettingsWriter)(nil).DeleteSettings), varargs...)
}

// UpdateSettings mocks base method
func (m *MockSettingsWriter) UpdateSettings(ctx context.Context, obj *v1.Settings, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSettings indicates an expected call of UpdateSettings
func (mr *MockSettingsWriterMockRecorder) UpdateSettings(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettings", reflect.TypeOf((*MockSettingsWriter)(nil).UpdateSettings), varargs...)
}

// PatchSettings mocks base method
func (m *MockSettingsWriter) PatchSettings(ctx context.Context, obj *v1.Settings, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchSettings indicates an expected call of PatchSettings
func (mr *MockSettingsWriterMockRecorder) PatchSettings(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchSettings", reflect.TypeOf((*MockSettingsWriter)(nil).PatchSettings), varargs...)
}

// DeleteAllOfSettings mocks base method
func (m *MockSettingsWriter) DeleteAllOfSettings(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfSettings indicates an expected call of DeleteAllOfSettings
func (mr *MockSettingsWriterMockRecorder) DeleteAllOfSettings(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfSettings", reflect.TypeOf((*MockSettingsWriter)(nil).DeleteAllOfSettings), varargs...)
}

// UpsertSettings mocks base method
func (m *MockSettingsWriter) UpsertSettings(ctx context.Context, obj *v1.Settings, transitionFuncs ...v1.SettingsTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertSettings indicates an expected call of UpsertSettings
func (mr *MockSettingsWriterMockRecorder) UpsertSettings(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertSettings", reflect.TypeOf((*MockSettingsWriter)(nil).UpsertSettings), varargs...)
}

// MockSettingsStatusWriter is a mock of SettingsStatusWriter interface
type MockSettingsStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsStatusWriterMockRecorder
}

// MockSettingsStatusWriterMockRecorder is the mock recorder for MockSettingsStatusWriter
type MockSettingsStatusWriterMockRecorder struct {
	mock *MockSettingsStatusWriter
}

// NewMockSettingsStatusWriter creates a new mock instance
func NewMockSettingsStatusWriter(ctrl *gomock.Controller) *MockSettingsStatusWriter {
	mock := &MockSettingsStatusWriter{ctrl: ctrl}
	mock.recorder = &MockSettingsStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsStatusWriter) EXPECT() *MockSettingsStatusWriterMockRecorder {
	return m.recorder
}

// UpdateSettingsStatus mocks base method
func (m *MockSettingsStatusWriter) UpdateSettingsStatus(ctx context.Context, obj *v1.Settings, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSettingsStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSettingsStatus indicates an expected call of UpdateSettingsStatus
func (mr *MockSettingsStatusWriterMockRecorder) UpdateSettingsStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettingsStatus", reflect.TypeOf((*MockSettingsStatusWriter)(nil).UpdateSettingsStatus), varargs...)
}

// PatchSettingsStatus mocks base method
func (m *MockSettingsStatusWriter) PatchSettingsStatus(ctx context.Context, obj *v1.Settings, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchSettingsStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchSettingsStatus indicates an expected call of PatchSettingsStatus
func (mr *MockSettingsStatusWriterMockRecorder) PatchSettingsStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchSettingsStatus", reflect.TypeOf((*MockSettingsStatusWriter)(nil).PatchSettingsStatus), varargs...)
}

// MockSettingsClient is a mock of SettingsClient interface
type MockSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsClientMockRecorder
}

// MockSettingsClientMockRecorder is the mock recorder for MockSettingsClient
type MockSettingsClientMockRecorder struct {
	mock *MockSettingsClient
}

// NewMockSettingsClient creates a new mock instance
func NewMockSettingsClient(ctrl *gomock.Controller) *MockSettingsClient {
	mock := &MockSettingsClient{ctrl: ctrl}
	mock.recorder = &MockSettingsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsClient) EXPECT() *MockSettingsClientMockRecorder {
	return m.recorder
}

// GetSettings mocks base method
func (m *MockSettingsClient) GetSettings(ctx context.Context, key client.ObjectKey) (*v1.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettings", ctx, key)
	ret0, _ := ret[0].(*v1.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettings indicates an expected call of GetSettings
func (mr *MockSettingsClientMockRecorder) GetSettings(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockSettingsClient)(nil).GetSettings), ctx, key)
}

// ListSettings mocks base method
func (m *MockSettingsClient) ListSettings(ctx context.Context, opts ...client.ListOption) (*v1.SettingsList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSettings", varargs...)
	ret0, _ := ret[0].(*v1.SettingsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSettings indicates an expected call of ListSettings
func (mr *MockSettingsClientMockRecorder) ListSettings(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSettings", reflect.TypeOf((*MockSettingsClient)(nil).ListSettings), varargs...)
}

// CreateSettings mocks base method
func (m *MockSettingsClient) CreateSettings(ctx context.Context, obj *v1.Settings, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSettings indicates an expected call of CreateSettings
func (mr *MockSettingsClientMockRecorder) CreateSettings(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSettings", reflect.TypeOf((*MockSettingsClient)(nil).CreateSettings), varargs...)
}

// DeleteSettings mocks base method
func (m *MockSettingsClient) DeleteSettings(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSettings indicates an expected call of DeleteSettings
func (mr *MockSettingsClientMockRecorder) DeleteSettings(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSettings", reflect.TypeOf((*MockSettingsClient)(nil).DeleteSettings), varargs...)
}

// UpdateSettings mocks base method
func (m *MockSettingsClient) UpdateSettings(ctx context.Context, obj *v1.Settings, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSettings indicates an expected call of UpdateSettings
func (mr *MockSettingsClientMockRecorder) UpdateSettings(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettings", reflect.TypeOf((*MockSettingsClient)(nil).UpdateSettings), varargs...)
}

// PatchSettings mocks base method
func (m *MockSettingsClient) PatchSettings(ctx context.Context, obj *v1.Settings, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchSettings indicates an expected call of PatchSettings
func (mr *MockSettingsClientMockRecorder) PatchSettings(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchSettings", reflect.TypeOf((*MockSettingsClient)(nil).PatchSettings), varargs...)
}

// DeleteAllOfSettings mocks base method
func (m *MockSettingsClient) DeleteAllOfSettings(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfSettings indicates an expected call of DeleteAllOfSettings
func (mr *MockSettingsClientMockRecorder) DeleteAllOfSettings(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfSettings", reflect.TypeOf((*MockSettingsClient)(nil).DeleteAllOfSettings), varargs...)
}

// UpsertSettings mocks base method
func (m *MockSettingsClient) UpsertSettings(ctx context.Context, obj *v1.Settings, transitionFuncs ...v1.SettingsTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertSettings indicates an expected call of UpsertSettings
func (mr *MockSettingsClientMockRecorder) UpsertSettings(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertSettings", reflect.TypeOf((*MockSettingsClient)(nil).UpsertSettings), varargs...)
}

// UpdateSettingsStatus mocks base method
func (m *MockSettingsClient) UpdateSettingsStatus(ctx context.Context, obj *v1.Settings, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSettingsStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSettingsStatus indicates an expected call of UpdateSettingsStatus
func (mr *MockSettingsClientMockRecorder) UpdateSettingsStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettingsStatus", reflect.TypeOf((*MockSettingsClient)(nil).UpdateSettingsStatus), varargs...)
}

// PatchSettingsStatus mocks base method
func (m *MockSettingsClient) PatchSettingsStatus(ctx context.Context, obj *v1.Settings, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchSettingsStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchSettingsStatus indicates an expected call of PatchSettingsStatus
func (mr *MockSettingsClientMockRecorder) PatchSettingsStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchSettingsStatus", reflect.TypeOf((*MockSettingsClient)(nil).PatchSettingsStatus), varargs...)
}

// MockMulticlusterSettingsClient is a mock of MulticlusterSettingsClient interface
type MockMulticlusterSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterSettingsClientMockRecorder
}

// MockMulticlusterSettingsClientMockRecorder is the mock recorder for MockMulticlusterSettingsClient
type MockMulticlusterSettingsClientMockRecorder struct {
	mock *MockMulticlusterSettingsClient
}

// NewMockMulticlusterSettingsClient creates a new mock instance
func NewMockMulticlusterSettingsClient(ctrl *gomock.Controller) *MockMulticlusterSettingsClient {
	mock := &MockMulticlusterSettingsClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterSettingsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterSettingsClient) EXPECT() *MockMulticlusterSettingsClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterSettingsClient) Cluster(cluster string) (v1.SettingsClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1.SettingsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterSettingsClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterSettingsClient)(nil).Cluster), cluster)
}