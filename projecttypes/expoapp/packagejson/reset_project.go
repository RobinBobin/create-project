package packagejson

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/robinbobin/create-project/utils"
)

func resetProject() (isProjectReset bool) {
	isProjectReset = utils.Confirm("Would you like to run `reset-project`?", true)

	if !isProjectReset {
		return
	}

	key := "reset-project"

	// Invoke reset
	utils.RunCmd(fmt.Sprintf("pnpm %v", key))

	// Remove from package.json
	jsonData := utils.ReadJSON(utils.PACKAGE_JSON)
	object := jsonData["scripts"].(map[string]any)

	delete(object, key)

	utils.WriteJSON(jsonData, utils.PACKAGE_JSON)

	// Move sources
	const app = "app"
	const src = "src"

	utils.PanicOnError(os.Mkdir(src, 0775))
	utils.PanicOnError(os.Rename(app, filepath.Join(src, app)))

	return
}
