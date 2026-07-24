// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a0f645f0276770ade95407cd051a4186a9c4cb0090a1846694aae5847de7f613"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "a44d74cef7ab4151b287dde52d43812ee423b599"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
