package assets

import (
	"io/fs"
	"path"
)

func ReadDir(name string) ([]fs.DirEntry, error) {
	return assetsFS.ReadDir(path.Join(ASSETS, name))
}
