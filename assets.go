// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8f67c538d922ce5ee166076ad75d2e36b7fbfd791d7b1d4bcd8ef5cc454be2a2"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "61d62e8376811651c61c9994a2d0db827895b1a3"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
