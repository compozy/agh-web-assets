// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "38660bc9f4b5b7de9861bf21c5871e3c2ee0fdbc7cbb330ee87154908e827941"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "118d594bed197636a5187318bbdf3503367a6778"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
