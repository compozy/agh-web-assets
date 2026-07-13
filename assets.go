// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "084a3c36804d3518a3bf6753859e994056d748f72a499ceccddb469ee8f93e17"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "08c1797bd11a8e2f8a51e99a8c8101d38b3386ba"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
