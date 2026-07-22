// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2e36840496099e94ba62f72ae9d6e4e37f87041bf1ad8054dd786a4c26157632"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "a2fe3098594cc664bc2c0dcb26c0da9d12dd4dd0"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
