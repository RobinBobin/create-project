package eslintconfig

import (
	"strings"

	"github.com/robinbobin/create-project/utils"
)

func run() {
	stdout := &strings.Builder{}

	defer func() {
		reason := recover()

		if reason == nil {
			return
		}

		errorMessage := stdout.String()

		if strings.Contains(errorMessage, `Cannot redefine plugin "react-hooks"`) {
			//

			return
		}

		panic(reason)
	}()

	utils.CaptureCmdOutput(&utils.CaptureCmdOutputOptions{
		CmdWithArgs: "pnpm eslint --fix",
		Stdout:      stdout,
	})
}
