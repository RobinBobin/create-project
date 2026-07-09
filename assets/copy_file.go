package assets

import (
	"github.com/robinbobin/create-project/utils"
)

func CopyFile(destinationName string, sourceName string) {
	source, sourceError := assets.Open(sourceName)

	utils.CopyFile(destinationName, source, sourceError)
}
