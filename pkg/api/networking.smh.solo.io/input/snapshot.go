// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./snapshot.go -destination mocks/snapshot.go

// The Input Snapshot contains the set of all:
// * MeshServices
// * MeshWorkloads
// * Meshes
// * TrafficPolicies
// * AccessPolicies
// * VirtualMeshes
// * FailoverServices
// * Secrets
// * KubernetesClusters
// read from a given cluster or set of clusters, across all namespaces.
//
// A snapshot can be constructed from either a single Manager (for a single cluster)
// or a ClusterWatcher (for multiple clusters) using the SnapshotBuilder.
//
// Resources in a MultiCluster snapshot will have their ClusterName set to the
// name of the cluster from which the resource was read.

package input

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/go-multierror"
	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"

	discovery_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	discovery_smh_solo_io_v1alpha2_sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"

	networking_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"
	networking_smh_solo_io_v1alpha2_sets "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2/sets"

	v1 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"

	multicluster_solo_io_v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	multicluster_solo_io_v1alpha1_sets "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/sets"
)

// the snapshot of input resources consumed by translation
type Snapshot interface {

	// return the set of input MeshServices
	MeshServices() discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet
	// return the set of input MeshWorkloads
	MeshWorkloads() discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet
	// return the set of input Meshes
	Meshes() discovery_smh_solo_io_v1alpha2_sets.MeshSet

	// return the set of input TrafficPolicies
	TrafficPolicies() networking_smh_solo_io_v1alpha2_sets.TrafficPolicySet
	// return the set of input AccessPolicies
	AccessPolicies() networking_smh_solo_io_v1alpha2_sets.AccessPolicySet
	// return the set of input VirtualMeshes
	VirtualMeshes() networking_smh_solo_io_v1alpha2_sets.VirtualMeshSet
	// return the set of input FailoverServices
	FailoverServices() networking_smh_solo_io_v1alpha2_sets.FailoverServiceSet

	// return the set of input Secrets
	Secrets() v1_sets.SecretSet

	// return the set of input KubernetesClusters
	KubernetesClusters() multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet
	// update the status of all input objects which support
	// the Status subresource (in the local cluster)
	SyncStatuses(ctx context.Context, c client.Client) error

	// update the status of all input objects which support
	// the Status subresource (across multiple clusters)
	SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client) error
	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

type snapshot struct {
	name string

	meshServices  discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet
	meshWorkloads discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet
	meshes        discovery_smh_solo_io_v1alpha2_sets.MeshSet

	trafficPolicies  networking_smh_solo_io_v1alpha2_sets.TrafficPolicySet
	accessPolicies   networking_smh_solo_io_v1alpha2_sets.AccessPolicySet
	virtualMeshes    networking_smh_solo_io_v1alpha2_sets.VirtualMeshSet
	failoverServices networking_smh_solo_io_v1alpha2_sets.FailoverServiceSet

	secrets v1_sets.SecretSet

	kubernetesClusters multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet
}

func NewSnapshot(
	name string,

	meshServices discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet,
	meshWorkloads discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet,
	meshes discovery_smh_solo_io_v1alpha2_sets.MeshSet,

	trafficPolicies networking_smh_solo_io_v1alpha2_sets.TrafficPolicySet,
	accessPolicies networking_smh_solo_io_v1alpha2_sets.AccessPolicySet,
	virtualMeshes networking_smh_solo_io_v1alpha2_sets.VirtualMeshSet,
	failoverServices networking_smh_solo_io_v1alpha2_sets.FailoverServiceSet,

	secrets v1_sets.SecretSet,

	kubernetesClusters multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet,

) Snapshot {
	return &snapshot{
		name: name,

		meshServices:       meshServices,
		meshWorkloads:      meshWorkloads,
		meshes:             meshes,
		trafficPolicies:    trafficPolicies,
		accessPolicies:     accessPolicies,
		virtualMeshes:      virtualMeshes,
		failoverServices:   failoverServices,
		secrets:            secrets,
		kubernetesClusters: kubernetesClusters,
	}
}

func (s snapshot) MeshServices() discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet {
	return s.meshServices
}

