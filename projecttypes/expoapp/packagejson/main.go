package packagejson

import (
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/projecttypes/expoapp/packagejson/addition"
	"github.com/robinbobin/create-project/utils"
)

func Process() {
	options.Options.IsProjectReset = resetProject()
	options.Options.IsESM = setType()

	isSorted := options.Options.IsESM || options.Options.IsProjectReset

	if !isSorted {
		utils.AskSortJSON(utils.PACKAGE_JSON)
	}

	uninstallPackages(options.Options.IsProjectReset)
	addition.Run()
}
