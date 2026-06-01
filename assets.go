// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6aaeaf053c75e2832cd4fa7933191e10c569438dd1492ebf4d55e13a29a40c20"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "d7920470a5cea6175700a77846a4dc4ef23913ad"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
