package addition

import (
	"path/filepath"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/utils"
)

func addESLint() {
	utils.RunCmd("pnpm install --save-dev eslint eslint-import-resolver-typescript")

	assets.CopyFile(
		utils.ESLINT_CONFIG_JS,
		filepath.Join(utils.ESLINT, utils.ESLINT_CONFIG_JS),
	)
}
