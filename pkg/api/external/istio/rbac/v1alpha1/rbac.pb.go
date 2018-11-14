// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rbac.proto

package v1alpha1 // import "github.com/solo-io/supergloo/pkg/api/external/istio/rbac/v1alpha1"

/*
Istio RBAC (Role Based Access Control) defines ServiceRole and ServiceRoleBinding
objects.

A ServiceRole specification includes a list of rules (permissions). Each rule has
the following standard fields:

  * services: a list of services.
  * methods: HTTP methods. In the case of gRPC, this field is ignored because the value is always "POST".
  * paths: HTTP paths or gRPC methods. Note that gRPC methods should be
    presented in the form of "/packageName.serviceName/methodName" and are case sensitive.

In addition to the standard fields, operators can also use custom keys in the `constraints` field,
the supported keys are listed in the "constraints and properties" page.

Below is an example of ServiceRole object "product-viewer", which has "read" ("GET" and "HEAD")
access to "products.svc.cluster.local" service at versions "v1" and "v2". "path" is not specified,
so it applies to any path in the service.

```yaml
apiVersion: "rbac.istio.io/v1alpha1"
kind: ServiceRole
metadata:
  name: products-viewer
  namespace: default
spec:
  rules:
  - services: ["products.svc.cluster.local"]
    methods: ["GET", "HEAD"]
    constraints:
    - key: "destination.labels[version]"
      value: ["v1", "v2"]
```

A ServiceRoleBinding specification includes two parts:

 * The `roleRef` field that refers to a ServiceRole object in the same namespace.
 * A list of `subjects` that are assigned the roles.

In addition to a simple `user` field, operators can also use custom keys in the `properties` field,
the supported keys are listed in the "constraints and properties" page.

Below is an example of ServiceRoleBinding object "test-binding-products", which binds two subjects
to ServiceRole "product-viewer":

  * User "alice@yahoo.com"
  * Services in "abc" namespace.

```yaml
apiVersion: "rbac.istio.io/v1alpha1"
kind: ServiceRoleBinding
metadata:
  name: test-binding-products
  namespace: default
spec:
  subjects:
  - user: alice@yahoo.com
  - properties:
      source.namespace: "abc"
  roleRef:
    kind: ServiceRole
    name: "products-viewer"
```
*/

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// $hide_from_docs
// RBAC ServiceRoleBinding enforcement mode, used to verify new ServiceRoleBinding
// configs work as expected before rolling to production. RBAC engine only logs results
// from configs that are in permissive mode, and discards result before returning
// to the user.
type EnforcementMode int32

const (
	// Policy in ENFORCED mode has impact on user experience.
	// Policy is in ENFORCED mode by default.
	EnforcementMode_ENFORCED EnforcementMode = 0
	// Policy in PERMISSIVE mode isn't enforced and has no impact on users.
	// RBAC engine run policies in PERMISSIVE mode and logs stats.
	EnforcementMode_PERMISSIVE EnforcementMode = 1
)

var EnforcementMode_name = map[int32]string{
	0: "ENFORCED",
	1: "PERMISSIVE",
}
var EnforcementMode_value = map[string]int32{
	"ENFORCED":   0,
	"PERMISSIVE": 1,
}

func (x EnforcementMode) String() string {
	return proto.EnumName(EnforcementMode_name, int32(x))
}
func (EnforcementMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rbac_20b3f085115066c7, []int{0}
}

type RbacConfig_Mode int32

const (
	// Disable Istio RBAC completely, any other config in RbacConfig will be ignored and Istio RBAC policies
	// will not be enforced.
	RbacConfig_OFF RbacConfig_Mode = 0
	// Enable Istio RBAC for all services and namespaces.
	RbacConfig_ON RbacConfig_Mode = 1
	// Enable Istio RBAC only for services and namespaces specified in the inclusion field. Any other
	// services and namespaces not in the inclusion field will not be enforced by Istio RBAC policies.
	RbacConfig_ON_WITH_INCLUSION RbacConfig_Mode = 2
	// Enable Istio RBAC for all services and namespaces except those specified in the exclusion field. Any other
	// services and namespaces not in the exclusion field will be enforced by Istio RBAC policies.
	RbacConfig_ON_WITH_EXCLUSION RbacConfig_Mode = 3
)

