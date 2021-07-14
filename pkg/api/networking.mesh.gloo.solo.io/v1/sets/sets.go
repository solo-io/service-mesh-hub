// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets

import (
	networking_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type TrafficPolicySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_mesh_gloo_solo_io_v1.TrafficPolicy) bool) []*networking_mesh_gloo_solo_io_v1.TrafficPolicy
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_mesh_gloo_solo_io_v1.TrafficPolicy) bool) []*networking_mesh_gloo_solo_io_v1.TrafficPolicy
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_mesh_gloo_solo_io_v1.TrafficPolicy
	// Insert a resource into the set.
	Insert(trafficPolicy ...*networking_mesh_gloo_solo_io_v1.TrafficPolicy)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(trafficPolicySet TrafficPolicySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(trafficPolicy ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(trafficPolicy ezkube.ResourceId)
	// Return the union with the provided set
	Union(set TrafficPolicySet) TrafficPolicySet
	// Return the difference with the provided set
	Difference(set TrafficPolicySet) TrafficPolicySet
	// Return the intersection with the provided set
	Intersection(set TrafficPolicySet) TrafficPolicySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_mesh_gloo_solo_io_v1.TrafficPolicy, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another TrafficPolicySet
	Delta(newSet TrafficPolicySet) sksets.ResourceDelta
	// Create a deep copy of the current TrafficPolicySet
	Clone() TrafficPolicySet
}

func makeGenericTrafficPolicySet(trafficPolicyList []*networking_mesh_gloo_solo_io_v1.TrafficPolicy) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range trafficPolicyList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type trafficPolicySet struct {
	set sksets.ResourceSet
}

func NewTrafficPolicySet(trafficPolicyList ...*networking_mesh_gloo_solo_io_v1.TrafficPolicy) TrafficPolicySet {
	return &trafficPolicySet{set: makeGenericTrafficPolicySet(trafficPolicyList)}
}

func NewTrafficPolicySetFromList(trafficPolicyList *networking_mesh_gloo_solo_io_v1.TrafficPolicyList) TrafficPolicySet {
	list := make([]*networking_mesh_gloo_solo_io_v1.TrafficPolicy, 0, len(trafficPolicyList.Items))
	for idx := range trafficPolicyList.Items {
		list = append(list, &trafficPolicyList.Items[idx])
	}
	return &trafficPolicySet{set: makeGenericTrafficPolicySet(list)}
}

func (s *trafficPolicySet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *trafficPolicySet) List(filterResource ...func(*networking_mesh_gloo_solo_io_v1.TrafficPolicy) bool) []*networking_mesh_gloo_solo_io_v1.TrafficPolicy {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_mesh_gloo_solo_io_v1.TrafficPolicy))
		})
	}

	objs := s.Generic().List(genericFilters...)
	trafficPolicyList := make([]*networking_mesh_gloo_solo_io_v1.TrafficPolicy, 0, len(objs))
	for _, obj := range objs {
		trafficPolicyList = append(trafficPolicyList, obj.(*networking_mesh_gloo_solo_io_v1.TrafficPolicy))
	}
	return trafficPolicyList
}

func (s *trafficPolicySet) UnsortedList(filterResource ...func(*networking_mesh_gloo_solo_io_v1.TrafficPolicy) bool) []*networking_mesh_gloo_solo_io_v1.TrafficPolicy {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_mesh_gloo_solo_io_v1.TrafficPolicy))
		})
	}

	var trafficPolicyList []*networking_mesh_gloo_solo_io_v1.TrafficPolicy
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		trafficPolicyList = append(trafficPolicyList, obj.(*networking_mesh_gloo_solo_io_v1.TrafficPolicy))
	}
	return trafficPolicyList
}

func (s *trafficPolicySet) Map() map[string]*networking_mesh_gloo_solo_io_v1.TrafficPolicy {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_mesh_gloo_solo_io_v1.TrafficPolicy{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_mesh_gloo_solo_io_v1.TrafficPolicy)
	}
	return newMap
}

func (s *trafficPolicySet) Insert(
	trafficPolicyList ...*networking_mesh_gloo_solo_io_v1.TrafficPolicy,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range trafficPolicyList {
		s.Generic().Insert(obj)
	}
}

func (s *trafficPolicySet) Has(trafficPolicy ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(trafficPolicy)
}

