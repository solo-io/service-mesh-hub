// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/discovery/v1alpha2/mesh_service.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//*
//The MeshService is an abstraction for a service which we have discovered to be part of a
//given mesh. The Mesh object has references to the MeshServices which belong to it.
type MeshServiceSpec struct {
	// the type of service backing the MeshService
	//
	// Types that are valid to be assigned to Type:
	//	*MeshServiceSpec_KubeService_
	Type isMeshServiceSpec_Type `protobuf_oneof:"type"`
	// The mesh with which this service is associated.
	Mesh                 *v1.ObjectRef `protobuf:"bytes,2,opt,name=mesh,proto3" json:"mesh,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MeshServiceSpec) Reset()         { *m = MeshServiceSpec{} }
func (m *MeshServiceSpec) String() string { return proto.CompactTextString(m) }
func (*MeshServiceSpec) ProtoMessage()    {}
func (*MeshServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_10703f4489f049f7, []int{0}
}
func (m *MeshServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceSpec.Unmarshal(m, b)
}
func (m *MeshServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceSpec.Marshal(b, m, deterministic)
}
func (m *MeshServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceSpec.Merge(m, src)
}
func (m *MeshServiceSpec) XXX_Size() int {
	return xxx_messageInfo_MeshServiceSpec.Size(m)
}
func (m *MeshServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceSpec proto.InternalMessageInfo

type isMeshServiceSpec_Type interface {
	isMeshServiceSpec_Type()
	Equal(interface{}) bool
}

type MeshServiceSpec_KubeService_ struct {
	KubeService *MeshServiceSpec_KubeService `protobuf:"bytes,1,opt,name=kube_service,json=kubeService,proto3,oneof" json:"kube_service,omitempty"`
}

func (*MeshServiceSpec_KubeService_) isMeshServiceSpec_Type() {}

func (m *MeshServiceSpec) GetType() isMeshServiceSpec_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *MeshServiceSpec) GetKubeService() *MeshServiceSpec_KubeService {
	if x, ok := m.GetType().(*MeshServiceSpec_KubeService_); ok {
		return x.KubeService
	}
	return nil
}

func (m *MeshServiceSpec) GetMesh() *v1.ObjectRef {
	if m != nil {
		return m.Mesh
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MeshServiceSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MeshServiceSpec_KubeService_)(nil),
	}
}

type MeshServiceSpec_KubeService struct {
	// A reference to the kube-native service that this MeshService represents.
	Ref *v1.ClusterObjectRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// Selectors for the set of pods targeted by the k8s Service.
	WorkloadSelectorLabels map[string]string `protobuf:"bytes,2,rep,name=workload_selector_labels,json=workloadSelectorLabels,proto3" json:"workload_selector_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Labels on the underlying k8s Service itself.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The ports exposed by the underlying service.
	Ports []*MeshServiceSpec_KubeService_KubeServicePort `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty"`
	// Subsets for routing, based on labels.
	Subsets              map[string]*MeshServiceSpec_KubeService_Subset `protobuf:"bytes,5,rep,name=subsets,proto3" json:"subsets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *MeshServiceSpec_KubeService) Reset()         { *m = MeshServiceSpec_KubeService{} }
func (m *MeshServiceSpec_KubeService) String() string { return proto.CompactTextString(m) }
func (*MeshServiceSpec_KubeService) ProtoMessage()    {}
func (*MeshServiceSpec_KubeService) Descriptor() ([]byte, []int) {
	return fileDescriptor_10703f4489f049f7, []int{0, 0}
}
func (m *MeshServiceSpec_KubeService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceSpec_KubeService.Unmarshal(m, b)
}
func (m *MeshServiceSpec_KubeService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceSpec_KubeService.Marshal(b, m, deterministic)
}
func (m *MeshServiceSpec_KubeService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceSpec_KubeService.Merge(m, src)
}
func (m *MeshServiceSpec_KubeService) XXX_Size() int {
	return xxx_messageInfo_MeshServiceSpec_KubeService.Size(m)
}
func (m *MeshServiceSpec_KubeService) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceSpec_KubeService.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceSpec_KubeService proto.InternalMessageInfo

func (m *MeshServiceSpec_KubeService) GetRef() *v1.ClusterObjectRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *MeshServiceSpec_KubeService) GetWorkloadSelectorLabels() map[string]string {
	if m != nil {
		return m.WorkloadSelectorLabels
	}
	return nil
}

func (m *MeshServiceSpec_KubeService) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MeshServiceSpec_KubeService) GetPorts() []*MeshServiceSpec_KubeService_KubeServicePort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *MeshServiceSpec_KubeService) GetSubsets() map[string]*MeshServiceSpec_KubeService_Subset {
	if m != nil {
		return m.Subsets
	}
	return nil
}

type MeshServiceSpec_KubeService_KubeServicePort struct {
	// external-facing port for this service (i.e., NOT the service's target port on the backing pods)
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Protocol             string   `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshServiceSpec_KubeService_KubeServicePort) Reset() {
	*m = MeshServiceSpec_KubeService_KubeServicePort{}
}
func (m *MeshServiceSpec_KubeService_KubeServicePort) String() string {
	return proto.CompactTextString(m)
}
func (*MeshServiceSpec_KubeService_KubeServicePort) ProtoMessage() {}
func (*MeshServiceSpec_KubeService_KubeServicePort) Descriptor() ([]byte, []int) {
	return fileDescriptor_10703f4489f049f7, []int{0, 0, 0}
}
func (m *MeshServiceSpec_KubeService_KubeServicePort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceSpec_KubeService_KubeServicePort.Unmarshal(m, b)
}
func (m *MeshServiceSpec_KubeService_KubeServicePort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceSpec_KubeService_KubeServicePort.Marshal(b, m, deterministic)
}
func (m *MeshServiceSpec_KubeService_KubeServicePort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceSpec_KubeService_KubeServicePort.Merge(m, src)
}
func (m *MeshServiceSpec_KubeService_KubeServicePort) XXX_Size() int {
	return xxx_messageInfo_MeshServiceSpec_KubeService_KubeServicePort.Size(m)
}
func (m *MeshServiceSpec_KubeService_KubeServicePort) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceSpec_KubeService_KubeServicePort.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceSpec_KubeService_KubeServicePort proto.InternalMessageInfo

