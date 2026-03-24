package eslintconfig

import (
	"os"
	"regexp"

	"github.com/robinbobin/create-project/utils"
)

func replaceRequire(isModule bool) {
	if !isModule {
		return
	}

	rawContent, err := os.ReadFile(eslintConfigJS)

	utils.PanicOnError(err)

	file, err := os.Create(eslintConfigJS)

	utils.PanicOnError(err)

	defer func() {
		_ = file.Close()
	}()

	// $1 = naming/destructuring, $2 = package path
	reImports := regexp.MustCompile(`const\s+([\w\s{},]+)\s+=\s+require\((?:'|")(.+?)(?:'|")\);?`)

	reExport := regexp.MustCompile(`(module.exports\s+=\s+)`)

	content := string(rawContent)
	content = reImports.ReplaceAllString(content, "import $1 from \"$2\";")
	content = reExport.ReplaceAllString(content, "export default ")

	_, err = file.WriteString(content)

	utils.PanicOnError(err)
}
