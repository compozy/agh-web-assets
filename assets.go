// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7269801fa0f7697d3d2b2b9776a72229511f3827e59eedaadd1d7462d35c244b"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "5cf77b83eebf7b6f3a4b1a6c6e34d8a6e3d5bc77"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
