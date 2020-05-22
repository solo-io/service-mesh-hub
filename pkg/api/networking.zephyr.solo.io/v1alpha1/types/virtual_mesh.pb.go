// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha1/virtual_mesh.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/gogo/protobuf/types"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
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

type VirtualMeshSpec_Federation_Mode int32

const (
	// All services in a VirtualMesh will be federated to all workloads in that Virtual Mesh.
	VirtualMeshSpec_Federation_PERMISSIVE VirtualMeshSpec_Federation_Mode = 0
)

var VirtualMeshSpec_Federation_Mode_name = map[int32]string{
	0: "PERMISSIVE",
}

var VirtualMeshSpec_Federation_Mode_value = map[string]int32{
	"PERMISSIVE": 0,
}

func (x VirtualMeshSpec_Federation_Mode) String() string {
	return proto.EnumName(VirtualMeshSpec_Federation_Mode_name, int32(x))
}

func (VirtualMeshSpec_Federation_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 1, 0}
}

type VirtualMeshSpec struct {
	// User-provided display name for the virtual mesh.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The meshes contained in this virtual mesh.
	Meshes               []*types.ResourceRef                  `protobuf:"bytes,2,rep,name=meshes,proto3" json:"meshes,omitempty"`
	CertificateAuthority *VirtualMeshSpec_CertificateAuthority `protobuf:"bytes,3,opt,name=certificate_authority,json=certificateAuthority,proto3" json:"certificate_authority,omitempty"`
	Federation           *VirtualMeshSpec_Federation           `protobuf:"bytes,4,opt,name=federation,proto3" json:"federation,omitempty"`
	// Types that are valid to be assigned to TrustModel:
	//	*VirtualMeshSpec_Shared
	//	*VirtualMeshSpec_Limited
	TrustModel isVirtualMeshSpec_TrustModel `protobuf_oneof:"trust_model"`
	//
	//If true, by default disallow traffic to all Services in the VirtualMesh unless explicitly allowed through AccessControlPolicies.
	//If false, by default allow traffic to all Services in the VirtualMesh.
	//If omitted, the default value depends on the type of the first mesh listed in `meshes`:
	//Istio: false
	//Appmesh: true
	EnforceAccessControl *types1.BoolValue `protobuf:"bytes,7,opt,name=enforce_access_control,json=enforceAccessControl,proto3" json:"enforce_access_control,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *VirtualMeshSpec) Reset()         { *m = VirtualMeshSpec{} }
func (m *VirtualMeshSpec) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec) ProtoMessage()    {}
func (*VirtualMeshSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0}
}
func (m *VirtualMeshSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec.Unmarshal(m, b)
}
func (m *VirtualMeshSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec.Merge(m, src)
}
func (m *VirtualMeshSpec) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec.Size(m)
}
func (m *VirtualMeshSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec proto.InternalMessageInfo

type isVirtualMeshSpec_TrustModel interface {
	isVirtualMeshSpec_TrustModel()
	Equal(interface{}) bool
}

type VirtualMeshSpec_Shared struct {
	Shared *VirtualMeshSpec_SharedTrust `protobuf:"bytes,5,opt,name=shared,proto3,oneof" json:"shared,omitempty"`
}
type VirtualMeshSpec_Limited struct {
	Limited *VirtualMeshSpec_LimitedTrust `protobuf:"bytes,6,opt,name=limited,proto3,oneof" json:"limited,omitempty"`
}

func (*VirtualMeshSpec_Shared) isVirtualMeshSpec_TrustModel()  {}
func (*VirtualMeshSpec_Limited) isVirtualMeshSpec_TrustModel() {}

func (m *VirtualMeshSpec) GetTrustModel() isVirtualMeshSpec_TrustModel {
	if m != nil {
		return m.TrustModel
	}
	return nil
}

func (m *VirtualMeshSpec) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *VirtualMeshSpec) GetMeshes() []*types.ResourceRef {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func (m *VirtualMeshSpec) GetCertificateAuthority() *VirtualMeshSpec_CertificateAuthority {
	if m != nil {
		return m.CertificateAuthority
	}
	return nil
}

func (m *VirtualMeshSpec) GetFederation() *VirtualMeshSpec_Federation {
	if m != nil {
		return m.Federation
	}
	return nil
}

func (m *VirtualMeshSpec) GetShared() *VirtualMeshSpec_SharedTrust {
	if x, ok := m.GetTrustModel().(*VirtualMeshSpec_Shared); ok {
		return x.Shared
	}
	return nil
}

func (m *VirtualMeshSpec) GetLimited() *VirtualMeshSpec_LimitedTrust {
	if x, ok := m.GetTrustModel().(*VirtualMeshSpec_Limited); ok {
		return x.Limited
	}
	return nil
}

func (m *VirtualMeshSpec) GetEnforceAccessControl() *types1.BoolValue {
	if m != nil {
		return m.EnforceAccessControl
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*VirtualMeshSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*VirtualMeshSpec_Shared)(nil),
		(*VirtualMeshSpec_Limited)(nil),
	}
}

type VirtualMeshSpec_CertificateAuthority struct {
	// If omitted, defaults to builtin.
	//
	// Types that are valid to be assigned to Type:
	//	*VirtualMeshSpec_CertificateAuthority_Builtin_
	//	*VirtualMeshSpec_CertificateAuthority_Provided_
	Type                 isVirtualMeshSpec_CertificateAuthority_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *VirtualMeshSpec_CertificateAuthority) Reset()         { *m = VirtualMeshSpec_CertificateAuthority{} }
func (m *VirtualMeshSpec_CertificateAuthority) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec_CertificateAuthority) ProtoMessage()    {}
func (*VirtualMeshSpec_CertificateAuthority) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 0}
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.Merge(m, src)
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.Size(m)
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_CertificateAuthority proto.InternalMessageInfo

type isVirtualMeshSpec_CertificateAuthority_Type interface {
	isVirtualMeshSpec_CertificateAuthority_Type()
	Equal(interface{}) bool
}

type VirtualMeshSpec_CertificateAuthority_Builtin_ struct {
	Builtin *VirtualMeshSpec_CertificateAuthority_Builtin `protobuf:"bytes,1,opt,name=builtin,proto3,oneof" json:"builtin,omitempty"`
}
type VirtualMeshSpec_CertificateAuthority_Provided_ struct {
	Provided *VirtualMeshSpec_CertificateAuthority_Provided `protobuf:"bytes,2,opt,name=provided,proto3,oneof" json:"provided,omitempty"`
}

func (*VirtualMeshSpec_CertificateAuthority_Builtin_) isVirtualMeshSpec_CertificateAuthority_Type() {}
func (*VirtualMeshSpec_CertificateAuthority_Provided_) isVirtualMeshSpec_CertificateAuthority_Type() {
}

func (m *VirtualMeshSpec_CertificateAuthority) GetType() isVirtualMeshSpec_CertificateAuthority_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *VirtualMeshSpec_CertificateAuthority) GetBuiltin() *VirtualMeshSpec_CertificateAuthority_Builtin {
	if x, ok := m.GetType().(*VirtualMeshSpec_CertificateAuthority_Builtin_); ok {
		return x.Builtin
	}
	return nil
}

func (m *VirtualMeshSpec_CertificateAuthority) GetProvided() *VirtualMeshSpec_CertificateAuthority_Provided {
	if x, ok := m.GetType().(*VirtualMeshSpec_CertificateAuthority_Provided_); ok {
		return x.Provided
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*VirtualMeshSpec_CertificateAuthority) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*VirtualMeshSpec_CertificateAuthority_Builtin_)(nil),
		(*VirtualMeshSpec_CertificateAuthority_Provided_)(nil),
	}
}

//
//Configuration for auto-generated root certificate unique to the VirtualMesh
//Uses the X.509 format, RFC5280
type VirtualMeshSpec_CertificateAuthority_Builtin struct {
	// Number of days before root cert expires. Defaults to 365.
	TtlDays uint32 `protobuf:"varint,1,opt,name=ttl_days,json=ttlDays,proto3" json:"ttl_days,omitempty"`
	// Size in bytes of the root cert's private key. Defaults to 4096
	RsaKeySizeBytes uint32 `protobuf:"varint,2,opt,name=rsa_key_size_bytes,json=rsaKeySizeBytes,proto3" json:"rsa_key_size_bytes,omitempty"`
	// Root cert organization name. Defaults to "service-mesh-hub"
	OrgName              string   `protobuf:"bytes,3,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMeshSpec_CertificateAuthority_Builtin) Reset() {
	*m = VirtualMeshSpec_CertificateAuthority_Builtin{}
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) String() string {
	return proto.CompactTextString(m)
}
func (*VirtualMeshSpec_CertificateAuthority_Builtin) ProtoMessage() {}
func (*VirtualMeshSpec_CertificateAuthority_Builtin) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 0, 0}
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.Merge(m, src)
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.Size(m)
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin proto.InternalMessageInfo

