package packagejson

import (
	"github.com/robinbobin/create-project/utils"
)

func Lint() {
	if !setModuleType() {
		utils.AskSortJSON(utils.PACKAGE_JSON)
	}

	uninstallPackages()
}
