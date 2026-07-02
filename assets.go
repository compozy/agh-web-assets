// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c5afa8f561a64180c74c58533970c3ee2dda73bd1076c4f43e75ee884792ddf1"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "00a42bb9a456418e5c0376b6f781c1e7f4d9275a"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
