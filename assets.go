// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "066638bf98f6a1691399c85e62c74a1453c92c3f11d5b7a43b015c292a279671"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "2efa801dbae736563c45e85abe4feb5b9485b586"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
