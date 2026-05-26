// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ff707feaa2e527eda55dc5b5f2f685376d1666b955bfb7f2ce4312c4db3e5826"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "6251d1dc663b478b10f5f9f0aa1d70c3aa043afb"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
