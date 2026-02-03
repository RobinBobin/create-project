package packageaddition

import "github.com/robinbobin/create-project/utils"

func addRobinBobinESLintConfigRN() {
	utils.RunCmd("pnpm i --save-dev @robinbobin/eslint-config-react-native")

	// cp node_modules/@robinbobin/eslint-config-react/eslint.config.js_ ./eslint.config.js
}
