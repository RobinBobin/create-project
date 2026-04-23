package packagejson

import (
	"encoding/json"
	"maps"
	"slices"
	"strings"

	"github.com/robinbobin/create-project/utils"
)

func GetInstalled() []string {
	stdout := strings.Builder{}

	utils.CaptureCmdOutput(&utils.CaptureCmdOutputOptions{
		CmdWithArgs: "pnpm list --json",
		Stdout:      &stdout,
	})

	jsonData := []map[string]any{}

	utils.PanicOnError(json.Unmarshal([]byte(stdout.String()), &jsonData))

	jsonDatum := jsonData[0]

	keys := []string{"dependencies", "devDependencies"}

	installed := make([]string, 0, 30)

	for _, key := range keys {
		installed = append(
			installed,
			slices.Collect(maps.Keys(jsonDatum[key].(map[string]any)))...,
		)
	}

	slices.Sort(installed)

	return installed
}