var RbacConfig_Mode_name = map[int32]string{
	0: "OFF",
	1: "ON",
	2: "ON_WITH_INCLUSION",
	3: "ON_WITH_EXCLUSION",
}
var RbacConfig_Mode_value = map[string]int32{
	"OFF":               0,
	"ON":                1,
	"ON_WITH_INCLUSION": 2,
	"ON_WITH_EXCLUSION": 3,
}

func (x RbacConfig_Mode) String() string {
	return proto.EnumName(RbacConfig_Mode_name, int32(x))
}
func (RbacConfig_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rbac_20b3f085115066c7, []int{5, 0}
}

// ServiceRole specification contains a list of access rules (permissions).
// This represent the "Spec" part of the ServiceRole object. The name and namespace
// of the ServiceRole is specified in "metadata" section of the ServiceRole object.
// @solo-kit:resource.short_name=svcrole
// @solo-kit:resource.plural_name=service_roles
// @solo-kit:resource.resource_groups=rbac.istio.io
type ServiceRole struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata" json:"metadata"`
	// Required. The set of access rules (permissions) that the role has.
	Rules                []*AccessRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ServiceRole) Reset()         { *m = ServiceRole{} }
func (m *ServiceRole) String() string { return proto.CompactTextString(m) }
func (*ServiceRole) ProtoMessage()    {}
func (*ServiceRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_20b3f085115066c7, []int{0}
}
func (m *ServiceRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceRole.Unmarshal(m, b)
}
func (m *ServiceRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceRole.Marshal(b, m, deterministic)
}
func (dst *ServiceRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceRole.Merge(dst, src)
}
func (m *ServiceRole) XXX_Size() int {
	return xxx_messageInfo_ServiceRole.Size(m)
}
func (m *ServiceRole) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceRole.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceRole proto.InternalMessageInfo

func (m *ServiceRole) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *ServiceRole) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *ServiceRole) GetRules() []*AccessRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// AccessRule defines a permission to access a list of services.
type AccessRule struct {
	// Required. A list of service names.
	// Exact match, prefix match, and suffix match are supported for service names.
	// For example, the service name "bookstore.mtv.cluster.local" matches
	// "bookstore.mtv.cluster.local" (exact match), or "bookstore*" (prefix match),
	// or "*.mtv.cluster.local" (suffix match).
	// If set to ["*"], it refers to all services in the namespace.
	Services []string `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
	// Optional. A list of HTTP paths or gRPC methods.
	// gRPC methods must be presented as fully-qualified name in the form of
	// "/packageName.serviceName/methodName" and are case sensitive.
	// Exact match, prefix match, and suffix match are supported for paths.
	// For example, the path "/books/review" matches
	// "/books/review" (exact match), or "/books/*" (prefix match),
	// or "*/review" (suffix match).
	// If not specified, it applies to any path.
	Paths []string `protobuf:"bytes,2,rep,name=paths" json:"paths,omitempty"`
	// Optional. A list of HTTP methods (e.g., "GET", "POST").
	// It is ignored in gRPC case because the value is always "POST".
	// If set to ["*"] or not specified, it applies to any method.
	Methods []string `protobuf:"bytes,3,rep,name=methods" json:"methods,omitempty"`
	// Optional. Extra constraints in the ServiceRole specification.
	// The above ServiceRole example shows an example of constraint "version".
	Constraints          []*AccessRule_Constraint `protobuf:"bytes,4,rep,name=constraints" json:"constraints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AccessRule) Reset()         { *m = AccessRule{} }
func (m *AccessRule) String() string { return proto.CompactTextString(m) }
func (*AccessRule) ProtoMessage()    {}
func (*AccessRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_20b3f085115066c7, []int{1}
}
func (m *AccessRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessRule.Unmarshal(m, b)
}
func (m *AccessRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessRule.Marshal(b, m, deterministic)
}
func (dst *AccessRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessRule.Merge(dst, src)
}
func (m *AccessRule) XXX_Size() int {
	return xxx_messageInfo_AccessRule.Size(m)
}
func (m *AccessRule) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessRule.DiscardUnknown(m)
}

