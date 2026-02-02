package expoapp

import (
	"os"
	"strings"

	"github.com/charmbracelet/huh"
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

	shouldDelete := true

	utils.PanicOnError(
		huh.NewConfirm().
			Title("Would you like to delete 'nodeLinker: hoisted' from 'pnpm-workspace.yaml'?").
			Value(&shouldDelete).
			Run(),
	)

	if !shouldDelete {
		return
	}

	utils.RunCmd("pnpm config --location project delete node-linker")

	utils.PanicOnError(os.RemoveAll("node_modules"))

	utils.RunCmd("pnpm install")
}
