package utils

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/acarl005/stripansi"
	"github.com/creack/pty"
	"golang.org/x/term"
)

type capturedOutput struct {
	capturedOutput strings.Builder
	file           *os.File
}

func (capturedOutput *capturedOutput) String() string {
	return capturedOutput.capturedOutput.String()
}

func (capturedOutput *capturedOutput) Write(p []byte) (n int, err error) {
	_, err = capturedOutput.capturedOutput.WriteString(stripansi.Strip(string(p)))
	_, err = capturedOutput.file.WriteString(capturedOutput.capturedOutput.String())
	_, err = capturedOutput.file.WriteString(strconv.FormatBool(strings.Contains(capturedOutput.capturedOutput.String(), "What is your app named? …")))

	return len(p), nil
}

func CaptureCmd(name string, arg ...string) {
	file, err := os.OpenFile("helpme", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	PanicOnError(err)

	defer file.Close()

	cmd := exec.Command(name, arg...)

	ptmx, err := pty.Start(cmd)
	PanicOnError(err)

	defer func() {
		err := ptmx.Close()

		if err == nil {
			fmt.Println("ptmx closed successfully.")
		} else {
			fmt.Println("ptmx failed to close:", err)
		}
	}()

	PanicOnError(pty.InheritSize(os.Stdin, ptmx))

	_, cancel := context.WithCancel(context.Background())

	defer cancel()

	oldTermState, err := term.MakeRaw(int(os.Stdin.Fd()))
	PanicOnError(err)

	defer func() {
		err := term.Restore(int(os.Stdin.Fd()), oldTermState)

		if err == nil {
			fmt.Println("terminal restored")
		} else {
			fmt.Println("terminal failed to be restored:", err)
		}
	}()

	go func() {
		defer RecoverFromPanic()

		io.Copy(ptmx, os.Stdin)

		// 	err := unix.SetNonblock(int(os.Stdin.Fd()), true)

		// 	fmt.Println("unix.SetNonblock(fd, true):", err)

		// OUTER:
		// 	for {
		// 		_, err = io.Copy(ptmx, os.Stdin)

		// 		if errors.Is(err, syscall.EAGAIN) || errors.Is(err, syscall.EWOULDBLOCK) {
		// 			select {
		// 			case <-context.Done():
		// 				break OUTER

		// 			default:
		// 				// if capture
		// 				time.Sleep(10 * time.Millisecond)
		// 			}
		// 		} else {
		// 			PanicOnError(err)
		// 		}
		// 	}

		// 	err = unix.SetNonblock(int(os.Stdin.Fd()), false)

		// 	fmt.Println("unix.SetNonblock(fd, false):", err)
	}()

	capturedOutput := capturedOutput{
		capturedOutput: strings.Builder{},
		file:           file,
	}

	multiWriter := io.MultiWriter(os.Stdout, &capturedOutput)

	_, _ = io.Copy(multiWriter, ptmx)

	fmt.Println("\t\t !!! hooray !!!")
	fmt.Println(capturedOutput.String())
}
