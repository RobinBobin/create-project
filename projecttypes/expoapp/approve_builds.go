package expoapp

import (
	"fmt"

	"github.com/robinbobin/create-project/utils"
)

func approveBuilds(preRunner utils.PreRunner) {
	fmt.Println()

	shouldApprove := utils.AskBool("There are ignored build scripts, run 'pnpm approve-builds'?")

	fmt.Println()

	if !shouldApprove {
		return
	}

	utils.RunCmdWithPreRunner("pnpm approve-builds", preRunner)

	fmt.Println()
}
