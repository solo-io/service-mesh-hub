// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/networking/v1alpha2/access_policy.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2/types"
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

//
//Access control policies apply ALLOW policies to communication in a mesh.
//Access control policies specify the following:
//ALLOW those requests that: originate from from **source workload**, target the **destination target**,
//and match the indicated request criteria (allowed_paths, allowed_methods, allowed_ports).
//Enforcement of access control is determined by the
//[VirtualMesh's GlobalAccessPolicy]({{% versioned_link_path fromRoot="/reference/api/virtual_mesh/#networking.mesh.gloo.solo.io.VirtualMeshSpec.GlobalAccessPolicy" %}})
type AccessPolicySpec struct {
	//
	//Requests originating from these pods will have the rule applied.
	//Leave empty to have all pods in the mesh apply these policies.
	//
	//Note that access control policies are mapped to source pods by their
	//service account. If other pods share the same service account,
	//this access control rule will apply to those pods as well.
	//
	//For fine-grained access control policies, ensure that your
	//service accounts properly reflect the desired
	//boundary for your access control policies.
	SourceSelector []*IdentitySelector `protobuf:"bytes,2,rep,name=source_selector,json=sourceSelector,proto3" json:"source_selector,omitempty"`
	//
	//Requests destined for these pods will have the rule applied.
	//Leave empty to apply to all destination pods in the mesh.
	DestinationSelector []*TrafficTargetSelector `protobuf:"bytes,3,rep,name=destination_selector,json=destinationSelector,proto3" json:"destination_selector,omitempty"`
	//
	//Optional. A list of HTTP paths or gRPC methods to allow.
	//gRPC methods must be presented as fully-qualified name in the form of
	//"/packageName.serviceName/methodName" and are case sensitive.
	//Exact match, prefix match, and suffix match are supported for paths.
	//For example, the path "/books/review" matches
	//"/books/review" (exact match), "*books/" (suffix match), or "/books*" (prefix match).
	//
	//If not specified, allow any path.
	AllowedPaths []string `protobuf:"bytes,4,rep,name=allowed_paths,json=allowedPaths,proto3" json:"allowed_paths,omitempty"`
	//
	//Optional. A list of HTTP methods to allow (e.g., "GET", "POST").
	//It is ignored in gRPC case because the value is always "POST".
	//If not specified, allows any method.
	AllowedMethods []types.HttpMethodValue `protobuf:"varint,5,rep,packed,name=allowed_methods,json=allowedMethods,proto3,enum=networking.mesh.gloo.solo.io.HttpMethodValue" json:"allowed_methods,omitempty"`
	//
	//Optional. A list of ports which to allow.
	//If not set any port is allowed.
	AllowedPorts         []uint32 `protobuf:"varint,6,rep,packed,name=allowed_ports,json=allowedPorts,proto3" json:"allowed_ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicySpec) Reset()         { *m = AccessPolicySpec{} }
func (m *AccessPolicySpec) String() string { return proto.CompactTextString(m) }
func (*AccessPolicySpec) ProtoMessage()    {}
func (*AccessPolicySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_36a0dece7ff0ff65, []int{0}
}
func (m *AccessPolicySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicySpec.Unmarshal(m, b)
}
func (m *AccessPolicySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicySpec.Marshal(b, m, deterministic)
}
func (m *AccessPolicySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicySpec.Merge(m, src)
}
func (m *AccessPolicySpec) XXX_Size() int {
	return xxx_messageInfo_AccessPolicySpec.Size(m)
}
func (m *AccessPolicySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicySpec.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicySpec proto.InternalMessageInfo

func (m *AccessPolicySpec) GetSourceSelector() []*IdentitySelector {
	if m != nil {
		return m.SourceSelector
	}
	return nil
}

func (m *AccessPolicySpec) GetDestinationSelector() []*TrafficTargetSelector {
	if m != nil {
		return m.DestinationSelector
	}
	return nil
}

func (m *AccessPolicySpec) GetAllowedPaths() []string {
	if m != nil {
		return m.AllowedPaths
	}
	return nil
}

func (m *AccessPolicySpec) GetAllowedMethods() []types.HttpMethodValue {
	if m != nil {
		return m.AllowedMethods
	}
	return nil
}

func (m *AccessPolicySpec) GetAllowedPorts() []uint32 {
	if m != nil {
		return m.AllowedPorts
	}
	return nil
}

type AccessPolicyStatus struct {
	// The most recent generation observed in the the AccessPolicy metadata.
	// If the observedGeneration does not match generation, the controller has not received the most
	// recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The state of the overall resource.
	// It will only show accepted if it has been successfully
	// applied to all target meshes.
	State ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=networking.mesh.gloo.solo.io.ApprovalState" json:"state,omitempty"`
	// The status of the AccessPolicy for each TrafficTarget to which it has been applied.
	// An AccessPolicy may be Accepted for some TrafficTargets and rejected for others.
	TrafficTargets map[string]*ApprovalStatus `protobuf:"bytes,3,rep,name=traffic_targets,json=trafficTargets,proto3" json:"traffic_targets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The list of Workloads to which this policy has been applied.
	Workloads []string `protobuf:"bytes,4,rep,name=workloads,proto3" json:"workloads,omitempty"`
	// Any errors found while processing this generation of the resource.
	Errors               []string `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicyStatus) Reset()         { *m = AccessPolicyStatus{} }
