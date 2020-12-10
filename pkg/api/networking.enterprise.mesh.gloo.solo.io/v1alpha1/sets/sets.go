// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha1sets

import (
	networking_enterprise_mesh_gloo_solo_io_v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1alpha1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type WasmDeploymentSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment) bool) []*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment
	// Insert a resource into the set.
	Insert(wasmDeployment ...*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(wasmDeploymentSet WasmDeploymentSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(wasmDeployment ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(wasmDeployment ezkube.ResourceId)
	// Return the union with the provided set
	Union(set WasmDeploymentSet) WasmDeploymentSet
	// Return the difference with the provided set
	Difference(set WasmDeploymentSet) WasmDeploymentSet
	// Return the intersection with the provided set
	Intersection(set WasmDeploymentSet) WasmDeploymentSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment, error)
	// Get the length of the set
	Length() int
}

func makeGenericWasmDeploymentSet(wasmDeploymentList []*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range wasmDeploymentList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type wasmDeploymentSet struct {
	set sksets.ResourceSet
}

func NewWasmDeploymentSet(wasmDeploymentList ...*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment) WasmDeploymentSet {
	return &wasmDeploymentSet{set: makeGenericWasmDeploymentSet(wasmDeploymentList)}
}

func NewWasmDeploymentSetFromList(wasmDeploymentList *networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeploymentList) WasmDeploymentSet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment, 0, len(wasmDeploymentList.Items))
	for idx := range wasmDeploymentList.Items {
		list = append(list, &wasmDeploymentList.Items[idx])
	}
	return &wasmDeploymentSet{set: makeGenericWasmDeploymentSet(list)}
}

func (s *wasmDeploymentSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.set.Keys()
}

func (s *wasmDeploymentSet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment) bool) []*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment))
		})
	}

	var wasmDeploymentList []*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment
	for _, obj := range s.set.List(genericFilters...) {
		wasmDeploymentList = append(wasmDeploymentList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment))
	}
	return wasmDeploymentList
}

func (s *wasmDeploymentSet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment)
	}
	return newMap
}

func (s *wasmDeploymentSet) Insert(
	wasmDeploymentList ...*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range wasmDeploymentList {
		s.set.Insert(obj)
	}
}

func (s *wasmDeploymentSet) Has(wasmDeployment ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.set.Has(wasmDeployment)
}

func (s *wasmDeploymentSet) Equal(
	wasmDeploymentSet WasmDeploymentSet,
) bool {
	if s == nil {
		return wasmDeploymentSet == nil
	}
	return s.set.Equal(makeGenericWasmDeploymentSet(wasmDeploymentSet.List()))
}

func (s *wasmDeploymentSet) Delete(WasmDeployment ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.set.Delete(WasmDeployment)
}

func (s *wasmDeploymentSet) Union(set WasmDeploymentSet) WasmDeploymentSet {
	if s == nil {
		return set
	}
	return NewWasmDeploymentSet(append(s.List(), set.List()...)...)
}

func (s *wasmDeploymentSet) Difference(set WasmDeploymentSet) WasmDeploymentSet {
	if s == nil {
		return set
	}
	newSet := s.set.Difference(makeGenericWasmDeploymentSet(set.List()))
	return &wasmDeploymentSet{set: newSet}
}

func (s *wasmDeploymentSet) Intersection(set WasmDeploymentSet) WasmDeploymentSet {
	if s == nil {
		return nil
	}
	newSet := s.set.Intersection(makeGenericWasmDeploymentSet(set.List()))
	var wasmDeploymentList []*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment
	for _, obj := range newSet.List() {
		wasmDeploymentList = append(wasmDeploymentList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment))
	}
	return NewWasmDeploymentSet(wasmDeploymentList...)
}

func (s *wasmDeploymentSet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find WasmDeployment %v", sksets.Key(id))
	}
	obj, err := s.set.Find(&networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment), nil
}

func (s *wasmDeploymentSet) Length() int {
	if s == nil {
		return 0
	}
	return s.set.Length()
}
