package expoapp

import (
	"fmt"

	"github.com/robinbobin/create-project/utils"
)

func approveBuilds() {
	shouldApprove := utils.AskBool("There are ignored build scripts, run 'pnpm approve-builds'?")

	fmt.Println()

	if !shouldApprove {
		return
	}

	utils.RunCmd("pnpm approve-builds")

	fmt.Println()
}
