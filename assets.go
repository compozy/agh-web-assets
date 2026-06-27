// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c5afa8f561a64180c74c58533970c3ee2dda73bd1076c4f43e75ee884792ddf1"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "1630bfe95c6ec34865fc8f29e06df90071f51814"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
