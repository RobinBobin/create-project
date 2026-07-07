package assets

import (
	"fmt"
)

func CreateVSCodeWorkspace(appName string) {
	fileName := "code-workspace"

	CopyFile(fmt.Sprintf("%v.%v", appName, fileName), fileName)
}
