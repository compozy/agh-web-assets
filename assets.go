// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "01a977135695ac9cd97491e358ed52858389c5f0f8e187882bc41e06f011205c"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "48d7696a55e145fcb0103f877ec2e34ef9ceb75b"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
