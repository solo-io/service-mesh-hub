// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./snapshot.go -destination mocks/snapshot.go

// The Input Snapshot contains the set of all:
// * IssuedCertificates
// * CertificateRequests
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

	"github.com/solo-io/skv2/pkg/verifier"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/hashicorp/go-multierror"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"

	certificates_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2"
	certificates_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2/sets"
)

// the snapshot of input resources consumed by translation
type Snapshot interface {

	// return the set of input IssuedCertificates
	IssuedCertificates() certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet
	// return the set of input CertificateRequests
	CertificateRequests() certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet
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

	issuedCertificates  certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet
	certificateRequests certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet
}

func NewSnapshot(
	name string,

	issuedCertificates certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet,
	certificateRequests certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet,

) Snapshot {
	return &snapshot{
		name: name,

		issuedCertificates:  issuedCertificates,
		certificateRequests: certificateRequests,
	}
}

func (s snapshot) IssuedCertificates() certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet {
	return s.issuedCertificates
}

func (s snapshot) CertificateRequests() certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet {
	return s.certificateRequests
}
func (s snapshot) SyncStatuses(ctx context.Context, c client.Client) error {

	for _, obj := range s.IssuedCertificates().List() {
		if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
			return err
		}
	}
	for _, obj := range s.CertificateRequests().List() {
		if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
			return err
		}
	}
	return nil
}

func (s snapshot) SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client) error {

	for _, obj := range s.IssuedCertificates().List() {
		clusterClient, err := mcClient.Cluster(obj.ClusterName)
		if err != nil {
			return err
		}
		if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
			return err
		}
	}
	for _, obj := range s.CertificateRequests().List() {
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

	snapshotMap["issuedCertificates"] = s.issuedCertificates.List()
	snapshotMap["certificateRequests"] = s.certificateRequests.List()
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

	// List options for composing a snapshot from IssuedCertificates
	IssuedCertificates ResourceBuildOptions
	// List options for composing a snapshot from CertificateRequests
	CertificateRequests ResourceBuildOptions
}

// Options for reading resources of a given type
type ResourceBuildOptions struct {

	// List options for composing a snapshot from a resource type
	ListOptions []client.ListOption

	// If provided, ensure the resource has been verified before adding it to snapshots
	Verifier verifier.ServerResourceVerifier
}

// build a snapshot from resources across multiple clusters
type multiClusterBuilder struct {
	clusters multicluster.Interface
	client   multicluster.Client
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewMultiClusterBuilder(
	clusters multicluster.Interface,
	client multicluster.Client,
) Builder {
	return &multiClusterBuilder{
		clusters: clusters,
		client:   client,
	}
}

func (b *multiClusterBuilder) BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error) {

	issuedCertificates := certificates_mesh_gloo_solo_io_v1alpha2_sets.NewIssuedCertificateSet()
	certificateRequests := certificates_mesh_gloo_solo_io_v1alpha2_sets.NewCertificateRequestSet()

	var errs error

	for _, cluster := range b.clusters.ListClusters() {

		if err := b.insertIssuedCertificatesFromCluster(ctx, cluster, issuedCertificates, opts.IssuedCertificates); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertCertificateRequestsFromCluster(ctx, cluster, certificateRequests, opts.CertificateRequests); err != nil {
			errs = multierror.Append(errs, err)
		}

	}

	outputSnap := NewSnapshot(
		name,

		issuedCertificates,
		certificateRequests,
	)

	return outputSnap, errs
}

func (b *multiClusterBuilder) insertIssuedCertificatesFromCluster(ctx context.Context, cluster string, issuedCertificates certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet, opts ResourceBuildOptions) error {
	issuedCertificateClient, err := certificates_mesh_gloo_solo_io_v1alpha2.NewMulticlusterIssuedCertificateClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "certificates.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "IssuedCertificate",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	issuedCertificateList, err := issuedCertificateClient.ListIssuedCertificate(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range issuedCertificateList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		issuedCertificates.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertCertificateRequestsFromCluster(ctx context.Context, cluster string, certificateRequests certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet, opts ResourceBuildOptions) error {
	certificateRequestClient, err := certificates_mesh_gloo_solo_io_v1alpha2.NewMulticlusterCertificateRequestClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "certificates.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "CertificateRequest",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	certificateRequestList, err := certificateRequestClient.ListCertificateRequest(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range certificateRequestList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		certificateRequests.Insert(&item)
	}

	return nil
}

// build a snapshot from resources in a single cluster
type singleClusterBuilder struct {
	mgr manager.Manager
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewSingleClusterBuilder(
	mgr manager.Manager,
) Builder {
	return &singleClusterBuilder{
		mgr: mgr,
	}
}

func (b *singleClusterBuilder) BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error) {

	issuedCertificates := certificates_mesh_gloo_solo_io_v1alpha2_sets.NewIssuedCertificateSet()
	certificateRequests := certificates_mesh_gloo_solo_io_v1alpha2_sets.NewCertificateRequestSet()

	var errs error

	if err := b.insertIssuedCertificates(ctx, issuedCertificates, opts.IssuedCertificates); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertCertificateRequests(ctx, certificateRequests, opts.CertificateRequests); err != nil {
		errs = multierror.Append(errs, err)
	}

	outputSnap := NewSnapshot(
		name,

		issuedCertificates,
		certificateRequests,
	)

	return outputSnap, errs
}

func (b *singleClusterBuilder) insertIssuedCertificates(ctx context.Context, issuedCertificates certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "certificates.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "IssuedCertificate",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	issuedCertificateList, err := certificates_mesh_gloo_solo_io_v1alpha2.NewIssuedCertificateClient(b.mgr.GetClient()).ListIssuedCertificate(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range issuedCertificateList.Items {
		item := item // pike
		issuedCertificates.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertCertificateRequests(ctx context.Context, certificateRequests certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "certificates.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "CertificateRequest",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	certificateRequestList, err := certificates_mesh_gloo_solo_io_v1alpha2.NewCertificateRequestClient(b.mgr.GetClient()).ListCertificateRequest(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range certificateRequestList.Items {
		item := item // pike
		certificateRequests.Insert(&item)
	}

	return nil
}