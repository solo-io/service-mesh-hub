// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh.proto

package v1 // import "github.com/solo-io/supergloo/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// TODO: Eventually we want to plug in new meshes easier, but for now it's useful to enumerate in the config
type MeshType int32

const (
	MeshType_ISTIO    MeshType = 0
	MeshType_LINKERD1 MeshType = 1
)

var MeshType_name = map[int32]string{
	0: "ISTIO",
	1: "LINKERD1",
}
var MeshType_value = map[string]int32{
	"ISTIO":    0,
	"LINKERD1": 1,
}

func (x MeshType) String() string {
	return proto.EnumName(MeshType_name, int32(x))
}
func (MeshType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mesh_ae9c8f46e151dfa7, []int{0}
}

//
// @solo-kit:resource.short_name=mesh
// @solo-kit:resource.plural_name=meshes
// @solo-kit:resource.resource_groups=translator.supergloo.solo.io
//
//
// // Initial API thoughts based on potential user workflows: https://github.com/solo-io/supergloo/wiki/User-Workflows
// // Notes on what Itsio and Linkerd1 support on wiki: https://github.com/solo-io/supergloo/wiki
//
// // Any user-configurable settings for a service mesh
// // This isn't meant to cover install or initial configuration of the mesh
// // The supergloo server would apply CRUD on MeshConfigs
type Mesh struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata  `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
	TargetMesh           *TargetMesh    `protobuf:"bytes,2,opt,name=target_mesh,json=targetMesh" json:"target_mesh,omitempty"`
	Routing              *Routing       `protobuf:"bytes,1,opt,name=routing" json:"routing,omitempty"`
	Encryption           *Encryption    `protobuf:"bytes,98,opt,name=encryption" json:"encryption,omitempty"`
	Observability        *Observability `protobuf:"bytes,99,opt,name=observability" json:"observability,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Mesh) Reset()         { *m = Mesh{} }
func (m *Mesh) String() string { return proto.CompactTextString(m) }
func (*Mesh) ProtoMessage()    {}
func (*Mesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_ae9c8f46e151dfa7, []int{0}
}
func (m *Mesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mesh.Unmarshal(m, b)
}
func (m *Mesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mesh.Marshal(b, m, deterministic)
}
func (dst *Mesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mesh.Merge(dst, src)
}
func (m *Mesh) XXX_Size() int {
	return xxx_messageInfo_Mesh.Size(m)
}
func (m *Mesh) XXX_DiscardUnknown() {
	xxx_messageInfo_Mesh.DiscardUnknown(m)
}

var xxx_messageInfo_Mesh proto.InternalMessageInfo

func (m *Mesh) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Mesh) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Mesh) GetTargetMesh() *TargetMesh {
	if m != nil {
		return m.TargetMesh
	}
	return nil
}

func (m *Mesh) GetRouting() *Routing {
	if m != nil {
		return m.Routing
	}
	return nil
}

func (m *Mesh) GetEncryption() *Encryption {
	if m != nil {
		return m.Encryption
	}
	return nil
}

func (m *Mesh) GetObservability() *Observability {
	if m != nil {
		return m.Observability
	}
	return nil
}

