// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets

import (
	observability_enterprise_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/observability.enterprise.mesh.gloo.solo.io/v1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type AccessLogRecordSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord) bool) []*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord) bool) []*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord
	// Return the Set as a map of key to resource.
	Map() map[string]*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord
	// Insert a resource into the set.
	Insert(accessLogRecord ...*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(accessLogRecordSet AccessLogRecordSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(accessLogRecord ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(accessLogRecord ezkube.ResourceId)
	// Return the union with the provided set
	Union(set AccessLogRecordSet) AccessLogRecordSet
	// Return the difference with the provided set
	Difference(set AccessLogRecordSet) AccessLogRecordSet
	// Return the intersection with the provided set
	Intersection(set AccessLogRecordSet) AccessLogRecordSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another AccessLogRecordSet
	Delta(newSet AccessLogRecordSet) sksets.ResourceDelta
}

func makeGenericAccessLogRecordSet(accessLogRecordList []*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range accessLogRecordList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type accessLogRecordSet struct {
	set sksets.ResourceSet
}

func NewAccessLogRecordSet(accessLogRecordList ...*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord) AccessLogRecordSet {
	return &accessLogRecordSet{set: makeGenericAccessLogRecordSet(accessLogRecordList)}
}

func NewAccessLogRecordSetFromList(accessLogRecordList *observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecordList) AccessLogRecordSet {
	list := make([]*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord, 0, len(accessLogRecordList.Items))
	for idx := range accessLogRecordList.Items {
		list = append(list, &accessLogRecordList.Items[idx])
	}
	return &accessLogRecordSet{set: makeGenericAccessLogRecordSet(list)}
}

func (s *accessLogRecordSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *accessLogRecordSet) List(filterResource ...func(*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord) bool) []*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord))
		})
	}

	var accessLogRecordList []*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord
	for _, obj := range s.Generic().List(genericFilters...) {
		accessLogRecordList = append(accessLogRecordList, obj.(*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord))
	}
	return accessLogRecordList
}

func (s *accessLogRecordSet) UnsortedList(filterResource ...func(*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord) bool) []*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord))
		})
	}

	var accessLogRecordList []*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		accessLogRecordList = append(accessLogRecordList, obj.(*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord))
	}
	return accessLogRecordList
}

func (s *accessLogRecordSet) Map() map[string]*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord {
	if s == nil {
		return nil
	}

	newMap := map[string]*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord)
	}
	return newMap
}

func (s *accessLogRecordSet) Insert(
	accessLogRecordList ...*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range accessLogRecordList {
		s.Generic().Insert(obj)
	}
}

func (s *accessLogRecordSet) Has(accessLogRecord ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(accessLogRecord)
}

func (s *accessLogRecordSet) Equal(
	accessLogRecordSet AccessLogRecordSet,
) bool {
	if s == nil {
		return accessLogRecordSet == nil
	}
	return s.Generic().Equal(accessLogRecordSet.Generic())
}

func (s *accessLogRecordSet) Delete(AccessLogRecord ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(AccessLogRecord)
}

func (s *accessLogRecordSet) Union(set AccessLogRecordSet) AccessLogRecordSet {
	if s == nil {
		return set
	}
	return NewAccessLogRecordSet(append(s.List(), set.List()...)...)
}

func (s *accessLogRecordSet) Difference(set AccessLogRecordSet) AccessLogRecordSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &accessLogRecordSet{set: newSet}
}

func (s *accessLogRecordSet) Intersection(set AccessLogRecordSet) AccessLogRecordSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var accessLogRecordList []*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord
	for _, obj := range newSet.List() {
		accessLogRecordList = append(accessLogRecordList, obj.(*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord))
	}
	return NewAccessLogRecordSet(accessLogRecordList...)
}

func (s *accessLogRecordSet) Find(id ezkube.ResourceId) (*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find AccessLogRecord %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord), nil
}

func (s *accessLogRecordSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *accessLogRecordSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *accessLogRecordSet) Delta(newSet AccessLogRecordSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}
