package tsconfig

import (
	"fmt"

	"github.com/robinbobin/create-project/options"
)

func processCompilerOptions(tsconfig map[string]any) {
	if !options.Options.IsProjectReset {
		return
	}

	const compilerOptionsKey = "compilerOptions"

	if !options.Options.AreSourcesCopied {
		delete(tsconfig, compilerOptionsKey)

		return
	}

	paths := map[string][]string{}

	for _, dirName := range []string{
		"commonComponents",
		"constants",
		"enums",
		"helpers",
		"hocs",
		"hooks",
		"mst",
		"routes",
		"types",
	} {
		paths[fmt.Sprint("@", dirName)] = []string{fmt.Sprint("./src/", dirName)}
	}

	compilerOptions := tsconfig[compilerOptionsKey].(map[string]any)

	delete(compilerOptions, "strict")
	compilerOptions["paths"] = paths
}
