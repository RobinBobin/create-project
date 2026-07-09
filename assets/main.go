package assets

import (
	"embed"
	"io/fs"

	"github.com/robinbobin/create-project/utils"
)

const ASSETS = "assets"
const BABEL_CONFIG_JS = "babel.config.js"

//go:embed all:assets/*
var assetsFS embed.FS

var assets fs.FS

func init() {
	var err error

	assets, err = fs.Sub(assetsFS, ASSETS)

	utils.PanicOnError(err)
}
