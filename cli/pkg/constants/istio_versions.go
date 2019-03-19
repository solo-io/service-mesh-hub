package constants

import (
	"github.com/solo-io/supergloo/pkg/install/mesh"
)

var SupportedIstioVersions = []string{
	mesh.IstioVersion103,
	mesh.IstioVersion105,
	mesh.IstioVersion106,
}

var SupportedGlooVersions = []string{
	"latest",
	"v0.11.1",
	"v0.10.5",
}
