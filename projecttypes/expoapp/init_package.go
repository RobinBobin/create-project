package expoapp

import (
	"strings"

	"github.com/robinbobin/create-project/utils"
)

func initPackage() {
	utils.CaptureCmd(
		"pnpm create expo-app --template",
		func(strippedOutput string) (needsMoreStdin bool) {
			return !strings.Contains(strippedOutput, "What is your app amed? …")
		},
	)

	// utils.UsePNPM()

	utils.AskSortJSON("package.json")
}
