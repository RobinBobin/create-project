package main

import (
	"fmt"
	"os"

	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
)

func main() {
	flags := parseFlags()

	action := getAction(*flags.projectType)

	isDone := false

	if action.Fn != nil {
		wd, err := os.Getwd()
		utils.PanicOnError(err)

		if action.Fn() {
			isDone = true
		}

		utils.PanicOnError(os.Chdir(wd))
	}

	if !isDone {
		fmt.Println("Bye.")

		return
	}

	fmt.Println("Done.")

	for _, hint := range options.Options.Hints {
		fmt.Print("\n", hint, "\n")
	}
}
