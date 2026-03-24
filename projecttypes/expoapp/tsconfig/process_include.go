package tsconfig

import (
	"fmt"
	"slices"
	"strings"
)

func processInclude(
	include []string,
	isProjectReset bool,
	tsconfig map[string]any,
) {
	rawInclude := tsconfig["include"].([]any)

	for index := range rawInclude {
		pattern := rawInclude[index].(string)

		if isProjectReset && strings.HasPrefix(pattern, "**") {
			pattern = fmt.Sprint("src/", pattern)
		}

		include = append(include, pattern)
	}

	slices.Sort(include)

	tsconfig["include"] = include
}
