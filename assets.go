// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ed912d5d6a98e32b8f1f88e671bea7adb3709a6d21d9ba2a0ba0c1b5c2b3b654"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "64b269f7ffff688ebf63f8802c5746923fbacae3"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
