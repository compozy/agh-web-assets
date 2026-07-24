// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7c59f3d95db8b9870ff227f5fd2f1ea242f36b6febd6f6cdadd08961d31dd9e6"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "545c0ca6ad6fbcf154c0c7b08e8dc382e08dd9de"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
