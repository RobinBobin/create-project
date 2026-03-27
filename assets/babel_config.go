package assets

import (
	"errors"
	"fmt"
	"os"

	"github.com/robinbobin/create-project/utils"
)

func CreateBabelConfig() {
	_, err := os.Stat(BABEL_CONFIG_JS)

	if err == nil {
		return
	}

	if !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	CopyFile(BABEL_CONFIG_JS, BABEL_CONFIG_JS)

	utils.RunCmd("pnpm install --save-dev @types/babel__core")

	fmt.Printf("'%v' created.\n", BABEL_CONFIG_JS)
}
