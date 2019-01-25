package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing path of image")
		os.Exit(2)
	}
	path := os.Args[1]
	if err := cat(path); err != nil {
		fmt.Fprintf(os.Stderr, "could not find the %s path: %v\n", path, err)
	}

}

func cat(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Printf("\033]1337;File=inline=1;")
	wc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	_, err = io.Copy(wc, f)
	if err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	fmt.Printf("\a\n")
	return nil
}
