package missingplugins

import (
	"path/filepath"

	"github.com/robinbobin/create-project/utils"
)

func AddMissingPlugins(appPath string) {
	jsonFile := filepath.Join(appPath, "app.json")
	jsonData := utils.ReadJSON(jsonFile)

	expo := jsonData["expo"].(map[string]any)
	plugins := expo["plugins"].([]any)

	pluginsToAdd := getPluginsToAdd(plugins)

	expo["plugins"] = mergePlugins(plugins, pluginsToAdd)

	utils.WriteJSON(jsonData, jsonFile)
}
