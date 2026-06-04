// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d3d65bcd6e51d2f9d6597a9cd18bbab358e7f2d5e40022d4a7d0dc3ac5c6f225"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "f8e5bd9ace6f99f7ded38a45c90114f049c8c822"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
