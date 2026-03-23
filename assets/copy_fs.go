package assets

import (
	"io/fs"
	"os"
)

func CopyFS(destinationDir string, sourceDir string) (err error) {
	subFS, err := fs.Sub(assets, sourceDir)

	if err == nil {
		err = os.CopyFS(destinationDir, subFS)
	}

	return
}
