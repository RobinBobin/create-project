package utils

import (
	"os"
	"os/exec"
)

func RunCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	PanicOnError(cmd.Run())
}
