package tsconfig

import (
	"github.com/robinbobin/create-project/utils"
)

const tsconfig_json = "tsconfig.json"

func Process(isProjectReset bool) {
	tsconfig := utils.ReadJSON(tsconfig_json)

	deleteCompilerOptions(isProjectReset, tsconfig)
	processExtends(tsconfig)
	addToFiles("eslint.config.js", tsconfig)
	processInclude(isProjectReset, tsconfig)

	utils.WriteJSON(tsconfig, tsconfig_json)
}
