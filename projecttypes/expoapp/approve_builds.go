package expoapp

import (
	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/utils"
)

func approveBuilds() {
	shouldApprove := true

	utils.PanicOnError(
		huh.NewConfirm().
			Title("There are ignored build scripts, run 'pnpm approve-builds'?").
			Value(&shouldApprove).
			Run(),
	)

	if !shouldApprove {
		return
	}

	utils.RunCmd("pnpm approve-builds")
}
