package web

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed assets
var assets embed.FS

func Assets() fs.FS {
	assetsFS, err := fs.Sub(assets, "assets")
	if err != nil {
		log.Fatal(err)
	}
	return assetsFS
}
