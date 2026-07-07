package eslintconfig

import (
	"strings"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/utils"
)

func getBaseConfigs() []string {
	entries, err := assets.ReadDir(utils.ESLINT)

	utils.PanicOnError(err)

	baseConfigs := []string{}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		baseConfigs = append(baseConfigs, strings.TrimSuffix(entry.Name(), ".ts"))
	}

	return baseConfigs
}
