package assets

import "github.com/robinbobin/create-project/utils"

func AddCommitizen() {
	utils.PnpmInstall("pnpm add --save-dev commitizen cz-conventional-changelog")

	const config = ".czrc"

	CopyFile(config, config)
}