func (m *AccessPolicyStatus) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyStatus) ProtoMessage()    {}
func (*AccessPolicyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_36a0dece7ff0ff65, []int{1}
}
func (m *AccessPolicyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyStatus.Unmarshal(m, b)
}
func (m *AccessPolicyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyStatus.Marshal(b, m, deterministic)
}
func (m *AccessPolicyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyStatus.Merge(m, src)
}
func (m *AccessPolicyStatus) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyStatus.Size(m)
}
func (m *AccessPolicyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyStatus proto.InternalMessageInfo

func (m *AccessPolicyStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *AccessPolicyStatus) GetState() ApprovalState {
	if m != nil {
		return m.State
	}
	return ApprovalState_PENDING
}

func (m *AccessPolicyStatus) GetTrafficTargets() map[string]*ApprovalStatus {
	if m != nil {
		return m.TrafficTargets
	}
	return nil
}

func (m *AccessPolicyStatus) GetWorkloads() []string {
	if m != nil {
		return m.Workloads
	}
	return nil
}

func (m *AccessPolicyStatus) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*AccessPolicySpec)(nil), "networking.mesh.gloo.solo.io.AccessPolicySpec")
	proto.RegisterType((*AccessPolicyStatus)(nil), "networking.mesh.gloo.solo.io.AccessPolicyStatus")
	proto.RegisterMapType((map[string]*ApprovalStatus)(nil), "networking.mesh.gloo.solo.io.AccessPolicyStatus.TrafficTargetsEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-mesh/api/networking/v1alpha2/access_policy.proto", fileDescriptor_36a0dece7ff0ff65)
}

