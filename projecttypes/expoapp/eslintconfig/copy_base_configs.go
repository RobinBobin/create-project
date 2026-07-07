package eslintconfig

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/utils"
)

func copyBaseConfigs(baseConfigs []string) {
	err := assets.CopyFS(utils.ESLINT, utils.ESLINT)

	utils.PanicOnError(err)

	file, err := os.OpenFile(path.Join(utils.ESLINT, "index.ts"), os.O_CREATE|os.O_WRONLY, 0664)

	utils.PanicOnError(err)

	defer func() {
		_ = file.Close()
	}()

	buffer := bytes.Buffer{}

	for _, baseConfig := range baseConfigs {
		fmt.Fprintf(&buffer, "export { default as %v } from './%v'\n", baseConfig, baseConfig)
	}

	_, err = file.WriteString(buffer.String())

	utils.PanicOnError(err)
}
