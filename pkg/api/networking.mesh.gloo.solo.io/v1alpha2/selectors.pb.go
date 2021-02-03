// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: github.com/solo-io/gloo-mesh/api/networking/v1alpha2/selectors.proto

package v1alpha2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//
//Select TrafficTargets using one or more platform-specific selection objects.
type TrafficTargetSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A KubeServiceMatcher matches kubernetes services by their labels, namespaces, and/or clusters.
	KubeServiceMatcher *TrafficTargetSelector_KubeServiceMatcher `protobuf:"bytes,1,opt,name=kube_service_matcher,json=kubeServiceMatcher,proto3" json:"kube_service_matcher,omitempty"`
	// Match individual k8s Services by direct reference.
	KubeServiceRefs *TrafficTargetSelector_KubeServiceRefs `protobuf:"bytes,2,opt,name=kube_service_refs,json=kubeServiceRefs,proto3" json:"kube_service_refs,omitempty"`
}

func (x *TrafficTargetSelector) Reset() {
	*x = TrafficTargetSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficTargetSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficTargetSelector) ProtoMessage() {}

func (x *TrafficTargetSelector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficTargetSelector.ProtoReflect.Descriptor instead.
func (*TrafficTargetSelector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescGZIP(), []int{0}
}

func (x *TrafficTargetSelector) GetKubeServiceMatcher() *TrafficTargetSelector_KubeServiceMatcher {
	if x != nil {
		return x.KubeServiceMatcher
	}
	return nil
}

func (x *TrafficTargetSelector) GetKubeServiceRefs() *TrafficTargetSelector_KubeServiceRefs {
	if x != nil {
		return x.KubeServiceRefs
	}
	return nil
}

//
//Select Kubernetes workloads directly using label namespace and/or cluster criteria. See comments on the fields for
//detailed semantics.
type WorkloadSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	//If specified, all labels must exist on k8s workload.
	//When used in a networking policy, omission matches any labels.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any label key and/or value.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//
	//If specified, match k8s workloads if they exist in one of the specified namespaces.
	//When used in a networking policy, omission matches any namespace.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any namespace.
	Namespaces []string `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	//
	//If specified, match k8s workloads if they exist in one of the specified clusters.
	//When used in a networking policy, omission matches any cluster.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any cluster.
	Clusters []string `protobuf:"bytes,3,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *WorkloadSelector) Reset() {
	*x = WorkloadSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadSelector) ProtoMessage() {}

func (x *WorkloadSelector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadSelector.ProtoReflect.Descriptor instead.
func (*WorkloadSelector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescGZIP(), []int{1}
}

func (x *WorkloadSelector) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *WorkloadSelector) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *WorkloadSelector) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

//
//Selector capable of selecting specific service identities. Useful for binding policy rules.
//Either (namespaces, cluster, service_account_names) or service_accounts can be specified.
//If all fields are omitted, any source identity is permitted.
type IdentitySelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A KubeIdentityMatcher matches request identities based on the k8s namespace and cluster.
	KubeIdentityMatcher *IdentitySelector_KubeIdentityMatcher `protobuf:"bytes,1,opt,name=kube_identity_matcher,json=kubeIdentityMatcher,proto3" json:"kube_identity_matcher,omitempty"`
	// KubeServiceAccountRefs matches request identities based on the k8s service account of request.
	KubeServiceAccountRefs *IdentitySelector_KubeServiceAccountRefs `protobuf:"bytes,2,opt,name=kube_service_account_refs,json=kubeServiceAccountRefs,proto3" json:"kube_service_account_refs,omitempty"`
}

func (x *IdentitySelector) Reset() {
	*x = IdentitySelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentitySelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentitySelector) ProtoMessage() {}

func (x *IdentitySelector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentitySelector.ProtoReflect.Descriptor instead.
func (*IdentitySelector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescGZIP(), []int{2}
}

func (x *IdentitySelector) GetKubeIdentityMatcher() *IdentitySelector_KubeIdentityMatcher {
	if x != nil {
		return x.KubeIdentityMatcher
	}
	return nil
}

func (x *IdentitySelector) GetKubeServiceAccountRefs() *IdentitySelector_KubeServiceAccountRefs {
	if x != nil {
		return x.KubeServiceAccountRefs
	}
	return nil
}

type TrafficTargetSelector_KubeServiceMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	//If specified, all labels must exist on k8s Service.
	//When used in a networking policy, omission matches any labels.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any label key and/or value.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//
	//If specified, match k8s Services if they exist in one of the specified namespaces.
	//When used in a networking policy, omission matches any namespace.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any namespace.
	Namespaces []string `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	//
	//If specified, match k8s Services if they exist in one of the specified clusters.
	//When used in a networking policy, omission matches any cluster.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any cluster.
	Clusters []string `protobuf:"bytes,3,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *TrafficTargetSelector_KubeServiceMatcher) Reset() {
	*x = TrafficTargetSelector_KubeServiceMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficTargetSelector_KubeServiceMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficTargetSelector_KubeServiceMatcher) ProtoMessage() {}

