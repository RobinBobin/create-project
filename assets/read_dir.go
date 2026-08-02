package assets

import (
	"io/fs"
	"path/filepath"
)

func ReadDir(name string) ([]fs.DirEntry, error) {
	return assetsFS.ReadDir(filepath.Join(ASSETS, name))
}
