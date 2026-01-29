package utils

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/acarl005/stripansi"
	"github.com/creack/pty"
	"golang.org/x/sys/unix"
	"golang.org/x/term"
)

type capturedOutput struct {
	cancel  context.CancelFunc
	context context.Context
	// file *os.File
	processCapturedOutput capturedOutputProcessor
}

func (capturedOutput *capturedOutput) Write(p []byte) (n int, err error) {
	stripped := stripansi.Strip(string(p))
	needsMoreStdin := capturedOutput.processCapturedOutput(stripped)

	if !needsMoreStdin {
		fmt.Println("\t\t!!! cancelling !!!")
		capturedOutput.cancel()
	}

	// _, err = fmt.Fprintf(capturedOutput.file, "stripped len: %v\n%v\n", len(stripped), stripped)

	// isFound := strings.Contains(stripped, "What is your app named? …")

	// if isFound {
	// 	close(capturedOutput.done)
	// }

	// _, err = fmt.Fprintf(capturedOutput.file, "Found: %v\n", strconv.FormatBool(isFound))

	return len(p), nil
}

type capturedOutputProcessor = func(strippedOutput string) (needsMoreStdin bool)

func CaptureCmd(
	cmdWithArgs string,
	processCapturedOutput capturedOutputProcessor,
) {
	// The command
	cmdArray := strings.Split(cmdWithArgs, " ")
	cmd := exec.Command(cmdArray[0], cmdArray[1:]...)

	// ptmx
	ptmx, closeptmx := openPTerminal(cmd)
	defer closeptmx()

	// Switch terminal to raw input mode
	switchToPrevious := switchToRaw()
	defer switchToPrevious()

	context, cancel := context.WithCancel(context.Background())

	capturedOutput := &capturedOutput{
		cancel:  cancel,
		context: context,
		// file: file,
		processCapturedOutput: processCapturedOutput,
	}

	// Start stdin -> ptmx goroutine
	go copyStdinToPTerminal(capturedOutput, ptmx)

	captureOutput(capturedOutput, ptmx)

	fmt.Println("\t\t!!! hooray !!!")
}

func captureOutput(capturedOutput *capturedOutput, ptmx *os.File) {
	file, err := os.OpenFile("helpme", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	PanicOnError(err)
	defer file.Close()

	multiWriter := io.MultiWriter(os.Stdout, capturedOutput)

	_, _ = io.Copy(multiWriter, ptmx)
}

func copyStdinToPTerminal(capturedOutput *capturedOutput, ptmx *os.File) {
	defer RecoverFromPanic()

	stdinfd := os.Stdin.Fd()

OUTER:
	for {
		PanicOnError(unix.SetNonblock(int(stdinfd), true))
		_, err := io.Copy(ptmx, os.Stdin)
		PanicOnError(unix.SetNonblock(int(os.Stdin.Fd()), false))

		if errors.Is(err, os.ErrClosed) {
			break OUTER
		}

		if errors.Is(err, syscall.EAGAIN) || errors.Is(err, syscall.EWOULDBLOCK) {
			select {
			case <-capturedOutput.context.Done():
				break OUTER

			default:
				time.Sleep(10 * time.Millisecond)

				continue
			}
		}

		PanicOnError(err)
	}
}

func openPTerminal(cmd *exec.Cmd) (ptmx *os.File, closeptmx func()) {
	ptmx, err := pty.Start(cmd)
	PanicOnError(err)

	closeptmx = func() {
		err := ptmx.Close()

		if err == nil {
			fmt.Println("ptmx closed successfully.")
		} else {
			fmt.Println("ptmx failed to close:", err)
		}
	}

	err = pty.InheritSize(os.Stdin, ptmx)

	if err != nil {
		closeptmx()

		PanicOnError(err)
	}

	return
}

func switchToRaw() (switchToPrevious func()) {
	stdinfd := int(os.Stdin.Fd())
	oldTermState, err := term.MakeRaw(stdinfd)

	PanicOnError(err)

	return func() {
		err := term.Restore(stdinfd, oldTermState)

		if err == nil {
			fmt.Println("\rterminal restored")
		} else {
			fmt.Println("\rterminal failed to be restored:", err)
		}
	}
}
