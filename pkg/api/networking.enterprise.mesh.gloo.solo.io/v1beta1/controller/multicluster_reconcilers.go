// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	networking_enterprise_mesh_gloo_solo_io_v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the WasmDeployment Resource across clusters.
// implemented by the user
type MulticlusterWasmDeploymentReconciler interface {
	ReconcileWasmDeployment(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) (reconcile.Result, error)
}

// Reconcile deletion events for the WasmDeployment Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterWasmDeploymentDeletionReconciler interface {
	ReconcileWasmDeploymentDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterWasmDeploymentReconcilerFuncs struct {
	OnReconcileWasmDeployment         func(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) (reconcile.Result, error)
	OnReconcileWasmDeploymentDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterWasmDeploymentReconcilerFuncs) ReconcileWasmDeployment(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) (reconcile.Result, error) {
	if f.OnReconcileWasmDeployment == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileWasmDeployment(clusterName, obj)
}

func (f *MulticlusterWasmDeploymentReconcilerFuncs) ReconcileWasmDeploymentDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileWasmDeploymentDeletion == nil {
		return nil
	}
	return f.OnReconcileWasmDeploymentDeletion(clusterName, req)
}

type MulticlusterWasmDeploymentReconcileLoop interface {
	// AddMulticlusterWasmDeploymentReconciler adds a MulticlusterWasmDeploymentReconciler to the MulticlusterWasmDeploymentReconcileLoop.
	AddMulticlusterWasmDeploymentReconciler(ctx context.Context, rec MulticlusterWasmDeploymentReconciler, predicates ...predicate.Predicate)
}

type multiclusterWasmDeploymentReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterWasmDeploymentReconcileLoop) AddMulticlusterWasmDeploymentReconciler(ctx context.Context, rec MulticlusterWasmDeploymentReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericWasmDeploymentMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterWasmDeploymentReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterWasmDeploymentReconcileLoop {
	return &multiclusterWasmDeploymentReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment{}, options)}
}

type genericWasmDeploymentMulticlusterReconciler struct {
	reconciler MulticlusterWasmDeploymentReconciler
}

func (g genericWasmDeploymentMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterWasmDeploymentDeletionReconciler); ok {
		return deletionReconciler.ReconcileWasmDeploymentDeletion(cluster, req)
	}
	return nil
}

func (g genericWasmDeploymentMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: WasmDeployment handler received event for %T", object)
	}
	return g.reconciler.ReconcileWasmDeployment(cluster, obj)
}

// Reconcile Upsert events for the VirtualDestination Resource across clusters.
// implemented by the user
type MulticlusterVirtualDestinationReconciler interface {
	ReconcileVirtualDestination(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualDestination Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualDestinationDeletionReconciler interface {
	ReconcileVirtualDestinationDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterVirtualDestinationReconcilerFuncs struct {
	OnReconcileVirtualDestination         func(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) (reconcile.Result, error)
	OnReconcileVirtualDestinationDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterVirtualDestinationReconcilerFuncs) ReconcileVirtualDestination(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) (reconcile.Result, error) {
	if f.OnReconcileVirtualDestination == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualDestination(clusterName, obj)
}

func (f *MulticlusterVirtualDestinationReconcilerFuncs) ReconcileVirtualDestinationDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileVirtualDestinationDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualDestinationDeletion(clusterName, req)
}

type MulticlusterVirtualDestinationReconcileLoop interface {
	// AddMulticlusterVirtualDestinationReconciler adds a MulticlusterVirtualDestinationReconciler to the MulticlusterVirtualDestinationReconcileLoop.
	AddMulticlusterVirtualDestinationReconciler(ctx context.Context, rec MulticlusterVirtualDestinationReconciler, predicates ...predicate.Predicate)
}

type multiclusterVirtualDestinationReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterVirtualDestinationReconcileLoop) AddMulticlusterVirtualDestinationReconciler(ctx context.Context, rec MulticlusterVirtualDestinationReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericVirtualDestinationMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterVirtualDestinationReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterVirtualDestinationReconcileLoop {
	return &multiclusterVirtualDestinationReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination{}, options)}
}

type genericVirtualDestinationMulticlusterReconciler struct {
	reconciler MulticlusterVirtualDestinationReconciler
}

