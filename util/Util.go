package util

import (
	"io"
	"log"
	"os"
)

func CheckIfError(err error) {
	if err != nil {
		log.Printf("Unhandled error: %v", err)
		panic(err)
	}
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() { CheckIfError(in.Close()) }()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() { CheckIfError(out.Close()) }()

	_, err = io.Copy(out, in)
	return err
}

func CountTrueValues(b ...bool) int {
	n := 0
	for _, v := range b {
		if v {
			n++
		}
	}
	return n
}
