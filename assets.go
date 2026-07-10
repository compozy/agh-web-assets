// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3d5a8a8824905052e6843190893f6706272219f57d8d920dbc96c160a551fa66"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "91690329ca144c6ac3902d92484ee79f15bea2a3"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