var xxx_messageInfo_AccessRule proto.InternalMessageInfo

func (m *AccessRule) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *AccessRule) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *AccessRule) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *AccessRule) GetConstraints() []*AccessRule_Constraint {
	if m != nil {
		return m.Constraints
	}
	return nil
}

// Definition of a custom constraint. The supported keys are listed in the "constraint and properties" page.
type AccessRule_Constraint struct {
	// Key of the constraint.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// List of valid values for the constraint.
	// Exact match, prefix match, and suffix match are supported for constraint values.
	// For example, the value "v1alpha2" matches
	// "v1alpha2" (exact match), or "v1*" (prefix match),
	// or "*alpha2" (suffix match).
	Values               []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessRule_Constraint) Reset()         { *m = AccessRule_Constraint{} }
func (m *AccessRule_Constraint) String() string { return proto.CompactTextString(m) }
func (*AccessRule_Constraint) ProtoMessage()    {}
func (*AccessRule_Constraint) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_20b3f085115066c7, []int{1, 0}
}
func (m *AccessRule_Constraint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessRule_Constraint.Unmarshal(m, b)
}
func (m *AccessRule_Constraint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessRule_Constraint.Marshal(b, m, deterministic)
}
func (dst *AccessRule_Constraint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessRule_Constraint.Merge(dst, src)
}
func (m *AccessRule_Constraint) XXX_Size() int {
	return xxx_messageInfo_AccessRule_Constraint.Size(m)
}
func (m *AccessRule_Constraint) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessRule_Constraint.DiscardUnknown(m)
}

var xxx_messageInfo_AccessRule_Constraint proto.InternalMessageInfo

