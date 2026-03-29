package expoapp

import (
	"fmt"
	"os"
	"strings"

	"github.com/robinbobin/create-project/assets"
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
		CmdWithArgs: fmt.Sprintf("pnpm list %v --parseable", strings.Join(requiredPackages, " ")),
		Stdout:      stdout,
	})

	lineCount := strings.Count(stdout.String(), "\n")

	if (lineCount - 1) != len(requiredPackages) {
		fmt.Println(stdout)
		fmt.Println(lineCount, requiredPackages, strings.Join(requiredPackages, " "))
		return
	}

	if !utils.Confirm("Would you like to copy some common sources?", true) {
		return
	}

	utils.PanicOnError(os.RemoveAll(utils.SRC))

	utils.PanicOnError(assets.CopyFS(utils.SRC, utils.SRC))

	options.Options.AreSourcesCopied = true
}
