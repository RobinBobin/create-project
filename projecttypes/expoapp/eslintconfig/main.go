package eslintconfig

import (
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/projecttypes/expoapp/eslintconfig/eslintconfigexpo"
	"github.com/robinbobin/create-project/utils"
)

func Process() {
	options.Options.AddFile(utils.ESLINT_CONFIG_JS)

	handleMissingTypes()
	useBaseConfigs()
	eslintconfigexpo.HandleOutdated()
	handleModuleType()
}
