// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "35283f2650f636a08902a3dc275002471e7a1b0542d5c26202a46d87a89df740"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "afdc459dc99ea6476dd0501534d3839baeca896e"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
