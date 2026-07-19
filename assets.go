// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "595c027f3c250697fcc13bc44374c1737b19980823641f8c46023e3e80c7c91b"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "658f2f67007099e676caf2ab1ee5c35585f56fca"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