func (m *VirtualMeshSpec_CertificateAuthority_Builtin) GetTtlDays() uint32 {
	if m != nil {
		return m.TtlDays
	}
	return 0
}

func (m *VirtualMeshSpec_CertificateAuthority_Builtin) GetRsaKeySizeBytes() uint32 {
	if m != nil {
		return m.RsaKeySizeBytes
	}
	return 0
}

func (m *VirtualMeshSpec_CertificateAuthority_Builtin) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

// Configuration for user-provided root certificate.
type VirtualMeshSpec_CertificateAuthority_Provided struct {
	// Reference to a Secret object containing the root certificate.
	Certificate          *types.ResourceRef `protobuf:"bytes,3,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *VirtualMeshSpec_CertificateAuthority_Provided) Reset() {
	*m = VirtualMeshSpec_CertificateAuthority_Provided{}
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) String() string {
	return proto.CompactTextString(m)
}
func (*VirtualMeshSpec_CertificateAuthority_Provided) ProtoMessage() {}
func (*VirtualMeshSpec_CertificateAuthority_Provided) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 0, 1}
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.Merge(m, src)
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.Size(m)
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided proto.InternalMessageInfo

func (m *VirtualMeshSpec_CertificateAuthority_Provided) GetCertificate() *types.ResourceRef {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type VirtualMeshSpec_Federation struct {
	Mode                 VirtualMeshSpec_Federation_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=networking.zephyr.solo.io.VirtualMeshSpec_Federation_Mode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *VirtualMeshSpec_Federation) Reset()         { *m = VirtualMeshSpec_Federation{} }
func (m *VirtualMeshSpec_Federation) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec_Federation) ProtoMessage()    {}
func (*VirtualMeshSpec_Federation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 1}
}
func (m *VirtualMeshSpec_Federation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_Federation.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_Federation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_Federation.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_Federation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_Federation.Merge(m, src)
}
func (m *VirtualMeshSpec_Federation) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_Federation.Size(m)
}
func (m *VirtualMeshSpec_Federation) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_Federation.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_Federation proto.InternalMessageInfo

func (m *VirtualMeshSpec_Federation) GetMode() VirtualMeshSpec_Federation_Mode {
	if m != nil {
		return m.Mode
	}
	return VirtualMeshSpec_Federation_PERMISSIVE
}

//
//Shared trust is a virtual mesh trust model requiring a shared root certificate, as well as shared identity
//between all entities which wish to communicate within the virtual mesh.
//
//The best current example of this would be the replicated control planes example from Istio:
//https://preliminary.istio.io/docs/setup/install/multicluster/gateways/
type VirtualMeshSpec_SharedTrust struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMeshSpec_SharedTrust) Reset()         { *m = VirtualMeshSpec_SharedTrust{} }
func (m *VirtualMeshSpec_SharedTrust) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec_SharedTrust) ProtoMessage()    {}
func (*VirtualMeshSpec_SharedTrust) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 2}
}
func (m *VirtualMeshSpec_SharedTrust) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_SharedTrust.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_SharedTrust) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_SharedTrust.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_SharedTrust) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_SharedTrust.Merge(m, src)
}
func (m *VirtualMeshSpec_SharedTrust) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_SharedTrust.Size(m)
}
func (m *VirtualMeshSpec_SharedTrust) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_SharedTrust.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_SharedTrust proto.InternalMessageInfo

//
//Limited trust is a virtual mesh trust model which does not require all meshes sharing the same root certificate
//or identity model. But rather, the limited trust creates trust between meshes running on different clusters
//by connecting their ingress/egress gateways with a common cert/identity. In this model all requests
//between different have the following request path when communicating between clusters
//
//cluster 1 MTLS               shared MTLS                  cluster 2 MTLS
//client/workload <-----------> egress gateway <----------> ingress gateway <--------------> server
//
//This approach has the downside of not maintaining identity from client to server, but allows for ad-hoc
//addition of additional clusters into a virtual mesh.
type VirtualMeshSpec_LimitedTrust struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMeshSpec_LimitedTrust) Reset()         { *m = VirtualMeshSpec_LimitedTrust{} }
func (m *VirtualMeshSpec_LimitedTrust) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec_LimitedTrust) ProtoMessage()    {}
func (*VirtualMeshSpec_LimitedTrust) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 3}
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_LimitedTrust.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_LimitedTrust.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_LimitedTrust.Merge(m, src)
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_LimitedTrust.Size(m)
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_LimitedTrust.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_LimitedTrust proto.InternalMessageInfo

type VirtualMeshStatus struct {
	// Status of the process writing federation decision metadata onto MeshServices.
	FederationStatus *types.Status `protobuf:"bytes,1,opt,name=federation_status,json=federationStatus,proto3" json:"federation_status,omitempty"`
	// Status of the process signing CSRs.
	CertificateStatus *types.Status `protobuf:"bytes,2,opt,name=certificate_status,json=certificateStatus,proto3" json:"certificate_status,omitempty"`
	// Overall validation status of this VirtualMesh.
	ConfigStatus *types.Status `protobuf:"bytes,3,opt,name=config_status,json=configStatus,proto3" json:"config_status,omitempty"`
	// Status of ensuring that access control is enforced within this VirtualMesh.
	AccessControlEnforcementStatus *types.Status `protobuf:"bytes,4,opt,name=access_control_enforcement_status,json=accessControlEnforcementStatus,proto3" json:"access_control_enforcement_status,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}      `json:"-"`
	XXX_unrecognized               []byte        `json:"-"`
	XXX_sizecache                  int32         `json:"-"`
}

