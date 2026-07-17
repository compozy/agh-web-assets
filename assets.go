// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2fe7b10de40e0316992c17a7b622aa40fa3e1ab474bcb2d7d89f91de72823564"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "d872360ed81308d4dcba6f2dca09003c7296ad40"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
