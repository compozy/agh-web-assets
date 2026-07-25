// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c2f79c7ad5f71c7d4a7e2858b857c2fd9b548256839fcbae3f33fa63453cc199"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "0f445fec5333d562edf1d19498ebf3026fc8b335"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