func (m *MeshServiceSpec_KubeService_KubeServicePort) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *MeshServiceSpec_KubeService_KubeServicePort) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MeshServiceSpec_KubeService_KubeServicePort) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type MeshServiceSpec_KubeService_Subset struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshServiceSpec_KubeService_Subset) Reset()         { *m = MeshServiceSpec_KubeService_Subset{} }
func (m *MeshServiceSpec_KubeService_Subset) String() string { return proto.CompactTextString(m) }
func (*MeshServiceSpec_KubeService_Subset) ProtoMessage()    {}
func (*MeshServiceSpec_KubeService_Subset) Descriptor() ([]byte, []int) {
	return fileDescriptor_10703f4489f049f7, []int{0, 0, 4}
}
func (m *MeshServiceSpec_KubeService_Subset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceSpec_KubeService_Subset.Unmarshal(m, b)
}
func (m *MeshServiceSpec_KubeService_Subset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceSpec_KubeService_Subset.Marshal(b, m, deterministic)
}
func (m *MeshServiceSpec_KubeService_Subset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceSpec_KubeService_Subset.Merge(m, src)
}
func (m *MeshServiceSpec_KubeService_Subset) XXX_Size() int {
	return xxx_messageInfo_MeshServiceSpec_KubeService_Subset.Size(m)
}
func (m *MeshServiceSpec_KubeService_Subset) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceSpec_KubeService_Subset.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceSpec_KubeService_Subset proto.InternalMessageInfo