func (g genericVirtualDestinationMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterVirtualDestinationDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualDestinationDeletion(cluster, req)
	}
	return nil
}

func (g genericVirtualDestinationMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualDestination handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualDestination(cluster, obj)
}

// Reconcile Upsert events for the VirtualGateway Resource across clusters.
// implemented by the user
type MulticlusterVirtualGatewayReconciler interface {
	ReconcileVirtualGateway(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualGateway Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualGatewayDeletionReconciler interface {
	ReconcileVirtualGatewayDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterVirtualGatewayReconcilerFuncs struct {
	OnReconcileVirtualGateway         func(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway) (reconcile.Result, error)
	OnReconcileVirtualGatewayDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterVirtualGatewayReconcilerFuncs) ReconcileVirtualGateway(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway) (reconcile.Result, error) {
	if f.OnReconcileVirtualGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualGateway(clusterName, obj)
}

func (f *MulticlusterVirtualGatewayReconcilerFuncs) ReconcileVirtualGatewayDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileVirtualGatewayDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualGatewayDeletion(clusterName, req)
}

type MulticlusterVirtualGatewayReconcileLoop interface {
	// AddMulticlusterVirtualGatewayReconciler adds a MulticlusterVirtualGatewayReconciler to the MulticlusterVirtualGatewayReconcileLoop.
	AddMulticlusterVirtualGatewayReconciler(ctx context.Context, rec MulticlusterVirtualGatewayReconciler, predicates ...predicate.Predicate)
}

type multiclusterVirtualGatewayReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterVirtualGatewayReconcileLoop) AddMulticlusterVirtualGatewayReconciler(ctx context.Context, rec MulticlusterVirtualGatewayReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericVirtualGatewayMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterVirtualGatewayReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterVirtualGatewayReconcileLoop {
	return &multiclusterVirtualGatewayReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway{}, options)}
}

type genericVirtualGatewayMulticlusterReconciler struct {
	reconciler MulticlusterVirtualGatewayReconciler
}

func (g genericVirtualGatewayMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterVirtualGatewayDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualGatewayDeletion(cluster, req)
	}
	return nil
}

func (g genericVirtualGatewayMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualGateway handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualGateway(cluster, obj)
}

// Reconcile Upsert events for the VirtualHost Resource across clusters.
// implemented by the user
type MulticlusterVirtualHostReconciler interface {
	ReconcileVirtualHost(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualHost Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualHostDeletionReconciler interface {
	ReconcileVirtualHostDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterVirtualHostReconcilerFuncs struct {
	OnReconcileVirtualHost         func(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost) (reconcile.Result, error)
	OnReconcileVirtualHostDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterVirtualHostReconcilerFuncs) ReconcileVirtualHost(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost) (reconcile.Result, error) {
	if f.OnReconcileVirtualHost == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualHost(clusterName, obj)
}

func (f *MulticlusterVirtualHostReconcilerFuncs) ReconcileVirtualHostDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileVirtualHostDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualHostDeletion(clusterName, req)
}

type MulticlusterVirtualHostReconcileLoop interface {
	// AddMulticlusterVirtualHostReconciler adds a MulticlusterVirtualHostReconciler to the MulticlusterVirtualHostReconcileLoop.
	AddMulticlusterVirtualHostReconciler(ctx context.Context, rec MulticlusterVirtualHostReconciler, predicates ...predicate.Predicate)
}

type multiclusterVirtualHostReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterVirtualHostReconcileLoop) AddMulticlusterVirtualHostReconciler(ctx context.Context, rec MulticlusterVirtualHostReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericVirtualHostMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterVirtualHostReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterVirtualHostReconcileLoop {
	return &multiclusterVirtualHostReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost{}, options)}
}

type genericVirtualHostMulticlusterReconciler struct {
	reconciler MulticlusterVirtualHostReconciler
}

func (g genericVirtualHostMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterVirtualHostDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualHostDeletion(cluster, req)
	}
	return nil
}

func (g genericVirtualHostMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualHost handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualHost(cluster, obj)
}

// Reconcile Upsert events for the RouteTable Resource across clusters.
// implemented by the user
type MulticlusterRouteTableReconciler interface {
	ReconcileRouteTable(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) (reconcile.Result, error)
}

