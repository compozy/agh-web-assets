// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "0771bbcabcf32fe6aaf52b45d5439bf413970e136423893200d64db788dbd7be"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "dc5e7b0cf1bc5850f35109738861e39f011dfb02"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
