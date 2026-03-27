package tsconfig

import "github.com/robinbobin/create-project/options"

func deleteCompilerOptions(tsconfig map[string]any) {
	if options.Options.IsProjectReset {
		delete(tsconfig, "compilerOptions")
	}
}