func (m *MeshServiceSpec_KubeService_Subset) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type MeshServiceStatus struct {
	// The most recent generation observed in the the TrafficPolicy metadata.
	// if the observedGeneration does not match generation, the controller has not received the most
	// recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The set of Traffic Policies that have been applied to this MeshService
	AppliedTrafficPolicies []*MeshServiceStatus_AppliedTrafficPolicy `protobuf:"bytes,3,rep,name=applied_traffic_policies,json=appliedTrafficPolicies,proto3" json:"applied_traffic_policies,omitempty"`
	// The set of Access Policies that have been applied to this MeshService
	AppliedAccessPolicies []*MeshServiceStatus_AppliedAccessPolicy `protobuf:"bytes,4,rep,name=applied_access_policies,json=appliedAccessPolicies,proto3" json:"applied_access_policies,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                                 `json:"-"`
	XXX_unrecognized      []byte                                   `json:"-"`
	XXX_sizecache         int32                                    `json:"-"`
}

func (m *MeshServiceStatus) Reset()         { *m = MeshServiceStatus{} }
func (m *MeshServiceStatus) String() string { return proto.CompactTextString(m) }
func (*MeshServiceStatus) ProtoMessage()    {}
func (*MeshServiceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_10703f4489f049f7, []int{1}
}
func (m *MeshServiceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceStatus.Unmarshal(m, b)
}
func (m *MeshServiceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceStatus.Marshal(b, m, deterministic)
}
func (m *MeshServiceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceStatus.Merge(m, src)
}
func (m *MeshServiceStatus) XXX_Size() int {
	return xxx_messageInfo_MeshServiceStatus.Size(m)
}
func (m *MeshServiceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceStatus proto.InternalMessageInfo

func (m *MeshServiceStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *MeshServiceStatus) GetAppliedTrafficPolicies() []*MeshServiceStatus_AppliedTrafficPolicy {
	if m != nil {
		return m.AppliedTrafficPolicies
	}
	return nil
}

func (m *MeshServiceStatus) GetAppliedAccessPolicies() []*MeshServiceStatus_AppliedAccessPolicy {
	if m != nil {
		return m.AppliedAccessPolicies
	}
	return nil
}

// AppliedTrafficPolicy represents a traffic policy that has been applied to the MeshService.
// if an existing Traffic Policy becomes invalid, the last applied policy will be used
type MeshServiceStatus_AppliedTrafficPolicy struct {
	// reference to the traffic policy
	Ref *v1.ObjectRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// the observed generation of the accepted traffic policy
	ObservedGeneration int64 `protobuf:"varint,2,opt,name=observedGeneration,proto3" json:"observedGeneration,omitempty"`
	// the last known valid spec of the traffic policy
	Spec                 *v1alpha2.TrafficPolicySpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MeshServiceStatus_AppliedTrafficPolicy) Reset() {
	*m = MeshServiceStatus_AppliedTrafficPolicy{}
}
func (m *MeshServiceStatus_AppliedTrafficPolicy) String() string { return proto.CompactTextString(m) }
func (*MeshServiceStatus_AppliedTrafficPolicy) ProtoMessage()    {}
func (*MeshServiceStatus_AppliedTrafficPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_10703f4489f049f7, []int{1, 0}
}
func (m *MeshServiceStatus_AppliedTrafficPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceStatus_AppliedTrafficPolicy.Unmarshal(m, b)
}
func (m *MeshServiceStatus_AppliedTrafficPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceStatus_AppliedTrafficPolicy.Marshal(b, m, deterministic)
}
func (m *MeshServiceStatus_AppliedTrafficPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceStatus_AppliedTrafficPolicy.Merge(m, src)
}
func (m *MeshServiceStatus_AppliedTrafficPolicy) XXX_Size() int {
	return xxx_messageInfo_MeshServiceStatus_AppliedTrafficPolicy.Size(m)
}
func (m *MeshServiceStatus_AppliedTrafficPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceStatus_AppliedTrafficPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceStatus_AppliedTrafficPolicy proto.InternalMessageInfo

func (m *MeshServiceStatus_AppliedTrafficPolicy) GetRef() *v1.ObjectRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *MeshServiceStatus_AppliedTrafficPolicy) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *MeshServiceStatus_AppliedTrafficPolicy) GetSpec() *v1alpha2.TrafficPolicySpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// AppliedAccessPolicy represents a access policy that has been applied to the MeshService.
// if an existing Access Policy becomes invalid, the last applied policy will be used
type MeshServiceStatus_AppliedAccessPolicy struct {
	// reference to the access policy
	Ref *v1.ObjectRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// the observed generation of the accepted access policy
	ObservedGeneration int32 `protobuf:"varint,2,opt,name=observedGeneration,proto3" json:"observedGeneration,omitempty"`
	// the last known valid spec of the access policy
	Spec                 *v1alpha2.AccessPolicySpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MeshServiceStatus_AppliedAccessPolicy) Reset()         { *m = MeshServiceStatus_AppliedAccessPolicy{} }
func (m *MeshServiceStatus_AppliedAccessPolicy) String() string { return proto.CompactTextString(m) }
func (*MeshServiceStatus_AppliedAccessPolicy) ProtoMessage()    {}
func (*MeshServiceStatus_AppliedAccessPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_10703f4489f049f7, []int{1, 1}
}
func (m *MeshServiceStatus_AppliedAccessPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceStatus_AppliedAccessPolicy.Unmarshal(m, b)
}
func (m *MeshServiceStatus_AppliedAccessPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceStatus_AppliedAccessPolicy.Marshal(b, m, deterministic)
}
func (m *MeshServiceStatus_AppliedAccessPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceStatus_AppliedAccessPolicy.Merge(m, src)
}
func (m *MeshServiceStatus_AppliedAccessPolicy) XXX_Size() int {
	return xxx_messageInfo_MeshServiceStatus_AppliedAccessPolicy.Size(m)
}
func (m *MeshServiceStatus_AppliedAccessPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceStatus_AppliedAccessPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceStatus_AppliedAccessPolicy proto.InternalMessageInfo

func (m *MeshServiceStatus_AppliedAccessPolicy) GetRef() *v1.ObjectRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *MeshServiceStatus_AppliedAccessPolicy) GetObservedGeneration() int32 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *MeshServiceStatus_AppliedAccessPolicy) GetSpec() *v1alpha2.AccessPolicySpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// Federation policy applied to this MeshService, allowing access
// to the service from other meshes/clusters.
type MeshServiceStatus_AppliedFederation struct {
	//*
	//For any workload that this service has federated to (i.e., any MeshWorkload whose ref appears in `federated_to_workloads`),
	//a client in that workload will be able to reach this service at this DNS name. This includes workloads on clusters other than
	//the one hosting this service.
	MulticlusterDnsName string `protobuf:"bytes,1,opt,name=multicluster_dns_name,json=multiclusterDnsName,proto3" json:"multicluster_dns_name,omitempty"`
	// The list of Meshes which are able to resolve this service's `multicluster_dns_name`.
	FederatedToMeshes    []*v1.ObjectRef `protobuf:"bytes,2,rep,name=federated_to_meshes,json=federatedToMeshes,proto3" json:"federated_to_meshes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MeshServiceStatus_AppliedFederation) Reset()         { *m = MeshServiceStatus_AppliedFederation{} }
func (m *MeshServiceStatus_AppliedFederation) String() string { return proto.CompactTextString(m) }
func (*MeshServiceStatus_AppliedFederation) ProtoMessage()    {}
func (*MeshServiceStatus_AppliedFederation) Descriptor() ([]byte, []int) {
	return fileDescriptor_10703f4489f049f7, []int{1, 2}
}
func (m *MeshServiceStatus_AppliedFederation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceStatus_AppliedFederation.Unmarshal(m, b)
}
func (m *MeshServiceStatus_AppliedFederation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceStatus_AppliedFederation.Marshal(b, m, deterministic)
}
func (m *MeshServiceStatus_AppliedFederation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceStatus_AppliedFederation.Merge(m, src)
}
func (m *MeshServiceStatus_AppliedFederation) XXX_Size() int {
	return xxx_messageInfo_MeshServiceStatus_AppliedFederation.Size(m)
}
func (m *MeshServiceStatus_AppliedFederation) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceStatus_AppliedFederation.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceStatus_AppliedFederation proto.InternalMessageInfo

func (m *MeshServiceStatus_AppliedFederation) GetMulticlusterDnsName() string {
	if m != nil {
		return m.MulticlusterDnsName
	}
	return ""
}

func (m *MeshServiceStatus_AppliedFederation) GetFederatedToMeshes() []*v1.ObjectRef {
	if m != nil {
		return m.FederatedToMeshes
	}
	return nil
}

func init() {
	proto.RegisterType((*MeshServiceSpec)(nil), "discovery.smh.solo.io.MeshServiceSpec")
	proto.RegisterType((*MeshServiceSpec_KubeService)(nil), "discovery.smh.solo.io.MeshServiceSpec.KubeService")
	proto.RegisterMapType((map[string]string)(nil), "discovery.smh.solo.io.MeshServiceSpec.KubeService.LabelsEntry")
	proto.RegisterMapType((map[string]*MeshServiceSpec_KubeService_Subset)(nil), "discovery.smh.solo.io.MeshServiceSpec.KubeService.SubsetsEntry")
	proto.RegisterMapType((map[string]string)(nil), "discovery.smh.solo.io.MeshServiceSpec.KubeService.WorkloadSelectorLabelsEntry")
	proto.RegisterType((*MeshServiceSpec_KubeService_KubeServicePort)(nil), "discovery.smh.solo.io.MeshServiceSpec.KubeService.KubeServicePort")
	proto.RegisterType((*MeshServiceSpec_KubeService_Subset)(nil), "discovery.smh.solo.io.MeshServiceSpec.KubeService.Subset")
	proto.RegisterType((*MeshServiceStatus)(nil), "discovery.smh.solo.io.MeshServiceStatus")
	proto.RegisterType((*MeshServiceStatus_AppliedTrafficPolicy)(nil), "discovery.smh.solo.io.MeshServiceStatus.AppliedTrafficPolicy")
	proto.RegisterType((*MeshServiceStatus_AppliedAccessPolicy)(nil), "discovery.smh.solo.io.MeshServiceStatus.AppliedAccessPolicy")
	proto.RegisterType((*MeshServiceStatus_AppliedFederation)(nil), "discovery.smh.solo.io.MeshServiceStatus.AppliedFederation")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/discovery/v1alpha2/mesh_service.proto", fileDescriptor_10703f4489f049f7)
}

var fileDescriptor_10703f4489f049f7 = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x86, 0x4b, 0xeb, 0xe2, 0xfa, 0xc8, 0x85, 0xeb, 0x91, 0xed, 0x12, 0x6c, 0x51, 0x08, 0xee,
	0x46, 0x5d, 0x98, 0xac, 0x59, 0x14, 0xa8, 0x0b, 0xbb, 0x85, 0xdd, 0x7b, 0xe2, 0x1b, 0x28, 0x27,
	0x4e, 0xb2, 0x21, 0x48, 0xea, 0x48, 0x62, 0x44, 0x71, 0x08, 0xce, 0x50, 0x86, 0xde, 0x20, 0x2f,
	0x90, 0x77, 0xc8, 0x3e, 0xc8, 0x22, 0xaf, 0x91, 0x47, 0xc8, 0x5b, 0x64, 0x17, 0xcc, 0x90, 0x94,
	0x28, 0x9b, 0x31, 0x6c, 0x21, 0x2b, 0xcd, 0xed, 0x7c, 0xff, 0x3f, 0x47, 0x87, 0x87, 0x84, 0x93,
	0xbe, 0xcf, 0x07, 0x89, 0xab, 0x7b, 0x74, 0x64, 0x30, 0x1a, 0xd0, 0x1d, 0x9f, 0x1a, 0x0c, 0xe3,
	0xb1, 0xef, 0xe1, 0xce, 0x08, 0xd9, 0x60, 0x67, 0x90, 0xb8, 0x86, 0x13, 0xf9, 0x46, 0xd7, 0x67,
	0x1e, 0x1d, 0x63, 0x3c, 0x31, 0xc6, 0xbb, 0x4e, 0x10, 0x0d, 0x1c, 0xd3, 0x10, 0xfb, 0x76, 0x76,
	0x58, 0x8f, 0x62, 0xca, 0x29, 0xd9, 0x9c, 0x1e, 0xd3, 0xd9, 0x68, 0xa0, 0x0b, 0xa2, 0xee, 0x53,
	0x4d, 0x2f, 0x53, 0x19, 0x8e, 0x4d, 0x49, 0xf6, 0x68, 0x8c, 0xc6, 0x78, 0x57, 0xfe, 0xa6, 0x18,
	0xed, 0xa0, 0xd4, 0x42, 0x88, 0xfc, 0x8a, 0xc6, 0x43, 0x3f, 0xec, 0xcf, 0x3c, 0xf0, 0xd8, 0xe9,
	0xf5, 0x7c, 0xcf, 0x8e, 0x68, 0xe0, 0x7b, 0x93, 0x2c, 0x7c, 0xff, 0xce, 0xe1, 0x8e, 0xe7, 0x21,
	0x63, 0xf3, 0xd1, 0x1b, 0x7d, 0xda, 0xa7, 0x72, 0x68, 0x88, 0x51, 0xba, 0xba, 0xfd, 0x61, 0x19,
	0xd6, 0x4e, 0x90, 0x0d, 0x3a, 0x29, 0xba, 0x13, 0xa1, 0x47, 0x2e, 0x61, 0x75, 0x98, 0xb8, 0x98,
	0xe7, 0x40, 0x55, 0x5a, 0x4a, 0xbb, 0x61, 0x9a, 0x7a, 0x69, 0x12, 0xf4, 0x6b, 0xd1, 0xfa, 0xc3,
	0xc4, 0xc5, 0x6c, 0xfe, 0xdf, 0x17, 0x56, 0x63, 0x38, 0x9b, 0x92, 0x9f, 0xa0, 0x2a, 0xac, 0xab,
	0x4b, 0x12, 0xf8, 0x9d, 0x2e, 0x53, 0x23, 0x12, 0x36, 0x85, 0x9d, 0xb9, 0xcf, 0xd1, 0xe3, 0x16,
	0xf6, 0x2c, 0x79, 0x52, 0x7b, 0x57, 0x87, 0x46, 0x01, 0x48, 0x7e, 0x81, 0x4a, 0x8c, 0xbd, 0xcc,
	0xd1, 0x0f, 0x25, 0x80, 0x3f, 0x83, 0x84, 0x71, 0x8c, 0x67, 0x1c, 0x71, 0x9e, 0xbc, 0x50, 0x40,
	0x15, 0x49, 0x0a, 0xa8, 0xd3, 0xb5, 0x19, 0x06, 0xe8, 0x71, 0x1a, 0xdb, 0x81, 0xe3, 0x62, 0xc0,
	0xd4, 0xa5, 0x56, 0xa5, 0xdd, 0x30, 0x4f, 0xef, 0x7f, 0x3d, 0xfd, 0x32, 0x43, 0x76, 0x32, 0xe2,
	0xb1, 0x04, 0xfe, 0x1d, 0xf2, 0x78, 0x62, 0x6d, 0x5d, 0x95, 0x6e, 0x92, 0xc7, 0x50, 0xcf, 0x74,
	0x2b, 0x52, 0xf7, 0xf7, 0x05, 0x74, 0x8b, 0x3a, 0x19, 0x8d, 0x3c, 0x81, 0x5a, 0x44, 0x63, 0xce,
	0xd4, 0xaa, 0xc4, 0x1e, 0x2d, 0x80, 0x2d, 0x8c, 0xcf, 0x69, 0xcc, 0xad, 0x14, 0x48, 0x9e, 0xc2,
	0x32, 0x4b, 0x5c, 0x86, 0x9c, 0xa9, 0x35, 0xc9, 0xfe, 0x63, 0x01, 0x76, 0x27, 0x25, 0xa4, 0x9e,
	0x73, 0x9e, 0xf6, 0x08, 0xd6, 0xae, 0x89, 0x12, 0x02, 0x55, 0x21, 0x2b, 0xff, 0xe2, 0xaf, 0x2c,
	0x39, 0x16, 0x6b, 0xa1, 0x33, 0x42, 0x59, 0x37, 0x2b, 0x96, 0x1c, 0x13, 0x0d, 0xbe, 0x94, 0x15,
	0xec, 0xd1, 0x40, 0xad, 0xc8, 0xf5, 0xe9, 0x5c, 0xfb, 0x1f, 0xbe, 0xbd, 0xe5, 0xaf, 0x21, 0x5f,
	0x43, 0x65, 0x88, 0x13, 0xa9, 0xb0, 0x62, 0x89, 0x21, 0xd9, 0x80, 0xda, 0xd8, 0x09, 0x92, 0x5c,
	0x21, 0x9d, 0xfc, 0xb6, 0xf4, 0xab, 0xa2, 0xed, 0x41, 0x63, 0xd1, 0xd0, 0x04, 0x56, 0x8b, 0xb7,
	0x2e, 0x89, 0x3d, 0x2b, 0xc6, 0x36, 0xcc, 0xbd, 0x85, 0xf3, 0x5a, 0x94, 0x6d, 0x41, 0x3d, 0x5d,
	0x24, 0x5b, 0x50, 0x97, 0xcb, 0x4c, 0x55, 0x5a, 0x95, 0xf6, 0x8a, 0x95, 0xcd, 0x8e, 0xea, 0x50,
	0xe5, 0x93, 0x08, 0xb7, 0xdf, 0xd6, 0x61, 0xbd, 0xc8, 0xe6, 0x0e, 0x4f, 0x18, 0x31, 0xa0, 0x49,
	0x5d, 0xf1, 0xe8, 0x63, 0xd7, 0xee, 0x63, 0x88, 0xb1, 0xc3, 0x7d, 0x1a, 0x4a, 0xdb, 0x15, 0x8b,
	0xe4, 0x5b, 0xff, 0x4e, 0x77, 0xc8, 0x15, 0xa8, 0x4e, 0x14, 0x05, 0x3e, 0x76, 0xed, 0xb9, 0xb6,
	0xe5, 0x63, 0x5e, 0xe3, 0x07, 0x77, 0xb8, 0x98, 0x14, 0xd7, 0x0f, 0x53, 0xd0, 0x45, 0xca, 0x39,
	0x97, 0xfd, 0xcb, 0xda, 0x72, 0x6e, 0xae, 0xfa, 0xc8, 0x08, 0x87, 0x6f, 0x72, 0xe1, 0x62, 0xc3,
	0x13, 0xba, 0xe9, 0x43, 0xb0, 0x7f, 0x5f, 0xdd, 0x43, 0x89, 0xc9, 0x64, 0x37, 0x9d, 0x1b, 0x8b,
	0x3e, 0x32, 0xed, 0x8d, 0x02, 0x1b, 0x65, 0x36, 0x89, 0x5e, 0xec, 0x4d, 0xb7, 0x37, 0x37, 0xd9,
	0x94, 0x74, 0x28, 0xc9, 0xa6, 0x2c, 0x85, 0xf2, 0x3c, 0x1f, 0x40, 0x95, 0x45, 0xe8, 0xc9, 0x6a,
	0x6f, 0x98, 0x3f, 0xea, 0xb3, 0xc6, 0x3f, 0x77, 0xb9, 0x39, 0x53, 0xa2, 0x5e, 0x2c, 0x19, 0xa6,
	0xbd, 0x56, 0xa0, 0x59, 0x72, 0xcd, 0xcf, 0x68, 0xbb, 0x56, 0x6a, 0x7b, 0x7f, 0xce, 0x76, 0xfb,
	0x53, 0xb6, 0x8b, 0x9e, 0x0a, 0xae, 0x5f, 0x2a, 0xb0, 0x9e, 0xb9, 0xfe, 0x07, 0xbb, 0x39, 0xd3,
	0x84, 0xcd, 0x51, 0x12, 0x70, 0xdf, 0x4b, 0xbb, 0xbd, 0xdd, 0x0d, 0x99, 0x2d, 0x3b, 0x44, 0xfa,
	0x70, 0x35, 0x8b, 0x9b, 0x7f, 0x85, 0xec, 0x54, 0x34, 0x8c, 0x63, 0x68, 0xf6, 0x52, 0x82, 0x28,
	0x54, 0x6a, 0x8b, 0xf7, 0x0b, 0xe6, 0xdd, 0xff, 0xf6, 0x7b, 0xaf, 0x4f, 0x03, 0x2f, 0xe8, 0x89,
	0x0c, 0x3b, 0x3a, 0x7f, 0xf5, 0xfe, 0x7b, 0xe5, 0xd9, 0x83, 0xbb, 0x7c, 0x66, 0x44, 0xc3, 0xfe,
	0xfc, 0xa7, 0x46, 0xf1, 0xde, 0xd3, 0x77, 0xb6, 0x5b, 0x97, 0xed, 0xeb, 0xe7, 0x8f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd3, 0x3b, 0x18, 0x6f, 0xbb, 0x08, 0x00, 0x00,
}

