// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d7362f36137c2210afc112ee5320ae5af29d1ad356d2474455c435cc6633bd2d"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "0c93ffb130718c5117c9e5f1f4f3d55440c0e5ec"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