var fileDescriptor_36a0dece7ff0ff65 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0x49, 0xbc, 0x16, 0xa2, 0xae, 0x69, 0x51, 0xca, 0x30, 0xa1, 0x0c, 0xd3, 0x5d, 0x0c,
	0x6b, 0x6c, 0x9a, 0x5e, 0xca, 0x6e, 0x29, 0x1b, 0xeb, 0x06, 0x83, 0xe0, 0x96, 0x0e, 0x76, 0x31,
	0x8a, 0xad, 0xd8, 0x22, 0x8a, 0x9f, 0x90, 0x5e, 0x52, 0xf2, 0x65, 0x76, 0xde, 0xe7, 0xda, 0x77,
	0xd8, 0x7d, 0x58, 0xb6, 0x93, 0x65, 0x1d, 0xc9, 0x6e, 0xf2, 0x5f, 0xfa, 0xff, 0xfe, 0x7a, 0xef,
	0x21, 0x93, 0xbb, 0x4c, 0x60, 0xbe, 0x98, 0x04, 0x09, 0xcc, 0x43, 0x03, 0x12, 0x06, 0x02, 0xc2,
	0x4c, 0x02, 0x0c, 0xe6, 0xdc, 0xe4, 0x21, 0x53, 0x22, 0x2c, 0x38, 0x3e, 0x81, 0x9e, 0x89, 0x22,
	0x0b, 0x97, 0x57, 0x4c, 0xaa, 0x9c, 0x0d, 0x43, 0x96, 0x24, 0xdc, 0x98, 0x58, 0x81, 0x14, 0xc9,
	0x2a, 0x50, 0x1a, 0x10, 0xe8, 0xf9, 0xe6, 0x60, 0x50, 0x9a, 0x83, 0x12, 0x13, 0x94, 0xcc, 0x40,
	0x40, 0xff, 0x72, 0x3f, 0x34, 0x47, 0x54, 0x15, 0xab, 0x7f, 0xb5, 0xff, 0xb4, 0xe1, 0x92, 0x27,
	0x08, 0xda, 0xd4, 0x96, 0x9b, 0xfd, 0x96, 0x25, 0x93, 0x22, 0x65, 0x28, 0xa0, 0x88, 0x0d, 0x32,
	0xe4, 0xb5, 0xf3, 0x2c, 0x83, 0x0c, 0xec, 0x32, 0x2c, 0x57, 0x95, 0x7a, 0xf1, 0xab, 0x4d, 0x4e,
	0x47, 0xb6, 0xcc, 0xb1, 0xad, 0xf2, 0x5e, 0xf1, 0x84, 0x7e, 0x25, 0x27, 0x06, 0x16, 0x3a, 0xe1,
	0x71, 0x13, 0xef, 0xb6, 0x3d, 0xc7, 0x3f, 0x1a, 0x06, 0xc1, 0xae, 0xea, 0x83, 0x4f, 0x29, 0x2f,
	0x50, 0xe0, 0xea, 0xbe, 0x76, 0x45, 0xdd, 0x0a, 0xd3, 0x7c, 0xd3, 0x29, 0x39, 0x4b, 0xb9, 0x41,
	0x51, 0xd4, 0xd7, 0x6b, 0xe8, 0x8e, 0xa5, 0x5f, 0xef, 0xa6, 0x3f, 0x68, 0x36, 0x9d, 0x8a, 0xe4,
	0x81, 0xe9, 0x8c, 0xe3, 0x3a, 0xa2, 0xf7, 0x07, 0x70, 0x9d, 0xf3, 0x86, 0x1c, 0x33, 0x29, 0xe1,
	0x89, 0xa7, 0xb1, 0x62, 0x98, 0x1b, 0xf7, 0x85, 0xe7, 0xf8, 0x9d, 0xe8, 0x65, 0x2d, 0x8e, 0x4b,
	0x8d, 0x3e, 0x92, 0x93, 0xe6, 0xd0, 0x9c, 0x63, 0x0e, 0xa9, 0x71, 0x0f, 0x3c, 0xc7, 0xef, 0x0e,
	0x07, 0xbb, 0xef, 0x71, 0x87, 0xa8, 0xbe, 0x58, 0xc3, 0x23, 0x93, 0x0b, 0x1e, 0x75, 0x6b, 0x4a,
	0xa5, 0x99, 0xad, 0x70, 0xd0, 0x68, 0xdc, 0x43, 0xcf, 0xf1, 0x8f, 0x37, 0xe1, 0xa5, 0x76, 0xf1,
	0xdd, 0x21, 0x74, 0xab, 0xef, 0xc8, 0x70, 0x61, 0x68, 0x48, 0x7a, 0x30, 0x31, 0x5c, 0x2f, 0x79,
	0x1a, 0x67, 0xbc, 0xe0, 0xda, 0xd6, 0xe5, 0xb6, 0xbc, 0x96, 0xef, 0x44, 0xb4, 0xd9, 0xfa, 0xb8,
	0xde, 0xa1, 0x23, 0x72, 0x60, 0x87, 0xec, 0xb6, 0xbd, 0x96, 0xdf, 0x1d, 0xbe, 0xdd, 0x7d, 0xf5,
	0x91, 0x52, 0x1a, 0x96, 0x4c, 0x96, 0x69, 0x3c, 0xaa, 0x9c, 0x74, 0x4e, 0x4e, 0xb0, 0x6a, 0x6d,
	0x8c, 0xb6, 0xb7, 0xa6, 0x9e, 0xc7, 0xfb, 0x3d, 0xb0, 0x67, 0xd7, 0xdf, 0x1e, 0x91, 0xf9, 0x50,
	0xa0, 0x5e, 0x45, 0x5d, 0xdc, 0x12, 0xe9, 0x39, 0xe9, 0x94, 0x4c, 0x09, 0x2c, 0x6d, 0xe6, 0xb2,
	0x11, 0xe8, 0x2b, 0x72, 0xc8, 0xb5, 0x06, 0x5d, 0xcd, 0xa2, 0x13, 0xd5, 0x5f, 0x7d, 0x20, 0xbd,
	0x7f, 0xc0, 0xe9, 0x29, 0x71, 0x66, 0x7c, 0x65, 0xfb, 0xd3, 0x89, 0xca, 0x25, 0xbd, 0x25, 0x07,
	0xcb, 0x72, 0x2c, 0xb6, 0x21, 0x47, 0xc3, 0xcb, 0xff, 0x6f, 0xc8, 0xc2, 0x44, 0x95, 0xf5, 0x5d,
	0xfb, 0xa6, 0x75, 0x3b, 0xfe, 0xf1, 0xf3, 0x75, 0xeb, 0xdb, 0xe7, 0x9d, 0xff, 0x0d, 0x35, 0xcb,
	0xfe, 0x7a, 0x85, 0xcf, 0x23, 0xd6, 0xef, 0x72, 0x72, 0x68, 0x5f, 0xdc, 0xf5, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd4, 0xa0, 0xdc, 0x16, 0x8c, 0x04, 0x00, 0x00,
}

