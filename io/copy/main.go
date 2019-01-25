package main

import (
	"io"
	"os"
	"strings"

	"github.com/google/logger"
)

func main() {
	r := strings.NewReader("ludin nento learning golang\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		logger.Error(err)
	}
}
