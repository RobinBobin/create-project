package utils

import (
	"io"
	"os"
)

func CopyFile(destinationName string, source io.ReadCloser, sourceError error) {
	destination, destinationError := os.Create(destinationName)

	defer func() {
		if destination != nil {
			_ = destination.Close()
		}

		if source != nil {
			_ = source.Close()
		}

		panicReason := recover()

		if panicReason == nil {
			return
		}

		_ = os.Remove(destinationName)

		panic(panicReason)
	}()

	PanicOnError(destinationError)
	PanicOnError(sourceError)

	_, err := io.Copy(destination, source)

	PanicOnError(err)
}
