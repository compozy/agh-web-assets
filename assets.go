// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "80206627024e4b34cf3f46342c82a291c78d497b7efd250081e3e5fc7779e425"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "d435cffaaa322e63f63219af985218c0200e446b"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
