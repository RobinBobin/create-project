package assets

import (
	"embed"
	"io/fs"

	"github.com/robinbobin/create-project/utils"
)

//go:embed all:assets/*
var assetsFS embed.FS

var assets fs.FS

func init() {
	var err error

	assets, err = fs.Sub(assetsFS, "assets")

	utils.PanicOnError(err)
}