func (x *TrafficTargetSelector_KubeServiceMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficTargetSelector_KubeServiceMatcher.ProtoReflect.Descriptor instead.
func (*TrafficTargetSelector_KubeServiceMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TrafficTargetSelector_KubeServiceMatcher) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *TrafficTargetSelector_KubeServiceMatcher) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *TrafficTargetSelector_KubeServiceMatcher) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type TrafficTargetSelector_KubeServiceRefs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	//Match k8s Services by direct reference.
	//When used in a networking policy, omission of any field (name, namespace, or clusterName) allows matching any value for that field.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any value for the given field.
	Services []*v1.ClusterObjectRef `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *TrafficTargetSelector_KubeServiceRefs) Reset() {
	*x = TrafficTargetSelector_KubeServiceRefs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficTargetSelector_KubeServiceRefs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficTargetSelector_KubeServiceRefs) ProtoMessage() {}

func (x *TrafficTargetSelector_KubeServiceRefs) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficTargetSelector_KubeServiceRefs.ProtoReflect.Descriptor instead.
func (*TrafficTargetSelector_KubeServiceRefs) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescGZIP(), []int{0, 1}
}

func (x *TrafficTargetSelector_KubeServiceRefs) GetServices() []*v1.ClusterObjectRef {
	if x != nil {
		return x.Services
	}
	return nil
}

type IdentitySelector_KubeIdentityMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	//If specified, match k8s identity if it exists in one of the specified namespaces.
	//When used in a networking policy, omission matches any namespace.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any namespace.
	Namespaces []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	//
	//If specified, match k8s identity if it exists in one of the specified clusters.
	//When used in a networking policy, omission matches any cluster.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any cluster.
	Clusters []string `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *IdentitySelector_KubeIdentityMatcher) Reset() {
	*x = IdentitySelector_KubeIdentityMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentitySelector_KubeIdentityMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentitySelector_KubeIdentityMatcher) ProtoMessage() {}

func (x *IdentitySelector_KubeIdentityMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentitySelector_KubeIdentityMatcher.ProtoReflect.Descriptor instead.
func (*IdentitySelector_KubeIdentityMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescGZIP(), []int{2, 0}
}

func (x *IdentitySelector_KubeIdentityMatcher) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *IdentitySelector_KubeIdentityMatcher) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type IdentitySelector_KubeServiceAccountRefs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	//Match k8s ServiceAccounts by direct reference.
	//When used in a networking policy, omission of any field (name, namespace, or clusterName) allows matching any value for that field.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any value for the given field.
	ServiceAccounts []*v1.ClusterObjectRef `protobuf:"bytes,1,rep,name=service_accounts,json=serviceAccounts,proto3" json:"service_accounts,omitempty"`
}

func (x *IdentitySelector_KubeServiceAccountRefs) Reset() {
	*x = IdentitySelector_KubeServiceAccountRefs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentitySelector_KubeServiceAccountRefs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentitySelector_KubeServiceAccountRefs) ProtoMessage() {}

func (x *IdentitySelector_KubeServiceAccountRefs) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentitySelector_KubeServiceAccountRefs.ProtoReflect.Descriptor instead.
func (*IdentitySelector_KubeServiceAccountRefs) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescGZIP(), []int{2, 1}
}