type TargetMesh struct {
	MeshType MeshType `protobuf:"varint,1,opt,name=mesh_type,json=meshType,proto3,enum=supergloo.solo.io.MeshType" json:"mesh_type,omitempty"`
	// Types that are valid to be assigned to DeploymentLocation:
	//	*TargetMesh_KubeLocation
	DeploymentLocation   isTargetMesh_DeploymentLocation `protobuf_oneof:"deployment_location"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *TargetMesh) Reset()         { *m = TargetMesh{} }
func (m *TargetMesh) String() string { return proto.CompactTextString(m) }
func (*TargetMesh) ProtoMessage()    {}
func (*TargetMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_ae9c8f46e151dfa7, []int{1}
}
func (m *TargetMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetMesh.Unmarshal(m, b)
}
func (m *TargetMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetMesh.Marshal(b, m, deterministic)
}
func (dst *TargetMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetMesh.Merge(dst, src)
}
func (m *TargetMesh) XXX_Size() int {
	return xxx_messageInfo_TargetMesh.Size(m)
}
func (m *TargetMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetMesh.DiscardUnknown(m)
}

var xxx_messageInfo_TargetMesh proto.InternalMessageInfo

type isTargetMesh_DeploymentLocation interface {
	isTargetMesh_DeploymentLocation()
	Equal(interface{}) bool
}

type TargetMesh_KubeLocation struct {
	KubeLocation *KubeLocation `protobuf:"bytes,2,opt,name=kube_location,json=kubeLocation,oneof"`
}

func (*TargetMesh_KubeLocation) isTargetMesh_DeploymentLocation() {}

func (m *TargetMesh) GetDeploymentLocation() isTargetMesh_DeploymentLocation {
	if m != nil {
		return m.DeploymentLocation
	}
	return nil
}

func (m *TargetMesh) GetMeshType() MeshType {
	if m != nil {
		return m.MeshType
	}
	return MeshType_ISTIO
}

func (m *TargetMesh) GetKubeLocation() *KubeLocation {
	if x, ok := m.GetDeploymentLocation().(*TargetMesh_KubeLocation); ok {
		return x.KubeLocation
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TargetMesh) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TargetMesh_OneofMarshaler, _TargetMesh_OneofUnmarshaler, _TargetMesh_OneofSizer, []interface{}{
		(*TargetMesh_KubeLocation)(nil),
	}
}

func _TargetMesh_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TargetMesh)
	// deployment_location
	switch x := m.DeploymentLocation.(type) {
	case *TargetMesh_KubeLocation:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KubeLocation); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TargetMesh.DeploymentLocation has unexpected type %T", x)
	}
	return nil
}

func _TargetMesh_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TargetMesh)
	switch tag {
	case 2: // deployment_location.kube_location
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KubeLocation)
		err := b.DecodeMessage(msg)
		m.DeploymentLocation = &TargetMesh_KubeLocation{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TargetMesh_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TargetMesh)
	// deployment_location
	switch x := m.DeploymentLocation.(type) {
	case *TargetMesh_KubeLocation:
		s := proto.Size(x.KubeLocation)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type KubeLocation struct {
	Kubeconfig           string   `protobuf:"bytes,1,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
	MasterAddress        string   `protobuf:"bytes,2,opt,name=master_address,json=masterAddress,proto3" json:"master_address,omitempty"`
	Namespace            string   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubeLocation) Reset()         { *m = KubeLocation{} }
func (m *KubeLocation) String() string { return proto.CompactTextString(m) }
func (*KubeLocation) ProtoMessage()    {}
func (*KubeLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_ae9c8f46e151dfa7, []int{2}
}
func (m *KubeLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubeLocation.Unmarshal(m, b)
}
func (m *KubeLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubeLocation.Marshal(b, m, deterministic)
}
func (dst *KubeLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubeLocation.Merge(dst, src)
}
func (m *KubeLocation) XXX_Size() int {
	return xxx_messageInfo_KubeLocation.Size(m)
}
func (m *KubeLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_KubeLocation.DiscardUnknown(m)
}

var xxx_messageInfo_KubeLocation proto.InternalMessageInfo

func (m *KubeLocation) GetKubeconfig() string {
	if m != nil {
		return m.Kubeconfig
	}
	return ""
}

func (m *KubeLocation) GetMasterAddress() string {
	if m != nil {
		return m.MasterAddress
	}
	return ""
}

