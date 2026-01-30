package utils

import (
	"os"
	"os/exec"
	"strings"
)

func RunCmd(cmdWithArgs string) {
	RunCmdWithPreRunner(cmdWithArgs, func(_ *exec.Cmd) {})
}

type PreRunner = func(cmd *exec.Cmd)

func RunCmdWithPreRunner(cmdWithArgs string, preRunner PreRunner) {
	cmdArray := strings.Split(cmdWithArgs, " ")
	cmd := exec.Command(cmdArray[0], cmdArray[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	preRunner(cmd)

	PanicOnError(cmd.Run())
}
