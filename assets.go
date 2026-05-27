// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2c56d76f58e084e028f0f83530a3154505343986e6c88805b18b35495537c200"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "ada037db7b6e9ec21eabaae283f0c6afd9927f03"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