// Reconcile deletion events for the RouteTable Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterRouteTableDeletionReconciler interface {
	ReconcileRouteTableDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterRouteTableReconcilerFuncs struct {
	OnReconcileRouteTable         func(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) (reconcile.Result, error)
	OnReconcileRouteTableDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterRouteTableReconcilerFuncs) ReconcileRouteTable(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) (reconcile.Result, error) {
	if f.OnReconcileRouteTable == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileRouteTable(clusterName, obj)
}

func (f *MulticlusterRouteTableReconcilerFuncs) ReconcileRouteTableDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileRouteTableDeletion == nil {
		return nil
	}
	return f.OnReconcileRouteTableDeletion(clusterName, req)
}

type MulticlusterRouteTableReconcileLoop interface {
	// AddMulticlusterRouteTableReconciler adds a MulticlusterRouteTableReconciler to the MulticlusterRouteTableReconcileLoop.
	AddMulticlusterRouteTableReconciler(ctx context.Context, rec MulticlusterRouteTableReconciler, predicates ...predicate.Predicate)
}

type multiclusterRouteTableReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterRouteTableReconcileLoop) AddMulticlusterRouteTableReconciler(ctx context.Context, rec MulticlusterRouteTableReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericRouteTableMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterRouteTableReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterRouteTableReconcileLoop {
	return &multiclusterRouteTableReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable{}, options)}
}

type genericRouteTableMulticlusterReconciler struct {
	reconciler MulticlusterRouteTableReconciler
}

func (g genericRouteTableMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterRouteTableDeletionReconciler); ok {
		return deletionReconciler.ReconcileRouteTableDeletion(cluster, req)
	}
	return nil
}

func (g genericRouteTableMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: RouteTable handler received event for %T", object)
	}
	return g.reconciler.ReconcileRouteTable(cluster, obj)
}

// Reconcile Upsert events for the ServiceDependency Resource across clusters.
// implemented by the user
type MulticlusterServiceDependencyReconciler interface {
	ReconcileServiceDependency(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.ServiceDependency) (reconcile.Result, error)
}

// Reconcile deletion events for the ServiceDependency Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterServiceDependencyDeletionReconciler interface {
	ReconcileServiceDependencyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterServiceDependencyReconcilerFuncs struct {
	OnReconcileServiceDependency         func(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.ServiceDependency) (reconcile.Result, error)
	OnReconcileServiceDependencyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterServiceDependencyReconcilerFuncs) ReconcileServiceDependency(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1beta1.ServiceDependency) (reconcile.Result, error) {
	if f.OnReconcileServiceDependency == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileServiceDependency(clusterName, obj)
}

func (f *MulticlusterServiceDependencyReconcilerFuncs) ReconcileServiceDependencyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileServiceDependencyDeletion == nil {
		return nil
	}
	return f.OnReconcileServiceDependencyDeletion(clusterName, req)
}

type MulticlusterServiceDependencyReconcileLoop interface {
	// AddMulticlusterServiceDependencyReconciler adds a MulticlusterServiceDependencyReconciler to the MulticlusterServiceDependencyReconcileLoop.
	AddMulticlusterServiceDependencyReconciler(ctx context.Context, rec MulticlusterServiceDependencyReconciler, predicates ...predicate.Predicate)
}

type multiclusterServiceDependencyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterServiceDependencyReconcileLoop) AddMulticlusterServiceDependencyReconciler(ctx context.Context, rec MulticlusterServiceDependencyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericServiceDependencyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterServiceDependencyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterServiceDependencyReconcileLoop {
	return &multiclusterServiceDependencyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_enterprise_mesh_gloo_solo_io_v1beta1.ServiceDependency{}, options)}
}

type genericServiceDependencyMulticlusterReconciler struct {
	reconciler MulticlusterServiceDependencyReconciler
}

func (g genericServiceDependencyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterServiceDependencyDeletionReconciler); ok {
		return deletionReconciler.ReconcileServiceDependencyDeletion(cluster, req)
	}
	return nil
}

func (g genericServiceDependencyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.ServiceDependency)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ServiceDependency handler received event for %T", object)
	}
	return g.reconciler.ReconcileServiceDependency(cluster, obj)
}