func (s *trafficPolicySet) Equal(
	trafficPolicySet TrafficPolicySet,
) bool {
	if s == nil {
		return trafficPolicySet == nil
	}
	return s.Generic().Equal(trafficPolicySet.Generic())
}

func (s *trafficPolicySet) Delete(TrafficPolicy ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(TrafficPolicy)
}

func (s *trafficPolicySet) Union(set TrafficPolicySet) TrafficPolicySet {
	if s == nil {
		return set
	}
	return NewTrafficPolicySet(append(s.List(), set.List()...)...)
}

func (s *trafficPolicySet) Difference(set TrafficPolicySet) TrafficPolicySet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &trafficPolicySet{set: newSet}
}

func (s *trafficPolicySet) Intersection(set TrafficPolicySet) TrafficPolicySet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var trafficPolicyList []*networking_mesh_gloo_solo_io_v1.TrafficPolicy
	for _, obj := range newSet.List() {
		trafficPolicyList = append(trafficPolicyList, obj.(*networking_mesh_gloo_solo_io_v1.TrafficPolicy))
	}
	return NewTrafficPolicySet(trafficPolicyList...)
}

func (s *trafficPolicySet) Find(id ezkube.ResourceId) (*networking_mesh_gloo_solo_io_v1.TrafficPolicy, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find TrafficPolicy %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_mesh_gloo_solo_io_v1.TrafficPolicy{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_mesh_gloo_solo_io_v1.TrafficPolicy), nil
}

func (s *trafficPolicySet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *trafficPolicySet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *trafficPolicySet) Delta(newSet TrafficPolicySet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *trafficPolicySet) Clone() TrafficPolicySet {
	if s == nil {
		return nil
	}
	return &trafficPolicySet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}

type AccessPolicySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_mesh_gloo_solo_io_v1.AccessPolicy) bool) []*networking_mesh_gloo_solo_io_v1.AccessPolicy
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_mesh_gloo_solo_io_v1.AccessPolicy) bool) []*networking_mesh_gloo_solo_io_v1.AccessPolicy
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_mesh_gloo_solo_io_v1.AccessPolicy
	// Insert a resource into the set.
	Insert(accessPolicy ...*networking_mesh_gloo_solo_io_v1.AccessPolicy)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(accessPolicySet AccessPolicySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(accessPolicy ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(accessPolicy ezkube.ResourceId)
	// Return the union with the provided set
	Union(set AccessPolicySet) AccessPolicySet
	// Return the difference with the provided set
	Difference(set AccessPolicySet) AccessPolicySet
	// Return the intersection with the provided set
	Intersection(set AccessPolicySet) AccessPolicySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_mesh_gloo_solo_io_v1.AccessPolicy, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another AccessPolicySet
	Delta(newSet AccessPolicySet) sksets.ResourceDelta
	// Create a deep copy of the current AccessPolicySet
	Clone() AccessPolicySet
}

func makeGenericAccessPolicySet(accessPolicyList []*networking_mesh_gloo_solo_io_v1.AccessPolicy) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range accessPolicyList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type accessPolicySet struct {
	set sksets.ResourceSet
}

func NewAccessPolicySet(accessPolicyList ...*networking_mesh_gloo_solo_io_v1.AccessPolicy) AccessPolicySet {
	return &accessPolicySet{set: makeGenericAccessPolicySet(accessPolicyList)}
}

func NewAccessPolicySetFromList(accessPolicyList *networking_mesh_gloo_solo_io_v1.AccessPolicyList) AccessPolicySet {
	list := make([]*networking_mesh_gloo_solo_io_v1.AccessPolicy, 0, len(accessPolicyList.Items))
	for idx := range accessPolicyList.Items {
		list = append(list, &accessPolicyList.Items[idx])
	}
	return &accessPolicySet{set: makeGenericAccessPolicySet(list)}
}

func (s *accessPolicySet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *accessPolicySet) List(filterResource ...func(*networking_mesh_gloo_solo_io_v1.AccessPolicy) bool) []*networking_mesh_gloo_solo_io_v1.AccessPolicy {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_mesh_gloo_solo_io_v1.AccessPolicy))
		})
	}

	objs := s.Generic().List(genericFilters...)
	accessPolicyList := make([]*networking_mesh_gloo_solo_io_v1.AccessPolicy, 0, len(objs))
	for _, obj := range objs {
		accessPolicyList = append(accessPolicyList, obj.(*networking_mesh_gloo_solo_io_v1.AccessPolicy))
	}
	return accessPolicyList
}