func (s snapshot) MeshWorkloads() discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet {
	return s.meshWorkloads
}

func (s snapshot) Meshes() discovery_smh_solo_io_v1alpha2_sets.MeshSet {
	return s.meshes
}

func (s snapshot) TrafficPolicies() networking_smh_solo_io_v1alpha2_sets.TrafficPolicySet {
	return s.trafficPolicies
}

func (s snapshot) AccessPolicies() networking_smh_solo_io_v1alpha2_sets.AccessPolicySet {
	return s.accessPolicies
}

func (s snapshot) VirtualMeshes() networking_smh_solo_io_v1alpha2_sets.VirtualMeshSet {
	return s.virtualMeshes
}

func (s snapshot) FailoverServices() networking_smh_solo_io_v1alpha2_sets.FailoverServiceSet {
	return s.failoverServices
}

func (s snapshot) Secrets() v1_sets.SecretSet {
	return s.secrets
}

func (s snapshot) KubernetesClusters() multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet {
	return s.kubernetesClusters
}
func (s snapshot) SyncStatuses(ctx context.Context, c client.Client) error {

	for _, obj := range s.MeshServices().List() {
		if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
			return err
		}
	}
	for _, obj := range s.MeshWorkloads().List() {
		if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
			return err
		}
	}
	for _, obj := range s.Meshes().List() {
		if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
			return err
		}
	}

	for _, obj := range s.TrafficPolicies().List() {
		if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
			return err
		}
	}
	for _, obj := range s.AccessPolicies().List() {
		if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
			return err
		}
	}
	for _, obj := range s.VirtualMeshes().List() {
		if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
			return err
		}
	}
	for _, obj := range s.FailoverServices().List() {
		if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
			return err
		}
	}

	for _, obj := range s.KubernetesClusters().List() {
		if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
			return err
		}
	}
	return nil
}

func (s snapshot) SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client) error {

	for _, obj := range s.MeshServices().List() {
		clusterClient, err := mcClient.Cluster(obj.ClusterName)
		if err != nil {
			return err
		}
		if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
			return err
		}
	}
	for _, obj := range s.MeshWorkloads().List() {
		clusterClient, err := mcClient.Cluster(obj.ClusterName)
		if err != nil {
			return err
		}
		if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
			return err
		}
	}
	for _, obj := range s.Meshes().List() {
		clusterClient, err := mcClient.Cluster(obj.ClusterName)
		if err != nil {
			return err
		}
		if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
			return err
		}
	}

	for _, obj := range s.TrafficPolicies().List() {
		clusterClient, err := mcClient.Cluster(obj.ClusterName)
		if err != nil {
			return err
		}
		if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
			return err
		}
	}
	for _, obj := range s.AccessPolicies().List() {
		clusterClient, err := mcClient.Cluster(obj.ClusterName)
		if err != nil {
			return err
		}
		if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
			return err
		}
	}
	for _, obj := range s.VirtualMeshes().List() {
		clusterClient, err := mcClient.Cluster(obj.ClusterName)
		if err != nil {
			return err
		}
		if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
			return err
		}
	}
	for _, obj := range s.FailoverServices().List() {
		clusterClient, err := mcClient.Cluster(obj.ClusterName)
		if err != nil {
			return err
		}
		if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
			return err
		}
	}

	for _, obj := range s.KubernetesClusters().List() {
		clusterClient, err := mcClient.Cluster(obj.ClusterName)
		if err != nil {
			return err
		}
		if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
			return err
		}
	}
	return nil
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	snapshotMap["meshServices"] = s.meshServices.List()
	snapshotMap["meshWorkloads"] = s.meshWorkloads.List()
	snapshotMap["meshes"] = s.meshes.List()
	snapshotMap["trafficPolicies"] = s.trafficPolicies.List()
	snapshotMap["accessPolicies"] = s.accessPolicies.List()
	snapshotMap["virtualMeshes"] = s.virtualMeshes.List()
	snapshotMap["failoverServices"] = s.failoverServices.List()
	snapshotMap["secrets"] = s.secrets.List()
	snapshotMap["kubernetesClusters"] = s.kubernetesClusters.List()
	return json.Marshal(snapshotMap)
}

