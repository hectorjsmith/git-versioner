package util

import "log"

func CheckIfError(err error) {
	if err != nil {
		log.Fatalf("unhandled error: %v", err)
	}
}