func (m *VirtualMeshStatus) Reset()         { *m = VirtualMeshStatus{} }
func (m *VirtualMeshStatus) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshStatus) ProtoMessage()    {}
func (*VirtualMeshStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{1}
}
func (m *VirtualMeshStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshStatus.Unmarshal(m, b)
}
func (m *VirtualMeshStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshStatus.Marshal(b, m, deterministic)
}
func (m *VirtualMeshStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshStatus.Merge(m, src)
}
func (m *VirtualMeshStatus) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshStatus.Size(m)
}
func (m *VirtualMeshStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshStatus.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshStatus proto.InternalMessageInfo

func (m *VirtualMeshStatus) GetFederationStatus() *types.Status {
	if m != nil {
		return m.FederationStatus
	}
	return nil
}

func (m *VirtualMeshStatus) GetCertificateStatus() *types.Status {
	if m != nil {
		return m.CertificateStatus
	}
	return nil
}

func (m *VirtualMeshStatus) GetConfigStatus() *types.Status {
	if m != nil {
		return m.ConfigStatus
	}
	return nil
}

func (m *VirtualMeshStatus) GetAccessControlEnforcementStatus() *types.Status {
	if m != nil {
		return m.AccessControlEnforcementStatus
	}
	return nil
}

func init() {
	proto.RegisterEnum("networking.zephyr.solo.io.VirtualMeshSpec_Federation_Mode", VirtualMeshSpec_Federation_Mode_name, VirtualMeshSpec_Federation_Mode_value)
	proto.RegisterType((*VirtualMeshSpec)(nil), "networking.zephyr.solo.io.VirtualMeshSpec")
	proto.RegisterType((*VirtualMeshSpec_CertificateAuthority)(nil), "networking.zephyr.solo.io.VirtualMeshSpec.CertificateAuthority")
	proto.RegisterType((*VirtualMeshSpec_CertificateAuthority_Builtin)(nil), "networking.zephyr.solo.io.VirtualMeshSpec.CertificateAuthority.Builtin")
	proto.RegisterType((*VirtualMeshSpec_CertificateAuthority_Provided)(nil), "networking.zephyr.solo.io.VirtualMeshSpec.CertificateAuthority.Provided")
	proto.RegisterType((*VirtualMeshSpec_Federation)(nil), "networking.zephyr.solo.io.VirtualMeshSpec.Federation")
	proto.RegisterType((*VirtualMeshSpec_SharedTrust)(nil), "networking.zephyr.solo.io.VirtualMeshSpec.SharedTrust")
	proto.RegisterType((*VirtualMeshSpec_LimitedTrust)(nil), "networking.zephyr.solo.io.VirtualMeshSpec.LimitedTrust")
	proto.RegisterType((*VirtualMeshStatus)(nil), "networking.zephyr.solo.io.VirtualMeshStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/networking/v1alpha1/virtual_mesh.proto", fileDescriptor_c28e03fd4cc9e166)
}

var fileDescriptor_c28e03fd4cc9e166 = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdf, 0x6e, 0xdb, 0x36,
	0x14, 0xc6, 0x93, 0xd8, 0xb3, 0xd3, 0xe3, 0x24, 0x6d, 0x08, 0xaf, 0x50, 0x35, 0x20, 0x70, 0x73,
	0x15, 0x60, 0x8b, 0x84, 0x7a, 0xd8, 0x1f, 0xec, 0x66, 0xab, 0xbb, 0x6c, 0xee, 0xb6, 0x18, 0x86,
	0xdc, 0xe5, 0xa2, 0x37, 0x02, 0x2d, 0x1d, 0x49, 0x44, 0x64, 0x51, 0x20, 0x29, 0x17, 0xca, 0xf5,
	0x1e, 0x66, 0x7b, 0x9e, 0xbd, 0xc1, 0x9e, 0x64, 0x10, 0x25, 0x55, 0x9a, 0xe1, 0xc1, 0x4d, 0x73,
	0x47, 0x53, 0xe7, 0xfb, 0x1d, 0x92, 0xfe, 0x3e, 0x12, 0x66, 0x21, 0x53, 0x51, 0xb6, 0xb4, 0x3c,
	0xbe, 0xb2, 0x25, 0x8f, 0xf9, 0x25, 0xe3, 0xb6, 0x44, 0xb1, 0x66, 0x1e, 0x5e, 0xae, 0x50, 0x46,
	0x97, 0x51, 0xb6, 0xb4, 0x69, 0xca, 0xec, 0x04, 0xd5, 0x3b, 0x2e, 0x6e, 0x59, 0x12, 0xda, 0xeb,
	0x17, 0x34, 0x4e, 0x23, 0xfa, 0xc2, 0x5e, 0x33, 0xa1, 0x32, 0x1a, 0xbb, 0x45, 0xa1, 0x95, 0x0a,
	0xae, 0x38, 0x79, 0xd6, 0xd4, 0x59, 0x77, 0x98, 0x46, 0xb9, 0xb0, 0x0a, 0xac, 0xc5, 0xb8, 0xf9,
	0xc5, 0x56, 0xae, 0xc7, 0x05, 0x36, 0x44, 0x81, 0x41, 0x09, 0x32, 0xed, 0x0f, 0xa8, 0x96, 0x8a,
	0xaa, 0x4c, 0x56, 0x82, 0xb3, 0x90, 0xf3, 0x30, 0x46, 0x5b, 0xff, 0x5a, 0x66, 0x81, 0xfd, 0x4e,
	0xd0, 0x34, 0x45, 0x51, 0x7f, 0x1f, 0x86, 0x3c, 0xe4, 0x7a, 0x68, 0x17, 0xa3, 0x72, 0xf6, 0xfc,
	0x8f, 0x47, 0xf0, 0xf8, 0xa6, 0xdc, 0xc6, 0x35, 0xca, 0x68, 0x91, 0xa2, 0x47, 0x9e, 0xc3, 0x91,
	0xcf, 0x64, 0x1a, 0xd3, 0xdc, 0x4d, 0xe8, 0x0a, 0x8d, 0xfd, 0xd1, 0xfe, 0xc5, 0x23, 0x67, 0x50,
	0xcd, 0xcd, 0xe8, 0x0a, 0xc9, 0xb7, 0xd0, 0x2b, 0xd6, 0x85, 0xd2, 0x38, 0x18, 0x75, 0x2e, 0x06,
	0xe3, 0x91, 0x55, 0xac, 0x6c, 0x63, 0xc7, 0x96, 0x83, 0x92, 0x67, 0xc2, 0x43, 0x07, 0x03, 0xa7,
	0xaa, 0x27, 0x0a, 0x3e, 0xf5, 0x50, 0x28, 0x16, 0x30, 0x8f, 0x2a, 0x74, 0x69, 0xa6, 0x22, 0x2e,
	0x98, 0xca, 0x8d, 0xce, 0x68, 0xff, 0x62, 0x30, 0xfe, 0xde, 0xfa, 0xdf, 0x03, 0xb4, 0x36, 0xd6,
	0x69, 0xbd, 0x6a, 0x38, 0x2f, 0x6b, 0x8c, 0x33, 0xf4, 0xb6, 0xcc, 0x92, 0xdf, 0x01, 0x02, 0xf4,
	0x51, 0x50, 0xc5, 0x78, 0x62, 0x74, 0x75, 0xab, 0xaf, 0xee, 0xd1, 0xea, 0xa7, 0xf7, 0x62, 0xa7,
	0x05, 0x22, 0x73, 0xe8, 0xc9, 0x88, 0x0a, 0xf4, 0x8d, 0x4f, 0x34, 0xf2, 0xeb, 0x7b, 0x20, 0x17,
	0x5a, 0xf8, 0x46, 0x64, 0x52, 0x4d, 0xf7, 0x9c, 0x8a, 0x43, 0x16, 0xd0, 0x8f, 0xd9, 0x8a, 0x29,
	0xf4, 0x8d, 0x9e, 0x46, 0x7e, 0x73, 0x0f, 0xe4, 0x6f, 0xa5, 0xb2, 0x66, 0xd6, 0x24, 0x32, 0x87,
	0xa7, 0x98, 0x04, 0x5c, 0x78, 0xe8, 0x52, 0xcf, 0x43, 0x29, 0x5d, 0x8f, 0x27, 0x4a, 0xf0, 0xd8,
	0xe8, 0xeb, 0x1e, 0xa6, 0x55, 0x7a, 0xc7, 0xaa, 0xbd, 0x63, 0x4d, 0x38, 0x8f, 0x6f, 0x68, 0x9c,
	0xa1, 0x33, 0xac, 0x94, 0x2f, 0xb5, 0xf0, 0x55, 0xa9, 0x33, 0xff, 0xea, 0xc0, 0x70, 0xdb, 0xf1,
	0x13, 0x0f, 0xfa, 0xcb, 0x8c, 0xc5, 0x8a, 0x25, 0xda, 0x36, 0x83, 0xf1, 0xcf, 0x0f, 0xfc, 0x43,
	0xad, 0x49, 0x89, 0x2b, 0xf6, 0x53, 0x91, 0x49, 0x00, 0x87, 0xa9, 0xe0, 0x6b, 0xe6, 0xa3, 0x6f,
	0x1c, 0xe8, 0x2e, 0xd3, 0x87, 0x76, 0x99, 0x57, 0xbc, 0xe9, 0x9e, 0xf3, 0x9e, 0x6d, 0xc6, 0xd0,
	0xaf, 0xba, 0x93, 0x67, 0x70, 0xa8, 0x54, 0xec, 0xfa, 0x34, 0x97, 0x7a, 0x63, 0xc7, 0x4e, 0x5f,
	0xa9, 0xf8, 0x47, 0x9a, 0x4b, 0xf2, 0x39, 0x10, 0x21, 0xa9, 0x7b, 0x8b, 0xb9, 0x2b, 0xd9, 0x1d,
	0xba, 0xcb, 0x5c, 0xe9, 0x5c, 0x14, 0x45, 0x8f, 0x85, 0xa4, 0xbf, 0x62, 0xbe, 0x60, 0x77, 0x38,
	0x29, 0xa6, 0x0b, 0x0e, 0x17, 0x61, 0x99, 0xab, 0x8e, 0xce, 0x55, 0x9f, 0x8b, 0xb0, 0xc8, 0x94,
	0x39, 0x83, 0xc3, 0x7a, 0x15, 0x64, 0x02, 0x83, 0x96, 0x8f, 0xab, 0x6c, 0xec, 0x0e, 0x59, 0x5b,
	0x34, 0xe9, 0x41, 0x57, 0xe5, 0x29, 0x9a, 0x0a, 0xa0, 0xb1, 0x2f, 0x99, 0x41, 0x77, 0xc5, 0xfd,
	0x32, 0xd4, 0x27, 0xe3, 0xef, 0x3e, 0x2a, 0x03, 0xd6, 0x35, 0xf7, 0xd1, 0xd1, 0x9c, 0xf3, 0xa7,
	0xd0, 0x2d, 0x7e, 0x91, 0x13, 0x80, 0xf9, 0x95, 0x73, 0xfd, 0x7a, 0xb1, 0x78, 0x7d, 0x73, 0xf5,
	0x64, 0xcf, 0x3c, 0x86, 0x41, 0xcb, 0xe1, 0xe6, 0x09, 0x1c, 0xb5, 0xdd, 0x39, 0x39, 0x86, 0x81,
	0x2a, 0x06, 0x6e, 0x01, 0x89, 0xcf, 0xff, 0x3e, 0x80, 0xd3, 0x76, 0x3f, 0x7d, 0xb1, 0x91, 0x29,
	0x9c, 0x36, 0x61, 0x73, 0xcb, 0xdb, 0xae, 0xb2, 0xd5, 0x67, 0x5b, 0xcf, 0xa2, 0xd4, 0x39, 0x4f,
	0x1a, 0x55, 0x45, 0xfa, 0x05, 0x48, 0xfb, 0xd6, 0xa9, 0x50, 0x07, 0xbb, 0x51, 0xa7, 0x2d, 0x59,
	0xc5, 0xfa, 0x01, 0x8e, 0x3d, 0x9e, 0x04, 0x2c, 0xac, 0x31, 0x9d, 0xdd, 0x98, 0xa3, 0x52, 0x51,
	0x11, 0x02, 0x78, 0xfe, 0xdf, 0x1c, 0xba, 0x55, 0xc8, 0x56, 0x98, 0xa8, 0x9a, 0xda, 0xdd, 0x4d,
	0x3d, 0xa3, 0xed, 0x50, 0x5e, 0x35, 0x8c, 0xf2, 0xfb, 0xe4, 0xed, 0x9f, 0xff, 0x9c, 0xed, 0xbf,
	0x7d, 0xf3, 0x21, 0x4f, 0x5c, 0x7a, 0x1b, 0x6e, 0x3c, 0x73, 0x1b, 0x3d, 0x9b, 0x57, 0xa7, 0x30,
	0x95, 0x5c, 0xf6, 0xf4, 0x5d, 0xf1, 0xe5, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x18, 0x78, 0xe6,
	0x88, 0x41, 0x07, 0x00, 0x00,
}

