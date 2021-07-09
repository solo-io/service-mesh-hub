// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/certificates/v1/issued_certificate.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
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
)

// Equal function
func (m *IssuedCertificateSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IssuedCertificateSpec)
	if !ok {
		that2, ok := that.(IssuedCertificateSpec)
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

	if len(m.GetHosts()) != len(target.GetHosts()) {
		return false
	}
	for idx, v := range m.GetHosts() {

		if strings.Compare(v, target.GetHosts()[idx]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetOrg(), target.GetOrg()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetSigningCertificateSecret()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSigningCertificateSecret()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSigningCertificateSecret(), target.GetSigningCertificateSecret()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIssuedCertificateSecret()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIssuedCertificateSecret()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIssuedCertificateSecret(), target.GetIssuedCertificateSecret()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPodBounceDirective()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPodBounceDirective()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPodBounceDirective(), target.GetPodBounceDirective()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCertOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCertOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCertOptions(), target.GetCertOptions()) {
			return false
		}
	}

	if m.GetRotationState() != target.GetRotationState() {
		return false
	}

	switch m.CertificateAuthority.(type) {

	case *IssuedCertificateSpec_GlooMeshCa:
		if _, ok := target.CertificateAuthority.(*IssuedCertificateSpec_GlooMeshCa); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGlooMeshCa()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGlooMeshCa()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGlooMeshCa(), target.GetGlooMeshCa()) {
				return false
			}
		}

	case *IssuedCertificateSpec_AgentCa:
		if _, ok := target.CertificateAuthority.(*IssuedCertificateSpec_AgentCa); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAgentCa()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAgentCa()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAgentCa(), target.GetAgentCa()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.CertificateAuthority != target.CertificateAuthority {
			return false
		}
	}

	return true
}

// Equal function
func (m *RootCertificateAuthority) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RootCertificateAuthority)
	if !ok {
		that2, ok := that.(RootCertificateAuthority)
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

	switch m.CertificateAuthority.(type) {

	case *RootCertificateAuthority_SigningCertificateSecret:
		if _, ok := target.CertificateAuthority.(*RootCertificateAuthority_SigningCertificateSecret); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSigningCertificateSecret()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSigningCertificateSecret()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSigningCertificateSecret(), target.GetSigningCertificateSecret()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.CertificateAuthority != target.CertificateAuthority {
			return false
		}
	}

	return true
}

// Equal function
func (m *IssuedCertificateStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IssuedCertificateStatus)
	if !ok {
		that2, ok := that.(IssuedCertificateStatus)
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

	if strings.Compare(m.GetError(), target.GetError()) != 0 {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	if len(m.GetConditions()) != len(target.GetConditions()) {
		return false
	}
	for idx, v := range m.GetConditions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetConditions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetConditions()[idx]) {
				return false
			}
		}

	}

	switch m.AppliedCertificateAuthority.(type) {

	case *IssuedCertificateStatus_AppliedGlooMeshCa:
		if _, ok := target.AppliedCertificateAuthority.(*IssuedCertificateStatus_AppliedGlooMeshCa); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAppliedGlooMeshCa()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedGlooMeshCa()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAppliedGlooMeshCa(), target.GetAppliedGlooMeshCa()) {
				return false
			}
		}

	case *IssuedCertificateStatus_AppliedAgentCa:
		if _, ok := target.AppliedCertificateAuthority.(*IssuedCertificateStatus_AppliedAgentCa); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAppliedAgentCa()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedAgentCa()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAppliedAgentCa(), target.GetAppliedAgentCa()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.AppliedCertificateAuthority != target.AppliedCertificateAuthority {
			return false
		}
	}

	return true
}
