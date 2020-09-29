// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha2sets

import (
	certificates_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type IssuedCertificateSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*certificates_smh_solo_io_v1alpha2.IssuedCertificate) bool) []*certificates_smh_solo_io_v1alpha2.IssuedCertificate
	// Return the Set as a map of key to resource.
	Map() map[string]*certificates_smh_solo_io_v1alpha2.IssuedCertificate
	// Insert a resource into the set.
	Insert(issuedCertificate ...*certificates_smh_solo_io_v1alpha2.IssuedCertificate)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(issuedCertificateSet IssuedCertificateSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(issuedCertificate ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(issuedCertificate ezkube.ResourceId)
	// Return the union with the provided set
	Union(set IssuedCertificateSet) IssuedCertificateSet
	// Return the difference with the provided set
	Difference(set IssuedCertificateSet) IssuedCertificateSet
	// Return the intersection with the provided set
	Intersection(set IssuedCertificateSet) IssuedCertificateSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*certificates_smh_solo_io_v1alpha2.IssuedCertificate, error)
	// Get the length of the set
	Length() int
}

func makeGenericIssuedCertificateSet(issuedCertificateList []*certificates_smh_solo_io_v1alpha2.IssuedCertificate) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range issuedCertificateList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type issuedCertificateSet struct {
	set sksets.ResourceSet
}

func NewIssuedCertificateSet(issuedCertificateList ...*certificates_smh_solo_io_v1alpha2.IssuedCertificate) IssuedCertificateSet {
	return &issuedCertificateSet{set: makeGenericIssuedCertificateSet(issuedCertificateList)}
}

func NewIssuedCertificateSetFromList(issuedCertificateList *certificates_smh_solo_io_v1alpha2.IssuedCertificateList) IssuedCertificateSet {
	list := make([]*certificates_smh_solo_io_v1alpha2.IssuedCertificate, 0, len(issuedCertificateList.Items))
	for idx := range issuedCertificateList.Items {
		list = append(list, &issuedCertificateList.Items[idx])
	}
	return &issuedCertificateSet{set: makeGenericIssuedCertificateSet(list)}
}

func (s *issuedCertificateSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.set.Keys()
}

func (s *issuedCertificateSet) List(filterResource ...func(*certificates_smh_solo_io_v1alpha2.IssuedCertificate) bool) []*certificates_smh_solo_io_v1alpha2.IssuedCertificate {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*certificates_smh_solo_io_v1alpha2.IssuedCertificate))
		})
	}

	var issuedCertificateList []*certificates_smh_solo_io_v1alpha2.IssuedCertificate
	for _, obj := range s.set.List(genericFilters...) {
		issuedCertificateList = append(issuedCertificateList, obj.(*certificates_smh_solo_io_v1alpha2.IssuedCertificate))
	}
	return issuedCertificateList
}

func (s *issuedCertificateSet) Map() map[string]*certificates_smh_solo_io_v1alpha2.IssuedCertificate {
	if s == nil {
		return nil
	}

	newMap := map[string]*certificates_smh_solo_io_v1alpha2.IssuedCertificate{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*certificates_smh_solo_io_v1alpha2.IssuedCertificate)
	}
	return newMap
}

func (s *issuedCertificateSet) Insert(
	issuedCertificateList ...*certificates_smh_solo_io_v1alpha2.IssuedCertificate,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range issuedCertificateList {
		s.set.Insert(obj)
	}
}

func (s *issuedCertificateSet) Has(issuedCertificate ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.set.Has(issuedCertificate)
}

func (s *issuedCertificateSet) Equal(
	issuedCertificateSet IssuedCertificateSet,
) bool {
	if s == nil {
		return issuedCertificateSet == nil
	}
	return s.set.Equal(makeGenericIssuedCertificateSet(issuedCertificateSet.List()))
}

func (s *issuedCertificateSet) Delete(IssuedCertificate ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.set.Delete(IssuedCertificate)
}

func (s *issuedCertificateSet) Union(set IssuedCertificateSet) IssuedCertificateSet {
	if s == nil {
		return set
	}
	return NewIssuedCertificateSet(append(s.List(), set.List()...)...)
}

