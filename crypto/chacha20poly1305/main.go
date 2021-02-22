// this example from golang.org website
// you can view full implementation three
package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/chacha20poly1305"
	"log"
	"os"
)

func main() {
	pass := "longlonglong"
	salt := "saltsaltsalt"
	key := argon2.Key([]byte(pass), []byte(salt), 3, 32*1024, 4, 32)

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Encryption.
	var encryptedMsg []byte

	msg := []byte("hello this is a lesson")

	// Select a random nonce, and leave capacity for the ciphertext.
	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(msg)+aead.Overhead())
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}

	// Encrypt the message and append the ciphertext to the nonce.
	encryptedMsg = aead.Seal(nonce, nonce, msg, nil)

	fmt.Printf("Base 64 string encode : %s\n", hex.EncodeToString(encryptedMsg))

	// Decryption.
	if len(encryptedMsg) < aead.NonceSize() {
		panic("ciphertext too short")
	}

	// Split nonce and ciphertext.
	nonce, ciphertext := encryptedMsg[:aead.NonceSize()], encryptedMsg[aead.NonceSize():]

	// Decrypt the message and check it wasn't tampered with.
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decode string : %s\n", plaintext)

}
