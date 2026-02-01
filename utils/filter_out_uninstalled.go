package utils

import (
	"fmt"
	"slices"
	"strings"
)

func FilterOutUninstalled(packageNames []string) []string {
	stdout := strings.Builder{}

	return slices.DeleteFunc(
		packageNames,
		func(pluginName string) bool {
			stdout.Reset()

			CaptureCmdOutput(&CaptureCmdOutputOptions{
				CmdWithArgs: fmt.Sprintf("pnpm list %v --parseable", pluginName),
				Stdout:      &stdout,
			})

			return len(stdout.String()) == 0
		},
	)
}
