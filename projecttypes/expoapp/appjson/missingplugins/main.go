package missingplugins

import (
	"github.com/robinbobin/create-project/utils"
)

func AddMissingPlugins() {
	jsonFile := "app.json"
	jsonData := utils.ReadJSON(jsonFile)

	expo := jsonData["expo"].(map[string]any)
	plugins, _ := expo["plugins"].([]any)

	if plugins == nil {
		plugins = []any{}
	}

	pluginsToAdd := getPluginsToAdd(plugins)

	expo["plugins"] = mergePlugins(plugins, pluginsToAdd)

	utils.WriteJSON(jsonData, jsonFile)
}
