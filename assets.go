// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d7362f36137c2210afc112ee5320ae5af29d1ad356d2474455c435cc6633bd2d"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "3b90b7d9b5fc0e03ad9c2be9c6512fc6eb73eea3"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
