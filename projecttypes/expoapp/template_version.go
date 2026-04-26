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

func getTemplateVersion() string {
	fmt.Println("Determining the template version...")

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

	if ar[0] == ar[1] {
		return "latest"
	}

	version := ar[1]

	if !options.Options.ShouldUseDefaults {
		utils.PanicOnError(
			huh.NewSelect[string]().
				Title("Which Expo version would you like to use?").
				Options(huh.NewOptions(ar...)...).
				Value(&version).
				Run(),
		)
	}

	return fmt.Sprint("sdk-", version)
}
