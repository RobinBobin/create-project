package expoapp

import (
	"path/filepath"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/utils"
)

func addPrettier() {
	utils.RunCmd("pnpm i --save-dev prettier")

	for _, fileName := range []string{".prettierignore", ".prettierrc.json"} {
		assets.CopyFile(fileName, filepath.Join("prettier", fileName))
	}
}
