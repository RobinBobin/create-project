package packagejson

import (
	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/utils"
)

func setModuleType() (shouldAdd bool) {
	shouldAdd = true

	utils.PanicOnError(
		huh.NewConfirm().
			Title("Would you like to add \"type\": \"module\" to package.json?").
			Value(&shouldAdd).
			Run(),
	)

	if !shouldAdd {
		return
	}

	fileName := "package.json"
	json := utils.ReadJSON(fileName)

	json["type"] = "module"

	utils.WriteJSON(json, fileName)

	return
}
