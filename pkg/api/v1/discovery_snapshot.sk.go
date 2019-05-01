// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"go.uber.org/zap"
)

type DiscoverySnapshot struct {
	Pods       github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodsByNamespace
	Configmaps github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigmapsByNamespace
}

func (s DiscoverySnapshot) Clone() DiscoverySnapshot {
	return DiscoverySnapshot{
		Pods:       s.Pods.Clone(),
		Configmaps: s.Configmaps.Clone(),
	}
}

func (s DiscoverySnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashPods(),
		s.hashConfigmaps(),
	)
}

func (s DiscoverySnapshot) hashPods() uint64 {
	return hashutils.HashAll(s.Pods.List().AsInterfaces()...)
}

func (s DiscoverySnapshot) hashConfigmaps() uint64 {
	return hashutils.HashAll(s.Configmaps.List().AsInterfaces()...)
}

func (s DiscoverySnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("pods", s.hashPods()))
	fields = append(fields, zap.Uint64("configmaps", s.hashConfigmaps()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type DiscoverySnapshotStringer struct {
	Version    uint64
	Pods       []string
	Configmaps []string
}

func (ss DiscoverySnapshotStringer) String() string {
	s := fmt.Sprintf("DiscoverySnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Pods %v\n", len(ss.Pods))
	for _, name := range ss.Pods {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Configmaps %v\n", len(ss.Configmaps))
	for _, name := range ss.Configmaps {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s DiscoverySnapshot) Stringer() DiscoverySnapshotStringer {
	return DiscoverySnapshotStringer{
		Version:    s.Hash(),
		Pods:       s.Pods.List().NamespacesDotNames(),
		Configmaps: s.Configmaps.List().NamespacesDotNames(),
	}
}
