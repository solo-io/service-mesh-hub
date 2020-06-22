// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	networking_smh_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the TrafficPolicy Resource across clusters.
// implemented by the user
type MulticlusterTrafficPolicyReconciler interface {
	ReconcileTrafficPolicy(clusterName string, obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the TrafficPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterTrafficPolicyDeletionReconciler interface {
	ReconcileTrafficPolicyDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterTrafficPolicyReconcilerFuncs struct {
	OnReconcileTrafficPolicy         func(clusterName string, obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) (reconcile.Result, error)
	OnReconcileTrafficPolicyDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterTrafficPolicyReconcilerFuncs) ReconcileTrafficPolicy(clusterName string, obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) (reconcile.Result, error) {
	if f.OnReconcileTrafficPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileTrafficPolicy(clusterName, obj)
}

func (f *MulticlusterTrafficPolicyReconcilerFuncs) ReconcileTrafficPolicyDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileTrafficPolicyDeletion == nil {
		return
	}
	f.OnReconcileTrafficPolicyDeletion(clusterName, req)
}

type MulticlusterTrafficPolicyReconcileLoop interface {
	// AddMulticlusterTrafficPolicyReconciler adds a MulticlusterTrafficPolicyReconciler to the MulticlusterTrafficPolicyReconcileLoop.
	AddMulticlusterTrafficPolicyReconciler(ctx context.Context, rec MulticlusterTrafficPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterTrafficPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterTrafficPolicyReconcileLoop) AddMulticlusterTrafficPolicyReconciler(ctx context.Context, rec MulticlusterTrafficPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericTrafficPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterTrafficPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterTrafficPolicyReconcileLoop {
	return &multiclusterTrafficPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_smh_solo_io_v1alpha1.TrafficPolicy{})}
}

type genericTrafficPolicyMulticlusterReconciler struct {
	reconciler MulticlusterTrafficPolicyReconciler
}

func (g genericTrafficPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterTrafficPolicyDeletionReconciler); ok {
		deletionReconciler.ReconcileTrafficPolicyDeletion(cluster, req)
	}
}

func (g genericTrafficPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.TrafficPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: TrafficPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileTrafficPolicy(cluster, obj)
}

// Reconcile Upsert events for the AccessControlPolicy Resource across clusters.
// implemented by the user
type MulticlusterAccessControlPolicyReconciler interface {
	ReconcileAccessControlPolicy(clusterName string, obj *networking_smh_solo_io_v1alpha1.AccessControlPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the AccessControlPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterAccessControlPolicyDeletionReconciler interface {
	ReconcileAccessControlPolicyDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterAccessControlPolicyReconcilerFuncs struct {
	OnReconcileAccessControlPolicy         func(clusterName string, obj *networking_smh_solo_io_v1alpha1.AccessControlPolicy) (reconcile.Result, error)
	OnReconcileAccessControlPolicyDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterAccessControlPolicyReconcilerFuncs) ReconcileAccessControlPolicy(clusterName string, obj *networking_smh_solo_io_v1alpha1.AccessControlPolicy) (reconcile.Result, error) {
	if f.OnReconcileAccessControlPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAccessControlPolicy(clusterName, obj)
}

func (f *MulticlusterAccessControlPolicyReconcilerFuncs) ReconcileAccessControlPolicyDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileAccessControlPolicyDeletion == nil {
		return
	}
	f.OnReconcileAccessControlPolicyDeletion(clusterName, req)
}

type MulticlusterAccessControlPolicyReconcileLoop interface {
	// AddMulticlusterAccessControlPolicyReconciler adds a MulticlusterAccessControlPolicyReconciler to the MulticlusterAccessControlPolicyReconcileLoop.
	AddMulticlusterAccessControlPolicyReconciler(ctx context.Context, rec MulticlusterAccessControlPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterAccessControlPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterAccessControlPolicyReconcileLoop) AddMulticlusterAccessControlPolicyReconciler(ctx context.Context, rec MulticlusterAccessControlPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericAccessControlPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterAccessControlPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterAccessControlPolicyReconcileLoop {
	return &multiclusterAccessControlPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_smh_solo_io_v1alpha1.AccessControlPolicy{})}
}

type genericAccessControlPolicyMulticlusterReconciler struct {
	reconciler MulticlusterAccessControlPolicyReconciler
}

func (g genericAccessControlPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterAccessControlPolicyDeletionReconciler); ok {
		deletionReconciler.ReconcileAccessControlPolicyDeletion(cluster, req)
	}
}

func (g genericAccessControlPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.AccessControlPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AccessControlPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileAccessControlPolicy(cluster, obj)
}

// Reconcile Upsert events for the VirtualMesh Resource across clusters.
// implemented by the user
type MulticlusterVirtualMeshReconciler interface {
	ReconcileVirtualMesh(clusterName string, obj *networking_smh_solo_io_v1alpha1.VirtualMesh) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualMesh Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualMeshDeletionReconciler interface {
	ReconcileVirtualMeshDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterVirtualMeshReconcilerFuncs struct {
	OnReconcileVirtualMesh         func(clusterName string, obj *networking_smh_solo_io_v1alpha1.VirtualMesh) (reconcile.Result, error)
	OnReconcileVirtualMeshDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterVirtualMeshReconcilerFuncs) ReconcileVirtualMesh(clusterName string, obj *networking_smh_solo_io_v1alpha1.VirtualMesh) (reconcile.Result, error) {
	if f.OnReconcileVirtualMesh == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualMesh(clusterName, obj)
}

func (f *MulticlusterVirtualMeshReconcilerFuncs) ReconcileVirtualMeshDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileVirtualMeshDeletion == nil {
		return
	}
	f.OnReconcileVirtualMeshDeletion(clusterName, req)
}

type MulticlusterVirtualMeshReconcileLoop interface {
	// AddMulticlusterVirtualMeshReconciler adds a MulticlusterVirtualMeshReconciler to the MulticlusterVirtualMeshReconcileLoop.
	AddMulticlusterVirtualMeshReconciler(ctx context.Context, rec MulticlusterVirtualMeshReconciler, predicates ...predicate.Predicate)
}

type multiclusterVirtualMeshReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterVirtualMeshReconcileLoop) AddMulticlusterVirtualMeshReconciler(ctx context.Context, rec MulticlusterVirtualMeshReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericVirtualMeshMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterVirtualMeshReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterVirtualMeshReconcileLoop {
	return &multiclusterVirtualMeshReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_smh_solo_io_v1alpha1.VirtualMesh{})}
}

type genericVirtualMeshMulticlusterReconciler struct {
	reconciler MulticlusterVirtualMeshReconciler
}

func (g genericVirtualMeshMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterVirtualMeshDeletionReconciler); ok {
		deletionReconciler.ReconcileVirtualMeshDeletion(cluster, req)
	}
}

