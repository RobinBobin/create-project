package packageaddition

import "github.com/robinbobin/create-project/utils"

func addRobinBobinPrettierConfig() {
	utils.RunCmd("pnpm i --save-dev @robinbobin/prettier-config prettier")

	// cp node_modules/@robinbobin/prettier-config/.prettierignore node_modules/@robinbobin/prettier-config/.prettierrc.json .
}
