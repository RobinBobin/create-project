package packageaddition

import "github.com/robinbobin/create-project/utils"

func addRobinBobinTSConfig() {
	utils.RunCmd("pnpm i --save-dev @robinbobin/ts-config")

	// "extends": ["expo/tsconfig.base", "@robinbobin/ts-config"],
}