// builds the input snapshot from API Clients.
// Two types of builders are available:
// a builder for snapshots of resources across multiple clusters
// a builder for snapshots of resources within a single cluster
type Builder interface {
	BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error)
}

// Options for building a snapshot
type BuildOptions struct {

	// List options for composing a snapshot from MeshServices
	MeshServices []client.ListOption
	// List options for composing a snapshot from MeshWorkloads
	MeshWorkloads []client.ListOption
	// List options for composing a snapshot from Meshes
	Meshes []client.ListOption

	// List options for composing a snapshot from TrafficPolicies
	TrafficPolicies []client.ListOption
	// List options for composing a snapshot from AccessPolicies
	AccessPolicies []client.ListOption
	// List options for composing a snapshot from VirtualMeshes
	VirtualMeshes []client.ListOption
	// List options for composing a snapshot from FailoverServices
	FailoverServices []client.ListOption

	// List options for composing a snapshot from Secrets
	Secrets []client.ListOption

	// List options for composing a snapshot from KubernetesClusters
	KubernetesClusters []client.ListOption
}

// build a snapshot from resources across multiple clusters
type multiClusterBuilder struct {
	clusters multicluster.ClusterSet

	meshServices  discovery_smh_solo_io_v1alpha2.MulticlusterMeshServiceClient
	meshWorkloads discovery_smh_solo_io_v1alpha2.MulticlusterMeshWorkloadClient
	meshes        discovery_smh_solo_io_v1alpha2.MulticlusterMeshClient

	trafficPolicies  networking_smh_solo_io_v1alpha2.MulticlusterTrafficPolicyClient
	accessPolicies   networking_smh_solo_io_v1alpha2.MulticlusterAccessPolicyClient
	virtualMeshes    networking_smh_solo_io_v1alpha2.MulticlusterVirtualMeshClient
	failoverServices networking_smh_solo_io_v1alpha2.MulticlusterFailoverServiceClient

	secrets v1.MulticlusterSecretClient

	kubernetesClusters multicluster_solo_io_v1alpha1.MulticlusterKubernetesClusterClient
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewMultiClusterBuilder(
	clusters multicluster.ClusterSet,
	client multicluster.Client,
) Builder {
	return &multiClusterBuilder{
		clusters: clusters,

		meshServices:  discovery_smh_solo_io_v1alpha2.NewMulticlusterMeshServiceClient(client),
		meshWorkloads: discovery_smh_solo_io_v1alpha2.NewMulticlusterMeshWorkloadClient(client),
		meshes:        discovery_smh_solo_io_v1alpha2.NewMulticlusterMeshClient(client),

		trafficPolicies:  networking_smh_solo_io_v1alpha2.NewMulticlusterTrafficPolicyClient(client),
		accessPolicies:   networking_smh_solo_io_v1alpha2.NewMulticlusterAccessPolicyClient(client),
		virtualMeshes:    networking_smh_solo_io_v1alpha2.NewMulticlusterVirtualMeshClient(client),
		failoverServices: networking_smh_solo_io_v1alpha2.NewMulticlusterFailoverServiceClient(client),

		secrets: v1.NewMulticlusterSecretClient(client),

		kubernetesClusters: multicluster_solo_io_v1alpha1.NewMulticlusterKubernetesClusterClient(client),
	}
}

func (b *multiClusterBuilder) BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error) {

	meshServices := discovery_smh_solo_io_v1alpha2_sets.NewMeshServiceSet()
	meshWorkloads := discovery_smh_solo_io_v1alpha2_sets.NewMeshWorkloadSet()
	meshes := discovery_smh_solo_io_v1alpha2_sets.NewMeshSet()

	trafficPolicies := networking_smh_solo_io_v1alpha2_sets.NewTrafficPolicySet()
	accessPolicies := networking_smh_solo_io_v1alpha2_sets.NewAccessPolicySet()
	virtualMeshes := networking_smh_solo_io_v1alpha2_sets.NewVirtualMeshSet()
	failoverServices := networking_smh_solo_io_v1alpha2_sets.NewFailoverServiceSet()

	secrets := v1_sets.NewSecretSet()

	kubernetesClusters := multicluster_solo_io_v1alpha1_sets.NewKubernetesClusterSet()

	var errs error

	for _, cluster := range b.clusters.ListClusters() {

		if err := b.insertMeshServicesFromCluster(ctx, cluster, meshServices, opts.MeshServices...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertMeshWorkloadsFromCluster(ctx, cluster, meshWorkloads, opts.MeshWorkloads...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertMeshesFromCluster(ctx, cluster, meshes, opts.Meshes...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertTrafficPoliciesFromCluster(ctx, cluster, trafficPolicies, opts.TrafficPolicies...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertAccessPoliciesFromCluster(ctx, cluster, accessPolicies, opts.AccessPolicies...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertVirtualMeshesFromCluster(ctx, cluster, virtualMeshes, opts.VirtualMeshes...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertFailoverServicesFromCluster(ctx, cluster, failoverServices, opts.FailoverServices...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertSecretsFromCluster(ctx, cluster, secrets, opts.Secrets...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertKubernetesClustersFromCluster(ctx, cluster, kubernetesClusters, opts.KubernetesClusters...); err != nil {
			errs = multierror.Append(errs, err)
		}

	}

	outputSnap := NewSnapshot(
		name,

		meshServices,
		meshWorkloads,
		meshes,
		trafficPolicies,
		accessPolicies,
		virtualMeshes,
		failoverServices,
		secrets,
		kubernetesClusters,
	)

	return outputSnap, errs
}

func (b *multiClusterBuilder) insertMeshServicesFromCluster(ctx context.Context, cluster string, meshServices discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet, opts ...client.ListOption) error {
	meshServiceClient, err := b.meshServices.Cluster(cluster)
	if err != nil {
		return err
	}

	meshServiceList, err := meshServiceClient.ListMeshService(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range meshServiceList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		meshServices.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertMeshWorkloadsFromCluster(ctx context.Context, cluster string, meshWorkloads discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet, opts ...client.ListOption) error {
	meshWorkloadClient, err := b.meshWorkloads.Cluster(cluster)
	if err != nil {
		return err
	}

	meshWorkloadList, err := meshWorkloadClient.ListMeshWorkload(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range meshWorkloadList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		meshWorkloads.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertMeshesFromCluster(ctx context.Context, cluster string, meshes discovery_smh_solo_io_v1alpha2_sets.MeshSet, opts ...client.ListOption) error {
	meshClient, err := b.meshes.Cluster(cluster)
	if err != nil {
		return err
	}

	meshList, err := meshClient.ListMesh(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range meshList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		meshes.Insert(&item)
	}

	return nil
}

func (b *multiClusterBuilder) insertTrafficPoliciesFromCluster(ctx context.Context, cluster string, trafficPolicies networking_smh_solo_io_v1alpha2_sets.TrafficPolicySet, opts ...client.ListOption) error {
	trafficPolicyClient, err := b.trafficPolicies.Cluster(cluster)
	if err != nil {
		return err
	}

	trafficPolicyList, err := trafficPolicyClient.ListTrafficPolicy(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range trafficPolicyList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		trafficPolicies.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertAccessPoliciesFromCluster(ctx context.Context, cluster string, accessPolicies networking_smh_solo_io_v1alpha2_sets.AccessPolicySet, opts ...client.ListOption) error {
	accessPolicyClient, err := b.accessPolicies.Cluster(cluster)
	if err != nil {
		return err
	}

	accessPolicyList, err := accessPolicyClient.ListAccessPolicy(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range accessPolicyList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		accessPolicies.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertVirtualMeshesFromCluster(ctx context.Context, cluster string, virtualMeshes networking_smh_solo_io_v1alpha2_sets.VirtualMeshSet, opts ...client.ListOption) error {
	virtualMeshClient, err := b.virtualMeshes.Cluster(cluster)
	if err != nil {
		return err
	}

	virtualMeshList, err := virtualMeshClient.ListVirtualMesh(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range virtualMeshList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		virtualMeshes.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertFailoverServicesFromCluster(ctx context.Context, cluster string, failoverServices networking_smh_solo_io_v1alpha2_sets.FailoverServiceSet, opts ...client.ListOption) error {
	failoverServiceClient, err := b.failoverServices.Cluster(cluster)
	if err != nil {
		return err
	}

	failoverServiceList, err := failoverServiceClient.ListFailoverService(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range failoverServiceList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		failoverServices.Insert(&item)
	}

	return nil
}

func (b *multiClusterBuilder) insertSecretsFromCluster(ctx context.Context, cluster string, secrets v1_sets.SecretSet, opts ...client.ListOption) error {
	secretClient, err := b.secrets.Cluster(cluster)
	if err != nil {
		return err
	}

	secretList, err := secretClient.ListSecret(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range secretList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		secrets.Insert(&item)
	}

	return nil
}

func (b *multiClusterBuilder) insertKubernetesClustersFromCluster(ctx context.Context, cluster string, kubernetesClusters multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet, opts ...client.ListOption) error {
	kubernetesClusterClient, err := b.kubernetesClusters.Cluster(cluster)
	if err != nil {
		return err
	}

	kubernetesClusterList, err := kubernetesClusterClient.ListKubernetesCluster(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range kubernetesClusterList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		kubernetesClusters.Insert(&item)
	}

	return nil
}

// build a snapshot from resources in a single cluster
type singleClusterBuilder struct {
	meshServices  discovery_smh_solo_io_v1alpha2.MeshServiceClient
	meshWorkloads discovery_smh_solo_io_v1alpha2.MeshWorkloadClient
	meshes        discovery_smh_solo_io_v1alpha2.MeshClient

	trafficPolicies  networking_smh_solo_io_v1alpha2.TrafficPolicyClient
	accessPolicies   networking_smh_solo_io_v1alpha2.AccessPolicyClient
	virtualMeshes    networking_smh_solo_io_v1alpha2.VirtualMeshClient
	failoverServices networking_smh_solo_io_v1alpha2.FailoverServiceClient

	secrets v1.SecretClient

	kubernetesClusters multicluster_solo_io_v1alpha1.KubernetesClusterClient
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewSingleClusterBuilder(
	client client.Client,
) Builder {
	return &singleClusterBuilder{

		meshServices:  discovery_smh_solo_io_v1alpha2.NewMeshServiceClient(client),
		meshWorkloads: discovery_smh_solo_io_v1alpha2.NewMeshWorkloadClient(client),
		meshes:        discovery_smh_solo_io_v1alpha2.NewMeshClient(client),

		trafficPolicies:  networking_smh_solo_io_v1alpha2.NewTrafficPolicyClient(client),
		accessPolicies:   networking_smh_solo_io_v1alpha2.NewAccessPolicyClient(client),
		virtualMeshes:    networking_smh_solo_io_v1alpha2.NewVirtualMeshClient(client),
		failoverServices: networking_smh_solo_io_v1alpha2.NewFailoverServiceClient(client),

		secrets: v1.NewSecretClient(client),

		kubernetesClusters: multicluster_solo_io_v1alpha1.NewKubernetesClusterClient(client),
	}
}

func (b *singleClusterBuilder) BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error) {

	meshServices := discovery_smh_solo_io_v1alpha2_sets.NewMeshServiceSet()
	meshWorkloads := discovery_smh_solo_io_v1alpha2_sets.NewMeshWorkloadSet()
	meshes := discovery_smh_solo_io_v1alpha2_sets.NewMeshSet()

	trafficPolicies := networking_smh_solo_io_v1alpha2_sets.NewTrafficPolicySet()
	accessPolicies := networking_smh_solo_io_v1alpha2_sets.NewAccessPolicySet()
	virtualMeshes := networking_smh_solo_io_v1alpha2_sets.NewVirtualMeshSet()
	failoverServices := networking_smh_solo_io_v1alpha2_sets.NewFailoverServiceSet()

	secrets := v1_sets.NewSecretSet()

	kubernetesClusters := multicluster_solo_io_v1alpha1_sets.NewKubernetesClusterSet()

	var errs error

	if err := b.insertMeshServices(ctx, meshServices, opts.MeshServices...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertMeshWorkloads(ctx, meshWorkloads, opts.MeshWorkloads...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertMeshes(ctx, meshes, opts.Meshes...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertTrafficPolicies(ctx, trafficPolicies, opts.TrafficPolicies...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertAccessPolicies(ctx, accessPolicies, opts.AccessPolicies...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertVirtualMeshes(ctx, virtualMeshes, opts.VirtualMeshes...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertFailoverServices(ctx, failoverServices, opts.FailoverServices...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertSecrets(ctx, secrets, opts.Secrets...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertKubernetesClusters(ctx, kubernetesClusters, opts.KubernetesClusters...); err != nil {
		errs = multierror.Append(errs, err)
	}

	outputSnap := NewSnapshot(
		name,

		meshServices,
		meshWorkloads,
		meshes,
		trafficPolicies,
		accessPolicies,
		virtualMeshes,
		failoverServices,
		secrets,
		kubernetesClusters,
	)

	return outputSnap, errs
}

func (b *singleClusterBuilder) insertMeshServices(ctx context.Context, meshServices discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet, opts ...client.ListOption) error {
	meshServiceList, err := b.meshServices.ListMeshService(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range meshServiceList.Items {
		item := item // pike
		meshServices.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertMeshWorkloads(ctx context.Context, meshWorkloads discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet, opts ...client.ListOption) error {
	meshWorkloadList, err := b.meshWorkloads.ListMeshWorkload(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range meshWorkloadList.Items {
		item := item // pike
		meshWorkloads.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertMeshes(ctx context.Context, meshes discovery_smh_solo_io_v1alpha2_sets.MeshSet, opts ...client.ListOption) error {
	meshList, err := b.meshes.ListMesh(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range meshList.Items {
		item := item // pike
		meshes.Insert(&item)
	}

	return nil
}

func (b *singleClusterBuilder) insertTrafficPolicies(ctx context.Context, trafficPolicies networking_smh_solo_io_v1alpha2_sets.TrafficPolicySet, opts ...client.ListOption) error {
	trafficPolicyList, err := b.trafficPolicies.ListTrafficPolicy(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range trafficPolicyList.Items {
		item := item // pike
		trafficPolicies.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertAccessPolicies(ctx context.Context, accessPolicies networking_smh_solo_io_v1alpha2_sets.AccessPolicySet, opts ...client.ListOption) error {
	accessPolicyList, err := b.accessPolicies.ListAccessPolicy(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range accessPolicyList.Items {
		item := item // pike
		accessPolicies.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertVirtualMeshes(ctx context.Context, virtualMeshes networking_smh_solo_io_v1alpha2_sets.VirtualMeshSet, opts ...client.ListOption) error {
	virtualMeshList, err := b.virtualMeshes.ListVirtualMesh(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range virtualMeshList.Items {
		item := item // pike
		virtualMeshes.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertFailoverServices(ctx context.Context, failoverServices networking_smh_solo_io_v1alpha2_sets.FailoverServiceSet, opts ...client.ListOption) error {
	failoverServiceList, err := b.failoverServices.ListFailoverService(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range failoverServiceList.Items {
		item := item // pike
		failoverServices.Insert(&item)
	}

	return nil
}

func (b *singleClusterBuilder) insertSecrets(ctx context.Context, secrets v1_sets.SecretSet, opts ...client.ListOption) error {
	secretList, err := b.secrets.ListSecret(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range secretList.Items {
		item := item // pike
		secrets.Insert(&item)
	}

	return nil
}

func (b *singleClusterBuilder) insertKubernetesClusters(ctx context.Context, kubernetesClusters multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet, opts ...client.ListOption) error {
	kubernetesClusterList, err := b.kubernetesClusters.ListKubernetesCluster(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range kubernetesClusterList.Items {
		item := item // pike
		kubernetesClusters.Insert(&item)
	}

	return nil
}
