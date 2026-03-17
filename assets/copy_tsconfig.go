package assets

import (
	"github.com/robinbobin/create-project/utils"
)

func CopyTSConfig() {
	source, sourceError := assets.Open("tsconfig.json")

	utils.CopyFile("tsconfig.base.json", source, sourceError)
}
