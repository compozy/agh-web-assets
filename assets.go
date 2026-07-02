// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4a5afef95508c57cab6eb0abd52292661c2ac299887497369ff418170abdb3e6"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "65410d6e4ede2a171e578bc58ac27170cf39c058"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
