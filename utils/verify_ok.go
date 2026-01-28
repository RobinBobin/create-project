package utils

import "log"

func VerifyOK(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
