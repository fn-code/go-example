package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 0, 16)

	if _, err := s.Seek(-4, io.SeekEnd); err != nil {
		log.Println(err)
	}

	buf := make([]byte, 4)
	if _, err := s.Read(buf); err != nil {
		log.Println(err)
	}
	fmt.Printf("%s\n", buf)
}
