// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v1alpha1

import (
	proto "github.com/gogo/protobuf/proto"
)

// DeepCopyInto for the XdsConfig.Spec
func (in *XdsConfigSpec) DeepCopyInto(out *XdsConfigSpec) {
	p := proto.Clone(in).(*XdsConfigSpec)
	*out = *p
}

// DeepCopyInto for the XdsConfig.Status
func (in *XdsConfigStatus) DeepCopyInto(out *XdsConfigStatus) {
	p := proto.Clone(in).(*XdsConfigStatus)
	*out = *p
}
