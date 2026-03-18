package expoapp

import (
	"os"
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

	if !utils.Confirm("Would you like to delete 'nodeLinker: hoisted' from 'pnpm-workspace.yaml'?", true) {
		return
	}

	utils.RunCmd("pnpm config --location project delete node-linker")

	utils.PanicOnError(os.RemoveAll("node_modules"))

	utils.RunCmd("pnpm install")
}