func (g genericVirtualMeshMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.VirtualMesh)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualMesh handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualMesh(cluster, obj)
}

// Reconcile Upsert events for the FailoverService Resource across clusters.
// implemented by the user
type MulticlusterFailoverServiceReconciler interface {
	ReconcileFailoverService(clusterName string, obj *networking_smh_solo_io_v1alpha1.FailoverService) (reconcile.Result, error)
}

// Reconcile deletion events for the FailoverService Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterFailoverServiceDeletionReconciler interface {
	ReconcileFailoverServiceDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterFailoverServiceReconcilerFuncs struct {
	OnReconcileFailoverService         func(clusterName string, obj *networking_smh_solo_io_v1alpha1.FailoverService) (reconcile.Result, error)
	OnReconcileFailoverServiceDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterFailoverServiceReconcilerFuncs) ReconcileFailoverService(clusterName string, obj *networking_smh_solo_io_v1alpha1.FailoverService) (reconcile.Result, error) {
	if f.OnReconcileFailoverService == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFailoverService(clusterName, obj)
}

func (f *MulticlusterFailoverServiceReconcilerFuncs) ReconcileFailoverServiceDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileFailoverServiceDeletion == nil {
		return
	}
	f.OnReconcileFailoverServiceDeletion(clusterName, req)
}

type MulticlusterFailoverServiceReconcileLoop interface {
	// AddMulticlusterFailoverServiceReconciler adds a MulticlusterFailoverServiceReconciler to the MulticlusterFailoverServiceReconcileLoop.
	AddMulticlusterFailoverServiceReconciler(ctx context.Context, rec MulticlusterFailoverServiceReconciler, predicates ...predicate.Predicate)
}

type multiclusterFailoverServiceReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterFailoverServiceReconcileLoop) AddMulticlusterFailoverServiceReconciler(ctx context.Context, rec MulticlusterFailoverServiceReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericFailoverServiceMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterFailoverServiceReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterFailoverServiceReconcileLoop {
	return &multiclusterFailoverServiceReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_smh_solo_io_v1alpha1.FailoverService{})}
}

type genericFailoverServiceMulticlusterReconciler struct {
	reconciler MulticlusterFailoverServiceReconciler
}

func (g genericFailoverServiceMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterFailoverServiceDeletionReconciler); ok {
		deletionReconciler.ReconcileFailoverServiceDeletion(cluster, req)
	}
}

func (g genericFailoverServiceMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.FailoverService)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FailoverService handler received event for %T", object)
	}
	return g.reconciler.ReconcileFailoverService(cluster, obj)
}
