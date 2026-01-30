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

type capturedOutputProcessor = func(strippedOutput string) (needsMoreStdin bool)

type capturedOutput struct {
	cancel                  context.CancelFunc
	context                 context.Context
	capturedOutputProcessor capturedOutputProcessor
}

func (capturedOutput *capturedOutput) Write(p []byte) (n int, err error) {
	n = len(p)
	err = nil

	if capturedOutput.capturedOutputProcessor == nil {
		return
	}

	stripped := stripansi.Strip(string(p))
	needsMoreStdin := capturedOutput.capturedOutputProcessor(stripped)

	if !needsMoreStdin {
		capturedOutput.cancel()
	}

	return
}

type CaptureCmdOutputOptions struct {
	CapturedOutputProcessor capturedOutputProcessor
	CmdWithArgs             string
	PreRunner               PreRunner
	Stdout                  io.Writer
}

func CaptureCmdOutput(options *CaptureCmdOutputOptions) {
	// The command
	cmdArray := strings.Split(options.CmdWithArgs, " ")
	cmd := exec.Command(cmdArray[0], cmdArray[1:]...)

	// ptmx
	ptmx, closeptmx := openPTerminal(cmd, options.PreRunner)
	defer closeptmx()

	// Switch terminal to raw input mode
	switchToPrevious := switchToRaw()
	defer switchToPrevious()

	// Captured output
	context, cancel := context.WithCancel(context.Background())

	capturedOutput := &capturedOutput{
		cancel:                  cancel,
		context:                 context,
		capturedOutputProcessor: options.CapturedOutputProcessor,
	}

	// Start stdin -> ptmx goroutine
	go copyStdinToPTerminal(capturedOutput, ptmx)

	stdout := options.Stdout

	if stdout == nil {
		stdout = os.Stdout
	}

	captureOutput(capturedOutput, ptmx, stdout)
}

func captureOutput(
	capturedOutput *capturedOutput,
	ptmx *os.File,
	stdout io.Writer,
) {
	multiWriter := io.MultiWriter(stdout, capturedOutput)

	_, _ = io.Copy(multiWriter, ptmx)
}

func copyStdinToPTerminal(capturedOutput *capturedOutput, ptmx *os.File) {
	defer RecoverFromPanic()

	stdinfd := int(os.Stdin.Fd())

OUTER:
	for {
		PanicOnError(unix.SetNonblock(stdinfd, true))
		_, err := io.Copy(ptmx, os.Stdin)
		PanicOnError(unix.SetNonblock(stdinfd, false))

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
	}
}

func openPTerminal(
	cmd *exec.Cmd,
	preRunner PreRunner,
) (ptmx *os.File, closeptmx func()) {
	if preRunner != nil {
		preRunner(cmd)
	}

	ptmx, err := pty.Start(cmd)
	PanicOnError(err)

	closeptmx = func() {
		err := ptmx.Close()

		if err != nil {
			fmt.Println("\rptmx failed to close:", err)
		}
	}

	err = pty.InheritSize(os.Stdin, ptmx)

	if err != nil {
		closeptmx()

		PanicOnError(err)
	}

	return ptmx, closeptmx
}

func switchToRaw() (switchToPrevious func()) {
	stdinfd := int(os.Stdin.Fd())
	oldTermState, err := term.MakeRaw(stdinfd)

	PanicOnError(err)

	return func() {
		err := term.Restore(stdinfd, oldTermState)

		if err != nil {
			fmt.Println("\rterminal failed to be restored:", err)
		}
	}
}
