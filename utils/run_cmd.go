package utils

import (
	"os"
	"os/exec"
	"strings"
)

func RunCmd(cmdWithArgs string) {
	cmdArray := strings.Split(cmdWithArgs, " ")
	cmd := exec.Command(cmdArray[0], cmdArray[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	PanicOnError(cmd.Run())
}
