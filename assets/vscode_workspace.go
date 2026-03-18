package assets

import (
	"fmt"

	"github.com/robinbobin/create-project/utils"
)

func AskCreateVSCodeWorkspace(appName string) {
	if !utils.Confirm("Would you like to create a VSCode workspace?", true) {
		return
	}

	fileName := "code-workspace"

	CopyFile(fmt.Sprintf("%v.%v", appName, fileName), fileName)
}
