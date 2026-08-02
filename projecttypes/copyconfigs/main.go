package copyconfigs

import (
	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/assets/eslintconfig"
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
)

func Create() bool {
	options.Options.CanInstallPackages = false

	// ESLint
	eslintconfig.Process()

	// Prettier
	assets.AddPrettier()

	// TS
	assets.CopyFile("tsconfig.base.json", utils.TSCONFIG_JSON)
	options.Options.Hints.Add("pnpm i --save-dev tsconfig@6")

	// VSCode workspace
	// assets.CreateVSCodeWorkspace()

	return true
}
