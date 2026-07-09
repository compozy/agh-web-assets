// Package webassets embeds the production AGH web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3666d92811cf817b9235ffddfaf00bb174e26f07faf0b60152a999ade02764a2"
	SourceRepository = "github.com/compozy/agh"
	SourceCommit = "23024457b3c17d04a2aa0db22780ca46f256faf4"
)

// DistFS embeds the generated production AGH web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
