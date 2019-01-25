package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/logger"
)

func main() {
	r := strings.NewReader("First Reader\n")
	r2 := strings.NewReader("Second Reader\n")
	buf := make([]byte, 5)

	n, err := io.CopyBuffer(os.Stdout, r, buf)
	if err != nil {
		logger.Error(err)
	}
	fmt.Println(n)

	k, err := io.CopyBuffer(os.Stdout, r2, buf)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(k)
}
