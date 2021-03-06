// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/certificates/v1/vault_ca.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

type VaultCA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `ca_path` is the mount path of the Vault PKI backend's `sign` endpoint, e.g:
	// "my_pki_mount/sign/my-role-name".
	CaPath string `protobuf:"bytes,1,opt,name=ca_path,json=caPath,proto3" json:"ca_path,omitempty"`
	// `csr_path` is the mount path of the Vault PKI backend's `generate` endpoint, e.g:
	// "my_pki_mount/intermediate/generate/exported".
	// "exported" is necessary here as istio needs access to the private key
	// See vault docs here: https://www.vaultproject.io/api-docs/secret/pki#parameters-4
	CsrPath string `protobuf:"bytes,2,opt,name=csr_path,json=csrPath,proto3" json:"csr_path,omitempty"`
	// Server is the connection address for the Vault server, e.g: "https://vault.example.com:8200".
	Server string `protobuf:"bytes,3,opt,name=server,proto3" json:"server,omitempty"`
	// PEM encoded CA bundle used to validate Vault server certificate. Only used
	// if the Server URL is using HTTPS protocol. This parameter is ignored for
	// plain HTTP protocol connection. If not set the system root certificates
	// are used to validate the TLS connection.
	CaBundle []byte `protobuf:"bytes,4,opt,name=ca_bundle,json=caBundle,proto3" json:"ca_bundle,omitempty"`
	// Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: "ns1"
	// More about namespaces can be found [here](https://www.vaultproject.io/docs/enterprise/namespaces)
	Namespace string `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Types that are assignable to AuthType:
	//	*VaultCA_TokenSecretRef
	//	*VaultCA_KubernetesAuth
	AuthType isVaultCA_AuthType `protobuf_oneof:"auth_type"`
}

func (x *VaultCA) Reset() {
	*x = VaultCA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultCA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultCA) ProtoMessage() {}

func (x *VaultCA) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultCA.ProtoReflect.Descriptor instead.
func (*VaultCA) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDescGZIP(), []int{0}
}

func (x *VaultCA) GetCaPath() string {
	if x != nil {
		return x.CaPath
	}
	return ""
}

func (x *VaultCA) GetCsrPath() string {
	if x != nil {
		return x.CsrPath
	}
	return ""
}

func (x *VaultCA) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *VaultCA) GetCaBundle() []byte {
	if x != nil {
		return x.CaBundle
	}
	return nil
}

func (x *VaultCA) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (m *VaultCA) GetAuthType() isVaultCA_AuthType {
	if m != nil {
		return m.AuthType
	}
	return nil
}

func (x *VaultCA) GetTokenSecretRef() *v1.ObjectRef {
	if x, ok := x.GetAuthType().(*VaultCA_TokenSecretRef); ok {
		return x.TokenSecretRef
	}
	return nil
}

func (x *VaultCA) GetKubernetesAuth() *VaultKubernetesAuth {
	if x, ok := x.GetAuthType().(*VaultCA_KubernetesAuth); ok {
		return x.KubernetesAuth
	}
	return nil
}

type isVaultCA_AuthType interface {
	isVaultCA_AuthType()
}

type VaultCA_TokenSecretRef struct {
	// TokenSecretRef authenticates with Vault by presenting a token.
	TokenSecretRef *v1.ObjectRef `protobuf:"bytes,6,opt,name=token_secret_ref,json=tokenSecretRef,proto3,oneof"`
}

type VaultCA_KubernetesAuth struct {
	// Kubernetes authenticates with Vault by passing the ServiceAccount
	// token stored in the named Secret resource to the Vault server.
	KubernetesAuth *VaultKubernetesAuth `protobuf:"bytes,8,opt,name=kubernetes_auth,json=kubernetesAuth,proto3,oneof"`
}

func (*VaultCA_TokenSecretRef) isVaultCA_AuthType() {}

func (*VaultCA_KubernetesAuth) isVaultCA_AuthType() {}

type VaultKubernetesAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Vault mountPath here is the mount path to use when authenticating with
	// Vault. For example, setting a value to `/v1/auth/foo`, will use the path
	// `/v1/auth/foo/login` to authenticate with Vault. If unspecified, the
	// default value "/v1/auth/kubernetes" will be used.
	MountPath string `protobuf:"bytes,1,opt,name=mount_path,json=mountPath,proto3" json:"mount_path,omitempty"`
	// A required field containing the Vault Role to assume. A Role binds a
	// Kubernetes ServiceAccount with a set of Vault policies.
	Role string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	// Key to search for the sa_token
	// Default to "token"
	SecretTokenKey string `protobuf:"bytes,3,opt,name=secret_token_key,json=secretTokenKey,proto3" json:"secret_token_key,omitempty"`
	// The method by which to get the service account token.
	// If unspecified will default to mounted_sa_path
	//
	// Types that are assignable to ServiceAccountLocation:
	//	*VaultKubernetesAuth_ServiceAccountRef
	//	*VaultKubernetesAuth_MountedSaPath
	ServiceAccountLocation isVaultKubernetesAuth_ServiceAccountLocation `protobuf_oneof:"service_account_location"`
}

func (x *VaultKubernetesAuth) Reset() {
	*x = VaultKubernetesAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultKubernetesAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultKubernetesAuth) ProtoMessage() {}

func (x *VaultKubernetesAuth) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultKubernetesAuth.ProtoReflect.Descriptor instead.
func (*VaultKubernetesAuth) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDescGZIP(), []int{1}
}

func (x *VaultKubernetesAuth) GetMountPath() string {
	if x != nil {
		return x.MountPath
	}
	return ""
}

func (x *VaultKubernetesAuth) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *VaultKubernetesAuth) GetSecretTokenKey() string {
	if x != nil {
		return x.SecretTokenKey
	}
	return ""
}

func (m *VaultKubernetesAuth) GetServiceAccountLocation() isVaultKubernetesAuth_ServiceAccountLocation {
	if m != nil {
		return m.ServiceAccountLocation
	}
	return nil
}

func (x *VaultKubernetesAuth) GetServiceAccountRef() *v1.ObjectRef {
	if x, ok := x.GetServiceAccountLocation().(*VaultKubernetesAuth_ServiceAccountRef); ok {
		return x.ServiceAccountRef
	}
	return nil
}

func (x *VaultKubernetesAuth) GetMountedSaPath() string {
	if x, ok := x.GetServiceAccountLocation().(*VaultKubernetesAuth_MountedSaPath); ok {
		return x.MountedSaPath
	}
	return ""
}

type isVaultKubernetesAuth_ServiceAccountLocation interface {
	isVaultKubernetesAuth_ServiceAccountLocation()
}

type VaultKubernetesAuth_ServiceAccountRef struct {
	// Reference to service account, other than the one mounted to the current pod.
	ServiceAccountRef *v1.ObjectRef `protobuf:"bytes,4,opt,name=service_account_ref,json=serviceAccountRef,proto3,oneof"`
}

type VaultKubernetesAuth_MountedSaPath struct {
	// File System path to grab the service account token from.
	// Defaults to /var/run/secrets/kubernetes.io/serviceaccount
	MountedSaPath string `protobuf:"bytes,5,opt,name=mounted_sa_path,json=mountedSaPath,proto3,oneof"`
}

func (*VaultKubernetesAuth_ServiceAccountRef) isVaultKubernetesAuth_ServiceAccountLocation() {}

func (*VaultKubernetesAuth_MountedSaPath) isVaultKubernetesAuth_ServiceAccountLocation() {}

var File_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x02, 0x0a, 0x07, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x43,
	0x41, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x50, 0x61, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x73,
	0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x73,
	0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x61, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x63, 0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x10, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66,
	0x48, 0x00, 0x52, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52,
	0x65, 0x66, 0x12, 0x5e, 0x0a, 0x0f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x41, 0x75, 0x74, 0x68,
	0x48, 0x00, 0x52, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x41, 0x75,
	0x74, 0x68, 0x42, 0x0b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x88, 0x02, 0x0a, 0x13, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x4e, 0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x48,
	0x00, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x66, 0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x5f,
	0x73, 0x61, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0d, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x53, 0x61, 0x50, 0x61, 0x74, 0x68, 0x42, 0x1a,
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x4c, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x76, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_goTypes = []interface{}{
	(*VaultCA)(nil),             // 0: certificates.mesh.gloo.solo.io.VaultCA
	(*VaultKubernetesAuth)(nil), // 1: certificates.mesh.gloo.solo.io.VaultKubernetesAuth
	(*v1.ObjectRef)(nil),        // 2: core.skv2.solo.io.ObjectRef
}
var file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_depIdxs = []int32{
	2, // 0: certificates.mesh.gloo.solo.io.VaultCA.token_secret_ref:type_name -> core.skv2.solo.io.ObjectRef
	1, // 1: certificates.mesh.gloo.solo.io.VaultCA.kubernetes_auth:type_name -> certificates.mesh.gloo.solo.io.VaultKubernetesAuth
	2, // 2: certificates.mesh.gloo.solo.io.VaultKubernetesAuth.service_account_ref:type_name -> core.skv2.solo.io.ObjectRef
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultCA); i {
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
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultKubernetesAuth); i {
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
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*VaultCA_TokenSecretRef)(nil),
		(*VaultCA_KubernetesAuth)(nil),
	}
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*VaultKubernetesAuth_ServiceAccountRef)(nil),
		(*VaultKubernetesAuth_MountedSaPath)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_depIdxs = nil
}
