package eslint

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/utils"
)

func copyBaseConfigs(configNames []string) {
	err := assets.CopyFS(utils.ESLINT, utils.ESLINT)

	utils.PanicOnError(err)

	file, err := os.OpenFile(filepath.Join(utils.ESLINT, "index.ts"), os.O_CREATE|os.O_WRONLY, 0664)

	utils.PanicOnError(err)

	defer func() {
		_ = file.Close()
	}()

	buffer := bytes.Buffer{}

	for _, configName := range configNames {
		fmt.Fprintf(&buffer, "export { default as %v } from './%v'\n", configName, configName)
	}

	_, err = file.WriteString(buffer.String())

	utils.PanicOnError(err)
}
