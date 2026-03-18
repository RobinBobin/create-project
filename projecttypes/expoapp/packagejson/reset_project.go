package packagejson

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/robinbobin/create-project/utils"
)

func resetProject(isTypeSet bool) bool {
	if !utils.Confirm("Would you like to run `reset-project`?", true) {
		return false
	}

	key := "reset-project"

	if isTypeSet {
		cjs := ".cjs"
		js := ".js"
		scripts := "scripts"
		pathFile := filepath.Join(scripts, key)

		utils.PanicOnError(os.Rename(fmt.Sprint(pathFile, js), fmt.Sprint(pathFile, cjs)))

		jsonData := utils.ReadJSON(utils.PACKAGE_JSON)
		object := jsonData[scripts].(map[string]any)

		var cmd string

		for objectKey, value := range object {
			if objectKey == key {
				cmd = strings.ReplaceAll(value.(string), js, cjs)

				break
			}
		}

		object[key] = cmd

		utils.WriteJSON(jsonData, utils.PACKAGE_JSON)
	}

	utils.RunCmd(fmt.Sprint("pnpm", " ", key))

	return true
}
