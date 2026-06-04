// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4cd206dc7883e6a738488c52aa5bfb64ac4e96c4ed04a0a29fdf8e28a0875156"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "3df924137035b4b1d5004095197893ad1955269f"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
