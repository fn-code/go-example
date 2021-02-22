package main

import (
	"fmt"
	"index/suffixarray"
)

func main() {
	data := []byte("banana")
	index := suffixarray.New([]byte("banana"))
	offsets := index.Lookup([]byte("ana"), -1)
	for _, off := range offsets {
		fmt.Println(off)
		fmt.Println(string(data[off:]))
	}

}
