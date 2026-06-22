// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6056d318a83a5590586d4141019ee8fa1086f58e243bb2c105322ac5c7e13880"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "4e7a6d2503f5bdb60e47729841610a4de675362d"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
