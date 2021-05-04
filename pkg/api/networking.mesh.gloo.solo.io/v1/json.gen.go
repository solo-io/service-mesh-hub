// Code generated by skv2. DO NOT EDIT.

// Generated json marshal and unmarshal functions

package v1

import (
    bytes "bytes"
    fmt "fmt"
    math "math"

    skv2jsonpb "github.com/solo-io/skv2/pkg/kube_jsonpb"
    jsonpb "github.com/golang/protobuf/jsonpb"
    proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var (
	marshaller = &skv2jsonpb.Marshaler{}
	unmarshaller = &jsonpb.Unmarshaler{}
)

// MarshalJSON is a custom marshaler for TrafficPolicySpec
func (this *TrafficPolicySpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec
func (this *TrafficPolicySpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}
// MarshalJSON is a custom marshaler for TrafficPolicyStatus
func (this *TrafficPolicyStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicyStatus
func (this *TrafficPolicyStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AccessPolicySpec
func (this *AccessPolicySpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AccessPolicySpec
func (this *AccessPolicySpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}
// MarshalJSON is a custom marshaler for AccessPolicyStatus
func (this *AccessPolicyStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AccessPolicyStatus
func (this *AccessPolicyStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec
func (this *VirtualMeshSpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec
func (this *VirtualMeshSpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}
// MarshalJSON is a custom marshaler for VirtualMeshStatus
func (this *VirtualMeshStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshStatus
func (this *VirtualMeshStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}
