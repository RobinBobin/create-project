package utils

import "fmt"

func RecoverFromPanic() {
	panicReason := recover()

	if panicReason != nil {
		fmt.Println(panicReason)
	}
}
