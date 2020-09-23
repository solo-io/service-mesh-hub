package configtarget

import (
	"fmt"
	"sort"

	discoveryv1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2/sets"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

/*
	Validate configuration target references in networking configuration resources, and report
	any errors (i.e. references to non-existent discovery entities) to the offending resource status.
*/
type ConfigTargetValidator interface {
	// Validate mesh references declared on FailoverServices.
	ValidateFailoverServices(
		failoverServices v1alpha2.FailoverServiceSlice,
	)

	// Validate traffic target references declared on TrafficPolicies.
	ValidateTrafficPolicies(
		trafficPolicies v1alpha2.TrafficPolicySlice,
	)

	// Validate mesh references declared on VirtualMeshes.
	// Also validate that all referenced meshes are contained in at most one virtual mesh.
	ValidateVirtualMeshes(
		virtualMeshes v1alpha2.VirtualMeshSlice,
	)

	// Validate traffic target references declared on AccessPolicies.
	ValidateAccessPolicies(
		accessPolicies v1alpha2.AccessPolicySlice,
	)
}

type configTargetValidator struct {
	meshes         discoveryv1alpha2sets.MeshSet
	trafficTargets discoveryv1alpha2sets.TrafficTargetSet
}

func NewConfigTargetValidator(
	meshes discoveryv1alpha2sets.MeshSet,
	trafficTargets discoveryv1alpha2sets.TrafficTargetSet,
) ConfigTargetValidator {
	return &configTargetValidator{
		meshes:         meshes,
		trafficTargets: trafficTargets,
	}
}

func (c *configTargetValidator) ValidateFailoverServices(failoverServices v1alpha2.FailoverServiceSlice) {
	for _, failoverService := range failoverServices {
		failoverService := failoverService // pike
		errs := c.validateMeshReferences(failoverService.Spec.Meshes)
		var errStrings []string
		for _, err := range errs {
			errStrings = append(errStrings, err.Error())
		}
		if errStrings != nil {
			failoverService.Status.State = v1alpha2.ApprovalState_INVALID
			failoverService.Status.ConfigTargetErrors = errStrings
		}
	}
}

func (c *configTargetValidator) ValidateVirtualMeshes(virtualMeshes v1alpha2.VirtualMeshSlice) {
	for _, virtualMesh := range virtualMeshes {
		virtualMesh := virtualMesh // pike
		errs := c.validateVirtualMesh(virtualMesh)
		var errStrings []string
		for _, err := range errs {
			errStrings = append(errStrings, err.Error())
		}
		if errStrings != nil {
			virtualMesh.Status.State = v1alpha2.ApprovalState_INVALID
			virtualMesh.Status.ConfigTargetErrors = errStrings
		}
	}
	validateOneVirtualMeshPerMesh(virtualMeshes)
}

func (c *configTargetValidator) ValidateTrafficPolicies(trafficPolicies v1alpha2.TrafficPolicySlice) {
	for _, trafficPolicy := range trafficPolicies {
		trafficPolicy := trafficPolicy // pike
		errs := c.validateTrafficTargetReferences(trafficPolicy.Spec.DestinationSelector)
		var errStrings []string
		for _, err := range errs {
			errStrings = append(errStrings, err.Error())
		}
		if errStrings != nil {
			trafficPolicy.Status.State = v1alpha2.ApprovalState_INVALID
			trafficPolicy.Status.ConfigTargetErrors = errStrings
		}
	}
}

func (c *configTargetValidator) ValidateAccessPolicies(accessPolicies v1alpha2.AccessPolicySlice) {
	for _, accessPolicy := range accessPolicies {
		accessPolicy := accessPolicy // pike
		errs := c.validateTrafficTargetReferences(accessPolicy.Spec.DestinationSelector)
		var errStrings []string
		for _, err := range errs {
			errStrings = append(errStrings, err.Error())
		}
		if errStrings != nil {
			accessPolicy.Status = v1alpha2.AccessPolicyStatus{
				ObservedGeneration: accessPolicy.Generation,
				State:              v1alpha2.ApprovalState_INVALID,
				ConfigTargetErrors: errStrings,
			}
		}
	}
}

func (c *configTargetValidator) validateMeshReferences(meshRefs []*v1.ObjectRef) []error {
	var errs []error
	for _, meshRef := range meshRefs {
		if _, err := c.meshes.Find(meshRef); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (c *configTargetValidator) validateTrafficTargetReferences(serviceSelectors []*v1alpha2.ServiceSelector) []error {
	var errs []error
	for _, destinationSelector := range serviceSelectors {
		kubeServiceRefs := destinationSelector.GetKubeServiceRefs()
		// only validate traffic targets selected by direct reference
		if kubeServiceRefs == nil {
			continue
		}
		for _, ref := range kubeServiceRefs.Services {
			if _, err := c.trafficTargets.Find(ref); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errs
}

func (c *configTargetValidator) validateVirtualMesh(virtualMesh *v1alpha2.VirtualMesh) []error {
	var errs []error
	meshRefErrors := c.validateMeshReferences(virtualMesh.Spec.Meshes)
	if meshRefErrors != nil {
		errs = append(errs, meshRefErrors...)
	}
	return errs
}

// For each VirtualMesh, sort them by accepted date, then invalidate if it applies to a Mesh that
// is already grouped into a VirtualMesh.
func validateOneVirtualMeshPerMesh(virtualMeshes []*v1alpha2.VirtualMesh) {
	sortVirtualMeshesByAcceptedDate(virtualMeshes)

	vMeshesPerMesh := map[string]*v1alpha2.VirtualMesh{}
	invalidVirtualMeshes := v1alpha2sets.NewVirtualMeshSet()

	// track accepted index
	var acceptedIndex uint32
	// Invalidate VirtualMesh if it applies to a Mesh that already has an applied VirtualMesh.
	for _, vMesh := range virtualMeshes {
		if vMesh.Status.State != v1alpha2.ApprovalState_ACCEPTED {
			continue
		}
		vMesh := vMesh
		for _, mesh := range vMesh.Spec.Meshes {
			// Ignore virtual mesh if previously invalidated.
			if invalidVirtualMeshes.Has(vMesh) {
				continue
			}
			meshKey := sets.Key(mesh)
			existingVirtualMesh, ok := vMeshesPerMesh[meshKey]
			vMesh.Status.ObservedGeneration = vMesh.Generation
			if !ok {
				vMeshesPerMesh[meshKey] = vMesh
				acceptedIndex++
			} else {
				vMesh.Status.State = v1alpha2.ApprovalState_INVALID
				vMesh.Status.ConfigTargetErrors = []string{
					fmt.Sprintf("Includes a Mesh (%s.%s) that already is grouped in a VirtualMesh (%s.%s)",
						mesh.Name, mesh.Namespace,
						existingVirtualMesh.Name, existingVirtualMesh.Namespace,
					)}
			}
			invalidVirtualMeshes.Insert(vMesh)
		}
	}
}

// sort the set of virtual meshes in the order in which they were accepted.
// VMeshes which were accepted first and have not changed (i.e. their observedGeneration is up-to-date) take precedence.
// Next are vMeshes that were previously accepted but whose observedGeneration is out of date. This permits vmeshes which were modified but formerly correct to maintain
// their acceptance status ahead of vmeshes which were unmodified and previously rejected.
// Next will be the vmeshes which have been modified and rejected.
// Finally, vmeshes which are rejected and modified
func sortVirtualMeshesByAcceptedDate(virtualMeshes v1alpha2.VirtualMeshSlice) {
	isUpToDate := func(vm *v1alpha2.VirtualMesh) bool {
		return vm.Status.ObservedGeneration == vm.Generation
	}

	sort.SliceStable(virtualMeshes, func(i, j int) bool {
		vMesh1, vMesh2 := virtualMeshes[i], virtualMeshes[j]

		state1 := vMesh1.Status.State
		state2 := vMesh2.Status.State

		switch {
		case state1 == v1alpha2.ApprovalState_ACCEPTED:
			if state2 != v1alpha2.ApprovalState_ACCEPTED {
				// accepted comes before non accepted
				return true
			}

			if vMesh1UpToDate := isUpToDate(vMesh1); vMesh1UpToDate != isUpToDate(vMesh2) {
				// up to date is validated before modified
				return vMesh1UpToDate
			}

			return true
		case state2 == v1alpha2.ApprovalState_ACCEPTED:
			// accepted comes before non accepted
			return false
		default:
			// neither policy has been accepted, we can simply sort by unique key
			return sets.Key(vMesh1) < sets.Key(vMesh2)
		}
	})
}