package copyconfigs

import (
	"os"
	"path/filepath"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/assets/eslint"
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
)

func Create() bool {
	options.Options.ShouldUseDefaults = false

	// VSCode workspace
	dir, err := os.Getwd()

	utils.PanicOnError(err)

	assets.CreateVSCodeWorkspace(filepath.Base(dir))

	if !utils.Confirm("VSCode workspace file created. Ready to continue?", true) {
		return true
	}

	utils.GitInit()

	assets.AddCommitizen()
	assets.AddCommitlint()
	assets.AddHusky()
	assets.AddLintStaged()
	assets.AddPrettier()
	assets.AddTypescript()

	eslint.Add()

	return true
}
