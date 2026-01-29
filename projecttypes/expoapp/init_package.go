package expoapp

import "github.com/robinbobin/create-project/utils"

func initPackage() {
	utils.CaptureCmd("pnpm", "create", "expo-app", "--template")

	// utils.UsePNPM()

	utils.AskSortJSON("package.json")
}
