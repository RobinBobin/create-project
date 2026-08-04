package assets

import (
	"github.com/robinbobin/create-project/utils"
)

func AddTypescript() {
	CopyFile("tsconfig.base.json", utils.TSCONFIG_JSON)

	utils.PnpmInstall("pnpm install --save-dev typescript@6")
}
