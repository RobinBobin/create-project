package addition

import (
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
)

func addTypeFest() {
	utils.RunCmd("pnpm install type-fest")

	options.Options.IsInstalled.TypeFest = true
}
