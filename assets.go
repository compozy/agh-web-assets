// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ca2b39686c0a083aee6647170f44f7eb1fd120d6b12fabbe80ab9e5a2d09103c"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "763968bdc7b04ee36f6318beea0d6d22f7a88718"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
