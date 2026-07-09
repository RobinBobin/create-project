package assets

import (
	"path"

	"github.com/robinbobin/create-project/utils"
)

func ReadFile(name string) []byte {
	buf, err := assetsFS.ReadFile(path.Join(ASSETS, name))

	utils.PanicOnError(err)

	return buf
}
