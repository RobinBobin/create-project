package expoapp

import (
	"os/exec"
	"strings"

	"github.com/robinbobin/create-project/utils"
)

func deleteNodeLinkerHoisted(preRunner utils.PreRunner) {
	nodeLinkerBuilder := &strings.Builder{}

	utils.CaptureCmdOutput(&utils.CaptureCmdOutputOptions{
		CmdWithArgs: "pnpm config --location project get node-linker",
		PreRunner:   preRunner,
		Stdout:      nodeLinkerBuilder,
	})

	nodeLinker := strings.TrimSpace(nodeLinkerBuilder.String())

	if nodeLinker != "hoisted" {
		return
	}

	if !utils.AskBool("Would you like to remove 'nodeLinker: hoisted' from 'pnpm-workspace.yaml'?") {
		return
	}

	utils.RunCmdWithPreRunner(
		"pnpm config --location project delete node-linker",
		func(cmd *exec.Cmd) {
			preRunner(cmd)

			// cmd.Stdout = nil
		},
	)
}
