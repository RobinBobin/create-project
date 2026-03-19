package tsconfig

import "github.com/robinbobin/create-project/utils"

func AddToFiles(fileName string) {
	tsconfig := utils.ReadJSON(tsconfig_json)

	addToFiles(fileName, tsconfig)

	utils.WriteJSON(tsconfig, tsconfig_json)
}
