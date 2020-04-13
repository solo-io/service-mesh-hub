// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	discovery_zephyr_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the KubernetesCluster Resource.
// implemented by the user
type KubernetesClusterReconciler interface {
	ReconcileKubernetesCluster(obj *discovery_zephyr_solo_io_v1alpha1.KubernetesCluster) (reconcile.Result, error)
}

// Reconcile deletion events for the KubernetesCluster Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type KubernetesClusterDeletionReconciler interface {
	ReconcileKubernetesClusterDeletion(req reconcile.Request)
}

type KubernetesClusterReconcilerFuncs struct {
	OnReconcileKubernetesCluster         func(obj *discovery_zephyr_solo_io_v1alpha1.KubernetesCluster) (reconcile.Result, error)
	OnReconcileKubernetesClusterDeletion func(req reconcile.Request)
}

func (f *KubernetesClusterReconcilerFuncs) ReconcileKubernetesCluster(obj *discovery_zephyr_solo_io_v1alpha1.KubernetesCluster) (reconcile.Result, error) {
	if f.OnReconcileKubernetesCluster == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileKubernetesCluster(obj)
}

func (f *KubernetesClusterReconcilerFuncs) ReconcileKubernetesClusterDeletion(req reconcile.Request) {
	if f.OnReconcileKubernetesClusterDeletion == nil {
		return
	}
	f.OnReconcileKubernetesClusterDeletion(req)
}

// Reconcile and finalize the KubernetesCluster Resource
// implemented by the user
type KubernetesClusterFinalizer interface {
	KubernetesClusterReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	KubernetesClusterFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeKubernetesCluster(obj *discovery_zephyr_solo_io_v1alpha1.KubernetesCluster) error
}

type KubernetesClusterReconcileLoop interface {
	RunKubernetesClusterReconciler(ctx context.Context, rec KubernetesClusterReconciler, predicates ...predicate.Predicate) error
}

type kubernetesClusterReconcileLoop struct {
	loop reconcile.Loop
}

func NewKubernetesClusterReconcileLoop(name string, mgr manager.Manager) KubernetesClusterReconcileLoop {
	return &kubernetesClusterReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &discovery_zephyr_solo_io_v1alpha1.KubernetesCluster{}),
	}
}

func (c *kubernetesClusterReconcileLoop) RunKubernetesClusterReconciler(ctx context.Context, reconciler KubernetesClusterReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericKubernetesClusterReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(KubernetesClusterFinalizer); ok {
		reconcilerWrapper = genericKubernetesClusterFinalizer{
			genericKubernetesClusterReconciler: genericReconciler,
			finalizingReconciler:               finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericKubernetesClusterHandler implements a generic reconcile.Reconciler
type genericKubernetesClusterReconciler struct {
	reconciler KubernetesClusterReconciler
}

func (r genericKubernetesClusterReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*discovery_zephyr_solo_io_v1alpha1.KubernetesCluster)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: KubernetesCluster handler received event for %T", object)
	}
	return r.reconciler.ReconcileKubernetesCluster(obj)
}

func (r genericKubernetesClusterReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(KubernetesClusterDeletionReconciler); ok {
		deletionReconciler.ReconcileKubernetesClusterDeletion(request)
	}
}

// genericKubernetesClusterFinalizer implements a generic reconcile.FinalizingReconciler
type genericKubernetesClusterFinalizer struct {
	genericKubernetesClusterReconciler
	finalizingReconciler KubernetesClusterFinalizer
}

func (r genericKubernetesClusterFinalizer) FinalizerName() string {
	return r.finalizingReconciler.KubernetesClusterFinalizerName()
}

func (r genericKubernetesClusterFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*discovery_zephyr_solo_io_v1alpha1.KubernetesCluster)
	if !ok {
		return errors.Errorf("internal error: KubernetesCluster handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeKubernetesCluster(obj)
}

// Reconcile Upsert events for the MeshService Resource.
// implemented by the user
type MeshServiceReconciler interface {
	ReconcileMeshService(obj *discovery_zephyr_solo_io_v1alpha1.MeshService) (reconcile.Result, error)
}

// Reconcile deletion events for the MeshService Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MeshServiceDeletionReconciler interface {
	ReconcileMeshServiceDeletion(req reconcile.Request)
}

type MeshServiceReconcilerFuncs struct {
	OnReconcileMeshService         func(obj *discovery_zephyr_solo_io_v1alpha1.MeshService) (reconcile.Result, error)
	OnReconcileMeshServiceDeletion func(req reconcile.Request)
}

func (f *MeshServiceReconcilerFuncs) ReconcileMeshService(obj *discovery_zephyr_solo_io_v1alpha1.MeshService) (reconcile.Result, error) {
	if f.OnReconcileMeshService == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileMeshService(obj)
}

func (f *MeshServiceReconcilerFuncs) ReconcileMeshServiceDeletion(req reconcile.Request) {
	if f.OnReconcileMeshServiceDeletion == nil {
		return
	}
	f.OnReconcileMeshServiceDeletion(req)
}

// Reconcile and finalize the MeshService Resource
// implemented by the user
type MeshServiceFinalizer interface {
	MeshServiceReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	MeshServiceFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeMeshService(obj *discovery_zephyr_solo_io_v1alpha1.MeshService) error
}

type MeshServiceReconcileLoop interface {
	RunMeshServiceReconciler(ctx context.Context, rec MeshServiceReconciler, predicates ...predicate.Predicate) error
}

type meshServiceReconcileLoop struct {
	loop reconcile.Loop
}

func NewMeshServiceReconcileLoop(name string, mgr manager.Manager) MeshServiceReconcileLoop {
	return &meshServiceReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &discovery_zephyr_solo_io_v1alpha1.MeshService{}),
	}
}

