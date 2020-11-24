package metautils

import (
	"encoding/json"
	"fmt"
	"strings"

	discoveryv1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/ezkube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// the key used to differentiate translated resources by
	// the GlooMesh instance which produced them
	OwnershipLabelKey = fmt.Sprintf("owner.%s", v1alpha2.SchemeGroupVersion.Group)

	// Annotation key indicating that the resource configures a federated traffic target
	FederationLabelKey = fmt.Sprintf("federation.%s", v1alpha2.SchemeGroupVersion.Group)

	// Annotation key for tracking the parent resources that were translated in the creation of a child resource
	ParentLabelkey = fmt.Sprintf("parents.%s", v1alpha2.SchemeGroupVersion.Group)
)

// construct an ObjectMeta for a discovered resource from a source object (the object from which the resource was discovered)
func TranslatedObjectMeta(sourceObj ezkube.ClusterResourceId, annotations map[string]string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:        sourceObj.GetName(),
		Namespace:   sourceObj.GetNamespace(),
		ClusterName: sourceObj.GetClusterName(),
		Labels:      TranslatedObjectLabels(),
		Annotations: annotations,
	}
}

// construct an ObjectMeta for a resource for a federated source object
// meshInstallation represents the mesh instance to which the object will be output
func FederatedObjectMeta(
	sourceObj ezkube.ClusterResourceId,
	meshInstallation *discoveryv1alpha2.MeshSpec_MeshInstallation,
	annotations map[string]string,
) metav1.ObjectMeta {
	if annotations == nil {
		annotations = map[string]string{}
	}
	annotations[FederationLabelKey] = sets.Key(sourceObj)

	return metav1.ObjectMeta{
		Name:        federatedObjectName(sourceObj),
		Namespace:   meshInstallation.Namespace,
		ClusterName: meshInstallation.Cluster,
		Labels:      TranslatedObjectLabels(),
		Annotations: annotations,
	}
}

func federatedObjectName(
	sourceObj ezkube.ClusterResourceId,
) string {
	return strings.Join([]string{sourceObj.GetName(), sourceObj.GetNamespace(), sourceObj.GetClusterName()}, "-")
}

// ownership label defaults to current namespace to allow multiple GlooMesh tenancy within a cluster.
func TranslatedObjectLabels() map[string]string {
	return map[string]string{OwnershipLabelKey: defaults.GetPodNamespace()}
}

// add a parent to the annotation for a given child object
func AppendParent(
	child metav1.Object,
	parentId ezkube.ResourceId,
	parentGVK schema.GroupVersionKind,
) error {
	annotations := child.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	parents := make(map[string][]*v1.ObjectRef)
	parentsStr, ok := annotations[ParentLabelkey]
	if ok {
		if err := json.Unmarshal([]byte(parentsStr), &parents); err != nil {
			return err
		}
	}

	curParents, ok := parents[parentGVK.String()]
	if !ok {
		curParents = make([]*v1.ObjectRef, 0, 1)
	}
	parentRef := ezkube.MakeObjectRef(parentId)
	for _, parent := range curParents {
		if parent.Equal(parentRef) {
			return nil
		}
	}
	parents[parentGVK.String()] = append(curParents, parentRef)

	b, err := json.Marshal(parents)
	if err != nil {
		return err
	}

	annotations[ParentLabelkey] = string(b)
	child.SetAnnotations(annotations)
	return nil
}
