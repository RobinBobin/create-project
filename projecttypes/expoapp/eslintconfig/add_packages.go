package eslintconfig

import (
	"github.com/robinbobin/create-project/utils"
	"github.com/robinbobin/create-project/utils/packagejson"
)

func addPackages() {
	// @types/node

	if packagejson.IsInstalled(utils.ESLINT) {
		return
	}

	utils.RunCmd("pnpm install --dangerously-allow-all-builds --save-dev eslint eslint-import-resolver-typescript")
}
