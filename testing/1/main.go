package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	errEmptyByte   = errors.New("empty header")
	errModeInvalid = errors.New("invalid header mode")
	errMaskInvalid = errors.New("invalid header mask")
)

var (
	mode = byte(0x81)
	mask = byte(0x80)
)

func checkHeader(r []byte) error {
	if len(r) == 0 {
		return errEmptyByte
	}
	if r[0] != mode {
		return errModeInvalid
	}
	if (r[1] & mask) != mask {
		return errMaskInvalid
	}
	return nil
}

func main() {
	frames := []byte{mode, mask | byte(32)}
	err := checkHeader(frames)
	if err != nil {
		log.Printf("error header : %v", err)
		return
	}
	fmt.Println("Success read header")
}
