package utils

import (
	"fmt"
	"strings"
)

func PnpmInstall(cmdWithArgs string) {
	stdout := strings.Builder{}

	CaptureCmdOutput(&CaptureCmdOutputOptions{
		CmdWithArgs: "pnpm config get packages",
		Stdout:      &stdout,
	})

	packages := strings.TrimSpace(stdout.String())

	cmd := cmdWithArgs

	if packages != "undefined" {
		cmd = fmt.Sprintf("%v --workspace-root", cmd)
	}

	RunCmd(cmd)
}
