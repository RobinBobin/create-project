package assets

import (
	"errors"
	"fmt"
	"os"

	"github.com/robinbobin/create-project/utils"
)

func CreateVSCodeWorkspace(appName string) {
	dst := fmt.Sprintf("%v.%v", appName, utils.VS_CODE_WORKSPACE)

	_, err := os.Stat(dst)

	if errors.Is(err, os.ErrExist) {
		return
	}

	if errors.Is(err, os.ErrNotExist) {
		CopyFile(dst, utils.VS_CODE_WORKSPACE)

		return
	}

	utils.PanicOnError(err)
}
