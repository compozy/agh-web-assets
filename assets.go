// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "18e4fd76032135d759be4b7d9d4aaaea7b2111a0d774e54122e10217156998a1"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "8ade56e5af44ad21618396ad0564d667106293b7"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
