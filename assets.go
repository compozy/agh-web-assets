// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2fe7b10de40e0316992c17a7b622aa40fa3e1ab474bcb2d7d89f91de72823564"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "bf32f245067765b02995b3a9e767957ef2d44164"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
