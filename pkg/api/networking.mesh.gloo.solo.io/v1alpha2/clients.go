// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha2

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the networking.mesh.gloo.solo.io/v1alpha2 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the networking.mesh.gloo.solo.io/v1alpha2 APIs
type Clientset interface {
	// clienset for the networking.mesh.gloo.solo.io/v1alpha2/v1alpha2 APIs
	TrafficPolicies() TrafficPolicyClient
	// clienset for the networking.mesh.gloo.solo.io/v1alpha2/v1alpha2 APIs
	AccessPolicies() AccessPolicyClient
	// clienset for the networking.mesh.gloo.solo.io/v1alpha2/v1alpha2 APIs
	VirtualMeshes() VirtualMeshClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the networking.mesh.gloo.solo.io/v1alpha2/v1alpha2 APIs
func (c *clientSet) TrafficPolicies() TrafficPolicyClient {
	return NewTrafficPolicyClient(c.client)
}

// clienset for the networking.mesh.gloo.solo.io/v1alpha2/v1alpha2 APIs
func (c *clientSet) AccessPolicies() AccessPolicyClient {
	return NewAccessPolicyClient(c.client)
}

// clienset for the networking.mesh.gloo.solo.io/v1alpha2/v1alpha2 APIs
func (c *clientSet) VirtualMeshes() VirtualMeshClient {
	return NewVirtualMeshClient(c.client)
}

// Reader knows how to read and list TrafficPolicys.
type TrafficPolicyReader interface {
	// Get retrieves a TrafficPolicy for the given object key
	GetTrafficPolicy(ctx context.Context, key client.ObjectKey) (*TrafficPolicy, error)

	// List retrieves list of TrafficPolicys for a given namespace and list options.
	ListTrafficPolicy(ctx context.Context, opts ...client.ListOption) (*TrafficPolicyList, error)
}

// TrafficPolicyTransitionFunction instructs the TrafficPolicyWriter how to transition between an existing
// TrafficPolicy object and a desired on an Upsert
type TrafficPolicyTransitionFunction func(existing, desired *TrafficPolicy) error

// Writer knows how to create, delete, and update TrafficPolicys.
type TrafficPolicyWriter interface {
	// Create saves the TrafficPolicy object.
	CreateTrafficPolicy(ctx context.Context, obj *TrafficPolicy, opts ...client.CreateOption) error

	// Delete deletes the TrafficPolicy object.
	DeleteTrafficPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given TrafficPolicy object.
	UpdateTrafficPolicy(ctx context.Context, obj *TrafficPolicy, opts ...client.UpdateOption) error

	// Patch patches the given TrafficPolicy object.
	PatchTrafficPolicy(ctx context.Context, obj *TrafficPolicy, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all TrafficPolicy objects matching the given options.
	DeleteAllOfTrafficPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the TrafficPolicy object.
	UpsertTrafficPolicy(ctx context.Context, obj *TrafficPolicy, transitionFuncs ...TrafficPolicyTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a TrafficPolicy object.
type TrafficPolicyStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given TrafficPolicy object.
	UpdateTrafficPolicyStatus(ctx context.Context, obj *TrafficPolicy, opts ...client.UpdateOption) error

	// Patch patches the given TrafficPolicy object's subresource.
	PatchTrafficPolicyStatus(ctx context.Context, obj *TrafficPolicy, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on TrafficPolicys.
type TrafficPolicyClient interface {
	TrafficPolicyReader
	TrafficPolicyWriter
	TrafficPolicyStatusWriter
}

type trafficPolicyClient struct {
	client client.Client
}

func NewTrafficPolicyClient(client client.Client) *trafficPolicyClient {
	return &trafficPolicyClient{client: client}
}

func (c *trafficPolicyClient) GetTrafficPolicy(ctx context.Context, key client.ObjectKey) (*TrafficPolicy, error) {
	obj := &TrafficPolicy{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *trafficPolicyClient) ListTrafficPolicy(ctx context.Context, opts ...client.ListOption) (*TrafficPolicyList, error) {
	list := &TrafficPolicyList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *trafficPolicyClient) CreateTrafficPolicy(ctx context.Context, obj *TrafficPolicy, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *trafficPolicyClient) DeleteTrafficPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &TrafficPolicy{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *trafficPolicyClient) UpdateTrafficPolicy(ctx context.Context, obj *TrafficPolicy, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *trafficPolicyClient) PatchTrafficPolicy(ctx context.Context, obj *TrafficPolicy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *trafficPolicyClient) DeleteAllOfTrafficPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &TrafficPolicy{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *trafficPolicyClient) UpsertTrafficPolicy(ctx context.Context, obj *TrafficPolicy, transitionFuncs ...TrafficPolicyTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*TrafficPolicy), desired.(*TrafficPolicy)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *trafficPolicyClient) UpdateTrafficPolicyStatus(ctx context.Context, obj *TrafficPolicy, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *trafficPolicyClient) PatchTrafficPolicyStatus(ctx context.Context, obj *TrafficPolicy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides TrafficPolicyClients for multiple clusters.
type MulticlusterTrafficPolicyClient interface {
	// Cluster returns a TrafficPolicyClient for the given cluster
	Cluster(cluster string) (TrafficPolicyClient, error)
}

type multiclusterTrafficPolicyClient struct {
	client multicluster.Client
}

func NewMulticlusterTrafficPolicyClient(client multicluster.Client) MulticlusterTrafficPolicyClient {
	return &multiclusterTrafficPolicyClient{client: client}
}

func (m *multiclusterTrafficPolicyClient) Cluster(cluster string) (TrafficPolicyClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewTrafficPolicyClient(client), nil
}

// Reader knows how to read and list AccessPolicys.
type AccessPolicyReader interface {
	// Get retrieves a AccessPolicy for the given object key
	GetAccessPolicy(ctx context.Context, key client.ObjectKey) (*AccessPolicy, error)

	// List retrieves list of AccessPolicys for a given namespace and list options.
	ListAccessPolicy(ctx context.Context, opts ...client.ListOption) (*AccessPolicyList, error)
}

// AccessPolicyTransitionFunction instructs the AccessPolicyWriter how to transition between an existing
// AccessPolicy object and a desired on an Upsert
type AccessPolicyTransitionFunction func(existing, desired *AccessPolicy) error

// Writer knows how to create, delete, and update AccessPolicys.
type AccessPolicyWriter interface {
	// Create saves the AccessPolicy object.
	CreateAccessPolicy(ctx context.Context, obj *AccessPolicy, opts ...client.CreateOption) error

	// Delete deletes the AccessPolicy object.
	DeleteAccessPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given AccessPolicy object.
	UpdateAccessPolicy(ctx context.Context, obj *AccessPolicy, opts ...client.UpdateOption) error

	// Patch patches the given AccessPolicy object.
	PatchAccessPolicy(ctx context.Context, obj *AccessPolicy, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all AccessPolicy objects matching the given options.
	DeleteAllOfAccessPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the AccessPolicy object.
	UpsertAccessPolicy(ctx context.Context, obj *AccessPolicy, transitionFuncs ...AccessPolicyTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a AccessPolicy object.
type AccessPolicyStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given AccessPolicy object.
	UpdateAccessPolicyStatus(ctx context.Context, obj *AccessPolicy, opts ...client.UpdateOption) error

	// Patch patches the given AccessPolicy object's subresource.
	PatchAccessPolicyStatus(ctx context.Context, obj *AccessPolicy, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on AccessPolicys.
type AccessPolicyClient interface {
	AccessPolicyReader
	AccessPolicyWriter
	AccessPolicyStatusWriter
}

type accessPolicyClient struct {
	client client.Client
}

func NewAccessPolicyClient(client client.Client) *accessPolicyClient {
	return &accessPolicyClient{client: client}
}

func (c *accessPolicyClient) GetAccessPolicy(ctx context.Context, key client.ObjectKey) (*AccessPolicy, error) {
	obj := &AccessPolicy{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *accessPolicyClient) ListAccessPolicy(ctx context.Context, opts ...client.ListOption) (*AccessPolicyList, error) {
	list := &AccessPolicyList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *accessPolicyClient) CreateAccessPolicy(ctx context.Context, obj *AccessPolicy, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *accessPolicyClient) DeleteAccessPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &AccessPolicy{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *accessPolicyClient) UpdateAccessPolicy(ctx context.Context, obj *AccessPolicy, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *accessPolicyClient) PatchAccessPolicy(ctx context.Context, obj *AccessPolicy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *accessPolicyClient) DeleteAllOfAccessPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &AccessPolicy{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *accessPolicyClient) UpsertAccessPolicy(ctx context.Context, obj *AccessPolicy, transitionFuncs ...AccessPolicyTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*AccessPolicy), desired.(*AccessPolicy)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *accessPolicyClient) UpdateAccessPolicyStatus(ctx context.Context, obj *AccessPolicy, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *accessPolicyClient) PatchAccessPolicyStatus(ctx context.Context, obj *AccessPolicy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides AccessPolicyClients for multiple clusters.
type MulticlusterAccessPolicyClient interface {
	// Cluster returns a AccessPolicyClient for the given cluster
	Cluster(cluster string) (AccessPolicyClient, error)
}

type multiclusterAccessPolicyClient struct {
	client multicluster.Client
}

func NewMulticlusterAccessPolicyClient(client multicluster.Client) MulticlusterAccessPolicyClient {
	return &multiclusterAccessPolicyClient{client: client}
}

func (m *multiclusterAccessPolicyClient) Cluster(cluster string) (AccessPolicyClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewAccessPolicyClient(client), nil
}

// Reader knows how to read and list VirtualMeshs.
type VirtualMeshReader interface {
	// Get retrieves a VirtualMesh for the given object key
	GetVirtualMesh(ctx context.Context, key client.ObjectKey) (*VirtualMesh, error)

	// List retrieves list of VirtualMeshs for a given namespace and list options.
	ListVirtualMesh(ctx context.Context, opts ...client.ListOption) (*VirtualMeshList, error)
}

// VirtualMeshTransitionFunction instructs the VirtualMeshWriter how to transition between an existing
// VirtualMesh object and a desired on an Upsert
type VirtualMeshTransitionFunction func(existing, desired *VirtualMesh) error

// Writer knows how to create, delete, and update VirtualMeshs.
type VirtualMeshWriter interface {
	// Create saves the VirtualMesh object.
	CreateVirtualMesh(ctx context.Context, obj *VirtualMesh, opts ...client.CreateOption) error

	// Delete deletes the VirtualMesh object.
	DeleteVirtualMesh(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given VirtualMesh object.
	UpdateVirtualMesh(ctx context.Context, obj *VirtualMesh, opts ...client.UpdateOption) error

	// Patch patches the given VirtualMesh object.
	PatchVirtualMesh(ctx context.Context, obj *VirtualMesh, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all VirtualMesh objects matching the given options.
	DeleteAllOfVirtualMesh(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the VirtualMesh object.
	UpsertVirtualMesh(ctx context.Context, obj *VirtualMesh, transitionFuncs ...VirtualMeshTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a VirtualMesh object.
type VirtualMeshStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given VirtualMesh object.
	UpdateVirtualMeshStatus(ctx context.Context, obj *VirtualMesh, opts ...client.UpdateOption) error

	// Patch patches the given VirtualMesh object's subresource.
	PatchVirtualMeshStatus(ctx context.Context, obj *VirtualMesh, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on VirtualMeshs.
type VirtualMeshClient interface {
	VirtualMeshReader
	VirtualMeshWriter
	VirtualMeshStatusWriter
}

type virtualMeshClient struct {
	client client.Client
}

func NewVirtualMeshClient(client client.Client) *virtualMeshClient {
	return &virtualMeshClient{client: client}
}

func (c *virtualMeshClient) GetVirtualMesh(ctx context.Context, key client.ObjectKey) (*VirtualMesh, error) {
	obj := &VirtualMesh{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *virtualMeshClient) ListVirtualMesh(ctx context.Context, opts ...client.ListOption) (*VirtualMeshList, error) {
	list := &VirtualMeshList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *virtualMeshClient) CreateVirtualMesh(ctx context.Context, obj *VirtualMesh, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *virtualMeshClient) DeleteVirtualMesh(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &VirtualMesh{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *virtualMeshClient) UpdateVirtualMesh(ctx context.Context, obj *VirtualMesh, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *virtualMeshClient) PatchVirtualMesh(ctx context.Context, obj *VirtualMesh, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *virtualMeshClient) DeleteAllOfVirtualMesh(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &VirtualMesh{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *virtualMeshClient) UpsertVirtualMesh(ctx context.Context, obj *VirtualMesh, transitionFuncs ...VirtualMeshTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*VirtualMesh), desired.(*VirtualMesh)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *virtualMeshClient) UpdateVirtualMeshStatus(ctx context.Context, obj *VirtualMesh, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *virtualMeshClient) PatchVirtualMeshStatus(ctx context.Context, obj *VirtualMesh, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides VirtualMeshClients for multiple clusters.
type MulticlusterVirtualMeshClient interface {
	// Cluster returns a VirtualMeshClient for the given cluster
	Cluster(cluster string) (VirtualMeshClient, error)
}

type multiclusterVirtualMeshClient struct {
	client multicluster.Client
}

func NewMulticlusterVirtualMeshClient(client multicluster.Client) MulticlusterVirtualMeshClient {
	return &multiclusterVirtualMeshClient{client: client}
}

func (m *multiclusterVirtualMeshClient) Cluster(cluster string) (VirtualMeshClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewVirtualMeshClient(client), nil
}
