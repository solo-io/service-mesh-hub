// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/extensions/v1alpha1/networking_extensions.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v1alpha3 "istio.io/api/networking/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

// Request patches to the resources created for a TrafficTarget
type TrafficTargetPatchRequest struct {
	// the object for which we are requesting patches
	TrafficTarget *TrafficTargetResource `protobuf:"bytes,1,opt,name=traffic_target,json=trafficTarget,proto3" json:"traffic_target,omitempty"`
	// the pre-patch resources generated by SMH for this object for a given discovery resource.
	// provided to the server to perform modifications.
	GeneratedResources   []*GeneratedResource `protobuf:"bytes,2,rep,name=generated_resources,json=generatedResources,proto3" json:"generated_resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TrafficTargetPatchRequest) Reset()         { *m = TrafficTargetPatchRequest{} }
func (m *TrafficTargetPatchRequest) String() string { return proto.CompactTextString(m) }
func (*TrafficTargetPatchRequest) ProtoMessage()    {}
func (*TrafficTargetPatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad002ba715800198, []int{0}
}
func (m *TrafficTargetPatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficTargetPatchRequest.Unmarshal(m, b)
}
func (m *TrafficTargetPatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficTargetPatchRequest.Marshal(b, m, deterministic)
}
func (m *TrafficTargetPatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficTargetPatchRequest.Merge(m, src)
}
func (m *TrafficTargetPatchRequest) XXX_Size() int {
	return xxx_messageInfo_TrafficTargetPatchRequest.Size(m)
}
func (m *TrafficTargetPatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficTargetPatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficTargetPatchRequest proto.InternalMessageInfo

func (m *TrafficTargetPatchRequest) GetTrafficTarget() *TrafficTargetResource {
	if m != nil {
		return m.TrafficTarget
	}
	return nil
}

func (m *TrafficTargetPatchRequest) GetGeneratedResources() []*GeneratedResource {
	if m != nil {
		return m.GeneratedResources
	}
	return nil
}

// a proto-serializable representation of a TrafficTarget object
type TrafficTargetResource struct {
	// metadata of the resource
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// the spec of the object
	Spec *v1alpha2.TrafficTargetSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// the status of the object
	Status               *v1alpha2.TrafficTargetStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *TrafficTargetResource) Reset()         { *m = TrafficTargetResource{} }
func (m *TrafficTargetResource) String() string { return proto.CompactTextString(m) }
func (*TrafficTargetResource) ProtoMessage()    {}
func (*TrafficTargetResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad002ba715800198, []int{1}
}
func (m *TrafficTargetResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficTargetResource.Unmarshal(m, b)
}
func (m *TrafficTargetResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficTargetResource.Marshal(b, m, deterministic)
}
func (m *TrafficTargetResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficTargetResource.Merge(m, src)
}
func (m *TrafficTargetResource) XXX_Size() int {
	return xxx_messageInfo_TrafficTargetResource.Size(m)
}
func (m *TrafficTargetResource) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficTargetResource.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficTargetResource proto.InternalMessageInfo

func (m *TrafficTargetResource) GetMetadata() *v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TrafficTargetResource) GetSpec() *v1alpha2.TrafficTargetSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *TrafficTargetResource) GetStatus() *v1alpha2.TrafficTargetStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Request patches to the resources created for a Workload
type WorkloadPatchRequest struct {
	// the object for which we are requesting patches
	Workload *WorkloadResource `protobuf:"bytes,1,opt,name=workload,proto3" json:"workload,omitempty"`
	// the pre-patch resources generated by SMH for this object for a given discovery resource.
	// provided to the server to perform modifications.
	GeneratedResources   []*GeneratedResource `protobuf:"bytes,2,rep,name=generated_resources,json=generatedResources,proto3" json:"generated_resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WorkloadPatchRequest) Reset()         { *m = WorkloadPatchRequest{} }
func (m *WorkloadPatchRequest) String() string { return proto.CompactTextString(m) }
func (*WorkloadPatchRequest) ProtoMessage()    {}
func (*WorkloadPatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad002ba715800198, []int{2}
}
func (m *WorkloadPatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadPatchRequest.Unmarshal(m, b)
}
func (m *WorkloadPatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadPatchRequest.Marshal(b, m, deterministic)
}
func (m *WorkloadPatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadPatchRequest.Merge(m, src)
}
func (m *WorkloadPatchRequest) XXX_Size() int {
	return xxx_messageInfo_WorkloadPatchRequest.Size(m)
}
func (m *WorkloadPatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadPatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadPatchRequest proto.InternalMessageInfo

func (m *WorkloadPatchRequest) GetWorkload() *WorkloadResource {
	if m != nil {
		return m.Workload
	}
	return nil
}

func (m *WorkloadPatchRequest) GetGeneratedResources() []*GeneratedResource {
	if m != nil {
		return m.GeneratedResources
	}
	return nil
}

// a proto-serializable representation of a Workload object
type WorkloadResource struct {
	// metadata of the resource
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// the spec of the object
	Spec *v1alpha2.WorkloadSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// the status of the object
	Status               *v1alpha2.WorkloadStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *WorkloadResource) Reset()         { *m = WorkloadResource{} }
func (m *WorkloadResource) String() string { return proto.CompactTextString(m) }
func (*WorkloadResource) ProtoMessage()    {}
func (*WorkloadResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad002ba715800198, []int{3}
}
func (m *WorkloadResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadResource.Unmarshal(m, b)
}
func (m *WorkloadResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadResource.Marshal(b, m, deterministic)
}
func (m *WorkloadResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadResource.Merge(m, src)
}
func (m *WorkloadResource) XXX_Size() int {
	return xxx_messageInfo_WorkloadResource.Size(m)
}
func (m *WorkloadResource) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadResource.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadResource proto.InternalMessageInfo

func (m *WorkloadResource) GetMetadata() *v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *WorkloadResource) GetSpec() *v1alpha2.WorkloadSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *WorkloadResource) GetStatus() *v1alpha2.WorkloadStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Request patches to the resources created for a Mesh
type MeshPatchRequest struct {
	// the object for which we are requesting patches
	Mesh *MeshResource `protobuf:"bytes,1,opt,name=mesh,proto3" json:"mesh,omitempty"`
	// the pre-patch resources generated by SMH for this object for a given discovery resource.
	// provided to the server to perform modifications.
	GeneratedResources   []*GeneratedResource `protobuf:"bytes,2,rep,name=generated_resources,json=generatedResources,proto3" json:"generated_resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MeshPatchRequest) Reset()         { *m = MeshPatchRequest{} }
func (m *MeshPatchRequest) String() string { return proto.CompactTextString(m) }
func (*MeshPatchRequest) ProtoMessage()    {}
func (*MeshPatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad002ba715800198, []int{4}
}
func (m *MeshPatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshPatchRequest.Unmarshal(m, b)
}
func (m *MeshPatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshPatchRequest.Marshal(b, m, deterministic)
}
func (m *MeshPatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshPatchRequest.Merge(m, src)
}
func (m *MeshPatchRequest) XXX_Size() int {
	return xxx_messageInfo_MeshPatchRequest.Size(m)
}
func (m *MeshPatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshPatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MeshPatchRequest proto.InternalMessageInfo

func (m *MeshPatchRequest) GetMesh() *MeshResource {
	if m != nil {
		return m.Mesh
	}
	return nil
}

func (m *MeshPatchRequest) GetGeneratedResources() []*GeneratedResource {
	if m != nil {
		return m.GeneratedResources
	}
	return nil
}

// a proto-serializable representation of a Mesh object
type MeshResource struct {
	// metadata of the resource
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// the spec of the object
	Spec *v1alpha2.MeshSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// the status of the object
	Status               *v1alpha2.MeshStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MeshResource) Reset()         { *m = MeshResource{} }
func (m *MeshResource) String() string { return proto.CompactTextString(m) }
func (*MeshResource) ProtoMessage()    {}
func (*MeshResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad002ba715800198, []int{5}
}
func (m *MeshResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshResource.Unmarshal(m, b)
}
func (m *MeshResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshResource.Marshal(b, m, deterministic)
}
func (m *MeshResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshResource.Merge(m, src)
}
func (m *MeshResource) XXX_Size() int {
	return xxx_messageInfo_MeshResource.Size(m)
}
func (m *MeshResource) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshResource.DiscardUnknown(m)
}

var xxx_messageInfo_MeshResource proto.InternalMessageInfo

func (m *MeshResource) GetMetadata() *v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MeshResource) GetSpec() *v1alpha2.MeshSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *MeshResource) GetStatus() *v1alpha2.MeshStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// a generated resource can be of any output type supported by SMH.
// the content of the type field should be used to determine
// the type of the output resource.
// TODO(ilackarms): consider parameterizing SMH to allow excluding GeneratedResources from patch requests in the case where an implementer only performs additions (no updates required).
type GeneratedResource struct {
	// metadata of the resource
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// type of the object.
	//
	// Types that are valid to be assigned to Type:
	//	*GeneratedResource_DestinationRule
	//	*GeneratedResource_EnvoyFilter
	//	*GeneratedResource_ServiceEntry
	//	*GeneratedResource_VirtualService
	Type                 isGeneratedResource_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GeneratedResource) Reset()         { *m = GeneratedResource{} }
func (m *GeneratedResource) String() string { return proto.CompactTextString(m) }
func (*GeneratedResource) ProtoMessage()    {}
func (*GeneratedResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad002ba715800198, []int{6}
}
func (m *GeneratedResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneratedResource.Unmarshal(m, b)
}
func (m *GeneratedResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneratedResource.Marshal(b, m, deterministic)
}
func (m *GeneratedResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneratedResource.Merge(m, src)
}
func (m *GeneratedResource) XXX_Size() int {
	return xxx_messageInfo_GeneratedResource.Size(m)
}
func (m *GeneratedResource) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneratedResource.DiscardUnknown(m)
}

var xxx_messageInfo_GeneratedResource proto.InternalMessageInfo

type isGeneratedResource_Type interface {
	isGeneratedResource_Type()
}

type GeneratedResource_DestinationRule struct {
	DestinationRule *v1alpha3.DestinationRule `protobuf:"bytes,2,opt,name=destination_rule,json=destinationRule,proto3,oneof" json:"destination_rule,omitempty"`
}
type GeneratedResource_EnvoyFilter struct {
	EnvoyFilter *v1alpha3.EnvoyFilter `protobuf:"bytes,3,opt,name=envoy_filter,json=envoyFilter,proto3,oneof" json:"envoy_filter,omitempty"`
}
type GeneratedResource_ServiceEntry struct {
	ServiceEntry *v1alpha3.ServiceEntry `protobuf:"bytes,4,opt,name=service_entry,json=serviceEntry,proto3,oneof" json:"service_entry,omitempty"`
}
type GeneratedResource_VirtualService struct {
	VirtualService *v1alpha3.VirtualService `protobuf:"bytes,5,opt,name=virtual_service,json=virtualService,proto3,oneof" json:"virtual_service,omitempty"`
}

func (*GeneratedResource_DestinationRule) isGeneratedResource_Type() {}
func (*GeneratedResource_EnvoyFilter) isGeneratedResource_Type()     {}
func (*GeneratedResource_ServiceEntry) isGeneratedResource_Type()    {}
func (*GeneratedResource_VirtualService) isGeneratedResource_Type()  {}

func (m *GeneratedResource) GetType() isGeneratedResource_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *GeneratedResource) GetMetadata() *v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *GeneratedResource) GetDestinationRule() *v1alpha3.DestinationRule {
	if x, ok := m.GetType().(*GeneratedResource_DestinationRule); ok {
		return x.DestinationRule
	}
	return nil
}

func (m *GeneratedResource) GetEnvoyFilter() *v1alpha3.EnvoyFilter {
	if x, ok := m.GetType().(*GeneratedResource_EnvoyFilter); ok {
		return x.EnvoyFilter
	}
	return nil
}

func (m *GeneratedResource) GetServiceEntry() *v1alpha3.ServiceEntry {
	if x, ok := m.GetType().(*GeneratedResource_ServiceEntry); ok {
		return x.ServiceEntry
	}
	return nil
}

func (m *GeneratedResource) GetVirtualService() *v1alpha3.VirtualService {
	if x, ok := m.GetType().(*GeneratedResource_VirtualService); ok {
		return x.VirtualService
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GeneratedResource) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GeneratedResource_DestinationRule)(nil),
		(*GeneratedResource_EnvoyFilter)(nil),
		(*GeneratedResource_ServiceEntry)(nil),
		(*GeneratedResource_VirtualService)(nil),
	}
}

// a list of patches to be returned for modifying & adding to the output SMH snapshot
type PatchList struct {
	// a list of generated resources to be inserted into to the snapshot.
	PatchedResources     []*GeneratedResource `protobuf:"bytes,1,rep,name=patched_resources,json=patchedResources,proto3" json:"patched_resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PatchList) Reset()         { *m = PatchList{} }
func (m *PatchList) String() string { return proto.CompactTextString(m) }
func (*PatchList) ProtoMessage()    {}
func (*PatchList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad002ba715800198, []int{7}
}
func (m *PatchList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchList.Unmarshal(m, b)
}
func (m *PatchList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchList.Marshal(b, m, deterministic)
}
func (m *PatchList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchList.Merge(m, src)
}
func (m *PatchList) XXX_Size() int {
	return xxx_messageInfo_PatchList.Size(m)
}
func (m *PatchList) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchList.DiscardUnknown(m)
}

