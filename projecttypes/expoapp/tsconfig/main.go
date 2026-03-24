package tsconfig

import (
	"github.com/robinbobin/create-project/projecttypes/expoapp/eslintconfig"
	"github.com/robinbobin/create-project/projecttypes/expoapp/packagejson"
	"github.com/robinbobin/create-project/utils"
)

const tsconfig_json = "tsconfig.json"

func Process(
	eslintConfigOptions *eslintconfig.Options,
	packageJsonOptions *packagejson.Options,
) {
	tsconfig := utils.ReadJSON(tsconfig_json)

	deleteCompilerOptions(packageJsonOptions.IsProjectReset, tsconfig)
	processExtends(tsconfig)
	processFiles(eslintConfigOptions.Files, tsconfig)

	processInclude(
		eslintConfigOptions.Include,
		packageJsonOptions.IsProjectReset,
		tsconfig,
	)

	utils.WriteJSON(tsconfig, tsconfig_json)
}
