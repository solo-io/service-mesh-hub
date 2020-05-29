// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_translation is a generated GoMock package.
package mock_translation

import (
	context "context"
	reflect "reflect"

	appmesh "github.com/aws/aws-sdk-go/service/appmesh"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	v1alpha10 "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	sets "github.com/solo-io/service-mesh-hub/pkg/collections/sets"
)

// MockAppmeshTranslationReconciler is a mock of AppmeshTranslationReconciler interface.
type MockAppmeshTranslationReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockAppmeshTranslationReconcilerMockRecorder
}

// MockAppmeshTranslationReconcilerMockRecorder is the mock recorder for MockAppmeshTranslationReconciler.
type MockAppmeshTranslationReconcilerMockRecorder struct {
	mock *MockAppmeshTranslationReconciler
}

// NewMockAppmeshTranslationReconciler creates a new mock instance.
func NewMockAppmeshTranslationReconciler(ctrl *gomock.Controller) *MockAppmeshTranslationReconciler {
	mock := &MockAppmeshTranslationReconciler{ctrl: ctrl}
	mock.recorder = &MockAppmeshTranslationReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppmeshTranslationReconciler) EXPECT() *MockAppmeshTranslationReconcilerMockRecorder {
	return m.recorder
}

// Reconcile mocks base method.
func (m *MockAppmeshTranslationReconciler) Reconcile(ctx context.Context, mesh *v1alpha1.Mesh, virtualMesh *v1alpha10.VirtualMesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", ctx, mesh, virtualMesh)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile.
func (mr *MockAppmeshTranslationReconcilerMockRecorder) Reconcile(ctx, mesh, virtualMesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockAppmeshTranslationReconciler)(nil).Reconcile), ctx, mesh, virtualMesh)
}

// MockAppmeshTranslator is a mock of AppmeshTranslator interface.
type MockAppmeshTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockAppmeshTranslatorMockRecorder
}

// MockAppmeshTranslatorMockRecorder is the mock recorder for MockAppmeshTranslator.
type MockAppmeshTranslatorMockRecorder struct {
	mock *MockAppmeshTranslator
}

// NewMockAppmeshTranslator creates a new mock instance.
func NewMockAppmeshTranslator(ctrl *gomock.Controller) *MockAppmeshTranslator {
	mock := &MockAppmeshTranslator{ctrl: ctrl}
	mock.recorder = &MockAppmeshTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppmeshTranslator) EXPECT() *MockAppmeshTranslatorMockRecorder {
	return m.recorder
}

// BuildVirtualNode mocks base method.
func (m *MockAppmeshTranslator) BuildVirtualNode(appmeshName *string, meshWorkload *v1alpha1.MeshWorkload, meshService *v1alpha1.MeshService, upstreamServices []*v1alpha1.MeshService) *appmesh.VirtualNodeData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildVirtualNode", appmeshName, meshWorkload, meshService, upstreamServices)
	ret0, _ := ret[0].(*appmesh.VirtualNodeData)
	return ret0
}

// BuildVirtualNode indicates an expected call of BuildVirtualNode.
func (mr *MockAppmeshTranslatorMockRecorder) BuildVirtualNode(appmeshName, meshWorkload, meshService, upstreamServices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildVirtualNode", reflect.TypeOf((*MockAppmeshTranslator)(nil).BuildVirtualNode), appmeshName, meshWorkload, meshService, upstreamServices)
}

// BuildRoute mocks base method.
func (m *MockAppmeshTranslator) BuildRoute(appmeshName *string, routeName string, priority int, meshService *v1alpha1.MeshService, targetedWorkloads []*v1alpha1.MeshWorkload) (*appmesh.RouteData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildRoute", appmeshName, routeName, priority, meshService, targetedWorkloads)
	ret0, _ := ret[0].(*appmesh.RouteData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildRoute indicates an expected call of BuildRoute.
func (mr *MockAppmeshTranslatorMockRecorder) BuildRoute(appmeshName, routeName, priority, meshService, targetedWorkloads interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildRoute", reflect.TypeOf((*MockAppmeshTranslator)(nil).BuildRoute), appmeshName, routeName, priority, meshService, targetedWorkloads)
}

// BuildVirtualService mocks base method.
func (m *MockAppmeshTranslator) BuildVirtualService(appmeshName *string, meshService *v1alpha1.MeshService) *appmesh.VirtualServiceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildVirtualService", appmeshName, meshService)
	ret0, _ := ret[0].(*appmesh.VirtualServiceData)
	return ret0
}

// BuildVirtualService indicates an expected call of BuildVirtualService.
func (mr *MockAppmeshTranslatorMockRecorder) BuildVirtualService(appmeshName, meshService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildVirtualService", reflect.TypeOf((*MockAppmeshTranslator)(nil).BuildVirtualService), appmeshName, meshService)
}

// BuildVirtualRouter mocks base method.
func (m *MockAppmeshTranslator) BuildVirtualRouter(appmeshName *string, meshService *v1alpha1.MeshService) *appmesh.VirtualRouterData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildVirtualRouter", appmeshName, meshService)
	ret0, _ := ret[0].(*appmesh.VirtualRouterData)
	return ret0
}

