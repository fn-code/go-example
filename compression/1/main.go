package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	// data := []byte("selamat datang di dunia ku yang indah ini terima kasih semua nya")
	data := []byte("selamat datang")

	fmt.Println("original : ", len(data))
	b := &bytes.Buffer{}
	gw := gzip.NewWriter(b)
	defer gw.Close()

	_, err := gw.Write(data)
	if err != nil {
		log.Println(err)
	}
	err = gw.Flush()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("compress : ", len(b.Bytes()))

	r := &bytes.Buffer{}
	gr, err := gzip.NewReader(bytes.NewBuffer(b.Bytes()))
	if err != nil {
		log.Println("1 : ", err)
	}
	defer gr.Close()
	buf, err := ioutil.ReadAll(gr)
	if err != nil {
		if err == io.EOF {
			log.Println("eof")
		}
	}
	r.Write(buf)
	fmt.Println("decompress : ", len(r.Bytes()))
}
