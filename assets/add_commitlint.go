package assets

import "github.com/robinbobin/create-project/utils"

func AddCommitlint() {
	utils.PnpmInstall("pnpm add --save-dev @commitlint/cli @commitlint/config-conventional")

	const config = "commitlint.config.js"

	CopyFile(config, config)
}
