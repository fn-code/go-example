package main

import (
	"io"
	"os"
	"strings"
)

type File struct {
	src io.Reader
}

func New(src io.Reader) *File {
	return &File{src}
}

func (a *File) Read(p []byte) (int, error) {
	count, err := a.src.Read(p)
	if err != nil {
		return count, err
	}
	for i := 0; i < len(p); i++ {
		if i == 0 {
			if p[i] >= 'a' && p[i] <= 'z' {
				p[i] = p[i] - 32
			}
		} else {
			if p[i] >= 'A' && p[i] <= 'Z' {
				p[i] = p[i] + 32
			}
		}
	}
	return count, io.EOF
}

func main() {
	var r io.Reader
	r = strings.NewReader("this IS ludin nEnTo")
	// rl := io.LimitReader(r, 12)
	r = New(r)

	w := os.Stdout
	io.Copy(w, r)
}