func (this *AccessPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessPolicySpec)
	if !ok {
		that2, ok := that.(AccessPolicySpec)
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
	if len(this.SourceSelector) != len(that1.SourceSelector) {
		return false
	}
	for i := range this.SourceSelector {
		if !this.SourceSelector[i].Equal(that1.SourceSelector[i]) {
			return false
		}
	}
	if len(this.DestinationSelector) != len(that1.DestinationSelector) {
		return false
	}
	for i := range this.DestinationSelector {
		if !this.DestinationSelector[i].Equal(that1.DestinationSelector[i]) {
			return false
		}
	}
	if len(this.AllowedPaths) != len(that1.AllowedPaths) {
		return false
	}
	for i := range this.AllowedPaths {
		if this.AllowedPaths[i] != that1.AllowedPaths[i] {
			return false
		}
	}
	if len(this.AllowedMethods) != len(that1.AllowedMethods) {
		return false
	}
	for i := range this.AllowedMethods {
		if this.AllowedMethods[i] != that1.AllowedMethods[i] {
			return false
		}
	}
	if len(this.AllowedPorts) != len(that1.AllowedPorts) {
		return false
	}
	for i := range this.AllowedPorts {
		if this.AllowedPorts[i] != that1.AllowedPorts[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AccessPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessPolicyStatus)
	if !ok {
		that2, ok := that.(AccessPolicyStatus)
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
	if this.State != that1.State {
		return false
	}
	if len(this.TrafficTargets) != len(that1.TrafficTargets) {
		return false
	}
	for i := range this.TrafficTargets {
		if !this.TrafficTargets[i].Equal(that1.TrafficTargets[i]) {
			return false
		}
	}
	if len(this.Workloads) != len(that1.Workloads) {
		return false
	}
	for i := range this.Workloads {
		if this.Workloads[i] != that1.Workloads[i] {
			return false
		}
	}
	if len(this.Errors) != len(that1.Errors) {
		return false
	}
	for i := range this.Errors {
		if this.Errors[i] != that1.Errors[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
