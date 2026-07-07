// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "80906339dcaf5564a44cc5bafc696aba6a30b3fae733067469d5a644393df399"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "53d4c44f007ab6c4e3a24694c4d91e0e20359fed"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
