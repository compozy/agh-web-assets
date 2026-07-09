// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8831b2155b84f37e3da8e35899e6a27fedbf539c26371dc05f2ca575911d57fd"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "220bac4ea0297ec027eeaeb2e9573233affe502c"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
