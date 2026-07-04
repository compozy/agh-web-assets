// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a64d85925ce6013d5ad12762ec6f5f7ff660bcdeaa55c9f499f44fd308d12763"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "a7f94851ac7ff86686621d516ea6213f82e425f9"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
