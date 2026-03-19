package tsconfig

import (
	"fmt"
	"os"

	"github.com/robinbobin/create-project/utils"
)

func backup() {
	tsconfigFile, err := os.Open(tsconfig_json)

	utils.CopyFile(fmt.Sprintf("%v.bak", tsconfig_json), tsconfigFile, err)
}
