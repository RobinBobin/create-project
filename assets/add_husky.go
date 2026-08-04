package assets

import (
	"os"

	"github.com/robinbobin/create-project/utils"
)

func AddHusky() {
	utils.PnpmInstall("pnpm install --save-dev husky")
	utils.RunCmd("pnpm exec husky init")

	const husky = ".husky"

	utils.PanicOnError(os.RemoveAll(husky))
	utils.PanicOnError(CopyFS(husky, husky))
}
