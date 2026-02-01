package expoapp

import (
	"fmt"
	"strings"

	"github.com/robinbobin/create-project/utils"
)

func deleteNodeLinkerHoisted() {
	nodeLinkerBuilder := &strings.Builder{}

	utils.CaptureCmdOutput(&utils.CaptureCmdOutputOptions{
		CmdWithArgs: "pnpm config --location project get node-linker",
		Stdout:      nodeLinkerBuilder,
	})

	nodeLinker := strings.TrimSpace(nodeLinkerBuilder.String())

	if nodeLinker != "hoisted" {
		return
	}

	fmt.Println()

	if !utils.AskBool("Would you like to remove 'nodeLinker: hoisted' from 'pnpm-workspace.yaml'?") {
		return
	}

	utils.RunCmd("pnpm config --location project delete node-linker")
}
