package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type data string

func main() {
	reader := strings.NewReader("selamat datang kawanku")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(p[:n])) //should handle any remainding bytes.
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(p[:n]))
	}
}
