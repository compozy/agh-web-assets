// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ce31192a4c1df12b29750197f5f4df6ab71aaa6a9387ef7093b56e969f931911"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "c1135fdb8da72151a575c3864f23450616523b86"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
