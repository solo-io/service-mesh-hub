// Code generated by skv2. DO NOT EDIT.

/*
	Utility for manually building input snapshots. Used primarily in tests.
*/
package input

import (
	settings_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2"
	settings_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2/sets"

	networking_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
	networking_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2/sets"
)

type InputSettingsSnapshotManualBuilder struct {
	name string

	settings settings_mesh_gloo_solo_io_v1alpha2_sets.SettingsSet

	virtualMeshes networking_mesh_gloo_solo_io_v1alpha2_sets.VirtualMeshSet
}

func NewInputSettingsSnapshotManualBuilder(name string) *InputSettingsSnapshotManualBuilder {
	return &InputSettingsSnapshotManualBuilder{
		name: name,

		settings: settings_mesh_gloo_solo_io_v1alpha2_sets.NewSettingsSet(),

		virtualMeshes: networking_mesh_gloo_solo_io_v1alpha2_sets.NewVirtualMeshSet(),
	}
}

func (i *InputSettingsSnapshotManualBuilder) Build() SettingsSnapshot {
	return NewSettingsSnapshot(
		i.name,

		i.settings,

		i.virtualMeshes,
	)
}
func (i *InputSettingsSnapshotManualBuilder) AddSettings(settings []*settings_mesh_gloo_solo_io_v1alpha2.Settings) *InputSettingsSnapshotManualBuilder {
	i.settings.Insert(settings...)
	return i
}
func (i *InputSettingsSnapshotManualBuilder) AddVirtualMeshes(virtualMeshes []*networking_mesh_gloo_solo_io_v1alpha2.VirtualMesh) *InputSettingsSnapshotManualBuilder {
	i.virtualMeshes.Insert(virtualMeshes...)
	return i
}