func (x *IdentitySelector_KubeServiceAccountRefs) GetServiceAccounts() []*v1.ClusterObjectRef {
	if x != nil {
		return x.ServiceAccounts
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x04, 0x0a, 0x15, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x78, 0x0a, 0x14, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x46, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x12, 0x6b, 0x75, 0x62, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x6f, 0x0a, 0x11,
	0x6b, 0x75, 0x62, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x66,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4b, 0x75, 0x62,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x66, 0x73, 0x52, 0x0f, 0x6b, 0x75,
	0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x66, 0x73, 0x1a, 0xf7, 0x01,
	0x0a, 0x12, 0x4b, 0x75, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x6a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x52, 0x0a, 0x0f, 0x4b, 0x75, 0x62, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x66, 0x73, 0x12, 0x3f, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x66, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x10,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x52, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xca, 0x03, 0x0a, 0x10,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x76, 0x0a, 0x15, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x42, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x52, 0x13, 0x6b, 0x75, 0x62, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x80, 0x01, 0x0a, 0x19, 0x6b, 0x75, 0x62,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x72, 0x65, 0x66, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4b, 0x75, 0x62,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x66, 0x73, 0x52, 0x16, 0x6b, 0x75, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x73, 0x1a, 0x51, 0x0a, 0x13, 0x4b,
	0x75, 0x62, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x68,
	0x0a, 0x16, 0x4b, 0x75, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x73, 0x12, 0x4e, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x50, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_goTypes = []interface{}{
	(*TrafficTargetSelector)(nil),                    // 0: networking.mesh.gloo.solo.io.TrafficTargetSelector
	(*WorkloadSelector)(nil),                         // 1: networking.mesh.gloo.solo.io.WorkloadSelector
	(*IdentitySelector)(nil),                         // 2: networking.mesh.gloo.solo.io.IdentitySelector
	(*TrafficTargetSelector_KubeServiceMatcher)(nil), // 3: networking.mesh.gloo.solo.io.TrafficTargetSelector.KubeServiceMatcher
	(*TrafficTargetSelector_KubeServiceRefs)(nil),    // 4: networking.mesh.gloo.solo.io.TrafficTargetSelector.KubeServiceRefs
	nil, // 5: networking.mesh.gloo.solo.io.TrafficTargetSelector.KubeServiceMatcher.LabelsEntry
	nil, // 6: networking.mesh.gloo.solo.io.WorkloadSelector.LabelsEntry
	(*IdentitySelector_KubeIdentityMatcher)(nil),    // 7: networking.mesh.gloo.solo.io.IdentitySelector.KubeIdentityMatcher
	(*IdentitySelector_KubeServiceAccountRefs)(nil), // 8: networking.mesh.gloo.solo.io.IdentitySelector.KubeServiceAccountRefs
	(*v1.ClusterObjectRef)(nil),                     // 9: core.skv2.solo.io.ClusterObjectRef
}
var file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_depIdxs = []int32{
	3, // 0: networking.mesh.gloo.solo.io.TrafficTargetSelector.kube_service_matcher:type_name -> networking.mesh.gloo.solo.io.TrafficTargetSelector.KubeServiceMatcher
	4, // 1: networking.mesh.gloo.solo.io.TrafficTargetSelector.kube_service_refs:type_name -> networking.mesh.gloo.solo.io.TrafficTargetSelector.KubeServiceRefs
	6, // 2: networking.mesh.gloo.solo.io.WorkloadSelector.labels:type_name -> networking.mesh.gloo.solo.io.WorkloadSelector.LabelsEntry
	7, // 3: networking.mesh.gloo.solo.io.IdentitySelector.kube_identity_matcher:type_name -> networking.mesh.gloo.solo.io.IdentitySelector.KubeIdentityMatcher
	8, // 4: networking.mesh.gloo.solo.io.IdentitySelector.kube_service_account_refs:type_name -> networking.mesh.gloo.solo.io.IdentitySelector.KubeServiceAccountRefs
	5, // 5: networking.mesh.gloo.solo.io.TrafficTargetSelector.KubeServiceMatcher.labels:type_name -> networking.mesh.gloo.solo.io.TrafficTargetSelector.KubeServiceMatcher.LabelsEntry
	9, // 6: networking.mesh.gloo.solo.io.TrafficTargetSelector.KubeServiceRefs.services:type_name -> core.skv2.solo.io.ClusterObjectRef
	9, // 7: networking.mesh.gloo.solo.io.IdentitySelector.KubeServiceAccountRefs.service_accounts:type_name -> core.skv2.solo.io.ClusterObjectRef
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficTargetSelector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadSelector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentitySelector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficTargetSelector_KubeServiceMatcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficTargetSelector_KubeServiceRefs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentitySelector_KubeIdentityMatcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentitySelector_KubeServiceAccountRefs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_selectors_proto_depIdxs = nil
}
