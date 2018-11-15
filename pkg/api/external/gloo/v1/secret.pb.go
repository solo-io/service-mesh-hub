// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: secret.proto

package v1 // import "github.com/solo-io/supergloo/pkg/api/external/gloo/v1"

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

//
// @solo-kit:resource.short_name=sec
// @solo-kit:resource.plural_name=secrets
// @solo-kit:resource.resource_groups=api.gloo.solo.io,discovery.gloo.solo.io,translator.supergloo.solo.io
//
// Certain plugins such as the AWS Lambda Plugin require the use of secrets for authentication, configuration of SSL Certificates, and other data that should not be stored in plaintext configuration.
//
// Gloo runs an independent (goroutine) controller to monitor secrets. Secrets are stored in their own secret storage layer. Gloo can monitor secrets stored in the following secret storage services:
//
// Kubernetes Secrets
// Hashicorp Vault
// Plaintext files (recommended only for testing)
// Secrets must adhere to a structure, specified by the plugin that requires them.
//
// Gloo's secret backend can be configured in Gloo's bootstrap options
type Secret struct {
	// Types that are valid to be assigned to Kind:
	//	*Secret_Aws
	//	*Secret_Azure
	//	*Secret_Tls
	//	*Secret_Cacerts
	Kind isSecret_Kind `protobuf_oneof:"kind"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_secret_ab614d07c9ef3898, []int{0}
}
func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (dst *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(dst, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

type isSecret_Kind interface {
	isSecret_Kind()
	Equal(interface{}) bool
}

type Secret_Aws struct {
	Aws *AwsSecret `protobuf:"bytes,1,opt,name=aws,oneof"`
}
type Secret_Azure struct {
	Azure *AzureSecret `protobuf:"bytes,2,opt,name=azure,oneof"`
}
type Secret_Tls struct {
	Tls *TlsSecret `protobuf:"bytes,3,opt,name=tls,oneof"`
}
type Secret_Cacerts struct {
	Cacerts *IstioCacertsSecret `protobuf:"bytes,4,opt,name=cacerts,oneof"`
}

func (*Secret_Aws) isSecret_Kind()     {}
func (*Secret_Azure) isSecret_Kind()   {}
func (*Secret_Tls) isSecret_Kind()     {}
func (*Secret_Cacerts) isSecret_Kind() {}

func (m *Secret) GetKind() isSecret_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Secret) GetAws() *AwsSecret {
	if x, ok := m.GetKind().(*Secret_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *Secret) GetAzure() *AzureSecret {
	if x, ok := m.GetKind().(*Secret_Azure); ok {
		return x.Azure
	}
	return nil
}

func (m *Secret) GetTls() *TlsSecret {
	if x, ok := m.GetKind().(*Secret_Tls); ok {
		return x.Tls
	}
	return nil
}

func (m *Secret) GetCacerts() *IstioCacertsSecret {
	if x, ok := m.GetKind().(*Secret_Cacerts); ok {
		return x.Cacerts
	}
	return nil
}

func (m *Secret) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Secret) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Secret_OneofMarshaler, _Secret_OneofUnmarshaler, _Secret_OneofSizer, []interface{}{
		(*Secret_Aws)(nil),
		(*Secret_Azure)(nil),
		(*Secret_Tls)(nil),
		(*Secret_Cacerts)(nil),
	}
}

func _Secret_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Secret)
	// kind
	switch x := m.Kind.(type) {
	case *Secret_Aws:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aws); err != nil {
			return err
		}
	case *Secret_Azure:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Azure); err != nil {
			return err
		}
	case *Secret_Tls:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tls); err != nil {
			return err
		}
	case *Secret_Cacerts:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Cacerts); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Secret.Kind has unexpected type %T", x)
	}
	return nil
}

func _Secret_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Secret)
	switch tag {
	case 1: // kind.aws
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AwsSecret)
		err := b.DecodeMessage(msg)
		m.Kind = &Secret_Aws{msg}
		return true, err
	case 2: // kind.azure
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AzureSecret)
		err := b.DecodeMessage(msg)
		m.Kind = &Secret_Azure{msg}
		return true, err
	case 3: // kind.tls
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TlsSecret)
		err := b.DecodeMessage(msg)
		m.Kind = &Secret_Tls{msg}
		return true, err
	case 4: // kind.cacerts
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IstioCacertsSecret)
		err := b.DecodeMessage(msg)
		m.Kind = &Secret_Cacerts{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Secret_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Secret)
	// kind
	switch x := m.Kind.(type) {
	case *Secret_Aws:
		s := proto.Size(x.Aws)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Secret_Azure:
		s := proto.Size(x.Azure)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Secret_Tls:
		s := proto.Size(x.Tls)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Secret_Cacerts:
		s := proto.Size(x.Cacerts)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AwsSecret struct {
	AccessKey            string   `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey            string   `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsSecret) Reset()         { *m = AwsSecret{} }
func (m *AwsSecret) String() string { return proto.CompactTextString(m) }
func (*AwsSecret) ProtoMessage()    {}
func (*AwsSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_secret_ab614d07c9ef3898, []int{1}
}
func (m *AwsSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsSecret.Unmarshal(m, b)
}
func (m *AwsSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsSecret.Marshal(b, m, deterministic)
}
func (dst *AwsSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsSecret.Merge(dst, src)
}
func (m *AwsSecret) XXX_Size() int {
	return xxx_messageInfo_AwsSecret.Size(m)
}
func (m *AwsSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsSecret.DiscardUnknown(m)
}

var xxx_messageInfo_AwsSecret proto.InternalMessageInfo

func (m *AwsSecret) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *AwsSecret) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

type AzureSecret struct {
	ApiKeys              map[string]string `protobuf:"bytes,1,rep,name=api_keys,json=apiKeys" json:"api_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AzureSecret) Reset()         { *m = AzureSecret{} }
func (m *AzureSecret) String() string { return proto.CompactTextString(m) }
func (*AzureSecret) ProtoMessage()    {}
func (*AzureSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_secret_ab614d07c9ef3898, []int{2}
}
func (m *AzureSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureSecret.Unmarshal(m, b)
}
func (m *AzureSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureSecret.Marshal(b, m, deterministic)
}
func (dst *AzureSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureSecret.Merge(dst, src)
}
func (m *AzureSecret) XXX_Size() int {
	return xxx_messageInfo_AzureSecret.Size(m)
}
func (m *AzureSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureSecret.DiscardUnknown(m)
}

var xxx_messageInfo_AzureSecret proto.InternalMessageInfo

func (m *AzureSecret) GetApiKeys() map[string]string {
	if m != nil {
		return m.ApiKeys
	}
	return nil
}

type TlsSecret struct {
	CertChain            string   `protobuf:"bytes,1,opt,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	PrivateKey           string   `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	RootCa               string   `protobuf:"bytes,3,opt,name=root_ca,json=rootCa,proto3" json:"root_ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TlsSecret) Reset()         { *m = TlsSecret{} }
func (m *TlsSecret) String() string { return proto.CompactTextString(m) }
func (*TlsSecret) ProtoMessage()    {}
func (*TlsSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_secret_ab614d07c9ef3898, []int{3}
}
func (m *TlsSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsSecret.Unmarshal(m, b)
}
func (m *TlsSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsSecret.Marshal(b, m, deterministic)
}
func (dst *TlsSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsSecret.Merge(dst, src)
}
func (m *TlsSecret) XXX_Size() int {
	return xxx_messageInfo_TlsSecret.Size(m)
}
func (m *TlsSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsSecret.DiscardUnknown(m)
}

var xxx_messageInfo_TlsSecret proto.InternalMessageInfo

func (m *TlsSecret) GetCertChain() string {
	if m != nil {
		return m.CertChain
	}
	return ""
}

func (m *TlsSecret) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *TlsSecret) GetRootCa() string {
	if m != nil {
		return m.RootCa
	}
	return ""
}

// Specifically structured cert that istio uses for non-default root certificates
type IstioCacertsSecret struct {
	RootCert             string   `protobuf:"bytes,1,opt,name=root_cert,json=root-cert.pem,proto3" json:"root_cert,omitempty"`
	CertChain            string   `protobuf:"bytes,2,opt,name=cert_chain,json=cert-chain.pem,proto3" json:"cert_chain,omitempty"`
	CaCert               string   `protobuf:"bytes,3,opt,name=ca_cert,json=ca-cert.pem,proto3" json:"ca_cert,omitempty"`
	CaKey                string   `protobuf:"bytes,4,opt,name=ca_key,json=ca-key.pem,proto3" json:"ca_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioCacertsSecret) Reset()         { *m = IstioCacertsSecret{} }
func (m *IstioCacertsSecret) String() string { return proto.CompactTextString(m) }
func (*IstioCacertsSecret) ProtoMessage()    {}
func (*IstioCacertsSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_secret_ab614d07c9ef3898, []int{4}
}
func (m *IstioCacertsSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioCacertsSecret.Unmarshal(m, b)
}
func (m *IstioCacertsSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioCacertsSecret.Marshal(b, m, deterministic)
}
func (dst *IstioCacertsSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioCacertsSecret.Merge(dst, src)
}
func (m *IstioCacertsSecret) XXX_Size() int {
	return xxx_messageInfo_IstioCacertsSecret.Size(m)
}
func (m *IstioCacertsSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioCacertsSecret.DiscardUnknown(m)
}

var xxx_messageInfo_IstioCacertsSecret proto.InternalMessageInfo

func (m *IstioCacertsSecret) GetRootCert() string {
	if m != nil {
		return m.RootCert
	}
	return ""
}

func (m *IstioCacertsSecret) GetCertChain() string {
	if m != nil {
		return m.CertChain
	}
	return ""
}

func (m *IstioCacertsSecret) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func (m *IstioCacertsSecret) GetCaKey() string {
	if m != nil {
		return m.CaKey
	}
	return ""
}

func init() {
	proto.RegisterType((*Secret)(nil), "gloo.solo.io.Secret")
	proto.RegisterType((*AwsSecret)(nil), "gloo.solo.io.AwsSecret")
	proto.RegisterType((*AzureSecret)(nil), "gloo.solo.io.AzureSecret")
	proto.RegisterMapType((map[string]string)(nil), "gloo.solo.io.AzureSecret.ApiKeysEntry")
	proto.RegisterType((*TlsSecret)(nil), "gloo.solo.io.TlsSecret")
	proto.RegisterType((*IstioCacertsSecret)(nil), "gloo.solo.io.IstioCacertsSecret")
}
func (this *Secret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret)
	if !ok {
		that2, ok := that.(Secret)
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
	if that1.Kind == nil {
		if this.Kind != nil {
			return false
		}
	} else if this.Kind == nil {
		return false
	} else if !this.Kind.Equal(that1.Kind) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Secret_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Aws)
	if !ok {
		that2, ok := that.(Secret_Aws)
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
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *Secret_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Azure)
	if !ok {
		that2, ok := that.(Secret_Azure)
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
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}
func (this *Secret_Tls) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Tls)
	if !ok {
		that2, ok := that.(Secret_Tls)
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
	if !this.Tls.Equal(that1.Tls) {
		return false
	}
	return true
}
func (this *Secret_Cacerts) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Cacerts)
	if !ok {
		that2, ok := that.(Secret_Cacerts)
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
	if !this.Cacerts.Equal(that1.Cacerts) {
		return false
	}
	return true
}
func (this *AwsSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsSecret)
	if !ok {
		that2, ok := that.(AwsSecret)
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
	if this.AccessKey != that1.AccessKey {
		return false
	}
	if this.SecretKey != that1.SecretKey {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AzureSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AzureSecret)
	if !ok {
		that2, ok := that.(AzureSecret)
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
	if len(this.ApiKeys) != len(that1.ApiKeys) {
		return false
	}
	for i := range this.ApiKeys {
		if this.ApiKeys[i] != that1.ApiKeys[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TlsSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TlsSecret)
	if !ok {
		that2, ok := that.(TlsSecret)
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
	if this.CertChain != that1.CertChain {
		return false
	}
	if this.PrivateKey != that1.PrivateKey {
		return false
	}
	if this.RootCa != that1.RootCa {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IstioCacertsSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IstioCacertsSecret)
	if !ok {
		that2, ok := that.(IstioCacertsSecret)
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
	if this.RootCert != that1.RootCert {
		return false
	}
	if this.CertChain != that1.CertChain {
		return false
	}
	if this.CaCert != that1.CaCert {
		return false
	}
	if this.CaKey != that1.CaKey {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("secret.proto", fileDescriptor_secret_ab614d07c9ef3898) }

var fileDescriptor_secret_ab614d07c9ef3898 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0x5f, 0xda, 0xae, 0x5d, 0xbe, 0x16, 0x84, 0xac, 0x89, 0x85, 0x0a, 0x58, 0x95, 0x03, 0x9a,
	0x84, 0x9a, 0xa8, 0x43, 0x48, 0xd3, 0xe0, 0xd2, 0x56, 0x48, 0x4c, 0x15, 0x97, 0xc0, 0x89, 0x4b,
	0xe5, 0xb9, 0x56, 0x66, 0x25, 0xad, 0x23, 0xdb, 0xed, 0x08, 0xcf, 0xc0, 0x11, 0xde, 0x81, 0x47,
	0xe1, 0x29, 0x38, 0xf0, 0x24, 0xe8, 0xb3, 0xdb, 0x2c, 0x83, 0xed, 0x14, 0xfb, 0xf7, 0xe7, 0xfb,
	0x1b, 0x43, 0x4f, 0x73, 0xa6, 0xb8, 0x89, 0x0a, 0x25, 0x8d, 0x24, 0xbd, 0x34, 0x97, 0x32, 0xd2,
	0x32, 0x97, 0x91, 0x90, 0xfd, 0xc3, 0x54, 0xa6, 0xd2, 0x12, 0x31, 0x9e, 0x9c, 0xa6, 0x3f, 0x4a,
	0x85, 0xb9, 0x5a, 0x5f, 0x46, 0x4c, 0x2e, 0x63, 0x54, 0x0e, 0x85, 0x74, 0xdf, 0x4c, 0x98, 0x98,
	0x16, 0x22, 0xde, 0x8c, 0xe2, 0x25, 0x37, 0x74, 0x41, 0x0d, 0x75, 0x96, 0xf0, 0x47, 0x03, 0xda,
	0x1f, 0x6d, 0x1e, 0xf2, 0x12, 0x9a, 0xf4, 0x5a, 0x07, 0xde, 0xc0, 0x3b, 0xe9, 0x9e, 0x1e, 0x45,
	0xf5, 0x7c, 0xd1, 0xf8, 0x5a, 0x3b, 0xd5, 0xfb, 0xbd, 0x04, 0x55, 0x64, 0x04, 0xfb, 0xf4, 0xeb,
	0x5a, 0xf1, 0xa0, 0x61, 0xe5, 0x4f, 0xfe, 0x91, 0x23, 0x55, 0x19, 0x9c, 0x12, 0xe3, 0x9b, 0x5c,
	0x07, 0xcd, 0xbb, 0xe2, 0x7f, 0xca, 0x6b, 0xf1, 0x4d, 0xae, 0xc9, 0x5b, 0xe8, 0x30, 0xca, 0xb8,
	0x32, 0x3a, 0x68, 0x59, 0xc3, 0xe0, 0xb6, 0xe1, 0x42, 0x1b, 0x21, 0xa7, 0x4e, 0x51, 0x39, 0x77,
	0x16, 0x72, 0x06, 0x07, 0xbb, 0x3e, 0x83, 0x8e, 0xb5, 0x3f, 0x8e, 0x98, 0x54, 0xbc, 0xb2, 0x7f,
	0xd8, 0xb2, 0x93, 0xd6, 0xaf, 0xdf, 0xc7, 0x7b, 0x49, 0xa5, 0x9e, 0xb4, 0xa1, 0x95, 0x89, 0xd5,
	0x22, 0xbc, 0x00, 0xbf, 0xea, 0x99, 0x3c, 0x03, 0xa0, 0x8c, 0x71, 0xad, 0xe7, 0x19, 0x2f, 0xed,
	0x80, 0xfc, 0xc4, 0x77, 0xc8, 0x8c, 0x97, 0x48, 0xbb, 0x55, 0x59, 0xba, 0xe1, 0x68, 0x87, 0xcc,
	0x78, 0x19, 0x7e, 0xf3, 0xa0, 0x5b, 0x1b, 0x08, 0x19, 0xc3, 0x01, 0x2d, 0x04, 0x6a, 0x71, 0xd8,
	0xcd, 0x93, 0xee, 0xe9, 0x8b, 0x7b, 0xa7, 0x17, 0x8d, 0x0b, 0x31, 0xe3, 0xa5, 0x7e, 0xb7, 0x32,
	0xaa, 0x4c, 0x3a, 0xd4, 0xdd, 0xfa, 0xe7, 0xd0, 0xab, 0x13, 0xe4, 0x11, 0x34, 0x6f, 0x2a, 0xc3,
	0x23, 0x39, 0x84, 0xfd, 0x0d, 0xcd, 0xd7, 0x7c, 0x5b, 0x8e, 0xbb, 0x9c, 0x37, 0xce, 0xbc, 0x70,
	0x01, 0x7e, 0x35, 0x6d, 0x2c, 0x1d, 0x27, 0x36, 0x67, 0x57, 0x54, 0xac, 0x76, 0x9d, 0x21, 0x32,
	0x45, 0x80, 0x1c, 0x43, 0xb7, 0x50, 0x62, 0x43, 0x0d, 0xaf, 0xb5, 0x06, 0x5b, 0x08, 0x5b, 0x3f,
	0x82, 0x8e, 0x92, 0xd2, 0xcc, 0x19, 0xb5, 0x7b, 0xf5, 0x93, 0x36, 0x5e, 0xa7, 0x34, 0xfc, 0xee,
	0x01, 0xf9, 0x7f, 0x47, 0x64, 0x00, 0xbe, 0xd3, 0x73, 0x65, 0xb6, 0xe9, 0x1e, 0x20, 0x30, 0x44,
	0x20, 0x2a, 0xf8, 0x92, 0x84, 0xb7, 0x2a, 0x72, 0x19, 0x1f, 0x22, 0x32, 0xb4, 0x88, 0xd5, 0x3c,
	0xc5, 0x9f, 0xc3, 0xc5, 0x70, 0x59, 0xbb, 0x8c, 0xde, 0x44, 0xe8, 0x43, 0x9b, 0x51, 0x5b, 0x6f,
	0xcb, 0xd5, 0xcb, 0xe8, 0x30, 0xe3, 0x25, 0x72, 0x93, 0x37, 0x3f, 0xff, 0x3c, 0xf7, 0x3e, 0xbf,
	0xbe, 0xeb, 0x9d, 0xac, 0x0b, 0xae, 0x70, 0x0d, 0x71, 0x91, 0xa5, 0xf6, 0xb1, 0xf0, 0x2f, 0x86,
	0xab, 0x15, 0xcd, 0x63, 0x8b, 0x6e, 0x46, 0x97, 0x6d, 0xfb, 0x64, 0x5e, 0xfd, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x32, 0xf3, 0xc8, 0xa5, 0x99, 0x03, 0x00, 0x00,
}
