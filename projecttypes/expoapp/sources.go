package expoapp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
)

func copySources() {
	if !options.Options.IsProjectReset {
		return
	}

	requiredPackages := []string{
		"mobx-state-tree",
		"radashi",
		"react-native-paper",
		"type-fest",
	}

	stdout := &strings.Builder{}

	utils.CaptureCmdOutput(&utils.CaptureCmdOutputOptions{
		CmdWithArgs: fmt.Sprintf("pnpm list %v | wc -l", strings.Join(requiredPackages, " ")),
		Stdout:      stdout,
	})

	lineCount, err := strconv.Atoi(stdout.String())

	utils.PanicOnError(err)

	if (lineCount - 1) != len(requiredPackages) {
		return
	}

	if !utils.Confirm("Would you like to copy some common sources?", true) {
		return
	}
}
