package copyconfigs

import (
	"os"
	"path/filepath"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/assets/eslint"
	"github.com/robinbobin/create-project/utils"
)

func Create() bool {
	// ESLint
	eslint.Add()

	// Prettier
	assets.AddPrettier()

	// TS
	assets.CopyFile("tsconfig.base.json", utils.TSCONFIG_JSON)
	utils.PnpmInstall("pnpm install --save-dev typescript@6")

	// VSCode workspace
	dir, err := os.Getwd()

	utils.PanicOnError(err)

	assets.CreateVSCodeWorkspace(filepath.Base(dir))

	// husky
	utils.PnpmInstall("pnpm install --save-dev husky")
	utils.RunCmd("pnpm exec husky init")

	// lint-staged
	const lintstagedrc = ".lintstagedrc.json"

	assets.CopyFile(lintstagedrc, lintstagedrc)
	utils.PnpmInstall("pnpm install --save-dev lint-staged")

	return true
}
