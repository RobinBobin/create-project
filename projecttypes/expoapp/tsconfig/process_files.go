package tsconfig

import (
	"fmt"
	"slices"

	"github.com/robinbobin/create-project/options"
)

func processFiles(tsconfig map[string]any) {
	if len(options.Options.TS.Files) == 0 {
		return
	}

	const key = "files"

	switch rawFiles := tsconfig[key].(type) {
	case []any:
		for _, rawFile := range rawFiles {
			options.Options.TS.AddFile(rawFile.(string))
		}

		slices.Sort(options.Options.TS.Files)

		options.Options.TS.Files = slices.Compact(options.Options.TS.Files)

	case nil:
		slices.Sort(options.Options.TS.Files)

	default:
		panic(fmt.Errorf("\"%v\": \"%v\" is of type \"%T\", equals \"%v\" and can't be parsed", tsconfig_json, key, rawFiles, rawFiles))
	}

	tsconfig[key] = options.Options.TS.Files
}
