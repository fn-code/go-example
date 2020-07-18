package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"strings"
)


// Checksum with string
func checksum1() {
	str := "This is my live"

	md5 := md5.Sum([]byte(str))
	sha1 := sha1.Sum([]byte(str))
	sha256 := sha256.Sum256([]byte(str))

	fmt.Printf("md5 -> %x\n", md5)
	fmt.Printf("sha1 -> %x\n", sha1)
	fmt.Printf("sha256 -> %x\n", sha256)
}

// Checksum with input stream data
func checksum2()  {
	input := strings.NewReader("This is my live")

	hash := sha256.New()
	if _, err := io.Copy(hash, input); err != nil {
		log.Fatal(err)
	}
	sum := hash.Sum(nil)

	fmt.Printf("sha256 -> %x\n", sum)
}

func main() {

	checksum1()
	checksum2()
}
