// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a921f087877373b636316b3e8314c715f3c76e64065e1a40a7573ddfcb1a8bca"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "f230872fcc46ba2bcd1f21cbdc404cb9ab2fb956"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
