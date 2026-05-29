// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d7362f36137c2210afc112ee5320ae5af29d1ad356d2474455c435cc6633bd2d"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "fe563e08a1a9e2d08a444cc3df90ca3b9bdfb62a"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
