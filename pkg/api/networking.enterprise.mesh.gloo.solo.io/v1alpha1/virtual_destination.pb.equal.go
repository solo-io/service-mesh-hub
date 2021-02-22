// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/enterprise/networking/v1alpha1/virtual_destination.proto

package v1alpha1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"

	v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)

	_ = v1alpha2.ApprovalState(0)
)

// Equal function
func (m *VirtualDestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualDestinationSpec)
	if !ok {
		that2, ok := that.(VirtualDestinationSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetHostname(), target.GetHostname()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetPort()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPort()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPort(), target.GetPort()) {
			return false
		}
	}

	switch m.ExportTo.(type) {

	case *VirtualDestinationSpec_VirtualMesh:

		if h, ok := interface{}(m.GetVirtualMesh()).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtualMesh()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetVirtualMesh(), target.GetVirtualMesh()) {
				return false
			}
		}

	case *VirtualDestinationSpec_MeshList_:

		if h, ok := interface{}(m.GetMeshList()).(equality.Equalizer); ok {
			if !h.Equal(target.GetMeshList()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetMeshList(), target.GetMeshList()) {
				return false
			}
		}

	}

	switch m.FailoverConfig.(type) {

	case *VirtualDestinationSpec_Static:

		if h, ok := interface{}(m.GetStatic()).(equality.Equalizer); ok {
			if !h.Equal(target.GetStatic()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetStatic(), target.GetStatic()) {
				return false
			}
		}

	case *VirtualDestinationSpec_Localized:

		if h, ok := interface{}(m.GetLocalized()).(equality.Equalizer); ok {
			if !h.Equal(target.GetLocalized()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetLocalized(), target.GetLocalized()) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *BackingService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*BackingService)
	if !ok {
		that2, ok := that.(BackingService)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.BackingServiceType.(type) {

	case *BackingService_KubeService:

		if h, ok := interface{}(m.GetKubeService()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKubeService()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKubeService(), target.GetKubeService()) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *VirtualDestinationStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualDestinationStatus)
	if !ok {
		that2, ok := that.(VirtualDestinationStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	if len(m.GetMeshes()) != len(target.GetMeshes()) {
		return false
	}
	for k, v := range m.GetMeshes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMeshes()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMeshes()[k]) {
				return false
			}
		}

	}

	if len(m.GetSelectedTrafficTargets()) != len(target.GetSelectedTrafficTargets()) {
		return false
	}
	for idx, v := range m.GetSelectedTrafficTargets() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedTrafficTargets()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedTrafficTargets()[idx]) {
				return false
			}
		}

	}

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if strings.Compare(v, target.GetErrors()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *VirtualDestinationSpec_Port) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualDestinationSpec_Port)
	if !ok {
		that2, ok := that.(VirtualDestinationSpec_Port)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetNumber() != target.GetNumber() {
		return false
	}

	if strings.Compare(m.GetProtocol(), target.GetProtocol()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *VirtualDestinationSpec_MeshList) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualDestinationSpec_MeshList)
	if !ok {
		that2, ok := that.(VirtualDestinationSpec_MeshList)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetMeshes()) != len(target.GetMeshes()) {
		return false
	}
	for idx, v := range m.GetMeshes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMeshes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMeshes()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *VirtualDestinationSpec_BackingServiceList) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualDestinationSpec_BackingServiceList)
	if !ok {
		that2, ok := that.(VirtualDestinationSpec_BackingServiceList)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetServices()) != len(target.GetServices()) {
		return false
	}
	for idx, v := range m.GetServices() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetServices()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetServices()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *VirtualDestinationSpec_LocalityConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualDestinationSpec_LocalityConfig)
	if !ok {
		that2, ok := that.(VirtualDestinationSpec_LocalityConfig)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetServiceSelectors()) != len(target.GetServiceSelectors()) {
		return false
	}
	for idx, v := range m.GetServiceSelectors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetServiceSelectors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetServiceSelectors()[idx]) {
				return false
			}
		}

	}

	if len(m.GetFailoverDirectives()) != len(target.GetFailoverDirectives()) {
		return false
	}
	for idx, v := range m.GetFailoverDirectives() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetFailoverDirectives()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetFailoverDirectives()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOutlierDetection()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOutlierDetection()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOutlierDetection(), target.GetOutlierDetection()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualDestinationSpec_LocalityConfig_LocalityFailoverDirective) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualDestinationSpec_LocalityConfig_LocalityFailoverDirective)
	if !ok {
		that2, ok := that.(VirtualDestinationSpec_LocalityConfig_LocalityFailoverDirective)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetFrom()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFrom()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFrom(), target.GetFrom()) {
			return false
		}
	}

	if len(m.GetTo()) != len(target.GetTo()) {
		return false
	}
	for idx, v := range m.GetTo() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetTo()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetTo()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *VirtualDestinationSpec_LocalityConfig_Locality) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualDestinationSpec_LocalityConfig_Locality)
	if !ok {
		that2, ok := that.(VirtualDestinationSpec_LocalityConfig_Locality)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetRegion(), target.GetRegion()) != 0 {
		return false
	}

	if strings.Compare(m.GetZone(), target.GetZone()) != 0 {
		return false
	}

	if strings.Compare(m.GetSubZone(), target.GetSubZone()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *VirtualDestinationStatus_SelectedTrafficTarget) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualDestinationStatus_SelectedTrafficTarget)
	if !ok {
		that2, ok := that.(VirtualDestinationStatus_SelectedTrafficTarget)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRef(), target.GetRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetService()).(equality.Equalizer); ok {
		if !h.Equal(target.GetService()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetService(), target.GetService()) {
			return false
		}
	}

	return true
}
