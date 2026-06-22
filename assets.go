// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9f837b3c8a85f6a04a494002279979883ce5d7fd32ad1945b58eb1b683be7a29"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "da71d0df0cb27a3e76d83f9053f2ad82ffff36e9"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
