// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "35283f2650f636a08902a3dc275002471e7a1b0542d5c26202a46d87a89df740"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "d4a8c92cc137d20aed111cc2b3caf57f1ed99f30"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
