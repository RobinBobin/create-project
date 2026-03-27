package tsconfig

import (
	"fmt"
	"slices"
	"strings"

	"github.com/robinbobin/create-project/options"
)

func processInclude(tsconfig map[string]any) {
	rawInclude := tsconfig["include"].([]any)

	for index := range rawInclude {
		pattern := rawInclude[index].(string)

		if options.Options.IsProjectReset && strings.HasPrefix(pattern, "**") {
			pattern = fmt.Sprint("src/", pattern)
		}

		options.Options.TS.Include = append(options.Options.TS.Include, pattern)
	}

	slices.Sort(options.Options.TS.Include)

	tsconfig["include"] = options.Options.TS.Include
}