var xxx_messageInfo_PatchList proto.InternalMessageInfo

func (m *PatchList) GetPatchedResources() []*GeneratedResource {
	if m != nil {
		return m.PatchedResources
	}
	return nil
}

// request to initiate push notifications
type WatchPushNotificationsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchPushNotificationsRequest) Reset()         { *m = WatchPushNotificationsRequest{} }
func (m *WatchPushNotificationsRequest) String() string { return proto.CompactTextString(m) }
func (*WatchPushNotificationsRequest) ProtoMessage()    {}
func (*WatchPushNotificationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad002ba715800198, []int{8}
}
func (m *WatchPushNotificationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchPushNotificationsRequest.Unmarshal(m, b)
}
func (m *WatchPushNotificationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchPushNotificationsRequest.Marshal(b, m, deterministic)
}
func (m *WatchPushNotificationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchPushNotificationsRequest.Merge(m, src)
}
func (m *WatchPushNotificationsRequest) XXX_Size() int {
	return xxx_messageInfo_WatchPushNotificationsRequest.Size(m)
}
func (m *WatchPushNotificationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchPushNotificationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchPushNotificationsRequest proto.InternalMessageInfo

// triggers a resync of SMH resources
type PushNotification struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotification) Reset()         { *m = PushNotification{} }
func (m *PushNotification) String() string { return proto.CompactTextString(m) }
func (*PushNotification) ProtoMessage()    {}
func (*PushNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad002ba715800198, []int{9}
}
func (m *PushNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotification.Unmarshal(m, b)
}
func (m *PushNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotification.Marshal(b, m, deterministic)
}
func (m *PushNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotification.Merge(m, src)
}
func (m *PushNotification) XXX_Size() int {
	return xxx_messageInfo_PushNotification.Size(m)
}
func (m *PushNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotification.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotification proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TrafficTargetPatchRequest)(nil), "extensions.networking.smh.solo.io.TrafficTargetPatchRequest")
	proto.RegisterType((*TrafficTargetResource)(nil), "extensions.networking.smh.solo.io.TrafficTargetResource")
	proto.RegisterType((*WorkloadPatchRequest)(nil), "extensions.networking.smh.solo.io.WorkloadPatchRequest")
	proto.RegisterType((*WorkloadResource)(nil), "extensions.networking.smh.solo.io.WorkloadResource")
	proto.RegisterType((*MeshPatchRequest)(nil), "extensions.networking.smh.solo.io.MeshPatchRequest")
	proto.RegisterType((*MeshResource)(nil), "extensions.networking.smh.solo.io.MeshResource")
	proto.RegisterType((*GeneratedResource)(nil), "extensions.networking.smh.solo.io.GeneratedResource")
	proto.RegisterType((*PatchList)(nil), "extensions.networking.smh.solo.io.PatchList")
	proto.RegisterType((*WatchPushNotificationsRequest)(nil), "extensions.networking.smh.solo.io.WatchPushNotificationsRequest")
	proto.RegisterType((*PushNotification)(nil), "extensions.networking.smh.solo.io.PushNotification")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/networking/extensions/v1alpha1/networking_extensions.proto", fileDescriptor_ad002ba715800198)
}

