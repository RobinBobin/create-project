package main

import (
	"fmt"
	"os"

	"github.com/robinbobin/create-project/utils"
)

func main() {
	flags := parseFlags()

	action := getAction(*flags.projectType)

	farewell := "Bye."

	if action.Fn != nil {
		wd, err := os.Getwd()
		utils.PanicOnError(err)

		if action.Fn() {
			farewell = "Done."
		}

		utils.PanicOnError(os.Chdir(wd))
	}

	fmt.Println(farewell)
}