func (c *meshServiceReconcileLoop) RunMeshServiceReconciler(ctx context.Context, reconciler MeshServiceReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericMeshServiceReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(MeshServiceFinalizer); ok {
		reconcilerWrapper = genericMeshServiceFinalizer{
			genericMeshServiceReconciler: genericReconciler,
			finalizingReconciler:         finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericMeshServiceHandler implements a generic reconcile.Reconciler
type genericMeshServiceReconciler struct {
	reconciler MeshServiceReconciler
}

func (r genericMeshServiceReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*discovery_zephyr_solo_io_v1alpha1.MeshService)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: MeshService handler received event for %T", object)
	}
	return r.reconciler.ReconcileMeshService(obj)
}

func (r genericMeshServiceReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(MeshServiceDeletionReconciler); ok {
		deletionReconciler.ReconcileMeshServiceDeletion(request)
	}
}

// genericMeshServiceFinalizer implements a generic reconcile.FinalizingReconciler
type genericMeshServiceFinalizer struct {
	genericMeshServiceReconciler
	finalizingReconciler MeshServiceFinalizer
}

func (r genericMeshServiceFinalizer) FinalizerName() string {
	return r.finalizingReconciler.MeshServiceFinalizerName()
}

func (r genericMeshServiceFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*discovery_zephyr_solo_io_v1alpha1.MeshService)
	if !ok {
		return errors.Errorf("internal error: MeshService handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeMeshService(obj)
}

// Reconcile Upsert events for the MeshWorkload Resource.
// implemented by the user
type MeshWorkloadReconciler interface {
	ReconcileMeshWorkload(obj *discovery_zephyr_solo_io_v1alpha1.MeshWorkload) (reconcile.Result, error)
}

// Reconcile deletion events for the MeshWorkload Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MeshWorkloadDeletionReconciler interface {
	ReconcileMeshWorkloadDeletion(req reconcile.Request)
}

type MeshWorkloadReconcilerFuncs struct {
	OnReconcileMeshWorkload         func(obj *discovery_zephyr_solo_io_v1alpha1.MeshWorkload) (reconcile.Result, error)
	OnReconcileMeshWorkloadDeletion func(req reconcile.Request)
}

func (f *MeshWorkloadReconcilerFuncs) ReconcileMeshWorkload(obj *discovery_zephyr_solo_io_v1alpha1.MeshWorkload) (reconcile.Result, error) {
	if f.OnReconcileMeshWorkload == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileMeshWorkload(obj)
}

func (f *MeshWorkloadReconcilerFuncs) ReconcileMeshWorkloadDeletion(req reconcile.Request) {
	if f.OnReconcileMeshWorkloadDeletion == nil {
		return
	}
	f.OnReconcileMeshWorkloadDeletion(req)
}

// Reconcile and finalize the MeshWorkload Resource
// implemented by the user
type MeshWorkloadFinalizer interface {
	MeshWorkloadReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	MeshWorkloadFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeMeshWorkload(obj *discovery_zephyr_solo_io_v1alpha1.MeshWorkload) error
}

type MeshWorkloadReconcileLoop interface {
	RunMeshWorkloadReconciler(ctx context.Context, rec MeshWorkloadReconciler, predicates ...predicate.Predicate) error
}

type meshWorkloadReconcileLoop struct {
	loop reconcile.Loop
}

func NewMeshWorkloadReconcileLoop(name string, mgr manager.Manager) MeshWorkloadReconcileLoop {
	return &meshWorkloadReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &discovery_zephyr_solo_io_v1alpha1.MeshWorkload{}),
	}
}

