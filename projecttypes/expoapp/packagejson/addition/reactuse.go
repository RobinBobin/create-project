package addition

import (
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
)

func addReactUse() {
	utils.RunCmd("pnpm install react-use")

	options.Options.IsInstalled.ReactUse = true
}