func (s *issuedCertificateSet) Difference(set IssuedCertificateSet) IssuedCertificateSet {
	if s == nil {
		return set
	}
	newSet := s.set.Difference(makeGenericIssuedCertificateSet(set.List()))
	return &issuedCertificateSet{set: newSet}
}

func (s *issuedCertificateSet) Intersection(set IssuedCertificateSet) IssuedCertificateSet {
	if s == nil {
		return nil
	}
	newSet := s.set.Intersection(makeGenericIssuedCertificateSet(set.List()))
	var issuedCertificateList []*certificates_smh_solo_io_v1alpha2.IssuedCertificate
	for _, obj := range newSet.List() {
		issuedCertificateList = append(issuedCertificateList, obj.(*certificates_smh_solo_io_v1alpha2.IssuedCertificate))
	}
	return NewIssuedCertificateSet(issuedCertificateList...)
}

func (s *issuedCertificateSet) Find(id ezkube.ResourceId) (*certificates_smh_solo_io_v1alpha2.IssuedCertificate, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find IssuedCertificate %v", sksets.Key(id))
	}
	obj, err := s.set.Find(&certificates_smh_solo_io_v1alpha2.IssuedCertificate{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*certificates_smh_solo_io_v1alpha2.IssuedCertificate), nil
}

func (s *issuedCertificateSet) Length() int {
	if s == nil {
		return 0
	}
	return s.set.Length()
}

type CertificateRequestSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*certificates_smh_solo_io_v1alpha2.CertificateRequest) bool) []*certificates_smh_solo_io_v1alpha2.CertificateRequest
	// Return the Set as a map of key to resource.
	Map() map[string]*certificates_smh_solo_io_v1alpha2.CertificateRequest
	// Insert a resource into the set.
	Insert(certificateRequest ...*certificates_smh_solo_io_v1alpha2.CertificateRequest)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(certificateRequestSet CertificateRequestSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(certificateRequest ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(certificateRequest ezkube.ResourceId)
	// Return the union with the provided set
	Union(set CertificateRequestSet) CertificateRequestSet
	// Return the difference with the provided set
	Difference(set CertificateRequestSet) CertificateRequestSet
	// Return the intersection with the provided set
	Intersection(set CertificateRequestSet) CertificateRequestSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*certificates_smh_solo_io_v1alpha2.CertificateRequest, error)
	// Get the length of the set
	Length() int
}

func makeGenericCertificateRequestSet(certificateRequestList []*certificates_smh_solo_io_v1alpha2.CertificateRequest) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range certificateRequestList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type certificateRequestSet struct {
	set sksets.ResourceSet
}

func NewCertificateRequestSet(certificateRequestList ...*certificates_smh_solo_io_v1alpha2.CertificateRequest) CertificateRequestSet {
	return &certificateRequestSet{set: makeGenericCertificateRequestSet(certificateRequestList)}
}

func NewCertificateRequestSetFromList(certificateRequestList *certificates_smh_solo_io_v1alpha2.CertificateRequestList) CertificateRequestSet {
	list := make([]*certificates_smh_solo_io_v1alpha2.CertificateRequest, 0, len(certificateRequestList.Items))
	for idx := range certificateRequestList.Items {
		list = append(list, &certificateRequestList.Items[idx])
	}
	return &certificateRequestSet{set: makeGenericCertificateRequestSet(list)}
}

func (s *certificateRequestSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.set.Keys()
}

func (s *certificateRequestSet) List(filterResource ...func(*certificates_smh_solo_io_v1alpha2.CertificateRequest) bool) []*certificates_smh_solo_io_v1alpha2.CertificateRequest {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*certificates_smh_solo_io_v1alpha2.CertificateRequest))
		})
	}

	var certificateRequestList []*certificates_smh_solo_io_v1alpha2.CertificateRequest
	for _, obj := range s.set.List(genericFilters...) {
		certificateRequestList = append(certificateRequestList, obj.(*certificates_smh_solo_io_v1alpha2.CertificateRequest))
	}
	return certificateRequestList
}