var fileDescriptor_ad002ba715800198 = []byte{
	// 883 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x41, 0x8f, 0xdb, 0x44,
	0x14, 0x8e, 0xbb, 0x61, 0x55, 0x66, 0xb7, 0xdb, 0x74, 0x28, 0x10, 0x56, 0x42, 0xed, 0x06, 0xd1,
	0x6e, 0x57, 0xec, 0xb8, 0x9b, 0x54, 0xb4, 0x48, 0x8b, 0x84, 0x16, 0x96, 0x44, 0xa2, 0xdd, 0x56,
	0xee, 0xaa, 0x8b, 0xb8, 0x58, 0x13, 0xe7, 0xc5, 0x1e, 0x92, 0x78, 0x8c, 0x67, 0xec, 0x92, 0x03,
	0x67, 0x7e, 0x00, 0x7f, 0x83, 0x2b, 0x57, 0x7e, 0x01, 0xe2, 0xcc, 0x9d, 0x03, 0x7f, 0x03, 0x8d,
	0x33, 0x76, 0x6c, 0xd7, 0x9b, 0xba, 0xda, 0x55, 0x8f, 0x7e, 0xfe, 0xbe, 0xef, 0xcd, 0x7b, 0xdf,
	0xcb, 0xbc, 0x18, 0x0d, 0x5d, 0x26, 0xbd, 0x68, 0x48, 0x1c, 0x3e, 0x33, 0x05, 0x9f, 0xf2, 0x7d,
	0xc6, 0x4d, 0x01, 0x61, 0xcc, 0x1c, 0xd8, 0x9f, 0x81, 0xf0, 0xf6, 0xbd, 0x68, 0x68, 0xd2, 0x80,
	0x99, 0x3e, 0xc8, 0x97, 0x3c, 0x9c, 0x30, 0xdf, 0x35, 0xe1, 0x67, 0x09, 0xbe, 0x60, 0xdc, 0x17,
	0x66, 0x7c, 0x40, 0xa7, 0x81, 0x47, 0x0f, 0x72, 0xaf, 0xed, 0xe5, 0x6b, 0x12, 0x84, 0x5c, 0x72,
	0xbc, 0x93, 0x8b, 0x2c, 0x71, 0x44, 0xcc, 0x3c, 0xa2, 0x52, 0x12, 0xc6, 0xb7, 0xbb, 0x95, 0x39,
	0x47, 0x4c, 0x38, 0x3c, 0x86, 0x70, 0x9e, 0xe6, 0xe9, 0x9a, 0xea, 0xfd, 0x42, 0x76, 0xfb, 0xb0,
	0x2e, 0x47, 0x86, 0x74, 0x3c, 0x66, 0x8e, 0x2d, 0x69, 0xe8, 0x82, 0xd4, 0xec, 0xcf, 0xeb, 0xb2,
	0xd5, 0x71, 0xa7, 0x9c, 0x8e, 0x34, 0xef, 0xc1, 0xe4, 0x91, 0x20, 0x8c, 0x2b, 0xf4, 0x8c, 0x3a,
	0x1e, 0xf3, 0x15, 0x38, 0x98, 0xb8, 0x2a, 0x20, 0xcc, 0x19, 0x48, 0x6a, 0xc6, 0x07, 0xa6, 0x0b,
	0x3e, 0x84, 0x54, 0x42, 0xca, 0xda, 0xcb, 0xb5, 0x4f, 0x2b, 0xf7, 0xcc, 0x11, 0x08, 0xc9, 0x7c,
	0x2a, 0x19, 0xf7, 0xed, 0x30, 0x9a, 0x82, 0xc6, 0xde, 0xa9, 0xc2, 0x82, 0x1f, 0xf3, 0xb9, 0x3d,
	0x66, 0x53, 0x09, 0xa1, 0xc6, 0xed, 0x54, 0xe1, 0x5c, 0x2a, 0xe1, 0x25, 0x9d, 0x6b, 0xc8, 0xdd,
	0x2a, 0x88, 0x2e, 0xdc, 0x06, 0x5f, 0x86, 0xf3, 0x55, 0x5a, 0x82, 0x8d, 0xc0, 0xa1, 0x69, 0xba,
	0x7b, 0x55, 0x90, 0x98, 0x85, 0x32, 0xa2, 0x53, 0x5b, 0x6b, 0x6a, 0xe8, 0x6e, 0x15, 0x34, 0xed,
	0x63, 0x3e, 0x6f, 0xe7, 0x5f, 0x03, 0x7d, 0x74, 0xba, 0xb0, 0xe7, 0x34, 0x71, 0xe7, 0x19, 0x95,
	0x8e, 0x67, 0xc1, 0x4f, 0x11, 0x08, 0x89, 0x6d, 0xb4, 0x55, 0xf4, 0xae, 0x6d, 0xdc, 0x36, 0x76,
	0x37, 0xba, 0x8f, 0xc8, 0x6b, 0x27, 0x8a, 0x14, 0x54, 0x2d, 0x10, 0x3c, 0x0a, 0x1d, 0xb0, 0xae,
	0xc9, 0x7c, 0x18, 0x03, 0x7a, 0x2f, 0x73, 0xca, 0x0e, 0x35, 0x48, 0xb4, 0xaf, 0xdc, 0x5e, 0xdb,
	0xdd, 0xe8, 0x3e, 0xa8, 0x91, 0xa5, 0x9f, 0xb2, 0xb3, 0x0c, 0xd8, 0x2d, 0x87, 0x44, 0xe7, 0x3f,
	0x03, 0xbd, 0x5f, 0x79, 0x1e, 0xfc, 0x18, 0x5d, 0x55, 0x23, 0x33, 0xa2, 0x92, 0xea, 0xda, 0xee,
	0x93, 0xc5, 0x80, 0x91, 0xfc, 0x80, 0x91, 0x60, 0xe2, 0xaa, 0x80, 0x20, 0x0a, 0x4d, 0xe2, 0x03,
	0xf2, 0x74, 0xf8, 0x23, 0x38, 0xf2, 0x09, 0x48, 0x6a, 0x65, 0x0a, 0xf8, 0x10, 0x35, 0x45, 0x00,
	0x4e, 0xfb, 0x4a, 0xa2, 0xb4, 0x4b, 0xb2, 0x69, 0x3e, 0xbf, 0x33, 0xcf, 0x03, 0x70, 0xac, 0x84,
	0x85, 0x8f, 0xd0, 0xba, 0x90, 0x54, 0x46, 0xa2, 0xbd, 0x96, 0xf0, 0xf7, 0x6a, 0xf1, 0x13, 0x86,
	0xa5, 0x99, 0x9d, 0xbf, 0x0d, 0x74, 0xf3, 0x4c, 0x1b, 0x5d, 0xb0, 0xf2, 0x29, 0xba, 0x9a, 0x0e,
	0x80, 0x2e, 0xb4, 0x57, 0xa3, 0xbd, 0xa9, 0x54, 0xd6, 0xdd, 0x4c, 0xe4, 0x6d, 0x59, 0xf7, 0x8f,
	0x81, 0x5a, 0xe5, 0x53, 0x5c, 0xb2, 0x6b, 0x0f, 0x0b, 0xae, 0x7d, 0x72, 0x4e, 0xd7, 0xd3, 0x43,
	0xe4, 0x0c, 0xfb, 0xb2, 0x64, 0xd8, 0xa7, 0xaf, 0xa3, 0x16, 0xbd, 0xfa, 0xd3, 0x40, 0xad, 0x27,
	0x20, 0xbc, 0x82, 0x4f, 0x5f, 0xa3, 0xa6, 0xba, 0x10, 0x75, 0x59, 0x66, 0x8d, 0x3e, 0x2a, 0x89,
	0xac, 0x85, 0x09, 0xf9, 0x6d, 0x79, 0xf3, 0x97, 0x81, 0x36, 0xf3, 0xd9, 0x2f, 0xd9, 0x97, 0x5e,
	0xc1, 0x97, 0x5b, 0xe7, 0x34, 0x57, 0x1d, 0x20, 0xe7, 0xc9, 0x17, 0x25, 0x4f, 0x76, 0x56, 0xd1,
	0x8a, 0x7e, 0xfc, 0xb1, 0x86, 0x6e, 0xbc, 0x52, 0xf8, 0x25, 0xd7, 0x74, 0x86, 0x5a, 0xe5, 0xad,
	0xa3, 0xeb, 0xdb, 0x23, 0x4c, 0x48, 0xc6, 0xf3, 0x8e, 0xa4, 0x57, 0x37, 0xf9, 0x66, 0x49, 0xb1,
	0xa2, 0x29, 0x0c, 0x1a, 0xd6, 0xf5, 0x51, 0x31, 0x84, 0xbf, 0x43, 0x9b, 0xf9, 0x15, 0xa5, 0xab,
	0xbf, 0xb3, 0x42, 0xf4, 0x58, 0xc1, 0xbf, 0x4d, 0xd0, 0x83, 0x86, 0xb5, 0x01, 0xcb, 0x47, 0x7c,
	0x82, 0xae, 0x15, 0x96, 0x54, 0xbb, 0x99, 0xa8, 0xdd, 0x5d, 0xa1, 0xf6, 0x7c, 0x81, 0x3f, 0x56,
	0xf0, 0x41, 0xc3, 0xda, 0x14, 0xb9, 0x67, 0x7c, 0x8a, 0xae, 0x97, 0x16, 0x55, 0xfb, 0x9d, 0x44,
	0xf1, 0xde, 0x0a, 0xc5, 0x17, 0x0b, 0x86, 0x16, 0x1e, 0x34, 0xac, 0xad, 0xb8, 0x10, 0x39, 0x5a,
	0x47, 0x4d, 0x39, 0x0f, 0xa0, 0xe3, 0xa3, 0x77, 0x93, 0x9f, 0xd0, 0x63, 0x26, 0x24, 0xa6, 0xe8,
	0x46, 0xa0, 0x1e, 0x0a, 0x83, 0x6f, 0x5c, 0x60, 0xf0, 0x5b, 0x5a, 0x6e, 0x39, 0xf6, 0xb7, 0xd0,
	0xc7, 0x67, 0x2a, 0xf6, 0x2c, 0x12, 0xde, 0x09, 0x97, 0x6c, 0xcc, 0x9c, 0xc4, 0x07, 0xa1, 0x7f,
	0xc3, 0x1d, 0x8c, 0x5a, 0xe5, 0x77, 0xdd, 0xdf, 0x9b, 0xe8, 0xe6, 0x49, 0x96, 0xf3, 0x38, 0x3b,
	0x08, 0xfe, 0xd5, 0x40, 0x1f, 0xf6, 0x41, 0xbe, 0xba, 0x84, 0x41, 0xe0, 0xc3, 0x37, 0xdd, 0xb3,
	0xf9, 0xab, 0x64, 0xfb, 0xb3, 0x1a, 0xec, 0xac, 0x71, 0x9d, 0x06, 0xfe, 0x05, 0xe1, 0x3e, 0xc8,
	0xc2, 0xf6, 0x00, 0x81, 0x1f, 0xbe, 0xc1, 0x9a, 0xb8, 0x50, 0xfa, 0x08, 0x6d, 0xf5, 0x41, 0x66,
	0x17, 0x22, 0x08, 0xdc, 0xab, 0x79, 0xfb, 0x5d, 0x28, 0xed, 0x6f, 0x06, 0xfa, 0xa0, 0xda, 0x4e,
	0xfc, 0x55, 0x9d, 0xd2, 0x57, 0x4d, 0xc2, 0x76, 0x9d, 0x0a, 0xca, 0xe4, 0x4e, 0xe3, 0xbe, 0x71,
	0xf4, 0xfd, 0x0f, 0x2f, 0xea, 0x7c, 0x18, 0xe8, 0xff, 0xbd, 0x66, 0xb5, 0x74, 0xd5, 0x87, 0xc2,
	0x70, 0x3d, 0xf9, 0xe3, 0xd7, 0xfb, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x48, 0xc2, 0xed, 0x79,
	0x0c, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkingExtensionsClient is the client API for NetworkingExtensions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkingExtensionsClient interface {
	// GetTrafficTargetPatches retrieves desired patches to the set of output service mesh config resources
	// generated for a given SMH TrafficTarget
	GetTrafficTargetPatches(ctx context.Context, in *TrafficTargetPatchRequest, opts ...grpc.CallOption) (*PatchList, error)
	// GetWorkloadPatches retrieves desired patches to the set of output Mesh config resources
	// generated for a given SMH Workload
	GetWorkloadPatches(ctx context.Context, in *WorkloadPatchRequest, opts ...grpc.CallOption) (*PatchList, error)
	// GetMeshPatches retrieves desired patches to the set of output Mesh config resources
	// generated for a given SMH Mesh
	GetMeshPatches(ctx context.Context, in *MeshPatchRequest, opts ...grpc.CallOption) (*PatchList, error)
	// WatchPushNotifications initiates a streaming connection which allows the NetworkingExtensions server
	// to push notifications to Service Mesh Hub telling it to resync its configuration.
	// This allows a NetworkingExtensions server to trigger SMH to resync its state for
	// events triggered by resources not watched by SMH.
	WatchPushNotifications(ctx context.Context, in *WatchPushNotificationsRequest, opts ...grpc.CallOption) (NetworkingExtensions_WatchPushNotificationsClient, error)
}

type networkingExtensionsClient struct {
	cc *grpc.ClientConn
}

func NewNetworkingExtensionsClient(cc *grpc.ClientConn) NetworkingExtensionsClient {
	return &networkingExtensionsClient{cc}
}

func (c *networkingExtensionsClient) GetTrafficTargetPatches(ctx context.Context, in *TrafficTargetPatchRequest, opts ...grpc.CallOption) (*PatchList, error) {
	out := new(PatchList)
	err := c.cc.Invoke(ctx, "/extensions.networking.smh.solo.io.NetworkingExtensions/GetTrafficTargetPatches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkingExtensionsClient) GetWorkloadPatches(ctx context.Context, in *WorkloadPatchRequest, opts ...grpc.CallOption) (*PatchList, error) {
	out := new(PatchList)
	err := c.cc.Invoke(ctx, "/extensions.networking.smh.solo.io.NetworkingExtensions/GetWorkloadPatches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkingExtensionsClient) GetMeshPatches(ctx context.Context, in *MeshPatchRequest, opts ...grpc.CallOption) (*PatchList, error) {
	out := new(PatchList)
	err := c.cc.Invoke(ctx, "/extensions.networking.smh.solo.io.NetworkingExtensions/GetMeshPatches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkingExtensionsClient) WatchPushNotifications(ctx context.Context, in *WatchPushNotificationsRequest, opts ...grpc.CallOption) (NetworkingExtensions_WatchPushNotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetworkingExtensions_serviceDesc.Streams[0], "/extensions.networking.smh.solo.io.NetworkingExtensions/WatchPushNotifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkingExtensionsWatchPushNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetworkingExtensions_WatchPushNotificationsClient interface {
	Recv() (*PushNotification, error)
	grpc.ClientStream
}

type networkingExtensionsWatchPushNotificationsClient struct {
	grpc.ClientStream
}

func (x *networkingExtensionsWatchPushNotificationsClient) Recv() (*PushNotification, error) {
	m := new(PushNotification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NetworkingExtensionsServer is the server API for NetworkingExtensions service.
type NetworkingExtensionsServer interface {
	// GetTrafficTargetPatches retrieves desired patches to the set of output service mesh config resources
	// generated for a given SMH TrafficTarget
	GetTrafficTargetPatches(context.Context, *TrafficTargetPatchRequest) (*PatchList, error)
	// GetWorkloadPatches retrieves desired patches to the set of output Mesh config resources
	// generated for a given SMH Workload
	GetWorkloadPatches(context.Context, *WorkloadPatchRequest) (*PatchList, error)
	// GetMeshPatches retrieves desired patches to the set of output Mesh config resources
	// generated for a given SMH Mesh
	GetMeshPatches(context.Context, *MeshPatchRequest) (*PatchList, error)
	// WatchPushNotifications initiates a streaming connection which allows the NetworkingExtensions server
	// to push notifications to Service Mesh Hub telling it to resync its configuration.
	// This allows a NetworkingExtensions server to trigger SMH to resync its state for
	// events triggered by resources not watched by SMH.
	WatchPushNotifications(*WatchPushNotificationsRequest, NetworkingExtensions_WatchPushNotificationsServer) error
}

// UnimplementedNetworkingExtensionsServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkingExtensionsServer struct {
}

func (*UnimplementedNetworkingExtensionsServer) GetTrafficTargetPatches(ctx context.Context, req *TrafficTargetPatchRequest) (*PatchList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrafficTargetPatches not implemented")
}
func (*UnimplementedNetworkingExtensionsServer) GetWorkloadPatches(ctx context.Context, req *WorkloadPatchRequest) (*PatchList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkloadPatches not implemented")
}
func (*UnimplementedNetworkingExtensionsServer) GetMeshPatches(ctx context.Context, req *MeshPatchRequest) (*PatchList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeshPatches not implemented")
}
func (*UnimplementedNetworkingExtensionsServer) WatchPushNotifications(req *WatchPushNotificationsRequest, srv NetworkingExtensions_WatchPushNotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchPushNotifications not implemented")
}

func RegisterNetworkingExtensionsServer(s *grpc.Server, srv NetworkingExtensionsServer) {
	s.RegisterService(&_NetworkingExtensions_serviceDesc, srv)
}

func _NetworkingExtensions_GetTrafficTargetPatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrafficTargetPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkingExtensionsServer).GetTrafficTargetPatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extensions.networking.smh.solo.io.NetworkingExtensions/GetTrafficTargetPatches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkingExtensionsServer).GetTrafficTargetPatches(ctx, req.(*TrafficTargetPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkingExtensions_GetWorkloadPatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkloadPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkingExtensionsServer).GetWorkloadPatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extensions.networking.smh.solo.io.NetworkingExtensions/GetWorkloadPatches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkingExtensionsServer).GetWorkloadPatches(ctx, req.(*WorkloadPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkingExtensions_GetMeshPatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeshPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkingExtensionsServer).GetMeshPatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extensions.networking.smh.solo.io.NetworkingExtensions/GetMeshPatches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkingExtensionsServer).GetMeshPatches(ctx, req.(*MeshPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkingExtensions_WatchPushNotifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchPushNotificationsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetworkingExtensionsServer).WatchPushNotifications(m, &networkingExtensionsWatchPushNotificationsServer{stream})
}

type NetworkingExtensions_WatchPushNotificationsServer interface {
	Send(*PushNotification) error
	grpc.ServerStream
}

type networkingExtensionsWatchPushNotificationsServer struct {
	grpc.ServerStream
}

func (x *networkingExtensionsWatchPushNotificationsServer) Send(m *PushNotification) error {
	return x.ServerStream.SendMsg(m)
}

var _NetworkingExtensions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "extensions.networking.smh.solo.io.NetworkingExtensions",
	HandlerType: (*NetworkingExtensionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTrafficTargetPatches",
			Handler:    _NetworkingExtensions_GetTrafficTargetPatches_Handler,
		},
		{
			MethodName: "GetWorkloadPatches",
			Handler:    _NetworkingExtensions_GetWorkloadPatches_Handler,
		},
		{
			MethodName: "GetMeshPatches",
			Handler:    _NetworkingExtensions_GetMeshPatches_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchPushNotifications",
			Handler:       _NetworkingExtensions_WatchPushNotifications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/solo-io/service-mesh-hub/api/networking/extensions/v1alpha1/networking_extensions.proto",
}
