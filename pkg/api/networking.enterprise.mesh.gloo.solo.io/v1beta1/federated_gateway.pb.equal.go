// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/enterprise/networking/v1beta1/federated_gateway.proto

package v1beta1

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
func (m *FederatedGatewaySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec)
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

	if strings.Compare(m.GetBindAddress(), target.GetBindAddress()) != 0 {
		return false
	}

	if m.GetBindPort() != target.GetBindPort() {
		return false
	}

	if len(m.GetConnectionHandlers()) != len(target.GetConnectionHandlers()) {
		return false
	}
	for idx, v := range m.GetConnectionHandlers() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetConnectionHandlers()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetConnectionHandlers()[idx]) {
				return false
			}
		}

	}

	if len(m.GetGatewayWorkloads()) != len(target.GetGatewayWorkloads()) {
		return false
	}
	for idx, v := range m.GetGatewayWorkloads() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetGatewayWorkloads()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetGatewayWorkloads()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetGatewayOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGatewayOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGatewayOptions(), target.GetGatewayOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewayStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewayStatus)
	if !ok {
		that2, ok := that.(FederatedGatewayStatus)
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

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if strings.Compare(v, target.GetErrors()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetWarnings()) != len(target.GetWarnings()) {
		return false
	}
	for idx, v := range m.GetWarnings() {

		if strings.Compare(v, target.GetWarnings()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetSelectedGateways()) != len(target.GetSelectedGateways()) {
		return false
	}
	for idx, v := range m.GetSelectedGateways() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedGateways()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedGateways()[idx]) {
				return false
			}
		}

	}

	if len(m.GetSelectedRouteTables()) != len(target.GetSelectedRouteTables()) {
		return false
	}
	for idx, v := range m.GetSelectedRouteTables() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedRouteTables()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedRouteTables()[idx]) {
				return false
			}
		}

	}

	if len(m.GetSelectedDelegatedRouteTables()) != len(target.GetSelectedDelegatedRouteTables()) {
		return false
	}
	for idx, v := range m.GetSelectedDelegatedRouteTables() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedDelegatedRouteTables()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedDelegatedRouteTables()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *SelectedGateway) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SelectedGateway)
	if !ok {
		that2, ok := that.(SelectedGateway)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	if strings.Compare(m.GetCluster(), target.GetCluster()) != 0 {
		return false
	}

	if strings.Compare(m.GetExternalUrl(), target.GetExternalUrl()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *SelectedRouteTable) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SelectedRouteTable)
	if !ok {
		that2, ok := that.(SelectedRouteTable)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	if strings.Compare(m.GetCluster(), target.GetCluster()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *SDSConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SDSConfig)
	if !ok {
		that2, ok := that.(SDSConfig)
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

	if strings.Compare(m.GetTargetUri(), target.GetTargetUri()) != 0 {
		return false
	}

	if strings.Compare(m.GetCertificatesSecretName(), target.GetCertificatesSecretName()) != 0 {
		return false
	}

	if strings.Compare(m.GetValidationContextName(), target.GetValidationContextName()) != 0 {
		return false
	}

	switch m.SdsBuilder.(type) {

	case *SDSConfig_CallCredentials_:
		if _, ok := target.SdsBuilder.(*SDSConfig_CallCredentials_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetCallCredentials()).(equality.Equalizer); ok {
			if !h.Equal(target.GetCallCredentials()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetCallCredentials(), target.GetCallCredentials()) {
				return false
			}
		}

	case *SDSConfig_ClusterName:
		if _, ok := target.SdsBuilder.(*SDSConfig_ClusterName); !ok {
			return false
		}

		if strings.Compare(m.GetClusterName(), target.GetClusterName()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.SdsBuilder != target.SdsBuilder {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler)
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

	if h, ok := interface{}(m.GetConnectionMatch()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConnectionMatch()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConnectionMatch(), target.GetConnectionMatch()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetConnectionOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConnectionOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConnectionOptions(), target.GetConnectionOptions()) {
			return false
		}
	}

	switch m.HandlerType.(type) {

	case *FederatedGatewaySpec_ConnectionHandler_Http:
		if _, ok := target.HandlerType.(*FederatedGatewaySpec_ConnectionHandler_Http); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHttp()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttp()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHttp(), target.GetHttp()) {
				return false
			}
		}

	case *FederatedGatewaySpec_ConnectionHandler_Tcp:
		if _, ok := target.HandlerType.(*FederatedGatewaySpec_ConnectionHandler_Tcp); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTcp()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcp()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTcp(), target.GetTcp()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.HandlerType != target.HandlerType {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_GatewayOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_GatewayOptions)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_GatewayOptions)
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

	if h, ok := interface{}(m.GetTrafficPolicy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTrafficPolicy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTrafficPolicy(), target.GetTrafficPolicy()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPerConnectionBufferLimitBytes()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPerConnectionBufferLimitBytes(), target.GetPerConnectionBufferLimitBytes()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_ConnectionMatch) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_ConnectionMatch)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_ConnectionMatch)
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

	if strings.Compare(m.GetTransportProtocol(), target.GetTransportProtocol()) != 0 {
		return false
	}

	if len(m.GetServerNames()) != len(target.GetServerNames()) {
		return false
	}
	for idx, v := range m.GetServerNames() {

		if strings.Compare(v, target.GetServerNames()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_ConnectionOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_ConnectionOptions)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_ConnectionOptions)
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

	if h, ok := interface{}(m.GetTlsContext()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTlsContext()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTlsContext(), target.GetTlsContext()) {
			return false
		}
	}

	if m.GetEnableProxyProtocol() != target.GetEnableProxyProtocol() {
		return false
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_HttpRoutes) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_HttpRoutes)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_HttpRoutes)
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

	if len(m.GetRouteConfig()) != len(target.GetRouteConfig()) {
		return false
	}
	for idx, v := range m.GetRouteConfig() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRouteConfig()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRouteConfig()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetRouteOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRouteOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRouteOptions(), target.GetRouteOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_TcpRoutes) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_TcpRoutes)
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

	if len(m.GetTcpHosts()) != len(target.GetTcpHosts()) {
		return false
	}
	for idx, v := range m.GetTcpHosts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpHosts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetTcpHosts()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_ConnectionOptions_TlsTerminationOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_ConnectionOptions_TlsTerminationOptions)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_ConnectionOptions_TlsTerminationOptions)
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

	if h, ok := interface{}(m.GetPresented()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPresented()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPresented(), target.GetPresented()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetValidated()).(equality.Equalizer); ok {
		if !h.Equal(target.GetValidated()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetValidated(), target.GetValidated()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier)
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

	switch m.RouteType.(type) {

	case *FederatedGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier_RouteSelector:
		if _, ok := target.RouteType.(*FederatedGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier_RouteSelector); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRouteSelector()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRouteSelector()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRouteSelector(), target.GetRouteSelector()) {
				return false
			}
		}

	case *FederatedGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier_RouteTable:
		if _, ok := target.RouteType.(*FederatedGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier_RouteTable); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRouteTable()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRouteTable()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRouteTable(), target.GetRouteTable()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.RouteType != target.RouteType {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_HttpRoutes_HttpOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_HttpRoutes_HttpOptions)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_HttpRoutes_HttpOptions)
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

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetSslConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSslConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSslConfig(), target.GetSslConfig()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDestination()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDestination()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDestination(), target.GetDestination()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions)
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

	if h, ok := interface{}(m.GetTcpProxySettings()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTcpProxySettings()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTcpProxySettings(), target.GetTcpProxySettings()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig)
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

	if len(m.GetSniDomains()) != len(target.GetSniDomains()) {
		return false
	}
	for idx, v := range m.GetSniDomains() {

		if strings.Compare(v, target.GetSniDomains()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetVerifySubjectAltName()) != len(target.GetVerifySubjectAltName()) {
		return false
	}
	for idx, v := range m.GetVerifySubjectAltName() {

		if strings.Compare(v, target.GetVerifySubjectAltName()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetParameters()).(equality.Equalizer); ok {
		if !h.Equal(target.GetParameters()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetParameters(), target.GetParameters()) {
			return false
		}
	}

	if len(m.GetAlpnProtocols()) != len(target.GetAlpnProtocols()) {
		return false
	}
	for idx, v := range m.GetAlpnProtocols() {

		if strings.Compare(v, target.GetAlpnProtocols()[idx]) != 0 {
			return false
		}

	}

	switch m.SslSecrets.(type) {

	case *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig_SecretRef:
		if _, ok := target.SslSecrets.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig_SecretRef); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSecretRef()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSecretRef()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSecretRef(), target.GetSecretRef()) {
				return false
			}
		}

	case *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig_SslFiles:
		if _, ok := target.SslSecrets.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig_SslFiles); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSslFiles()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSslFiles()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSslFiles(), target.GetSslFiles()) {
				return false
			}
		}

	case *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig_Sds:
		if _, ok := target.SslSecrets.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig_Sds); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSds()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSds()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSds(), target.GetSds()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.SslSecrets != target.SslSecrets {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction)
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

	if m.GetWeight() != target.GetWeight() {
		return false
	}

	switch m.Destination.(type) {

	case *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_Static:
		if _, ok := target.Destination.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_Static); !ok {
			return false
		}

		if h, ok := interface{}(m.GetStatic()).(equality.Equalizer); ok {
			if !h.Equal(target.GetStatic()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetStatic(), target.GetStatic()) {
				return false
			}
		}

	case *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_Virtual:
		if _, ok := target.Destination.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_Virtual); !ok {
			return false
		}

		if h, ok := interface{}(m.GetVirtual()).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtual()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetVirtual(), target.GetVirtual()) {
				return false
			}
		}

	case *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_Kube:
		if _, ok := target.Destination.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_Kube); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKube()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKube()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKube(), target.GetKube()) {
				return false
			}
		}

	case *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_ForwardSniClusterName:
		if _, ok := target.Destination.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_ForwardSniClusterName); !ok {
			return false
		}

		if h, ok := interface{}(m.GetForwardSniClusterName()).(equality.Equalizer); ok {
			if !h.Equal(target.GetForwardSniClusterName()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetForwardSniClusterName(), target.GetForwardSniClusterName()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Destination != target.Destination {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig_SSLFiles) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig_SSLFiles)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig_SSLFiles)
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

	if strings.Compare(m.GetTlsCert(), target.GetTlsCert()) != 0 {
		return false
	}

	if strings.Compare(m.GetTlsKey(), target.GetTlsKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetRootCa(), target.GetRootCa()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig_SslParameters) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig_SslParameters)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_SslConfig_SslParameters)
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

	if m.GetMinimumProtocolVersion() != target.GetMinimumProtocolVersion() {
		return false
	}

	if m.GetMaximumProtocolVersion() != target.GetMaximumProtocolVersion() {
		return false
	}

	if len(m.GetCipherSuites()) != len(target.GetCipherSuites()) {
		return false
	}
	for idx, v := range m.GetCipherSuites() {

		if strings.Compare(v, target.GetCipherSuites()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetEcdhCurves()) != len(target.GetEcdhCurves()) {
		return false
	}
	for idx, v := range m.GetEcdhCurves() {

		if strings.Compare(v, target.GetEcdhCurves()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions_TcpProxySettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions_TcpProxySettings)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions_TcpProxySettings)
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

	if h, ok := interface{}(m.GetMaxConnectAttempts()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxConnectAttempts()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxConnectAttempts(), target.GetMaxConnectAttempts()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIdleTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIdleTimeout(), target.GetIdleTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTunnelingConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTunnelingConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTunnelingConfig(), target.GetTunnelingConfig()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions_TcpProxySettings_TunnelingConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions_TcpProxySettings_TunnelingConfig)
	if !ok {
		that2, ok := that.(FederatedGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions_TcpProxySettings_TunnelingConfig)
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

	return true
}

// Equal function
func (m *SDSConfig_CallCredentials) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SDSConfig_CallCredentials)
	if !ok {
		that2, ok := that.(SDSConfig_CallCredentials)
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

	if h, ok := interface{}(m.GetFileCredentialSource()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFileCredentialSource()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFileCredentialSource(), target.GetFileCredentialSource()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *SDSConfig_CallCredentials_FileCredentialSource) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SDSConfig_CallCredentials_FileCredentialSource)
	if !ok {
		that2, ok := that.(SDSConfig_CallCredentials_FileCredentialSource)
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

	if strings.Compare(m.GetTokenFileName(), target.GetTokenFileName()) != 0 {
		return false
	}

	if strings.Compare(m.GetHeader(), target.GetHeader()) != 0 {
		return false
	}

	return true
}
