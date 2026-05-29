// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d7362f36137c2210afc112ee5320ae5af29d1ad356d2474455c435cc6633bd2d"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "7cc2e89b8c353e79be65cec25fb05f12e0f08465"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
