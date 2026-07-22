// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a908f9a6e228ad156d5208f51d38636526d059be95b974bd5dbf04e422f84ed6"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "36eb9715944ba48fca5112fc87e598edcfe0d704"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
