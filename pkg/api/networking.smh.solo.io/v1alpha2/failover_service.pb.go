// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/failover_service.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	math "math"
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
//A FailoverService creates a new hostname to which services can send requests.
//Requests will be routed based on a list of backing services ordered by
//decreasing priority. When outlier detection detects that a service in the list is
//in an unhealthy state, requests sent to the FailoverService will be routed
//to the next healthy service in the list. For each service referenced in the
//failover services list, outlier detection must be configured using a TrafficPolicy.
//
//Currently this feature only supports Services backed by Istio.
type FailoverServiceSpec struct {
	//
	//The DNS name of the failover service. Must be unique within the service mesh instance
	//since it is used as the hostname with which clients communicate.
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// The port on which the failover service listens.
	Port *FailoverServiceSpec_Port `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	// The meshes that this failover service will be visible to.
	Meshes []*v1.ObjectRef `protobuf:"bytes,3,rep,name=meshes,proto3" json:"meshes,omitempty"`
	//
	//A list of services ordered by decreasing priority for failover.
	//All services must be backed by either the same service mesh instance or
	//backed by service meshes that are grouped under a common VirtualMesh.
	FailoverServices     []*v1.ClusterObjectRef `protobuf:"bytes,4,rep,name=failover_services,json=failoverServices,proto3" json:"failover_services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FailoverServiceSpec) Reset()         { *m = FailoverServiceSpec{} }
func (m *FailoverServiceSpec) String() string { return proto.CompactTextString(m) }
func (*FailoverServiceSpec) ProtoMessage()    {}
func (*FailoverServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2a3822bd950167, []int{0}
}
func (m *FailoverServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverServiceSpec.Unmarshal(m, b)
}
func (m *FailoverServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverServiceSpec.Marshal(b, m, deterministic)
}
func (m *FailoverServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverServiceSpec.Merge(m, src)
}
func (m *FailoverServiceSpec) XXX_Size() int {
	return xxx_messageInfo_FailoverServiceSpec.Size(m)
}
func (m *FailoverServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverServiceSpec proto.InternalMessageInfo

func (m *FailoverServiceSpec) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *FailoverServiceSpec) GetPort() *FailoverServiceSpec_Port {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *FailoverServiceSpec) GetMeshes() []*v1.ObjectRef {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func (m *FailoverServiceSpec) GetFailoverServices() []*v1.ClusterObjectRef {
	if m != nil {
		return m.FailoverServices
	}
	return nil
}

// The port on which the failover service listens.
type FailoverServiceSpec_Port struct {
	// Port number.
	Port uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// Protocol of the requests sent to the failover service, must be one of HTTP, HTTPS, GRPC, HTTP2, MONGO, TCP, TLS.
	Protocol             string   `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FailoverServiceSpec_Port) Reset()         { *m = FailoverServiceSpec_Port{} }
func (m *FailoverServiceSpec_Port) String() string { return proto.CompactTextString(m) }
func (*FailoverServiceSpec_Port) ProtoMessage()    {}
func (*FailoverServiceSpec_Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2a3822bd950167, []int{0, 0}
}
func (m *FailoverServiceSpec_Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverServiceSpec_Port.Unmarshal(m, b)
}
func (m *FailoverServiceSpec_Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverServiceSpec_Port.Marshal(b, m, deterministic)
}
func (m *FailoverServiceSpec_Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverServiceSpec_Port.Merge(m, src)
}
func (m *FailoverServiceSpec_Port) XXX_Size() int {
	return xxx_messageInfo_FailoverServiceSpec_Port.Size(m)
}
func (m *FailoverServiceSpec_Port) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverServiceSpec_Port.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverServiceSpec_Port proto.InternalMessageInfo

func (m *FailoverServiceSpec_Port) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *FailoverServiceSpec_Port) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type FailoverServiceStatus struct {
	//
	//The most recent generation observed in the the FailoverService metadata.
	//If the observedGeneration does not match generation, the controller has not received the most
	//recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	//
	//The state of the overall resource, will only show accepted if it has been successfully
	//applied to all target meshes.
	State ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=networking.smh.solo.io.ApprovalState" json:"state,omitempty"`
	// The status of the FailoverService for each Mesh to which it has been applied.
	Meshes map[string]*ApprovalStatus `protobuf:"bytes,3,rep,name=meshes,proto3" json:"meshes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// any errors observed which prevented the resource from being Accepted.
	ValidationErrors     []string `protobuf:"bytes,4,rep,name=validation_errors,json=validationErrors,proto3" json:"validation_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FailoverServiceStatus) Reset()         { *m = FailoverServiceStatus{} }
func (m *FailoverServiceStatus) String() string { return proto.CompactTextString(m) }
func (*FailoverServiceStatus) ProtoMessage()    {}
func (*FailoverServiceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2a3822bd950167, []int{1}
}
func (m *FailoverServiceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverServiceStatus.Unmarshal(m, b)
}
func (m *FailoverServiceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverServiceStatus.Marshal(b, m, deterministic)
}
func (m *FailoverServiceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverServiceStatus.Merge(m, src)
}
func (m *FailoverServiceStatus) XXX_Size() int {
	return xxx_messageInfo_FailoverServiceStatus.Size(m)
}
func (m *FailoverServiceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverServiceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverServiceStatus proto.InternalMessageInfo

func (m *FailoverServiceStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *FailoverServiceStatus) GetState() ApprovalState {
	if m != nil {
		return m.State
	}
	return ApprovalState_PENDING
}

func (m *FailoverServiceStatus) GetMeshes() map[string]*ApprovalStatus {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func (m *FailoverServiceStatus) GetValidationErrors() []string {
	if m != nil {
		return m.ValidationErrors
	}
	return nil
}

func init() {
	proto.RegisterType((*FailoverServiceSpec)(nil), "networking.smh.solo.io.FailoverServiceSpec")
	proto.RegisterType((*FailoverServiceSpec_Port)(nil), "networking.smh.solo.io.FailoverServiceSpec.Port")
	proto.RegisterType((*FailoverServiceStatus)(nil), "networking.smh.solo.io.FailoverServiceStatus")
	proto.RegisterMapType((map[string]*ApprovalStatus)(nil), "networking.smh.solo.io.FailoverServiceStatus.MeshesEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/failover_service.proto", fileDescriptor_4c2a3822bd950167)
}

var fileDescriptor_4c2a3822bd950167 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x26, 0x6d, 0x77, 0xd9, 0x4e, 0x51, 0xba, 0xb3, 0x2a, 0x25, 0x88, 0x94, 0x15, 0xa5, 0x20,
	0x9d, 0xb8, 0x51, 0xc4, 0x3f, 0x10, 0x7f, 0x56, 0x2f, 0x44, 0xdc, 0x9d, 0xbd, 0xf3, 0xa6, 0x4c,
	0xb2, 0xa7, 0xc9, 0xd8, 0x34, 0x27, 0xcc, 0x4c, 0x22, 0xfb, 0x46, 0xbe, 0x81, 0xef, 0xe3, 0x95,
	0x8f, 0x21, 0x33, 0xc9, 0xa6, 0xb5, 0x54, 0xe8, 0x55, 0x26, 0x73, 0xe6, 0x7c, 0xe7, 0xfb, 0xbe,
	0x73, 0x0e, 0xe1, 0x89, 0x34, 0x69, 0x19, 0xb1, 0x18, 0x97, 0x81, 0xc6, 0x0c, 0xa7, 0x12, 0x03,
	0x0d, 0xaa, 0x92, 0x31, 0x4c, 0x97, 0xa0, 0xd3, 0x69, 0x5a, 0x46, 0x81, 0x28, 0x64, 0x90, 0x83,
	0xf9, 0x81, 0x6a, 0x21, 0xf3, 0x24, 0xa8, 0x4e, 0x44, 0x56, 0xa4, 0x22, 0x0c, 0xe6, 0x42, 0x66,
	0x58, 0x81, 0x9a, 0x35, 0x19, 0xac, 0x50, 0x68, 0x90, 0xde, 0x59, 0xbd, 0x65, 0x7a, 0x99, 0x32,
	0x8b, 0xcb, 0x24, 0xfa, 0x6c, 0x5b, 0xad, 0x45, 0x15, 0x3a, 0xfc, 0x18, 0x15, 0x04, 0xd5, 0x89,
	0xfb, 0xd6, 0x38, 0xfe, 0x9b, 0x9d, 0x89, 0x54, 0x22, 0x93, 0x97, 0xc2, 0x48, 0xcc, 0x67, 0xda,
	0x08, 0x73, 0x0d, 0x70, 0x2b, 0xc1, 0x04, 0xdd, 0x31, 0xb0, 0xa7, 0xfa, 0xf6, 0xf8, 0x57, 0x87,
	0x1c, 0x7d, 0x6c, 0x98, 0x5f, 0xd4, 0x15, 0x2e, 0x0a, 0x88, 0xa9, 0x4f, 0x0e, 0x52, 0xd4, 0x26,
	0x17, 0x4b, 0x18, 0x79, 0x63, 0x6f, 0xd2, 0xe7, 0xed, 0x3f, 0xfd, 0x40, 0x7a, 0x05, 0x2a, 0x33,
	0xea, 0x8c, 0xbd, 0xc9, 0x20, 0x7c, 0xcc, 0xb6, 0x2b, 0x64, 0x5b, 0x60, 0xd9, 0x19, 0x2a, 0xc3,
	0x5d, 0x36, 0x7d, 0x4a, 0xf6, 0xad, 0x14, 0xd0, 0xa3, 0xee, 0xb8, 0x3b, 0x19, 0x84, 0x77, 0x99,
	0x53, 0x6b, 0x3d, 0x68, 0x21, 0xbe, 0x46, 0xdf, 0x21, 0x36, 0x1c, 0xe6, 0xbc, 0x79, 0x4b, 0xcf,
	0xc8, 0xe1, 0xa6, 0xd1, 0x7a, 0xd4, 0x73, 0x00, 0xf7, 0xb7, 0x00, 0xbc, 0xcf, 0x4a, 0x6d, 0x40,
	0xad, 0x70, 0x86, 0xf3, 0x7f, 0x59, 0x69, 0xff, 0x19, 0xe9, 0x59, 0x56, 0x94, 0x36, 0xaa, 0xac,
	0xda, 0x1b, 0x0d, 0x47, 0x9f, 0x1c, 0x38, 0x9b, 0x62, 0xcc, 0x9c, 0xda, 0x3e, 0x6f, 0xff, 0x8f,
	0xff, 0x74, 0xc8, 0xed, 0x4d, 0x89, 0x46, 0x98, 0x52, 0xd3, 0x80, 0x1c, 0x61, 0x64, 0xc9, 0xc1,
	0xe5, 0x2c, 0x81, 0x1c, 0x94, 0x6b, 0x86, 0x03, 0xee, 0x72, 0x7a, 0x1d, 0xfa, 0xd4, 0x46, 0xe8,
	0x2b, 0xb2, 0xe7, 0x3a, 0xe5, 0x6a, 0xdc, 0x0c, 0x1f, 0xfc, 0xcf, 0xd1, 0xb7, 0x45, 0xa1, 0xb0,
	0x12, 0x99, 0xad, 0x03, 0xbc, 0xce, 0xa1, 0xe7, 0x1b, 0x3e, 0xbe, 0xd8, 0xb5, 0x1f, 0x8e, 0x2c,
	0xfb, 0xe2, 0x72, 0x4f, 0x73, 0xa3, 0xae, 0x5a, 0x93, 0x1f, 0x91, 0xc3, 0xb5, 0x21, 0x02, 0xa5,
	0x50, 0xd5, 0x26, 0xf7, 0xf9, 0x70, 0x15, 0x38, 0x75, 0xf7, 0xbe, 0x20, 0x83, 0x35, 0x0c, 0x3a,
	0x24, 0xdd, 0x05, 0x5c, 0x35, 0x33, 0x63, 0x8f, 0xf4, 0x35, 0xd9, 0xab, 0x44, 0x56, 0x42, 0x33,
	0x2f, 0x0f, 0x77, 0x51, 0x57, 0x6a, 0x5e, 0x27, 0xbd, 0xec, 0x3c, 0xf7, 0xde, 0x9d, 0xff, 0xfc,
	0x7d, 0xcf, 0xfb, 0xf6, 0x79, 0x97, 0xed, 0x2c, 0x16, 0xc9, 0xc6, 0x62, 0xac, 0xd7, 0x68, 0x97,
	0x24, 0xda, 0x77, 0x7d, 0x7c, 0xf2, 0x37, 0x00, 0x00, 0xff, 0xff, 0x69, 0xb2, 0xa3, 0x2f, 0xf3,
	0x03, 0x00, 0x00,
}

func (this *FailoverServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceSpec)
	if !ok {
		that2, ok := that.(FailoverServiceSpec)
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
	if this.Hostname != that1.Hostname {
		return false
	}
	if !this.Port.Equal(that1.Port) {
		return false
	}
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if len(this.FailoverServices) != len(that1.FailoverServices) {
		return false
	}
	for i := range this.FailoverServices {
		if !this.FailoverServices[i].Equal(that1.FailoverServices[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailoverServiceSpec_Port) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceSpec_Port)
	if !ok {
		that2, ok := that.(FailoverServiceSpec_Port)
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
	if this.Protocol != that1.Protocol {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailoverServiceStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceStatus)
	if !ok {
		that2, ok := that.(FailoverServiceStatus)
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
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if len(this.ValidationErrors) != len(that1.ValidationErrors) {
		return false
	}
	for i := range this.ValidationErrors {
		if this.ValidationErrors[i] != that1.ValidationErrors[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
