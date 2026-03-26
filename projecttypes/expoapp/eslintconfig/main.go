package eslintconfig

import (
	"github.com/robinbobin/create-project/projecttypes/expoapp/eslintconfig/eslintconfigexpo"
	"github.com/robinbobin/create-project/projecttypes/expoapp/packagejson"
	"github.com/robinbobin/create-project/utils"
)

func Process(packageJsonOptions *packagejson.Options) *Options {
	options := &Options{Files: []string{utils.ESLINT_CONFIG_JS}}

	handleMissingTypes(packageJsonOptions.IsESM, options)
	useBaseConfigs(options)
	eslintconfigexpo.HandleOutdated()
	handleModuleType(packageJsonOptions.IsESM)

	return options
}
