package assets

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/utils"
)

func AskCreateVSCodeWorkspace(appName string) {
	shouldCreate := true

	utils.PanicOnError(
		huh.NewConfirm().
			Title("Would you like to create a VSCode workspace?").
			Value(&shouldCreate).
			Run(),
	)

	if !shouldCreate {
		return
	}

	fileName := "code-workspace"

	CopyFile(fmt.Sprintf("%v.%v", appName, fileName), fileName)
}
