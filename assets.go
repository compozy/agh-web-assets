// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2c56d76f58e084e028f0f83530a3154505343986e6c88805b18b35495537c200"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "dba9069da7882306d593e99dac9ffa448af9ccaf"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