func (this *MeshServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceSpec)
	if !ok {
		that2, ok := that.(MeshServiceSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Type == nil {
		if this.Type != nil {
			return false
		}
	} else if this.Type == nil {
		return false
	} else if !this.Type.Equal(that1.Type) {
		return false
	}
	if !this.Mesh.Equal(that1.Mesh) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceSpec_KubeService_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceSpec_KubeService_)
	if !ok {
		that2, ok := that.(MeshServiceSpec_KubeService_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.KubeService.Equal(that1.KubeService) {
		return false
	}
	return true
}
func (this *MeshServiceSpec_KubeService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceSpec_KubeService)
	if !ok {
		that2, ok := that.(MeshServiceSpec_KubeService)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if len(this.WorkloadSelectorLabels) != len(that1.WorkloadSelectorLabels) {
		return false
	}
	for i := range this.WorkloadSelectorLabels {
		if this.WorkloadSelectorLabels[i] != that1.WorkloadSelectorLabels[i] {
			return false
		}
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if len(this.Ports) != len(that1.Ports) {
		return false
	}
	for i := range this.Ports {
		if !this.Ports[i].Equal(that1.Ports[i]) {
			return false
		}
	}
	if len(this.Subsets) != len(that1.Subsets) {
		return false
	}
	for i := range this.Subsets {
		if !this.Subsets[i].Equal(that1.Subsets[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceSpec_KubeService_KubeServicePort) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceSpec_KubeService_KubeServicePort)
	if !ok {
		that2, ok := that.(MeshServiceSpec_KubeService_KubeServicePort)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Protocol != that1.Protocol {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceSpec_KubeService_Subset) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceSpec_KubeService_Subset)
	if !ok {
		that2, ok := that.(MeshServiceSpec_KubeService_Subset)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Values) != len(that1.Values) {
		return false
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceStatus)
	if !ok {
		that2, ok := that.(MeshServiceStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if len(this.AppliedTrafficPolicies) != len(that1.AppliedTrafficPolicies) {
		return false
	}
	for i := range this.AppliedTrafficPolicies {
		if !this.AppliedTrafficPolicies[i].Equal(that1.AppliedTrafficPolicies[i]) {
			return false
		}
	}
	if len(this.AppliedAccessPolicies) != len(that1.AppliedAccessPolicies) {
		return false
	}
	for i := range this.AppliedAccessPolicies {
		if !this.AppliedAccessPolicies[i].Equal(that1.AppliedAccessPolicies[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceStatus_AppliedTrafficPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceStatus_AppliedTrafficPolicy)
	if !ok {
		that2, ok := that.(MeshServiceStatus_AppliedTrafficPolicy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceStatus_AppliedAccessPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceStatus_AppliedAccessPolicy)
	if !ok {
		that2, ok := that.(MeshServiceStatus_AppliedAccessPolicy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceStatus_AppliedFederation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceStatus_AppliedFederation)
	if !ok {
		that2, ok := that.(MeshServiceStatus_AppliedFederation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MulticlusterDnsName != that1.MulticlusterDnsName {
		return false
	}
	if len(this.FederatedToMeshes) != len(that1.FederatedToMeshes) {
		return false
	}
	for i := range this.FederatedToMeshes {
		if !this.FederatedToMeshes[i].Equal(that1.FederatedToMeshes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
