// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b0ab6e11ea61f1063c024989cb8897cffbfea84ec40bf5d4fe2e122c67ff7f39"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "1099164be7f7d4398129db062a464530e7b49751"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
