package expoapp

import (
	"fmt"

	"github.com/robinbobin/create-project/utils"
)

func approveBuilds(preRunner utils.PreRunner) {
	fmt.Println()

	utils.RunCmdWithPreRunner("pnpm approve-builds", preRunner)

	fmt.Println()
}