func (c *meshWorkloadReconcileLoop) RunMeshWorkloadReconciler(ctx context.Context, reconciler MeshWorkloadReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericMeshWorkloadReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(MeshWorkloadFinalizer); ok {
		reconcilerWrapper = genericMeshWorkloadFinalizer{
			genericMeshWorkloadReconciler: genericReconciler,
			finalizingReconciler:          finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericMeshWorkloadHandler implements a generic reconcile.Reconciler
type genericMeshWorkloadReconciler struct {
	reconciler MeshWorkloadReconciler
}

func (r genericMeshWorkloadReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*discovery_zephyr_solo_io_v1alpha1.MeshWorkload)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: MeshWorkload handler received event for %T", object)
	}
	return r.reconciler.ReconcileMeshWorkload(obj)
}

func (r genericMeshWorkloadReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(MeshWorkloadDeletionReconciler); ok {
		deletionReconciler.ReconcileMeshWorkloadDeletion(request)
	}
}

// genericMeshWorkloadFinalizer implements a generic reconcile.FinalizingReconciler
type genericMeshWorkloadFinalizer struct {
	genericMeshWorkloadReconciler
	finalizingReconciler MeshWorkloadFinalizer
}

func (r genericMeshWorkloadFinalizer) FinalizerName() string {
	return r.finalizingReconciler.MeshWorkloadFinalizerName()
}

func (r genericMeshWorkloadFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*discovery_zephyr_solo_io_v1alpha1.MeshWorkload)
	if !ok {
		return errors.Errorf("internal error: MeshWorkload handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeMeshWorkload(obj)
}

// Reconcile Upsert events for the Mesh Resource.
// implemented by the user
type MeshReconciler interface {
	ReconcileMesh(obj *discovery_zephyr_solo_io_v1alpha1.Mesh) (reconcile.Result, error)
}

// Reconcile deletion events for the Mesh Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MeshDeletionReconciler interface {
	ReconcileMeshDeletion(req reconcile.Request)
}

type MeshReconcilerFuncs struct {
	OnReconcileMesh         func(obj *discovery_zephyr_solo_io_v1alpha1.Mesh) (reconcile.Result, error)
	OnReconcileMeshDeletion func(req reconcile.Request)
}

func (f *MeshReconcilerFuncs) ReconcileMesh(obj *discovery_zephyr_solo_io_v1alpha1.Mesh) (reconcile.Result, error) {
	if f.OnReconcileMesh == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileMesh(obj)
}

func (f *MeshReconcilerFuncs) ReconcileMeshDeletion(req reconcile.Request) {
	if f.OnReconcileMeshDeletion == nil {
		return
	}
	f.OnReconcileMeshDeletion(req)
}

// Reconcile and finalize the Mesh Resource
// implemented by the user
type MeshFinalizer interface {
	MeshReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	MeshFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeMesh(obj *discovery_zephyr_solo_io_v1alpha1.Mesh) error
}

type MeshReconcileLoop interface {
	RunMeshReconciler(ctx context.Context, rec MeshReconciler, predicates ...predicate.Predicate) error
}

type meshReconcileLoop struct {
	loop reconcile.Loop
}

func NewMeshReconcileLoop(name string, mgr manager.Manager) MeshReconcileLoop {
	return &meshReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &discovery_zephyr_solo_io_v1alpha1.Mesh{}),
	}
}

func (c *meshReconcileLoop) RunMeshReconciler(ctx context.Context, reconciler MeshReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericMeshReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(MeshFinalizer); ok {
		reconcilerWrapper = genericMeshFinalizer{
			genericMeshReconciler: genericReconciler,
			finalizingReconciler:  finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericMeshHandler implements a generic reconcile.Reconciler
type genericMeshReconciler struct {
	reconciler MeshReconciler
}

func (r genericMeshReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*discovery_zephyr_solo_io_v1alpha1.Mesh)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Mesh handler received event for %T", object)
	}
	return r.reconciler.ReconcileMesh(obj)
}

func (r genericMeshReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(MeshDeletionReconciler); ok {
		deletionReconciler.ReconcileMeshDeletion(request)
	}
}

// genericMeshFinalizer implements a generic reconcile.FinalizingReconciler
type genericMeshFinalizer struct {
	genericMeshReconciler
	finalizingReconciler MeshFinalizer
}

func (r genericMeshFinalizer) FinalizerName() string {
	return r.finalizingReconciler.MeshFinalizerName()
}

func (r genericMeshFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*discovery_zephyr_solo_io_v1alpha1.Mesh)
	if !ok {
		return errors.Errorf("internal error: Mesh handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeMesh(obj)
}
