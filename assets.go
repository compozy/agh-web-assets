// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c6c5dcb60a6d3a94957d8beb91679985cea874aa2193ae5b63dbdd5a2e2e683c"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "dadf35f1f1ef00d56bea293df0f868ff52345bc3"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
