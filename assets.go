// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b7a101555518ee4455811f2248ac189bc5d99e1d7043e846c87f09eb318a8432"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "49fbc39067c5724311e5ee0739453d596d87e8d7"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
