package assets

import (
	"github.com/robinbobin/create-project/utils"
)

func AddLintStaged() {
	utils.PnpmInstall("pnpm install --save-dev lint-staged")

	const lintstagedrc = ".lintstagedrc.json"

	CopyFile(lintstagedrc, lintstagedrc)
}
