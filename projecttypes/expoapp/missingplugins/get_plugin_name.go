package missingplugins

import (
	"fmt"

	"github.com/robinbobin/create-project/utils"
)

func getPluginName(rawPlugin any) string {
	switch plugin := rawPlugin.(type) {
	case string:
		return plugin

	case []any:
		return plugin[0].(string)

	default:
		utils.PanicOnError(fmt.Errorf("'getPluginName()': %T, %+v", plugin, plugin))
	}

	panic("unreachable")
}
