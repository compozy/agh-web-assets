// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "75b003bbc98c7d9d79832bb6f72782c679ed2ca8a0b2a33befbf3c72f118fed1"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "fd57d66ca34a2e681c7a7f235c5b0a0e638d400c"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
