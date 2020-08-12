// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/validation_state.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for ApprovalStatus
func (this *ApprovalStatus) MarshalJSON() ([]byte, error) {
	str, err := ValidationStateMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ApprovalStatus
func (this *ApprovalStatus) UnmarshalJSON(b []byte) error {
	return ValidationStateUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ValidationStateMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ValidationStateUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
