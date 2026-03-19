package tsconfig

func deleteCompilerOptions(isProjectReset bool, tsconfig map[string]any) {
	if isProjectReset {
		delete(tsconfig, "compilerOptions")
	}
}
