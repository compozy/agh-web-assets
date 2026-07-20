// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6df47188e7cb0d39e6cdf55182a2130261f3be3f8db62c8cf68ad4ad1766dbf9"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "a87ed21ab75ff5c9c3681c6e11dc18076e75ef63"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