func (s *certificateRequestSet) Map() map[string]*certificates_smh_solo_io_v1alpha2.CertificateRequest {
	if s == nil {
		return nil
	}

	newMap := map[string]*certificates_smh_solo_io_v1alpha2.CertificateRequest{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*certificates_smh_solo_io_v1alpha2.CertificateRequest)
	}
	return newMap
}

func (s *certificateRequestSet) Insert(
	certificateRequestList ...*certificates_smh_solo_io_v1alpha2.CertificateRequest,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range certificateRequestList {
		s.set.Insert(obj)
	}
}

func (s *certificateRequestSet) Has(certificateRequest ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.set.Has(certificateRequest)
}

func (s *certificateRequestSet) Equal(
	certificateRequestSet CertificateRequestSet,
) bool {
	if s == nil {
		return certificateRequestSet == nil
	}
	return s.set.Equal(makeGenericCertificateRequestSet(certificateRequestSet.List()))
}

func (s *certificateRequestSet) Delete(CertificateRequest ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.set.Delete(CertificateRequest)
}

func (s *certificateRequestSet) Union(set CertificateRequestSet) CertificateRequestSet {
	if s == nil {
		return set
	}
	return NewCertificateRequestSet(append(s.List(), set.List()...)...)
}

func (s *certificateRequestSet) Difference(set CertificateRequestSet) CertificateRequestSet {
	if s == nil {
		return set
	}
	newSet := s.set.Difference(makeGenericCertificateRequestSet(set.List()))
	return &certificateRequestSet{set: newSet}
}

func (s *certificateRequestSet) Intersection(set CertificateRequestSet) CertificateRequestSet {
	if s == nil {
		return nil
	}
	newSet := s.set.Intersection(makeGenericCertificateRequestSet(set.List()))
	var certificateRequestList []*certificates_smh_solo_io_v1alpha2.CertificateRequest
	for _, obj := range newSet.List() {
		certificateRequestList = append(certificateRequestList, obj.(*certificates_smh_solo_io_v1alpha2.CertificateRequest))
	}
	return NewCertificateRequestSet(certificateRequestList...)
}

func (s *certificateRequestSet) Find(id ezkube.ResourceId) (*certificates_smh_solo_io_v1alpha2.CertificateRequest, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find CertificateRequest %v", sksets.Key(id))
	}
	obj, err := s.set.Find(&certificates_smh_solo_io_v1alpha2.CertificateRequest{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*certificates_smh_solo_io_v1alpha2.CertificateRequest), nil
}

func (s *certificateRequestSet) Length() int {
	if s == nil {
		return 0
	}
	return s.set.Length()
}

type PodBounceDirectiveSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*certificates_smh_solo_io_v1alpha2.PodBounceDirective) bool) []*certificates_smh_solo_io_v1alpha2.PodBounceDirective
	// Return the Set as a map of key to resource.
	Map() map[string]*certificates_smh_solo_io_v1alpha2.PodBounceDirective
	// Insert a resource into the set.
	Insert(podBounceDirective ...*certificates_smh_solo_io_v1alpha2.PodBounceDirective)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(podBounceDirectiveSet PodBounceDirectiveSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(podBounceDirective ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(podBounceDirective ezkube.ResourceId)
	// Return the union with the provided set
	Union(set PodBounceDirectiveSet) PodBounceDirectiveSet
	// Return the difference with the provided set
	Difference(set PodBounceDirectiveSet) PodBounceDirectiveSet
	// Return the intersection with the provided set
	Intersection(set PodBounceDirectiveSet) PodBounceDirectiveSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*certificates_smh_solo_io_v1alpha2.PodBounceDirective, error)
	// Get the length of the set
	Length() int
}

func makeGenericPodBounceDirectiveSet(podBounceDirectiveList []*certificates_smh_solo_io_v1alpha2.PodBounceDirective) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range podBounceDirectiveList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type podBounceDirectiveSet struct {
	set sksets.ResourceSet
}

func NewPodBounceDirectiveSet(podBounceDirectiveList ...*certificates_smh_solo_io_v1alpha2.PodBounceDirective) PodBounceDirectiveSet {
	return &podBounceDirectiveSet{set: makeGenericPodBounceDirectiveSet(podBounceDirectiveList)}
}

