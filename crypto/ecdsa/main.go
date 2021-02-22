package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

func main() {
	publicKeyCurve := elliptic.P256()

	privateKey := &ecdsa.PrivateKey{}
	privateKey, err := ecdsa.GenerateKey(publicKeyCurve, rand.Reader)
	if err != nil {
		log.Println(err)
	}

	var publicKey ecdsa.PublicKey
	publicKey = privateKey.PublicKey

	// print private key
	//fmt.Println("Private Key :")
	//fmt.Printf("%x \n", privateKey)

	// print public keys
	//fmt.Println("Public Key :")
	//fmt.Printf("%x \n", publicKey)



	// Sign ecdsa style
	var h hash.Hash
	h = sha256.New()

	io.WriteString(h, "This is to be signed and verified by ECDSA")
	signHash := h.Sum(nil)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, signHash)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	fmt.Printf("Signature : %x\n", signature)
	fmt.Printf("r : %x\n",string(r.Bytes()))
	fmt.Printf("s : %x\n", string(s.Bytes()))

	// Verify with public key
	verifyStatus := ecdsa.Verify(&publicKey, signHash, r, s)
	fmt.Println("verify result : ", verifyStatus)

}