func (this *VirtualMeshSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec)
	if !ok {
		that2, ok := that.(VirtualMeshSpec)
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
	if this.DisplayName != that1.DisplayName {
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
	if !this.CertificateAuthority.Equal(that1.CertificateAuthority) {
		return false
	}
	if !this.Federation.Equal(that1.Federation) {
		return false
	}
	if that1.TrustModel == nil {
		if this.TrustModel != nil {
			return false
		}
	} else if this.TrustModel == nil {
		return false
	} else if !this.TrustModel.Equal(that1.TrustModel) {
		return false
	}
	if !this.EnforceAccessControl.Equal(that1.EnforceAccessControl) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_Shared) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_Shared)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_Shared)
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
	if !this.Shared.Equal(that1.Shared) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_Limited) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_Limited)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_Limited)
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
	if !this.Limited.Equal(that1.Limited) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority_Builtin_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority_Builtin_)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority_Builtin_)
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
	if !this.Builtin.Equal(that1.Builtin) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority_Provided_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority_Provided_)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority_Provided_)
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
	if !this.Provided.Equal(that1.Provided) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority_Builtin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority_Builtin)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority_Builtin)
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
	if this.TtlDays != that1.TtlDays {
		return false
	}
	if this.RsaKeySizeBytes != that1.RsaKeySizeBytes {
		return false
	}
	if this.OrgName != that1.OrgName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority_Provided) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority_Provided)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority_Provided)
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
	if !this.Certificate.Equal(that1.Certificate) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_Federation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_Federation)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_Federation)
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
	if this.Mode != that1.Mode {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_SharedTrust) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_SharedTrust)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_SharedTrust)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_LimitedTrust) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_LimitedTrust)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_LimitedTrust)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshStatus)
	if !ok {
		that2, ok := that.(VirtualMeshStatus)
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
	if !this.FederationStatus.Equal(that1.FederationStatus) {
		return false
	}
	if !this.CertificateStatus.Equal(that1.CertificateStatus) {
		return false
	}
	if !this.ConfigStatus.Equal(that1.ConfigStatus) {
		return false
	}
	if !this.AccessControlEnforcementStatus.Equal(that1.AccessControlEnforcementStatus) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
