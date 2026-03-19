package tsconfig

import (
	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/utils"
)

const tsconfig_json = "tsconfig.json"

func Process(isProjectReset bool) {
	if !utils.Confirm("Would you like to use a custom tsconfig?", true) {
		return
	}

	assets.CopyFile("tsconfig.base.json", tsconfig_json)

	backup()

	tsconfig := utils.ReadJSON(tsconfig_json)

	deleteCompilerOptions(isProjectReset, tsconfig)
	processExtends(tsconfig)
	processFiles(tsconfig)
	processInclude(isProjectReset, tsconfig)

	utils.WriteJSON(tsconfig, tsconfig_json)
}