func (s *accessPolicySet) UnsortedList(filterResource ...func(*networking_mesh_gloo_solo_io_v1.AccessPolicy) bool) []*networking_mesh_gloo_solo_io_v1.AccessPolicy {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_mesh_gloo_solo_io_v1.AccessPolicy))
		})
	}

	var accessPolicyList []*networking_mesh_gloo_solo_io_v1.AccessPolicy
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		accessPolicyList = append(accessPolicyList, obj.(*networking_mesh_gloo_solo_io_v1.AccessPolicy))
	}
	return accessPolicyList
}

func (s *accessPolicySet) Map() map[string]*networking_mesh_gloo_solo_io_v1.AccessPolicy {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_mesh_gloo_solo_io_v1.AccessPolicy{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_mesh_gloo_solo_io_v1.AccessPolicy)
	}
	return newMap
}

func (s *accessPolicySet) Insert(
	accessPolicyList ...*networking_mesh_gloo_solo_io_v1.AccessPolicy,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range accessPolicyList {
		s.Generic().Insert(obj)
	}
}

func (s *accessPolicySet) Has(accessPolicy ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(accessPolicy)
}

func (s *accessPolicySet) Equal(
	accessPolicySet AccessPolicySet,
) bool {
	if s == nil {
		return accessPolicySet == nil
	}
	return s.Generic().Equal(accessPolicySet.Generic())
}

func (s *accessPolicySet) Delete(AccessPolicy ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(AccessPolicy)
}

func (s *accessPolicySet) Union(set AccessPolicySet) AccessPolicySet {
	if s == nil {
		return set
	}
	return NewAccessPolicySet(append(s.List(), set.List()...)...)
}

func (s *accessPolicySet) Difference(set AccessPolicySet) AccessPolicySet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &accessPolicySet{set: newSet}
}

func (s *accessPolicySet) Intersection(set AccessPolicySet) AccessPolicySet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var accessPolicyList []*networking_mesh_gloo_solo_io_v1.AccessPolicy
	for _, obj := range newSet.List() {
		accessPolicyList = append(accessPolicyList, obj.(*networking_mesh_gloo_solo_io_v1.AccessPolicy))
	}
	return NewAccessPolicySet(accessPolicyList...)
}

func (s *accessPolicySet) Find(id ezkube.ResourceId) (*networking_mesh_gloo_solo_io_v1.AccessPolicy, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find AccessPolicy %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_mesh_gloo_solo_io_v1.AccessPolicy{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_mesh_gloo_solo_io_v1.AccessPolicy), nil
}

func (s *accessPolicySet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *accessPolicySet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *accessPolicySet) Delta(newSet AccessPolicySet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *accessPolicySet) Clone() AccessPolicySet {
	if s == nil {
		return nil
	}
	return &accessPolicySet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}

type VirtualMeshSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_mesh_gloo_solo_io_v1.VirtualMesh) bool) []*networking_mesh_gloo_solo_io_v1.VirtualMesh
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_mesh_gloo_solo_io_v1.VirtualMesh) bool) []*networking_mesh_gloo_solo_io_v1.VirtualMesh
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_mesh_gloo_solo_io_v1.VirtualMesh
	// Insert a resource into the set.
	Insert(virtualMesh ...*networking_mesh_gloo_solo_io_v1.VirtualMesh)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualMeshSet VirtualMeshSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualMesh ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualMesh ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualMeshSet) VirtualMeshSet
	// Return the difference with the provided set
	Difference(set VirtualMeshSet) VirtualMeshSet
	// Return the intersection with the provided set
	Intersection(set VirtualMeshSet) VirtualMeshSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_mesh_gloo_solo_io_v1.VirtualMesh, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another VirtualMeshSet
	Delta(newSet VirtualMeshSet) sksets.ResourceDelta
	// Create a deep copy of the current VirtualMeshSet
	Clone() VirtualMeshSet
}

func makeGenericVirtualMeshSet(virtualMeshList []*networking_mesh_gloo_solo_io_v1.VirtualMesh) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualMeshList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualMeshSet struct {
	set sksets.ResourceSet
}

