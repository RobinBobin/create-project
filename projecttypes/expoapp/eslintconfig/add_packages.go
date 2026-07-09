package eslintconfig

import (
	"fmt"
	"strings"

	"github.com/robinbobin/create-project/utils"
)

func addPackages(dependencies []string) {
	stdout := strings.Builder{}

	utils.CaptureCmdOutput(&utils.CaptureCmdOutputOptions{
		CmdWithArgs: "node -v",
		Stdout:      &stdout,
	})

	nodeVersion := strings.TrimPrefix(strings.SplitN(stdout.String(), ".", 2)[0], "v")

	deps := append(
		dependencies,
		"eslint-import-resolver-typescript",
		"jiti",
		fmt.Sprintf("@types/node@%v", nodeVersion),
	)

	utils.RunCmd(
		fmt.Sprintf(
			"pnpm install --dangerously-allow-all-builds --save-dev %v",
			strings.Join(deps, " "),
		),
	)
}
