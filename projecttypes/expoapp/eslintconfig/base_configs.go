package eslintconfig

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/utils"
)

type config struct {
	fileName string
}

func (config config) String() string {
	parts := strings.Split(
		strings.TrimSuffix(config.fileName, ".js"),
		".",
	)

	for index := range parts {
		parts[index] = strings.ToUpper(parts[index][:1]) + parts[index][1:]
	}

	return strings.Join(parts, " ")
}

func UseBaseConfigs() {
	const eslint = "eslint"

	// Get the file list from `assets/eslint`
	entries, err := assets.ReadDir(eslint)

	utils.PanicOnError(err)

	configs := []config{}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		configs = append(configs, config{
			fileName: entry.Name(),
		})
	}

	utils.PanicOnError(
		huh.NewMultiSelect[config]().
			Title("Would you like to add any of these custom ESLint configs?").
			Options(huh.NewOptions(configs...)...).
			Value(&configs).
			Run(),
	)

	if len(configs) == 0 {
		return
	}

	utils.PanicOnError(os.Mkdir(eslint, 0775))

	re := regexp.MustCompile(`(?s)import\s+.*?\s+from\s+['"]([^'"]+)['"]`)

	for _, config := range configs {
		fileName := filepath.Join(eslint, config.fileName)

		// Copy the file
		assets.CopyFile(fileName, fileName)

		// Read the file to get the imports
		rawData, err := os.ReadFile(fileName)

		utils.PanicOnError(err)

		data := string(rawData)

		// Get package names from imports
		matches := re.FindAllStringSubmatch(data, -1)
		packages := []string{}

		for _, match := range matches {
			parts := strings.Split(match[1], "/")

			if strings.HasPrefix(parts[0], ".") {
				// Relative import
				dir := filepath.Join(eslint, parts[1])
				err := assets.CopyFS(dir, dir)

				if !errors.Is(err, os.ErrExist) {
					utils.PanicOnError(err)
				}

				continue
			}

			var packageName string

			if strings.HasPrefix(parts[0], "@") && len(parts) >= 2 {
				// Scoped package
				packageName = strings.Join(parts[:2], "/")
			} else {
				// Unscoped package
				packageName = parts[0]
			}

			packages = append(packages, packageName)
		}

		utils.RunCmd(
			fmt.Sprint(
				"pnpm install --save-dev ",
				strings.Join(packages, " "),
			),
		)
	}
}