// BuildVirtualRouter indicates an expected call of BuildVirtualRouter.
func (mr *MockAppmeshTranslatorMockRecorder) BuildVirtualRouter(appmeshName, meshService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildVirtualRouter", reflect.TypeOf((*MockAppmeshTranslator)(nil).BuildVirtualRouter), appmeshName, meshService)
}

// MockAppmeshTranslationDao is a mock of AppmeshTranslationDao interface.
type MockAppmeshTranslationDao struct {
	ctrl     *gomock.Controller
	recorder *MockAppmeshTranslationDaoMockRecorder
}

// MockAppmeshTranslationDaoMockRecorder is the mock recorder for MockAppmeshTranslationDao.
type MockAppmeshTranslationDaoMockRecorder struct {
	mock *MockAppmeshTranslationDao
}

// NewMockAppmeshTranslationDao creates a new mock instance.
func NewMockAppmeshTranslationDao(ctrl *gomock.Controller) *MockAppmeshTranslationDao {
	mock := &MockAppmeshTranslationDao{ctrl: ctrl}
	mock.recorder = &MockAppmeshTranslationDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppmeshTranslationDao) EXPECT() *MockAppmeshTranslationDaoMockRecorder {
	return m.recorder
}

// GetAllServiceWorkloadPairsForMesh mocks base method.
func (m *MockAppmeshTranslationDao) GetAllServiceWorkloadPairsForMesh(ctx context.Context, mesh *v1alpha1.Mesh) (map[*v1alpha1.MeshService][]*v1alpha1.MeshWorkload, map[*v1alpha1.MeshWorkload][]*v1alpha1.MeshService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllServiceWorkloadPairsForMesh", ctx, mesh)
	ret0, _ := ret[0].(map[*v1alpha1.MeshService][]*v1alpha1.MeshWorkload)
	ret1, _ := ret[1].(map[*v1alpha1.MeshWorkload][]*v1alpha1.MeshService)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllServiceWorkloadPairsForMesh indicates an expected call of GetAllServiceWorkloadPairsForMesh.
func (mr *MockAppmeshTranslationDaoMockRecorder) GetAllServiceWorkloadPairsForMesh(ctx, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllServiceWorkloadPairsForMesh", reflect.TypeOf((*MockAppmeshTranslationDao)(nil).GetAllServiceWorkloadPairsForMesh), ctx, mesh)
}

// GetWorkloadsToAllUpstreamServices mocks base method.
func (m *MockAppmeshTranslationDao) GetWorkloadsToAllUpstreamServices(ctx context.Context, mesh *v1alpha1.Mesh) (map[*v1alpha1.MeshWorkload]sets.MeshServiceSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkloadsToAllUpstreamServices", ctx, mesh)
	ret0, _ := ret[0].(map[*v1alpha1.MeshWorkload]sets.MeshServiceSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkloadsToAllUpstreamServices indicates an expected call of GetWorkloadsToAllUpstreamServices.
func (mr *MockAppmeshTranslationDaoMockRecorder) GetWorkloadsToAllUpstreamServices(ctx, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkloadsToAllUpstreamServices", reflect.TypeOf((*MockAppmeshTranslationDao)(nil).GetWorkloadsToAllUpstreamServices), ctx, mesh)
}

// GetServicesWithACP mocks base method.
func (m *MockAppmeshTranslationDao) GetServicesWithACP(ctx context.Context, mesh *v1alpha1.Mesh) (sets.MeshServiceSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServicesWithACP", ctx, mesh)
	ret0, _ := ret[0].(sets.MeshServiceSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServicesWithACP indicates an expected call of GetServicesWithACP.
func (mr *MockAppmeshTranslationDaoMockRecorder) GetServicesWithACP(ctx, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServicesWithACP", reflect.TypeOf((*MockAppmeshTranslationDao)(nil).GetServicesWithACP), ctx, mesh)
}

// GetWorkloadsToUpstreamServicesWithACP mocks base method.
func (m *MockAppmeshTranslationDao) GetWorkloadsToUpstreamServicesWithACP(ctx context.Context, mesh *v1alpha1.Mesh) (sets.MeshWorkloadSet, map[*v1alpha1.MeshWorkload]sets.MeshServiceSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkloadsToUpstreamServicesWithACP", ctx, mesh)
	ret0, _ := ret[0].(sets.MeshWorkloadSet)
	ret1, _ := ret[1].(map[*v1alpha1.MeshWorkload]sets.MeshServiceSet)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetWorkloadsToUpstreamServicesWithACP indicates an expected call of GetWorkloadsToUpstreamServicesWithACP.
func (mr *MockAppmeshTranslationDaoMockRecorder) GetWorkloadsToUpstreamServicesWithACP(ctx, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkloadsToUpstreamServicesWithACP", reflect.TypeOf((*MockAppmeshTranslationDao)(nil).GetWorkloadsToUpstreamServicesWithACP), ctx, mesh)
}

// EnsureVirtualService mocks base method.
func (m *MockAppmeshTranslationDao) EnsureVirtualService(mesh *v1alpha1.Mesh, virtualServiceData *appmesh.VirtualServiceData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureVirtualService", mesh, virtualServiceData)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureVirtualService indicates an expected call of EnsureVirtualService.
func (mr *MockAppmeshTranslationDaoMockRecorder) EnsureVirtualService(mesh, virtualServiceData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureVirtualService", reflect.TypeOf((*MockAppmeshTranslationDao)(nil).EnsureVirtualService), mesh, virtualServiceData)
}

// EnsureVirtualRouter mocks base method.
func (m *MockAppmeshTranslationDao) EnsureVirtualRouter(mesh *v1alpha1.Mesh, virtualRouter *appmesh.VirtualRouterData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureVirtualRouter", mesh, virtualRouter)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureVirtualRouter indicates an expected call of EnsureVirtualRouter.
func (mr *MockAppmeshTranslationDaoMockRecorder) EnsureVirtualRouter(mesh, virtualRouter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureVirtualRouter", reflect.TypeOf((*MockAppmeshTranslationDao)(nil).EnsureVirtualRouter), mesh, virtualRouter)
}

// EnsureRoute mocks base method.
func (m *MockAppmeshTranslationDao) EnsureRoute(mesh *v1alpha1.Mesh, route *appmesh.RouteData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureRoute", mesh, route)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureRoute indicates an expected call of EnsureRoute.
func (mr *MockAppmeshTranslationDaoMockRecorder) EnsureRoute(mesh, route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureRoute", reflect.TypeOf((*MockAppmeshTranslationDao)(nil).EnsureRoute), mesh, route)
}

// EnsureVirtualNode mocks base method.
func (m *MockAppmeshTranslationDao) EnsureVirtualNode(mesh *v1alpha1.Mesh, virtualNode *appmesh.VirtualNodeData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureVirtualNode", mesh, virtualNode)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureVirtualNode indicates an expected call of EnsureVirtualNode.
func (mr *MockAppmeshTranslationDaoMockRecorder) EnsureVirtualNode(mesh, virtualNode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureVirtualNode", reflect.TypeOf((*MockAppmeshTranslationDao)(nil).EnsureVirtualNode), mesh, virtualNode)
}

// ReconcileVirtualServices mocks base method.
func (m *MockAppmeshTranslationDao) ReconcileVirtualServices(ctx context.Context, mesh *v1alpha1.Mesh, virtualServices []*appmesh.VirtualServiceData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualServices", ctx, mesh, virtualServices)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualServices indicates an expected call of ReconcileVirtualServices.
func (mr *MockAppmeshTranslationDaoMockRecorder) ReconcileVirtualServices(ctx, mesh, virtualServices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualServices", reflect.TypeOf((*MockAppmeshTranslationDao)(nil).ReconcileVirtualServices), ctx, mesh, virtualServices)
}

// ReconcileVirtualRouters mocks base method.
func (m *MockAppmeshTranslationDao) ReconcileVirtualRouters(ctx context.Context, mesh *v1alpha1.Mesh, virtualRouters []*appmesh.VirtualRouterData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualRouters", ctx, mesh, virtualRouters)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualRouters indicates an expected call of ReconcileVirtualRouters.
func (mr *MockAppmeshTranslationDaoMockRecorder) ReconcileVirtualRouters(ctx, mesh, virtualRouters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualRouters", reflect.TypeOf((*MockAppmeshTranslationDao)(nil).ReconcileVirtualRouters), ctx, mesh, virtualRouters)
}

// ReconcileRoutes mocks base method.
func (m *MockAppmeshTranslationDao) ReconcileRoutes(ctx context.Context, mesh *v1alpha1.Mesh, routes []*appmesh.RouteData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRoutes", ctx, mesh, routes)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileRoutes indicates an expected call of ReconcileRoutes.
func (mr *MockAppmeshTranslationDaoMockRecorder) ReconcileRoutes(ctx, mesh, routes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRoutes", reflect.TypeOf((*MockAppmeshTranslationDao)(nil).ReconcileRoutes), ctx, mesh, routes)
}

// ReconcileVirtualNodes mocks base method.
func (m *MockAppmeshTranslationDao) ReconcileVirtualNodes(ctx context.Context, mesh *v1alpha1.Mesh, virtualNodes []*appmesh.VirtualNodeData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualNodes", ctx, mesh, virtualNodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualNodes indicates an expected call of ReconcileVirtualNodes.
func (mr *MockAppmeshTranslationDaoMockRecorder) ReconcileVirtualNodes(ctx, mesh, virtualNodes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualNodes", reflect.TypeOf((*MockAppmeshTranslationDao)(nil).ReconcileVirtualNodes), ctx, mesh, virtualNodes)
}
