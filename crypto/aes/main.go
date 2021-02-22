// simple AES encryption/decryption example with PBKDF2 key derivation in Go
// from https://gist.github.com/tscholl2/dc7dc15dc132ea70a98e8542fefffa28
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"strings"
)

func getKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		salt = make([]byte, 8)
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
}

func encrypt(passphrase, plaintext string) string {
	key, salt := getKey(passphrase, nil)
	// iv or nonce
	iv := make([]byte, 12)
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	rand.Read(iv)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, []byte(plaintext), nil)
	return fmt.Sprintf("%s-%s-%s", hex.EncodeToString(salt), hex.EncodeToString(iv), hex.EncodeToString(data))
}


func decrypt(passphrase, ciphertext string) string {
	arr := strings.Split(ciphertext, "-")
	salt, _ := hex.DecodeString(arr[0])
	iv, _ := hex.DecodeString(arr[1])
	data, _ := hex.DecodeString(arr[2])
	key, _ := getKey(passphrase, salt)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data, _ = aesgcm.Open(nil, iv, data, nil)
	return string(data)
}

func main() {
	c := encrypt("mykey", "hello world")
	fmt.Println(c)
	fmt.Println(decrypt("mykey", c))
	fmt.Println(decrypt("mykey", "b9a34e828c6c2f7d-095b75208bef13f11ee7a8fb-4a76a09496f37eebe9dc2400dd9bfcc84a5bda598d87aa5a0a3cd3"))
}