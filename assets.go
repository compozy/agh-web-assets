// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8901a5245c9bd554bf5623e94cf532443db623006bb54a3684c53ebf5dea125d"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "ac0b709184d2a867aa6b4eeb88732483570911c4"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
