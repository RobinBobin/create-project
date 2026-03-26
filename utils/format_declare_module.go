package utils

import "fmt"

func FormatDeclareModule(moduleName string) string {
	return fmt.Sprintf("declare module '%v'", moduleName)
}
