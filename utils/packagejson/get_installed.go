package packagejson

import (
	"regexp"
	"strings"

	"github.com/robinbobin/create-project/utils"
)

func GetInstalled() []string {
	stdout := strings.Builder{}

	utils.CaptureCmdOutput(&utils.CaptureCmdOutputOptions{
		CmdWithArgs: "pnpm list --parseable",
		Stdout:      &stdout,
	})

	re := regexp.MustCompile(`\r?\n`)

	return re.Split(stdout.String(), -1)
}
