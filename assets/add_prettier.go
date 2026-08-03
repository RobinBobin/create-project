package assets

import (
	"path"

	"github.com/robinbobin/create-project/utils"
)

func AddPrettier() {
	utils.RunCmd("pnpm install --save-dev prettier")

	for _, fileName := range []string{".prettierignore", ".prettierrc.json"} {
		CopyFile(fileName, path.Join("prettier", fileName))
	}
}
