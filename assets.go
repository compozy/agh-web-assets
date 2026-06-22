// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c5afa8f561a64180c74c58533970c3ee2dda73bd1076c4f43e75ee884792ddf1"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "86ea1bd2c0155f1f617684a79ea96e6c0670a348"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
