package eslintconfig

import (
	"fmt"
	"strings"

	"github.com/robinbobin/create-project/options"
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

	const dangerouslyAllowAllBuilds = "--dangerously-allow-all-builds"

	cmd := fmt.Sprintf(
		"pnpm install %v --save-dev %v",
		dangerouslyAllowAllBuilds,
		strings.Join(deps, " "),
	)

	if options.Options.CanInstallPackages {
		utils.RunCmd(cmd)

		return
	}

	cmd = strings.Replace(cmd, dangerouslyAllowAllBuilds, "", 1)

	options.Options.Hints.Add(cmd)
}
