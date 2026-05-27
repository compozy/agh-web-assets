// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "35283f2650f636a08902a3dc275002471e7a1b0542d5c26202a46d87a89df740"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "85b1e8aa015773839a1387871bca5c2ed86eafef"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
