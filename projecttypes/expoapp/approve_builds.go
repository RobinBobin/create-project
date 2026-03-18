package expoapp

import (
	"github.com/robinbobin/create-project/utils"
)

func approveBuilds() {
	if !utils.Confirm(
		"There are ignored build scripts, run 'pnpm approve-builds'?",
		true,
	) {
		return
	}

	utils.RunCmd("pnpm approve-builds")
}
