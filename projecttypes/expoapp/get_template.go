package expoapp

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
)

func getTemplate() string {
	fmt.Print("Determining the template... ")

	ar := []string{"expo-template-default", "expo"}

	stdout := strings.Builder{}

	for index, packageName := range ar {
		stdout.Reset()

		utils.CaptureCmdOutput(&utils.CaptureCmdOutputOptions{
			CmdWithArgs: fmt.Sprintf("pnpm show %v dist-tags.latest", packageName),
			PreProcessCmd: func(cmd *exec.Cmd) {
				cmd.Env = append(os.Environ(), "npm_config_progress=false")
			},
			Stdout: &stdout,
		})

		ar[index] = strings.Split(stdout.String(), ".")[0]
	}

	version := ar[1]

	if ar[0] != ar[1] && !options.Options.ShouldUseDefaults {
		utils.PanicOnError(
			huh.NewSelect[string]().
				Title("Which Expo version would you like to use?").
				Options(huh.NewOptions(ar...)...).
				Value(&version).
				Run(),
		)
	}

	template := fmt.Sprintf("default@sdk-%v", version)

	fmt.Println(template)

	return fmt.Sprintf("--template %v", template)
}
