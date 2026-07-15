// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a921f087877373b636316b3e8314c715f3c76e64065e1a40a7573ddfcb1a8bca"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "9d27b9d37a44de19a8896e5b59e356150bf85254"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
