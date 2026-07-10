// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "12adb89221cd474a71298e24a49e14865f7d4bc8bfe6d68fa25dcdd24e2cea3d"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "a7a5397facef9a64596b94e0a733440923fd50a7"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
