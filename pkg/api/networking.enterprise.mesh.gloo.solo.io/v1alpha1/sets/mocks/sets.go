// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha1sets is a generated GoMock package.
package mock_v1alpha1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1alpha1"
	v1alpha1sets "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1alpha1/sets"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockWasmDeploymentSet is a mock of WasmDeploymentSet interface
type MockWasmDeploymentSet struct {
	ctrl     *gomock.Controller
	recorder *MockWasmDeploymentSetMockRecorder
}

// MockWasmDeploymentSetMockRecorder is the mock recorder for MockWasmDeploymentSet
type MockWasmDeploymentSetMockRecorder struct {
	mock *MockWasmDeploymentSet
}

// NewMockWasmDeploymentSet creates a new mock instance
func NewMockWasmDeploymentSet(ctrl *gomock.Controller) *MockWasmDeploymentSet {
	mock := &MockWasmDeploymentSet{ctrl: ctrl}
	mock.recorder = &MockWasmDeploymentSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWasmDeploymentSet) EXPECT() *MockWasmDeploymentSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockWasmDeploymentSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockWasmDeploymentSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Keys))
}

// List mocks base method
func (m *MockWasmDeploymentSet) List(filterResource ...func(*v1alpha1.WasmDeployment) bool) []*v1alpha1.WasmDeployment {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1alpha1.WasmDeployment)
	return ret0
}

// List indicates an expected call of List
func (mr *MockWasmDeploymentSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWasmDeploymentSet)(nil).List), filterResource...)
}

// Map mocks base method
func (m *MockWasmDeploymentSet) Map() map[string]*v1alpha1.WasmDeployment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.WasmDeployment)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockWasmDeploymentSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Map))
}

// Insert mocks base method
func (m *MockWasmDeploymentSet) Insert(wasmDeployment ...*v1alpha1.WasmDeployment) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range wasmDeployment {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockWasmDeploymentSetMockRecorder) Insert(wasmDeployment ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Insert), wasmDeployment...)
}

// Equal mocks base method
func (m *MockWasmDeploymentSet) Equal(wasmDeploymentSet v1alpha1sets.WasmDeploymentSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", wasmDeploymentSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockWasmDeploymentSetMockRecorder) Equal(wasmDeploymentSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Equal), wasmDeploymentSet)
}

// Has mocks base method
func (m *MockWasmDeploymentSet) Has(wasmDeployment ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", wasmDeployment)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockWasmDeploymentSetMockRecorder) Has(wasmDeployment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Has), wasmDeployment)
}

// Delete mocks base method
func (m *MockWasmDeploymentSet) Delete(wasmDeployment ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", wasmDeployment)
}

// Delete indicates an expected call of Delete
func (mr *MockWasmDeploymentSetMockRecorder) Delete(wasmDeployment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Delete), wasmDeployment)
}

// Union mocks base method
func (m *MockWasmDeploymentSet) Union(set v1alpha1sets.WasmDeploymentSet) v1alpha1sets.WasmDeploymentSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.WasmDeploymentSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockWasmDeploymentSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockWasmDeploymentSet) Difference(set v1alpha1sets.WasmDeploymentSet) v1alpha1sets.WasmDeploymentSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.WasmDeploymentSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockWasmDeploymentSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockWasmDeploymentSet) Intersection(set v1alpha1sets.WasmDeploymentSet) v1alpha1sets.WasmDeploymentSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.WasmDeploymentSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockWasmDeploymentSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockWasmDeploymentSet) Find(id ezkube.ResourceId) (*v1alpha1.WasmDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha1.WasmDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockWasmDeploymentSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockWasmDeploymentSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockWasmDeploymentSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Length))
}

// Generic mocks base method
func (m *MockWasmDeploymentSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockWasmDeploymentSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Generic))
}

// Delta mocks base method
func (m *MockWasmDeploymentSet) Delta(newSet v1alpha1sets.WasmDeploymentSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta
func (mr *MockWasmDeploymentSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockWasmDeploymentSet)(nil).Delta), newSet)
}

// MockVirtualDestinationSet is a mock of VirtualDestinationSet interface
type MockVirtualDestinationSet struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualDestinationSetMockRecorder
}

// MockVirtualDestinationSetMockRecorder is the mock recorder for MockVirtualDestinationSet
type MockVirtualDestinationSetMockRecorder struct {
	mock *MockVirtualDestinationSet
}

// NewMockVirtualDestinationSet creates a new mock instance
func NewMockVirtualDestinationSet(ctrl *gomock.Controller) *MockVirtualDestinationSet {
	mock := &MockVirtualDestinationSet{ctrl: ctrl}
	mock.recorder = &MockVirtualDestinationSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualDestinationSet) EXPECT() *MockVirtualDestinationSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockVirtualDestinationSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockVirtualDestinationSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Keys))
}

// List mocks base method
func (m *MockVirtualDestinationSet) List(filterResource ...func(*v1alpha1.VirtualDestination) bool) []*v1alpha1.VirtualDestination {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1alpha1.VirtualDestination)
	return ret0
}

// List indicates an expected call of List
func (mr *MockVirtualDestinationSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualDestinationSet)(nil).List), filterResource...)
}

// Map mocks base method
func (m *MockVirtualDestinationSet) Map() map[string]*v1alpha1.VirtualDestination {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.VirtualDestination)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockVirtualDestinationSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Map))
}

// Insert mocks base method
func (m *MockVirtualDestinationSet) Insert(virtualDestination ...*v1alpha1.VirtualDestination) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range virtualDestination {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockVirtualDestinationSetMockRecorder) Insert(virtualDestination ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Insert), virtualDestination...)
}

// Equal mocks base method
func (m *MockVirtualDestinationSet) Equal(virtualDestinationSet v1alpha1sets.VirtualDestinationSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", virtualDestinationSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockVirtualDestinationSetMockRecorder) Equal(virtualDestinationSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Equal), virtualDestinationSet)
}

// Has mocks base method
func (m *MockVirtualDestinationSet) Has(virtualDestination ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", virtualDestination)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockVirtualDestinationSetMockRecorder) Has(virtualDestination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Has), virtualDestination)
}

// Delete mocks base method
func (m *MockVirtualDestinationSet) Delete(virtualDestination ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", virtualDestination)
}

// Delete indicates an expected call of Delete
func (mr *MockVirtualDestinationSetMockRecorder) Delete(virtualDestination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Delete), virtualDestination)
}

// Union mocks base method
func (m *MockVirtualDestinationSet) Union(set v1alpha1sets.VirtualDestinationSet) v1alpha1sets.VirtualDestinationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.VirtualDestinationSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockVirtualDestinationSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockVirtualDestinationSet) Difference(set v1alpha1sets.VirtualDestinationSet) v1alpha1sets.VirtualDestinationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.VirtualDestinationSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockVirtualDestinationSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockVirtualDestinationSet) Intersection(set v1alpha1sets.VirtualDestinationSet) v1alpha1sets.VirtualDestinationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.VirtualDestinationSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockVirtualDestinationSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockVirtualDestinationSet) Find(id ezkube.ResourceId) (*v1alpha1.VirtualDestination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha1.VirtualDestination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockVirtualDestinationSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockVirtualDestinationSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockVirtualDestinationSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Length))
}

// Generic mocks base method
func (m *MockVirtualDestinationSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockVirtualDestinationSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Generic))
}

// Delta mocks base method
func (m *MockVirtualDestinationSet) Delta(newSet v1alpha1sets.VirtualDestinationSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta
func (mr *MockVirtualDestinationSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockVirtualDestinationSet)(nil).Delta), newSet)
}