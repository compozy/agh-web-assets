// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3041bee7813886d9b8e02a09bd229d3fec50b3f726b0e4925c105d85d64b2743"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "4504cfa0bde0465b76b727404f194d7406fc2fd4"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
