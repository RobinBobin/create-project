package tsconfig

import (
	"github.com/robinbobin/create-project/utils"
)

func Process() {
	tsconfig := utils.ReadJSON(utils.TSCONFIG_JSON)

	processCompilerOptions(tsconfig)
	processExtends(tsconfig)
	processFiles(tsconfig)
	processInclude(tsconfig)

	utils.WriteJSON(tsconfig, utils.TSCONFIG_JSON)
}
