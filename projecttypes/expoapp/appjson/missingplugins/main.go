package missingplugins

import (
	"github.com/robinbobin/create-project/utils"
)

func AddMissingPlugins() {
	jsonData := utils.ReadJSON(utils.APP_JSON)

	expo := jsonData["expo"].(map[string]any)
	plugins, _ := expo["plugins"].([]any)

	if plugins == nil {
		plugins = []any{}
	}

	pluginsToAdd := getPluginsToAdd(plugins)

	expo["plugins"] = mergePlugins(plugins, pluginsToAdd)

	utils.WriteJSON(jsonData, utils.APP_JSON)
}
