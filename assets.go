// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7c59f3d95db8b9870ff227f5fd2f1ea242f36b6febd6f6cdadd08961d31dd9e6"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "51e065575f9e0e6e8b58511a445f7d741941b705"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
