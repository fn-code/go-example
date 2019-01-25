package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		_, err := fmt.Fprintln(pw, "hello")
		if err != nil {
			log.Fatal(err)
		}
	}()

	io.Copy(os.Stdout, pr)
}
