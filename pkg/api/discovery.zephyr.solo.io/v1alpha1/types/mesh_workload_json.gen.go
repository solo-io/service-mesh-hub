// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/mesh_workload.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for MeshWorkloadSpec
func (this *MeshWorkloadSpec) MarshalJSON() ([]byte, error) {
	str, err := MeshWorkloadMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshWorkloadSpec
func (this *MeshWorkloadSpec) UnmarshalJSON(b []byte) error {
	return MeshWorkloadUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshWorkloadSpec_Appmesh
func (this *MeshWorkloadSpec_Appmesh) MarshalJSON() ([]byte, error) {
	str, err := MeshWorkloadMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshWorkloadSpec_Appmesh
func (this *MeshWorkloadSpec_Appmesh) UnmarshalJSON(b []byte) error {
	return MeshWorkloadUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshWorkloadSpec_Appmesh_ContainerPort
func (this *MeshWorkloadSpec_Appmesh_ContainerPort) MarshalJSON() ([]byte, error) {
	str, err := MeshWorkloadMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshWorkloadSpec_Appmesh_ContainerPort
func (this *MeshWorkloadSpec_Appmesh_ContainerPort) UnmarshalJSON(b []byte) error {
	return MeshWorkloadUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshWorkloadSpec_KubeController
func (this *MeshWorkloadSpec_KubeController) MarshalJSON() ([]byte, error) {
	str, err := MeshWorkloadMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshWorkloadSpec_KubeController
func (this *MeshWorkloadSpec_KubeController) UnmarshalJSON(b []byte) error {
	return MeshWorkloadUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshWorkloadStatus
func (this *MeshWorkloadStatus) MarshalJSON() ([]byte, error) {
	str, err := MeshWorkloadMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshWorkloadStatus
func (this *MeshWorkloadStatus) UnmarshalJSON(b []byte) error {
	return MeshWorkloadUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	MeshWorkloadMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	MeshWorkloadUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
