// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b02682db605b49cfa9b847a84797840a92988d535c8a5d6b640d35ff9880c3d4"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "e2c73b0e8dc98b5ef0d805eb106c7b73bb6f3b0c"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
