package packagejson

import (
	"github.com/robinbobin/create-project/utils"
)

func setModuleType() (shouldAdd bool) {
	if !utils.Confirm(
		"Would you like to add \"type\": \"module\" to package.json?",
		true,
	) {
		return
	}

	fileName := "package.json"
	json := utils.ReadJSON(fileName)

	json["type"] = "module"

	utils.WriteJSON(json, fileName)

	return
}