func NewPodBounceDirectiveSetFromList(podBounceDirectiveList *certificates_smh_solo_io_v1alpha2.PodBounceDirectiveList) PodBounceDirectiveSet {
	list := make([]*certificates_smh_solo_io_v1alpha2.PodBounceDirective, 0, len(podBounceDirectiveList.Items))
	for idx := range podBounceDirectiveList.Items {
		list = append(list, &podBounceDirectiveList.Items[idx])
	}
	return &podBounceDirectiveSet{set: makeGenericPodBounceDirectiveSet(list)}
}

func (s *podBounceDirectiveSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.set.Keys()
}

func (s *podBounceDirectiveSet) List(filterResource ...func(*certificates_smh_solo_io_v1alpha2.PodBounceDirective) bool) []*certificates_smh_solo_io_v1alpha2.PodBounceDirective {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*certificates_smh_solo_io_v1alpha2.PodBounceDirective))
		})
	}

	var podBounceDirectiveList []*certificates_smh_solo_io_v1alpha2.PodBounceDirective
	for _, obj := range s.set.List(genericFilters...) {
		podBounceDirectiveList = append(podBounceDirectiveList, obj.(*certificates_smh_solo_io_v1alpha2.PodBounceDirective))
	}
	return podBounceDirectiveList
}

func (s *podBounceDirectiveSet) Map() map[string]*certificates_smh_solo_io_v1alpha2.PodBounceDirective {
	if s == nil {
		return nil
	}

	newMap := map[string]*certificates_smh_solo_io_v1alpha2.PodBounceDirective{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*certificates_smh_solo_io_v1alpha2.PodBounceDirective)
	}
	return newMap
}

func (s *podBounceDirectiveSet) Insert(
	podBounceDirectiveList ...*certificates_smh_solo_io_v1alpha2.PodBounceDirective,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range podBounceDirectiveList {
		s.set.Insert(obj)
	}
}

func (s *podBounceDirectiveSet) Has(podBounceDirective ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.set.Has(podBounceDirective)
}

func (s *podBounceDirectiveSet) Equal(
	podBounceDirectiveSet PodBounceDirectiveSet,
) bool {
	if s == nil {
		return podBounceDirectiveSet == nil
	}
	return s.set.Equal(makeGenericPodBounceDirectiveSet(podBounceDirectiveSet.List()))
}

func (s *podBounceDirectiveSet) Delete(PodBounceDirective ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.set.Delete(PodBounceDirective)
}

func (s *podBounceDirectiveSet) Union(set PodBounceDirectiveSet) PodBounceDirectiveSet {
	if s == nil {
		return set
	}
	return NewPodBounceDirectiveSet(append(s.List(), set.List()...)...)
}

func (s *podBounceDirectiveSet) Difference(set PodBounceDirectiveSet) PodBounceDirectiveSet {
	if s == nil {
		return set
	}
	newSet := s.set.Difference(makeGenericPodBounceDirectiveSet(set.List()))
	return &podBounceDirectiveSet{set: newSet}
}

func (s *podBounceDirectiveSet) Intersection(set PodBounceDirectiveSet) PodBounceDirectiveSet {
	if s == nil {
		return nil
	}
	newSet := s.set.Intersection(makeGenericPodBounceDirectiveSet(set.List()))
	var podBounceDirectiveList []*certificates_smh_solo_io_v1alpha2.PodBounceDirective
	for _, obj := range newSet.List() {
		podBounceDirectiveList = append(podBounceDirectiveList, obj.(*certificates_smh_solo_io_v1alpha2.PodBounceDirective))
	}
	return NewPodBounceDirectiveSet(podBounceDirectiveList...)
}

func (s *podBounceDirectiveSet) Find(id ezkube.ResourceId) (*certificates_smh_solo_io_v1alpha2.PodBounceDirective, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find PodBounceDirective %v", sksets.Key(id))
	}
	obj, err := s.set.Find(&certificates_smh_solo_io_v1alpha2.PodBounceDirective{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*certificates_smh_solo_io_v1alpha2.PodBounceDirective), nil
}

func (s *podBounceDirectiveSet) Length() int {
	if s == nil {
		return 0
	}
	return s.set.Length()
}
