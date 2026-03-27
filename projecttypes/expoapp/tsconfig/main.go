package tsconfig

import (
	"github.com/robinbobin/create-project/utils"
)

const tsconfig_json = "tsconfig.json"

func Process() {
	tsconfig := utils.ReadJSON(tsconfig_json)

	deleteCompilerOptions(tsconfig)
	processExtends(tsconfig)
	processFiles(tsconfig)
	processInclude(tsconfig)

	utils.WriteJSON(tsconfig, tsconfig_json)
}
