package eslintconfigexpo

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/robinbobin/create-project/utils"
)

func commentOut() {
	// Process eslint.config.js
	rawContent, err := os.ReadFile(utils.ESLINT_CONFIG_JS)

	utils.PanicOnError(err)

	content := string(rawContent)

	re := regexp.MustCompile(fmt.Sprintf(`const\s+([\w\s{},]+)\s+=\s+require\((?:'|")%v.*`, eslintConfigExpo))

	matches := re.FindStringSubmatch(content)

	if len(matches) < 2 {
		fmt.Printf("Failed to comment out '%v' usage.\n", eslintConfigExpo)

		return
	}

	configUsage := fmt.Sprintf("%v,", matches[1])

	content = strings.Replace(
		content,
		matches[0],
		fmt.Sprintf("// %v", matches[0]),
		1,
	)

	content = strings.Replace(
		content,
		configUsage,
		fmt.Sprintf("// %v", configUsage),
		1,
	)

	eslintConfig, err := os.Create(utils.ESLINT_CONFIG_JS)

	utils.PanicOnError(err)

	defer func() {
		_ = eslintConfig.Close()
	}()

	_, err = eslintConfig.WriteString(content)

	utils.PanicOnError(err)

	// Process custom.d.ts if necessary
	rawContent, err = os.ReadFile(utils.CUSTOM_D_TS)

	utils.PanicOnError(err)

	content = string(rawContent)

	re = regexp.MustCompile(
		utils.FormatDeclareModule(
			fmt.Sprintf("%v.*", eslintConfigExpo),
		),
	)

	matches = re.FindStringSubmatch(content)

	if len(matches) < 1 {
		return
	}

	content = strings.Replace(
		content,
		matches[0],
		fmt.Sprintf("// %v", matches[0]),
		1,
	)

	customDTS, err := os.Create(utils.CUSTOM_D_TS)

	utils.PanicOnError(err)

	defer func() {
		_ = customDTS.Close()
	}()

	_, err = customDTS.WriteString(content)

	utils.PanicOnError(err)

	fmt.Println("Fixed.")
}
