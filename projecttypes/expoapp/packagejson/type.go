package packagejson

import (
	"fmt"

	"github.com/robinbobin/create-project/utils"
)

func setType() (isESM bool) {
	isESM = utils.Confirm(
		fmt.Sprintf("Would you like to add \"type\": \"module\" to your %v?", utils.PACKAGE_JSON),
		true,
	)

	if !isESM {
		return
	}

	json := utils.ReadJSON(utils.PACKAGE_JSON)

	json["type"] = "module"

	utils.WriteJSON(json, utils.PACKAGE_JSON)

	return
}
