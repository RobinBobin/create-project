package copyconfigs

import (
	"os"
	"path/filepath"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/assets/eslint"
	"github.com/robinbobin/create-project/utils"
)

func Create() bool {
	utils.GitInit()

	assets.AddHusky()
	assets.AddLintStaged()
	assets.AddPrettier()
	assets.AddTypescript()

	eslint.Add()

	// VSCode workspace
	dir, err := os.Getwd()

	utils.PanicOnError(err)

	assets.CreateVSCodeWorkspace(filepath.Base(dir))

	return true
}