func (m *AccessRule_Constraint) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AccessRule_Constraint) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// ServiceRoleBinding assigns a ServiceRole to a list of subjects.
// This represents the "Spec" part of the ServiceRoleBinding object. The name and namespace
// of the ServiceRoleBinding is specified in "metadata" section of the ServiceRoleBinding
// object.
// @solo-kit:resource.short_name=svcrolebind
// @solo-kit:resource.plural_name=service_role_bindings
// @solo-kit:resource.resource_groups=rbac.istio.io
type ServiceRoleBinding struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata" json:"metadata"`
	// Required. List of subjects that are assigned the ServiceRole object.
	Subjects []*Subject `protobuf:"bytes,1,rep,name=subjects" json:"subjects,omitempty"`
	// Required. Reference to the ServiceRole object.
	RoleRef *RoleRef `protobuf:"bytes,2,opt,name=roleRef" json:"roleRef,omitempty"`
	// $hide_from_docs
	// Indicates enforcement mode of the ServiceRoleBinding.
	Mode                 EnforcementMode `protobuf:"varint,3,opt,name=mode,proto3,enum=istio.rbac.v1alpha1.EnforcementMode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ServiceRoleBinding) Reset()         { *m = ServiceRoleBinding{} }
func (m *ServiceRoleBinding) String() string { return proto.CompactTextString(m) }
func (*ServiceRoleBinding) ProtoMessage()    {}
func (*ServiceRoleBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_20b3f085115066c7, []int{2}
}
func (m *ServiceRoleBinding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceRoleBinding.Unmarshal(m, b)
}
func (m *ServiceRoleBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceRoleBinding.Marshal(b, m, deterministic)
}
func (dst *ServiceRoleBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceRoleBinding.Merge(dst, src)
}
func (m *ServiceRoleBinding) XXX_Size() int {
	return xxx_messageInfo_ServiceRoleBinding.Size(m)
}
func (m *ServiceRoleBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceRoleBinding.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceRoleBinding proto.InternalMessageInfo

func (m *ServiceRoleBinding) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *ServiceRoleBinding) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *ServiceRoleBinding) GetSubjects() []*Subject {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *ServiceRoleBinding) GetRoleRef() *RoleRef {
	if m != nil {
		return m.RoleRef
	}
	return nil
}

func (m *ServiceRoleBinding) GetMode() EnforcementMode {
	if m != nil {
		return m.Mode
	}
	return EnforcementMode_ENFORCED
}

// Subject defines an identity. The identity is either a user or identified by a set of `properties`.
// The supported keys in `properties` are listed in "constraint and properties" page.
type Subject struct {
	// Optional. The user name/ID that the subject represents.
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// $hide_from_docs
	// Optional. The group that the subject belongs to.
	Group string `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	// Optional. The set of properties that identify the subject.
	// The above ServiceRoleBinding example shows an example of property "source.namespace".
	Properties           map[string]string `protobuf:"bytes,3,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Subject) Reset()         { *m = Subject{} }
func (m *Subject) String() string { return proto.CompactTextString(m) }
func (*Subject) ProtoMessage()    {}
func (*Subject) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_20b3f085115066c7, []int{3}
}
func (m *Subject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subject.Unmarshal(m, b)
}
func (m *Subject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subject.Marshal(b, m, deterministic)
}
func (dst *Subject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subject.Merge(dst, src)
}
func (m *Subject) XXX_Size() int {
	return xxx_messageInfo_Subject.Size(m)
}
func (m *Subject) XXX_DiscardUnknown() {
	xxx_messageInfo_Subject.DiscardUnknown(m)
}

var xxx_messageInfo_Subject proto.InternalMessageInfo

func (m *Subject) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Subject) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Subject) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

// RoleRef refers to a role object.
type RoleRef struct {
	// Required. The type of the role being referenced.
	// Currently, "ServiceRole" is the only supported value for "kind".
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Required. The name of the ServiceRole object being referenced.
	// The ServiceRole object must be in the same namespace as the ServiceRoleBinding
	// object.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleRef) Reset()         { *m = RoleRef{} }
func (m *RoleRef) String() string { return proto.CompactTextString(m) }
func (*RoleRef) ProtoMessage()    {}
func (*RoleRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_20b3f085115066c7, []int{4}
}
func (m *RoleRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleRef.Unmarshal(m, b)
}
func (m *RoleRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleRef.Marshal(b, m, deterministic)
}
func (dst *RoleRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleRef.Merge(dst, src)
}
func (m *RoleRef) XXX_Size() int {
	return xxx_messageInfo_RoleRef.Size(m)
}
func (m *RoleRef) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleRef.DiscardUnknown(m)
}

var xxx_messageInfo_RoleRef proto.InternalMessageInfo

func (m *RoleRef) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *RoleRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// RbacConfig defines the global config to control Istio RBAC behavior.
// This Custom Resource is a singleton where only one Custom Resource should be created globally in
// the mesh and the namespace should be the same to other Istio components, which usually is istio-system.
// Note: This is enforced in both istioctl and server side, new Custom Resource will be rejected if found any
// existing one, the user should either delete the existing one or change the existing one directly.
//
// Below is an example of RbacConfig object "istio-rbac-config" which enables Istio RBAC for all
// services in the default namespace.
//
// ```yaml
// apiVersion: "rbac.istio.io/v1alpha1"
// kind: RbacConfig
// metadata:
//   name: default
//   namespace: istio-system
// spec:
//   mode: ON_WITH_INCLUSION
//   inclusion:
//     namespaces: [ "default" ]
// ```
// @solo-kit:resource.short_name=rbaccfg
// @solo-kit:resource.plural_name=rbac_configs
// @solo-kit:resource.resource_groups=rbac.istio.io
type RbacConfig struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata" json:"metadata"`
	// Istio RBAC mode.
	Mode RbacConfig_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=istio.rbac.v1alpha1.RbacConfig_Mode" json:"mode,omitempty"`
	// A list of services or namespaces that should be enforced by Istio RBAC policies. Note: This field have
	// effect only when mode is ON_WITH_INCLUSION and will be ignored for any other modes.
	Inclusion *RbacConfig_Target `protobuf:"bytes,2,opt,name=inclusion" json:"inclusion,omitempty"`
	// A list of services or namespaces that should not be enforced by Istio RBAC policies. Note: This field have
	// effect only when mode is ON_WITH_EXCLUSION and will be ignored for any other modes.
	Exclusion *RbacConfig_Target `protobuf:"bytes,3,opt,name=exclusion" json:"exclusion,omitempty"`
	// $hide_from_docs
	// Indicates enforcement mode of the RbacConfig, in ENFORCED mode by default.
	// It's used to verify new RbacConfig work as expected before rolling to production.
	// When setting as PERMISSIVE, RBAC isn't enforced and has no impact on users.
	// RBAC engine run RbacConfig in PERMISSIVE mode and logs stats.
	// Invalid to set RbacConfig in PERMISSIVE and ServiceRoleBinding in ENFORCED mode.
	EnforcementMode      EnforcementMode `protobuf:"varint,4,opt,name=enforcement_mode,json=enforcementMode,proto3,enum=istio.rbac.v1alpha1.EnforcementMode" json:"enforcement_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RbacConfig) Reset()         { *m = RbacConfig{} }
func (m *RbacConfig) String() string { return proto.CompactTextString(m) }
func (*RbacConfig) ProtoMessage()    {}
func (*RbacConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_20b3f085115066c7, []int{5}
}
func (m *RbacConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RbacConfig.Unmarshal(m, b)
}
func (m *RbacConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RbacConfig.Marshal(b, m, deterministic)
}
func (dst *RbacConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RbacConfig.Merge(dst, src)
}
func (m *RbacConfig) XXX_Size() int {
	return xxx_messageInfo_RbacConfig.Size(m)
}
func (m *RbacConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RbacConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RbacConfig proto.InternalMessageInfo

func (m *RbacConfig) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *RbacConfig) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *RbacConfig) GetMode() RbacConfig_Mode {
	if m != nil {
		return m.Mode
	}
	return RbacConfig_OFF
}

func (m *RbacConfig) GetInclusion() *RbacConfig_Target {
	if m != nil {
		return m.Inclusion
	}
	return nil
}

func (m *RbacConfig) GetExclusion() *RbacConfig_Target {
	if m != nil {
		return m.Exclusion
	}
	return nil
}

func (m *RbacConfig) GetEnforcementMode() EnforcementMode {
	if m != nil {
		return m.EnforcementMode
	}
	return EnforcementMode_ENFORCED
}

// Target defines a list of services or namespaces.
type RbacConfig_Target struct {
	// A list of services.
	Services []string `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
	// A list of namespaces.
	Namespaces           []string `protobuf:"bytes,2,rep,name=namespaces" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RbacConfig_Target) Reset()         { *m = RbacConfig_Target{} }
func (m *RbacConfig_Target) String() string { return proto.CompactTextString(m) }
func (*RbacConfig_Target) ProtoMessage()    {}
func (*RbacConfig_Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_20b3f085115066c7, []int{5, 0}
}
func (m *RbacConfig_Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RbacConfig_Target.Unmarshal(m, b)
}
func (m *RbacConfig_Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RbacConfig_Target.Marshal(b, m, deterministic)
}
func (dst *RbacConfig_Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RbacConfig_Target.Merge(dst, src)
}
func (m *RbacConfig_Target) XXX_Size() int {
	return xxx_messageInfo_RbacConfig_Target.Size(m)
}
func (m *RbacConfig_Target) XXX_DiscardUnknown() {
	xxx_messageInfo_RbacConfig_Target.DiscardUnknown(m)
}

var xxx_messageInfo_RbacConfig_Target proto.InternalMessageInfo

func (m *RbacConfig_Target) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *RbacConfig_Target) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceRole)(nil), "istio.rbac.v1alpha1.ServiceRole")
	proto.RegisterType((*AccessRule)(nil), "istio.rbac.v1alpha1.AccessRule")
	proto.RegisterType((*AccessRule_Constraint)(nil), "istio.rbac.v1alpha1.AccessRule.Constraint")
	proto.RegisterType((*ServiceRoleBinding)(nil), "istio.rbac.v1alpha1.ServiceRoleBinding")
	proto.RegisterType((*Subject)(nil), "istio.rbac.v1alpha1.Subject")
	proto.RegisterMapType((map[string]string)(nil), "istio.rbac.v1alpha1.Subject.PropertiesEntry")
	proto.RegisterType((*RoleRef)(nil), "istio.rbac.v1alpha1.RoleRef")
	proto.RegisterType((*RbacConfig)(nil), "istio.rbac.v1alpha1.RbacConfig")
	proto.RegisterType((*RbacConfig_Target)(nil), "istio.rbac.v1alpha1.RbacConfig.Target")
	proto.RegisterEnum("istio.rbac.v1alpha1.EnforcementMode", EnforcementMode_name, EnforcementMode_value)
	proto.RegisterEnum("istio.rbac.v1alpha1.RbacConfig_Mode", RbacConfig_Mode_name, RbacConfig_Mode_value)
}

func init() { proto.RegisterFile("rbac.proto", fileDescriptor_rbac_20b3f085115066c7) }

var fileDescriptor_rbac_20b3f085115066c7 = []byte{
	// 752 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdb, 0x8e, 0xda, 0x46,
	0x18, 0x5e, 0x03, 0xcb, 0xe1, 0xa7, 0xca, 0x92, 0x29, 0x8d, 0x5c, 0x54, 0x75, 0x91, 0x55, 0x55,
	0x28, 0x6a, 0x6d, 0x41, 0xd5, 0x08, 0x45, 0xea, 0x45, 0x60, 0xbd, 0x2d, 0xd2, 0x2e, 0x44, 0x43,
	0x7a, 0x50, 0x6f, 0x56, 0x83, 0x3d, 0x98, 0x29, 0xc6, 0x63, 0xcd, 0x8c, 0x51, 0xf2, 0x5c, 0x7d,
	0x81, 0x4a, 0xbd, 0x68, 0x9f, 0x62, 0x1f, 0xa2, 0x4f, 0x50, 0x79, 0x6c, 0x03, 0x89, 0xd0, 0x6e,
	0xf6, 0x6e, 0xaf, 0xf8, 0x4f, 0xdf, 0xc7, 0xf7, 0x1f, 0xe4, 0x01, 0x10, 0x0b, 0xe2, 0xd9, 0xb1,
	0xe0, 0x8a, 0xa3, 0x4f, 0x99, 0x54, 0x8c, 0xdb, 0x3a, 0xb2, 0xed, 0x93, 0x30, 0x5e, 0x91, 0x7e,
	0xa7, 0x1f, 0x30, 0xb5, 0x4a, 0x16, 0xb6, 0xc7, 0x37, 0x8e, 0xe4, 0x21, 0xff, 0x96, 0xf1, 0xec,
	0x77, 0xcd, 0x94, 0x43, 0x62, 0xe6, 0x6c, 0xfb, 0xce, 0x86, 0x2a, 0xe2, 0x13, 0x45, 0x32, 0x9e,
	0x8e, 0xf3, 0x11, 0x10, 0xa9, 0x88, 0x4a, 0x64, 0x0e, 0x68, 0x07, 0x3c, 0xe0, 0xda, 0x74, 0x52,
	0x2b, 0x8b, 0x5a, 0xff, 0x18, 0xd0, 0x9c, 0x53, 0xb1, 0x65, 0x1e, 0xc5, 0x3c, 0xa4, 0xe8, 0x47,
	0xa8, 0x66, 0x28, 0xd3, 0xef, 0x1a, 0xbd, 0xe6, 0xa0, 0x6d, 0x7b, 0x5c, 0x50, 0x3b, 0x65, 0xb6,
	0x19, 0xb7, 0xe7, 0x3a, 0x37, 0xfa, 0xfc, 0xdf, 0xdb, 0xf3, 0x93, 0xff, 0x6e, 0xcf, 0x9f, 0x2a,
	0x2a, 0x95, 0xcf, 0x96, 0xcb, 0x97, 0x16, 0x0b, 0x22, 0x2e, 0xa8, 0x85, 0x73, 0x38, 0x1a, 0x42,
	0xbd, 0x50, 0x6c, 0x52, 0x4d, 0xf5, 0xec, 0x7d, 0xaa, 0xeb, 0x3c, 0x3b, 0xaa, 0xa4, 0x64, 0x78,
	0x57, 0x8d, 0xbe, 0x87, 0x53, 0x91, 0x84, 0x54, 0x9a, 0x46, 0xb7, 0xdc, 0x6b, 0x0e, 0xce, 0xed,
	0x23, 0x13, 0xb3, 0x5f, 0x79, 0x1e, 0x95, 0x12, 0x27, 0x21, 0xc5, 0x59, 0xb5, 0x75, 0x6b, 0x00,
	0xec, 0xa3, 0xa8, 0x03, 0x75, 0x99, 0xf5, 0x95, 0x11, 0x35, 0xf0, 0xce, 0x47, 0x6d, 0x38, 0x8d,
	0x89, 0x5a, 0x49, 0xb3, 0xa4, 0x13, 0x99, 0x83, 0x4c, 0xa8, 0x6d, 0xa8, 0x5a, 0x71, 0x5f, 0x9a,
	0x65, 0x1d, 0x2f, 0x5c, 0x74, 0x05, 0x4d, 0x8f, 0x47, 0x52, 0x09, 0xc2, 0x22, 0x25, 0xcd, 0x8a,
	0xd6, 0xf5, 0xfc, 0x1e, 0x5d, 0xf6, 0x78, 0x07, 0xc1, 0x87, 0xf0, 0xce, 0x0b, 0x80, 0x7d, 0x0a,
	0xb5, 0xa0, 0xbc, 0xa6, 0xef, 0x4c, 0xa3, 0x6b, 0xf4, 0x1a, 0x38, 0x35, 0xd1, 0x33, 0xa8, 0x6e,
	0x49, 0x98, 0xd0, 0x42, 0x5e, 0xee, 0x59, 0x7f, 0x97, 0x00, 0x1d, 0xac, 0x6a, 0xc4, 0x22, 0x9f,
	0x45, 0xc1, 0x63, 0xd8, 0xd8, 0x10, 0xea, 0x32, 0x59, 0xfc, 0x41, 0x3d, 0x55, 0x2c, 0xed, 0x8b,
	0xa3, 0xc3, 0x99, 0x67, 0x45, 0x78, 0x57, 0x8d, 0x5e, 0x40, 0x4d, 0xf0, 0x90, 0x62, 0xba, 0x34,
	0x4b, 0xfa, 0x2f, 0x8f, 0x03, 0x71, 0x56, 0x83, 0x8b, 0x62, 0x34, 0x84, 0xca, 0x86, 0xfb, 0xd4,
	0x2c, 0x77, 0x8d, 0xde, 0x93, 0xc1, 0x57, 0x47, 0x41, 0x6e, 0xb4, 0xe4, 0xc2, 0xa3, 0x1b, 0x1a,
	0xa9, 0x6b, 0xee, 0x53, 0xac, 0x11, 0xd6, 0x5f, 0x06, 0xd4, 0x72, 0x1d, 0x08, 0x41, 0x25, 0x91,
	0x54, 0xe4, 0xc3, 0xd7, 0x76, 0x7a, 0x1b, 0x81, 0xe0, 0x49, 0xac, 0xf5, 0x34, 0x70, 0xe6, 0xa0,
	0x2b, 0x80, 0x58, 0xf0, 0x98, 0x0a, 0xc5, 0x68, 0x76, 0x1e, 0xcd, 0xc1, 0x37, 0x77, 0xf5, 0x68,
	0xbf, 0xde, 0x95, 0xbb, 0x91, 0x12, 0xef, 0xf0, 0x01, 0xbe, 0xf3, 0x03, 0x9c, 0x7d, 0x90, 0x3e,
	0x72, 0x06, 0x6d, 0x38, 0xd5, 0x8b, 0x2f, 0x84, 0x68, 0xe7, 0x65, 0x69, 0x68, 0x58, 0x7d, 0xa8,
	0xe5, 0x03, 0x49, 0x3b, 0x58, 0xb3, 0xc8, 0x2f, 0x3a, 0x48, 0xed, 0x34, 0x16, 0x91, 0x4d, 0x81,
	0xd3, 0xb6, 0xf5, 0x67, 0x05, 0x00, 0x2f, 0x88, 0x37, 0xe6, 0xd1, 0x92, 0x3d, 0x92, 0x9b, 0xc9,
	0x36, 0x68, 0xdc, 0xb1, 0xc1, 0xbd, 0x62, 0x7b, 0xbf, 0x41, 0x74, 0x01, 0x0d, 0x16, 0x79, 0x61,
	0x22, 0x19, 0x8f, 0xf2, 0xab, 0xf9, 0xfa, 0x3e, 0xf8, 0x1b, 0x22, 0x02, 0xaa, 0xf0, 0x1e, 0x98,
	0xb2, 0xd0, 0xb7, 0x05, 0x4b, 0xf9, 0x61, 0x2c, 0x3b, 0x20, 0x9a, 0x41, 0x8b, 0xee, 0xcf, 0xec,
	0x46, 0x77, 0x54, 0x79, 0xc0, 0x4d, 0x9e, 0xd1, 0xf7, 0x03, 0x9d, 0x0b, 0xa8, 0x66, 0xff, 0x72,
	0xe7, 0x07, 0xec, 0x4b, 0x80, 0x74, 0xad, 0x32, 0x26, 0xde, 0xee, 0x33, 0x71, 0x10, 0xb1, 0x5c,
	0xa8, 0xa4, 0x6c, 0xa8, 0x06, 0xe5, 0xd9, 0xe5, 0x65, 0xeb, 0x04, 0x55, 0xa1, 0x34, 0x9b, 0xb6,
	0x0c, 0xf4, 0x19, 0x3c, 0x9d, 0x4d, 0x6f, 0x7e, 0x9d, 0xbc, 0xf9, 0xe9, 0x66, 0x32, 0x1d, 0x5f,
	0xfd, 0x3c, 0x9f, 0xcc, 0xa6, 0xad, 0xd2, 0x61, 0xd8, 0xfd, 0xad, 0x08, 0x97, 0x9f, 0x3b, 0x70,
	0xf6, 0x81, 0x60, 0xf4, 0x09, 0xd4, 0xdd, 0xe9, 0xe5, 0x0c, 0x8f, 0xdd, 0x8b, 0xd6, 0x09, 0x7a,
	0x02, 0xf0, 0xda, 0xc5, 0xd7, 0x93, 0xf9, 0x7c, 0xf2, 0x8b, 0xdb, 0x32, 0x46, 0xe3, 0xdf, 0x5f,
	0x1d, 0x7b, 0x96, 0x92, 0x98, 0x8a, 0x20, 0xe4, 0xdc, 0x89, 0xd7, 0x81, 0x7e, 0x9b, 0xe8, 0x5b,
	0x45, 0x45, 0x44, 0x42, 0x47, 0xcf, 0xc8, 0x49, 0x67, 0xe4, 0x14, 0x33, 0x5a, 0x54, 0xf5, 0xcb,
	0xf4, 0xdd, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x92, 0x1e, 0x60, 0x76, 0x36, 0x07, 0x00, 0x00,
}
