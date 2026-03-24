package eslintconfig

import (
	"regexp"
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

		directoryImportRE := regexp.MustCompile(`Directory import '[^']*node_modules/([^']+)`)

		matches := directoryImportRE.FindStringSubmatch(errorMessage)

		if len(matches) > 1 {
			run()

			return
		}

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
