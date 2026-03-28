package addition

import (
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
)

func addMST() {
	utils.RunCmd("pnpm install mobx mobx-react-lite mobx-state-tree")

	options.Options.IsInstalled.MST = true
}
