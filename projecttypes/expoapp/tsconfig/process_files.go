package tsconfig

import "slices"

func processFiles(tsconfig map[string]any) {
	const eslintConfig = "eslint.config.js"
	const key = "files"

	if value, ok := tsconfig[key].([]any); !ok {
		tsconfig[key] = []string{eslintConfig}
	} else {
		value = append(value, eslintConfig)

		slices.SortFunc(value, func(a any, b any) int {
			s1 := a.(string)
			s2 := b.(string)

			if s1 < s2 {
				return -1
			}

			if s1 > s2 {
				return 1
			}

			return 0
		})

		tsconfig[key] = value
	}
}