func NewVirtualMeshSet(virtualMeshList ...*networking_mesh_gloo_solo_io_v1.VirtualMesh) VirtualMeshSet {
	return &virtualMeshSet{set: makeGenericVirtualMeshSet(virtualMeshList)}
}

func NewVirtualMeshSetFromList(virtualMeshList *networking_mesh_gloo_solo_io_v1.VirtualMeshList) VirtualMeshSet {
	list := make([]*networking_mesh_gloo_solo_io_v1.VirtualMesh, 0, len(virtualMeshList.Items))
	for idx := range virtualMeshList.Items {
		list = append(list, &virtualMeshList.Items[idx])
	}
	return &virtualMeshSet{set: makeGenericVirtualMeshSet(list)}
}

func (s *virtualMeshSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *virtualMeshSet) List(filterResource ...func(*networking_mesh_gloo_solo_io_v1.VirtualMesh) bool) []*networking_mesh_gloo_solo_io_v1.VirtualMesh {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_mesh_gloo_solo_io_v1.VirtualMesh))
		})
	}

	objs := s.Generic().List(genericFilters...)
	virtualMeshList := make([]*networking_mesh_gloo_solo_io_v1.VirtualMesh, 0, len(objs))
	for _, obj := range objs {
		virtualMeshList = append(virtualMeshList, obj.(*networking_mesh_gloo_solo_io_v1.VirtualMesh))
	}
	return virtualMeshList
}

func (s *virtualMeshSet) UnsortedList(filterResource ...func(*networking_mesh_gloo_solo_io_v1.VirtualMesh) bool) []*networking_mesh_gloo_solo_io_v1.VirtualMesh {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_mesh_gloo_solo_io_v1.VirtualMesh))
		})
	}

	var virtualMeshList []*networking_mesh_gloo_solo_io_v1.VirtualMesh
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		virtualMeshList = append(virtualMeshList, obj.(*networking_mesh_gloo_solo_io_v1.VirtualMesh))
	}
	return virtualMeshList
}

func (s *virtualMeshSet) Map() map[string]*networking_mesh_gloo_solo_io_v1.VirtualMesh {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_mesh_gloo_solo_io_v1.VirtualMesh{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_mesh_gloo_solo_io_v1.VirtualMesh)
	}
	return newMap
}

func (s *virtualMeshSet) Insert(
	virtualMeshList ...*networking_mesh_gloo_solo_io_v1.VirtualMesh,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualMeshList {
		s.Generic().Insert(obj)
	}
}

func (s *virtualMeshSet) Has(virtualMesh ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(virtualMesh)
}

func (s *virtualMeshSet) Equal(
	virtualMeshSet VirtualMeshSet,
) bool {
	if s == nil {
		return virtualMeshSet == nil
	}
	return s.Generic().Equal(virtualMeshSet.Generic())
}

func (s *virtualMeshSet) Delete(VirtualMesh ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(VirtualMesh)
}

func (s *virtualMeshSet) Union(set VirtualMeshSet) VirtualMeshSet {
	if s == nil {
		return set
	}
	return NewVirtualMeshSet(append(s.List(), set.List()...)...)
}

func (s *virtualMeshSet) Difference(set VirtualMeshSet) VirtualMeshSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &virtualMeshSet{set: newSet}
}

func (s *virtualMeshSet) Intersection(set VirtualMeshSet) VirtualMeshSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var virtualMeshList []*networking_mesh_gloo_solo_io_v1.VirtualMesh
	for _, obj := range newSet.List() {
		virtualMeshList = append(virtualMeshList, obj.(*networking_mesh_gloo_solo_io_v1.VirtualMesh))
	}
	return NewVirtualMeshSet(virtualMeshList...)
}

func (s *virtualMeshSet) Find(id ezkube.ResourceId) (*networking_mesh_gloo_solo_io_v1.VirtualMesh, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualMesh %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_mesh_gloo_solo_io_v1.VirtualMesh{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_mesh_gloo_solo_io_v1.VirtualMesh), nil
}

func (s *virtualMeshSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *virtualMeshSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *virtualMeshSet) Delta(newSet VirtualMeshSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *virtualMeshSet) Clone() VirtualMeshSet {
	if s == nil {
		return nil
	}
	return &virtualMeshSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
