// Code generated by skv2. DO NOT EDIT.

// Definitions for the Kubernetes types
package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// TrafficPolicy is the Schema for the trafficPolicy API
type TrafficPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrafficPolicySpec   `json:"spec,omitempty"`
	Status TrafficPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (TrafficPolicy) GVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "networking.mesh.gloo.solo.io/v1alpha2",
		Version: "v1alpha2",
		Kind:    "TrafficPolicy",
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficPolicyList contains a list of TrafficPolicy
type TrafficPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// AccessPolicy is the Schema for the accessPolicy API
type AccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessPolicySpec   `json:"spec,omitempty"`
	Status AccessPolicyStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (AccessPolicy) GVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "networking.mesh.gloo.solo.io/v1alpha2",
		Version: "v1alpha2",
		Kind:    "AccessPolicy",
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessPolicyList contains a list of AccessPolicy
type AccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// VirtualMesh is the Schema for the virtualMesh API
type VirtualMesh struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMeshSpec   `json:"spec,omitempty"`
	Status VirtualMeshStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (VirtualMesh) GVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "networking.mesh.gloo.solo.io/v1alpha2",
		Version: "v1alpha2",
		Kind:    "VirtualMesh",
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualMeshList contains a list of VirtualMesh
type VirtualMeshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMesh `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// FailoverService is the Schema for the failoverService API
type FailoverService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FailoverServiceSpec   `json:"spec,omitempty"`
	Status FailoverServiceStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (FailoverService) GVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "networking.mesh.gloo.solo.io/v1alpha2",
		Version: "v1alpha2",
		Kind:    "FailoverService",
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FailoverServiceList contains a list of FailoverService
type FailoverServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FailoverService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrafficPolicy{}, &TrafficPolicyList{})
	SchemeBuilder.Register(&AccessPolicy{}, &AccessPolicyList{})
	SchemeBuilder.Register(&VirtualMesh{}, &VirtualMeshList{})
	SchemeBuilder.Register(&FailoverService{}, &FailoverServiceList{})
}
