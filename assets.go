// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "efb4640713a230ebac6ea4548dcd615c32599a93a43678bb281aaa477b0b838d"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "2150b84e31f04020035604fbb5d82e8839a4783c"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
