package eslintconfig

import (
	"os"
	"regexp"
	"strings"

	"github.com/robinbobin/create-project/utils"
)

func handleModuleType() {
	rawContent, err := os.ReadFile(utils.ESLINT_CONFIG_JS)

	utils.PanicOnError(err)

	file, err := os.Create(utils.ESLINT_CONFIG_JS)

	utils.PanicOnError(err)

	defer func() {
		_ = file.Close()
	}()

	// $1 = naming/destructuring, $2 = package path
	reImports := regexp.MustCompile(`const\s+([\w\s{},]+)\s+=\s+require\((?:'|")(.+?)(?:'|")\);?`)

	reExport := regexp.MustCompile(`module.exports\s+=\s+`)

	content := string(rawContent)
	content = reImports.ReplaceAllString(content, "import $1 from \"$2\";")
	content = reExport.ReplaceAllString(content, "export default ")

	content = strings.Replace(
		content,
		"eslint-config-expo/flat",
		"eslint-config-expo/flat.js",
		1,
	)

	_, err = file.WriteString(content)

	utils.PanicOnError(err)
}
