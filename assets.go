// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b02682db605b49cfa9b847a84797840a92988d535c8a5d6b640d35ff9880c3d4"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "8cc8ff39b52f4c74cb34cb9447c8cda6d0bea10c"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
