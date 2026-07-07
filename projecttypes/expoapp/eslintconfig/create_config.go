package eslintconfig

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
)

func createConfig(baseConfigs []string) {
	buffer := bytes.Buffer{}
	buffer.Grow(500)

	data, err := os.ReadFile(utils.ESLINT_CONFIG_TS)

	if err == nil {
		buffer.WriteString("Previous ESLint config:\n")
		buffer.Write(data)

		options.Options.Hints.Add(buffer.String())
	} else if !errors.Is(err, os.ErrNotExist) {
		utils.PanicOnError(err)
	}

	file, err := os.OpenFile(utils.ESLINT_CONFIG_TS, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)

	utils.PanicOnError(err)

	defer func() {
		_ = file.Close()
	}()

	buffer.Reset()

	buffer.WriteString("import { defineConfig, includeIgnoreFile } from 'eslint/config'\n")
	buffer.WriteString("import { fileURLToPath } from 'node:url'\n")
	buffer.WriteString("\n")
	buffer.WriteString("import { ")

	for _, baseConfig := range baseConfigs {
		fmt.Fprintf(&buffer, "%v,", baseConfig)
	}

	buffer.WriteString(" } from './eslint'\n")
	buffer.WriteString("\n")
	buffer.WriteString("const gitignore = fileURLToPath(new URL('.gitignore', import.meta.url))\n")
	buffer.WriteString("\n")
	buffer.WriteString("export default defineConfig(\n")
	buffer.WriteString("includeIgnoreFile(gitignore),\n")

	for _, baseConfig := range baseConfigs {
		fmt.Fprintf(&buffer, "%v,\n", baseConfig)
	}

	buffer.WriteString(")\n")

	_, err = file.WriteString(buffer.String())

	utils.PanicOnError(err)
}
