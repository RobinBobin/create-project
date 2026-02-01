package missingplugins

import (
	"fmt"
	"slices"
	"strings"

	"github.com/robinbobin/create-project/utils"
)

func getPluginsToCheck() []string {
	pluginsToCheck := []string{
		"expo-font",
		"expo-splash-screen",
		"expo-system-ui",
	}

	stdout := strings.Builder{}

	return slices.DeleteFunc(
		pluginsToCheck,
		func(pluginName string) bool {
			stdout.Reset()

			utils.CaptureCmdOutput(&utils.CaptureCmdOutputOptions{
				CmdWithArgs: fmt.Sprintf("pnpm list %v --parseable", pluginName),
				Stdout:      &stdout,
			})

			return len(stdout.String()) == 0
		},
	)
}
