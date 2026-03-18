package packagejson

import (
	"fmt"

	"github.com/robinbobin/create-project/utils"
)

func setModuleType() (shouldAdd bool) {
	if !utils.Confirm(
		fmt.Sprintf("Would you like to add \"type\": \"module\" to %v?", utils.PACKAGE_JSON),
		true,
	) {
		return
	}

	json := utils.ReadJSON(utils.PACKAGE_JSON)

	json["type"] = "module"

	utils.WriteJSON(json, utils.PACKAGE_JSON)

	return
}
