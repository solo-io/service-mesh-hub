// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	xds_agent_enterprise_mesh_gloo_solo_io_v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1beta1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the XdsConfig Resource across clusters.
// implemented by the user
type MulticlusterXdsConfigReconciler interface {
	ReconcileXdsConfig(clusterName string, obj *xds_agent_enterprise_mesh_gloo_solo_io_v1beta1.XdsConfig) (reconcile.Result, error)
}

// Reconcile deletion events for the XdsConfig Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterXdsConfigDeletionReconciler interface {
	ReconcileXdsConfigDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterXdsConfigReconcilerFuncs struct {
	OnReconcileXdsConfig         func(clusterName string, obj *xds_agent_enterprise_mesh_gloo_solo_io_v1beta1.XdsConfig) (reconcile.Result, error)
	OnReconcileXdsConfigDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterXdsConfigReconcilerFuncs) ReconcileXdsConfig(clusterName string, obj *xds_agent_enterprise_mesh_gloo_solo_io_v1beta1.XdsConfig) (reconcile.Result, error) {
	if f.OnReconcileXdsConfig == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileXdsConfig(clusterName, obj)
}

func (f *MulticlusterXdsConfigReconcilerFuncs) ReconcileXdsConfigDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileXdsConfigDeletion == nil {
		return nil
	}
	return f.OnReconcileXdsConfigDeletion(clusterName, req)
}

type MulticlusterXdsConfigReconcileLoop interface {
	// AddMulticlusterXdsConfigReconciler adds a MulticlusterXdsConfigReconciler to the MulticlusterXdsConfigReconcileLoop.
	AddMulticlusterXdsConfigReconciler(ctx context.Context, rec MulticlusterXdsConfigReconciler, predicates ...predicate.Predicate)
}

type multiclusterXdsConfigReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterXdsConfigReconcileLoop) AddMulticlusterXdsConfigReconciler(ctx context.Context, rec MulticlusterXdsConfigReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericXdsConfigMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterXdsConfigReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterXdsConfigReconcileLoop {
	return &multiclusterXdsConfigReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &xds_agent_enterprise_mesh_gloo_solo_io_v1beta1.XdsConfig{}, options)}
}

type genericXdsConfigMulticlusterReconciler struct {
	reconciler MulticlusterXdsConfigReconciler
}

func (g genericXdsConfigMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterXdsConfigDeletionReconciler); ok {
		return deletionReconciler.ReconcileXdsConfigDeletion(cluster, req)
	}
	return nil
}

func (g genericXdsConfigMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*xds_agent_enterprise_mesh_gloo_solo_io_v1beta1.XdsConfig)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: XdsConfig handler received event for %T", object)
	}
	return g.reconciler.ReconcileXdsConfig(cluster, obj)
}