func (m *KubeLocation) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// message MeshMetadata {
//    string name = 1;
//    MeshType meshType = 2;
// }
//
// // Defines mesh-level configuration for encryption. Supports communication within a mesh and through ingress.
// // Communicating between TLS and non-TLS enabled services is not supported. TODO: Is this ok?
// // TODO: What do we need to support communication across mesh?
type Encryption struct {
	// If set to true, TLS is enabled across the entire mesh.
	TlsEnabled bool `protobuf:"varint,1,opt,name=tlsEnabled,proto3" json:"tlsEnabled,omitempty"`
	// If TLS is enabled, this is the name of the secret containing the certs.
	// When using Istio, this should either be "istio.default", meaning Istio is using the default Citadel cert
	// generation, or "cacert", which is a custom-uploaded Kubernetes secret containing all the cert files.
	// When using Linkerd, this is the name of a secret that will be mounted into the linkerd Kubernetes DaemonSet.
	Secret               *core.ResourceRef `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Encryption) Reset()         { *m = Encryption{} }
func (m *Encryption) String() string { return proto.CompactTextString(m) }
func (*Encryption) ProtoMessage()    {}
func (*Encryption) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_ae9c8f46e151dfa7, []int{3}
}
func (m *Encryption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Encryption.Unmarshal(m, b)
}
func (m *Encryption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Encryption.Marshal(b, m, deterministic)
}
func (dst *Encryption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Encryption.Merge(dst, src)
}
func (m *Encryption) XXX_Size() int {
	return xxx_messageInfo_Encryption.Size(m)
}
func (m *Encryption) XXX_DiscardUnknown() {
	xxx_messageInfo_Encryption.DiscardUnknown(m)
}

var xxx_messageInfo_Encryption proto.InternalMessageInfo

func (m *Encryption) GetTlsEnabled() bool {
	if m != nil {
		return m.TlsEnabled
	}
	return false
}

func (m *Encryption) GetSecret() *core.ResourceRef {
	if m != nil {
		return m.Secret
	}
	return nil
}

func init() {
	proto.RegisterType((*Mesh)(nil), "supergloo.solo.io.Mesh")
	proto.RegisterType((*TargetMesh)(nil), "supergloo.solo.io.TargetMesh")
	proto.RegisterType((*KubeLocation)(nil), "supergloo.solo.io.KubeLocation")
	proto.RegisterType((*Encryption)(nil), "supergloo.solo.io.Encryption")
	proto.RegisterEnum("supergloo.solo.io.MeshType", MeshType_name, MeshType_value)
}
func (this *Mesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh)
	if !ok {
		that2, ok := that.(Mesh)
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
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.TargetMesh.Equal(that1.TargetMesh) {
		return false
	}
	if !this.Routing.Equal(that1.Routing) {
		return false
	}
	if !this.Encryption.Equal(that1.Encryption) {
		return false
	}
	if !this.Observability.Equal(that1.Observability) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TargetMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TargetMesh)
	if !ok {
		that2, ok := that.(TargetMesh)
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
	if this.MeshType != that1.MeshType {
		return false
	}
	if that1.DeploymentLocation == nil {
		if this.DeploymentLocation != nil {
			return false
		}
	} else if this.DeploymentLocation == nil {
		return false
	} else if !this.DeploymentLocation.Equal(that1.DeploymentLocation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TargetMesh_KubeLocation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TargetMesh_KubeLocation)
	if !ok {
		that2, ok := that.(TargetMesh_KubeLocation)
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
	if !this.KubeLocation.Equal(that1.KubeLocation) {
		return false
	}
	return true
}
func (this *KubeLocation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KubeLocation)
	if !ok {
		that2, ok := that.(KubeLocation)
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
	if this.Kubeconfig != that1.Kubeconfig {
		return false
	}
	if this.MasterAddress != that1.MasterAddress {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Encryption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Encryption)
	if !ok {
		that2, ok := that.(Encryption)
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
	if this.TlsEnabled != that1.TlsEnabled {
		return false
	}
	if !this.Secret.Equal(that1.Secret) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("mesh.proto", fileDescriptor_mesh_ae9c8f46e151dfa7) }

var fileDescriptor_mesh_ae9c8f46e151dfa7 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x57, 0x36, 0xba, 0xf6, 0xec, 0x8f, 0x36, 0x6f, 0xa0, 0x6c, 0xc0, 0x3a, 0x05, 0x21,
	0x10, 0x62, 0x89, 0x3a, 0xb8, 0x98, 0x90, 0x40, 0xa2, 0xa2, 0x40, 0xb5, 0x8d, 0x49, 0x5e, 0xaf,
	0xb8, 0xa9, 0x9c, 0xe4, 0x34, 0xb5, 0x9a, 0xc4, 0x91, 0xed, 0x4c, 0xea, 0x0b, 0xf0, 0x26, 0x48,
	0x3c, 0x0a, 0x4f, 0xb1, 0x0b, 0x1e, 0x81, 0x27, 0x40, 0x71, 0xd2, 0x36, 0x15, 0x45, 0xda, 0x55,
	0xed, 0xe3, 0xef, 0x77, 0x8e, 0xbf, 0xaf, 0x31, 0x40, 0x8c, 0x6a, 0xe4, 0xa4, 0x52, 0x68, 0x41,
	0x76, 0x55, 0x96, 0xa2, 0x0c, 0x23, 0x21, 0x1c, 0x25, 0x22, 0xe1, 0x70, 0x71, 0xb8, 0x1f, 0x8a,
	0x50, 0x98, 0x53, 0x37, 0x5f, 0x15, 0xc2, 0xc3, 0x76, 0xc8, 0xf5, 0x28, 0xf3, 0x1c, 0x5f, 0xc4,
	0x6e, 0xae, 0x3c, 0xe1, 0xa2, 0xf8, 0x1d, 0x73, 0xed, 0xb2, 0x94, 0xbb, 0x37, 0x6d, 0x37, 0x46,
	0xcd, 0x02, 0xa6, 0x59, 0x89, 0xb8, 0x77, 0x40, 0x94, 0x66, 0x3a, 0x53, 0x25, 0xf0, 0xea, 0x0e,
	0x80, 0xc4, 0x61, 0xa9, 0xde, 0x92, 0x22, 0xd3, 0x3c, 0x09, 0xcb, 0xed, 0x9e, 0xf0, 0x14, 0xca,
	0x1b, 0xe6, 0xf1, 0x88, 0xeb, 0x49, 0x51, 0xb4, 0xbf, 0xaf, 0xc2, 0xda, 0x25, 0xaa, 0x11, 0xf9,
	0x0c, 0xf5, 0x62, 0x94, 0x55, 0x3f, 0xae, 0xbd, 0xd8, 0x38, 0xdd, 0x77, 0x7c, 0x21, 0x71, 0xea,
	0xd9, 0xb9, 0x36, 0x67, 0x9d, 0x83, 0x5f, 0xb7, 0xad, 0x95, 0x3f, 0xb7, 0xad, 0x5d, 0x8d, 0x4a,
	0x07, 0x7c, 0x38, 0x7c, 0x6b, 0xf3, 0x30, 0x11, 0x12, 0x6d, 0x5a, 0xe2, 0xe4, 0x0c, 0x1a, 0x53,
	0x9b, 0xd6, 0xba, 0x69, 0xf5, 0x70, 0xb1, 0xd5, 0x65, 0x79, 0xda, 0x59, 0xcb, 0x9b, 0xd1, 0x99,
	0x9a, 0xbc, 0x87, 0x0d, 0xcd, 0x64, 0x88, 0x7a, 0x90, 0xe7, 0x6f, 0xdd, 0x33, 0xf0, 0x13, 0xe7,
	0x9f, 0x3f, 0xc0, 0xe9, 0x1b, 0x55, 0x7e, 0x6d, 0x0a, 0x7a, 0xb6, 0x26, 0x6f, 0x60, 0xbd, 0x74,
	0x6c, 0xd5, 0x0c, 0x7b, 0xb8, 0x84, 0xa5, 0x85, 0x82, 0x4e, 0xa5, 0xe4, 0x1d, 0x00, 0x26, 0xbe,
	0x9c, 0xa4, 0x9a, 0x8b, 0xc4, 0xf2, 0xfe, 0x3b, 0xb4, 0x3b, 0x13, 0xd1, 0x0a, 0x40, 0x3e, 0xc1,
	0xd6, 0x42, 0xae, 0x96, 0x6f, 0x3a, 0x1c, 0x2f, 0xe9, 0x70, 0x55, 0xd5, 0xd1, 0x45, 0xcc, 0xfe,
	0x51, 0x03, 0x98, 0xfb, 0x22, 0x67, 0xd0, 0xcc, 0x43, 0x18, 0xe8, 0x49, 0x8a, 0xc6, 0xcd, 0xf6,
	0xe9, 0xa3, 0x25, 0x2d, 0x73, 0x6d, 0x7f, 0x92, 0x62, 0x9e, 0x62, 0xb1, 0xca, 0x2f, 0x34, 0xce,
	0x3c, 0x1c, 0x44, 0xc2, 0x67, 0xc6, 0x52, 0x91, 0x63, 0x6b, 0x09, 0x7d, 0x9e, 0x79, 0x78, 0x51,
	0xca, 0xbe, 0xac, 0xd0, 0xcd, 0x71, 0x65, 0xdf, 0x79, 0x00, 0x7b, 0x01, 0xa6, 0x91, 0x98, 0xc4,
	0x98, 0xe8, 0x59, 0x37, 0x5b, 0xc1, 0x66, 0x15, 0x23, 0x47, 0x00, 0x39, 0xe6, 0x8b, 0x64, 0xc8,
	0x8b, 0xdc, 0x9b, 0xb4, 0x52, 0x21, 0xcf, 0x60, 0x3b, 0x66, 0x4a, 0xa3, 0x1c, 0xb0, 0x20, 0x90,
	0xa8, 0x94, 0xb9, 0x4f, 0x93, 0x6e, 0x15, 0xd5, 0x0f, 0x45, 0x91, 0x3c, 0x86, 0x66, 0xc2, 0x62,
	0x54, 0x29, 0xf3, 0xd1, 0x5a, 0x35, 0x8a, 0x79, 0xc1, 0x1e, 0x00, 0xcc, 0xe3, 0xcf, 0x47, 0xea,
	0x48, 0x75, 0x13, 0xe6, 0x45, 0x18, 0x98, 0x91, 0x0d, 0x5a, 0xa9, 0x90, 0x36, 0xd4, 0x15, 0xfa,
	0x12, 0x75, 0x69, 0xfd, 0x60, 0xf1, 0xfb, 0xa3, 0xa8, 0x44, 0x26, 0x7d, 0xa4, 0x38, 0xa4, 0xa5,
	0xf0, 0xe5, 0x53, 0x68, 0x4c, 0xa3, 0x24, 0x4d, 0xb8, 0xdf, 0xbb, 0xee, 0xf7, 0xae, 0x76, 0x56,
	0xc8, 0x26, 0x34, 0x2e, 0x7a, 0x5f, 0xcf, 0xbb, 0xf4, 0x63, 0x7b, 0xa7, 0xd6, 0x39, 0xf9, 0xf9,
	0xfb, 0xa8, 0xf6, 0xed, 0xf9, 0xb2, 0x37, 0x38, 0x8d, 0xd6, 0x4d, 0xc7, 0x61, 0xf9, 0x10, 0xbd,
	0xba, 0x79, 0x61, 0xaf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xe8, 0x65, 0xd4, 0x4e, 0x04,
	0x00, 0x00,
}
