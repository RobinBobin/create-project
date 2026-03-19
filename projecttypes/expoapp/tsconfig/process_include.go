package tsconfig

import (
	"fmt"
	"slices"
	"strings"
)

func processInclude(isProjectReset bool, tsconfig map[string]any) {
	rawInclude := tsconfig["include"].([]any)
	include := make([]string, len(rawInclude))

	for index := range include {
		pattern := rawInclude[index].(string)

		if isProjectReset && strings.HasPrefix(pattern, "**") {
			pattern = fmt.Sprint("src/", pattern)
		}

		include[index] = pattern
	}

	slices.Sort(include)

	tsconfig["include"] = include
}
