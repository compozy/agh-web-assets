// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8f67c538d922ce5ee166076ad75d2e36b7fbfd791d7b1d4bcd8ef5cc454be2a2"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "8ac994d52fcf0c50f754aa577563e1d80aae3fa2"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
