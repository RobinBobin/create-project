package packagejson

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/robinbobin/create-project/utils"
)

func resetProject() {
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

	err := os.Mkdir(utils.SRC, 0775)

	if errors.Is(err, os.ErrExist) {
		return
	}

	utils.PanicOnError(err)
	utils.PanicOnError(os.Rename(app, filepath.Join(utils.SRC, app)))
}
