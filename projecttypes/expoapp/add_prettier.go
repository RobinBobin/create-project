package expoapp

import (
	"path"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/utils"
)

func addPrettier() {
	utils.RunCmd("pnpm i --save-dev prettier")

	for _, fileName := range []string{".prettierignore", ".prettierrc.json"} {
		assets.CopyFile(fileName, path.Join("prettier", fileName))
	}
}
