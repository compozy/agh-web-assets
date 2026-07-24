// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "179723fedcfd67801226e58f5824d5b4673adeb1109ca6fa27d0aec79bfe6bf5"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "8561f9be56959c73d5b2caebcb689cffa41a8e77"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
