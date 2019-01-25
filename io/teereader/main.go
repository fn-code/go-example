package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	s := strings.NewReader("hello")
	bufr := new(bytes.Buffer)
	r := io.TeeReader(s, bufr)

	readerMap := make([]byte, s.Len())
	length, err := r.Read(readerMap)
	fmt.Printf("\nBufferRead: %s", bufr)
	fmt.Printf("\nRead: %s", readerMap)
	fmt.Printf("\nLength: %d, Error:%v", length, err)
}
