package missingplugins

import (
	"slices"
	"strings"
)

func mergePlugins(plugins []any, pluginsToAdd []string) []any {
	for _, plugin := range pluginsToAdd {
		plugins = append(plugins, plugin)
	}

	slices.SortFunc(plugins, func(rawPluginA any, rawPluginB any) int {
		rawPluginAName := getPluginName(rawPluginA)
		rawPluginBName := getPluginName(rawPluginB)

		return strings.Compare(rawPluginAName, rawPluginBName)
	})

	return plugins
}
