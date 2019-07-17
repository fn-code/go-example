package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

func main() {
	pipeline := BuildPipline(Decrypt{}, Authenticate{}, Charge{})

	time.Sleep(5 * time.Second)
	go func() {
		od := GetOrders()
		for _, order := range od {
			pipeline.Send(*order)
		}
		fmt.Println("Close pipeline")
		pipeline.Close()
	}()

	pipeline.Receive(func(o Order) {
		log.Printf("Received: %v", o)
	})
}

type Order struct {
	OrderNumber  int
	IsValid      bool
	IsDecrypted  bool
	Credentials  string
	CCardNumber  string
	CCardExpDate string
	LineItems    []LineItem
}

type LineItem struct {
	Descriptions string
	Count        int
}

func GetOrders() []*Order {
	order1 := &Order{
		10001,
		false,
		false,
		"alice,secret",
		"7b/HWvtIB9a16AYk+Yv6WWwer3GFbxpjoR+GO9iHIYY=",
		"0922",
		[]LineItem{
			{"Apples", 1},
			{"Oranges", 4},
		},
	}

	order2 := &Order{
		10002,
		false,
		false,
		"bob,secret",
		"EOc3kF/OmxY+dRCaYRrey8h24QoGzVU0/T2QKVCHb1Q=",
		"0123",
		[]LineItem{
			{"Milk", 2},
			{"Sugar", 1},
			{"Salt", 3},
		},
	}

	orders := []*Order{order1, order2}
	return orders

}

var EncryptionKey = "a very very very very secret key"

func encrypt(plain string) (string, error) {
	rwBytes := []byte(plain)
	block, err := aes.NewCipher([]byte(EncryptionKey))
	if err != nil {
		return "", err
	}

	if len(rwBytes)%aes.BlockSize != 0 {
		padding := aes.BlockSize - len(rwBytes)%aes.BlockSize
		padText := bytes.Repeat([]byte{byte(0)}, padding)
		rwBytes = append(rwBytes, padText...)
	}

	chiperText := make([]byte, aes.BlockSize+len(rwBytes))
	iv := chiperText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(chiperText[aes.BlockSize:], rwBytes)
	return base64.StdEncoding.EncodeToString(chiperText), nil

}

func decrypt(encodedValue string) (string, error) {
	block, err := aes.NewCipher([]byte(EncryptionKey))
	if err != nil {
		return "", err
	}
	b, err := base64.StdEncoding.DecodeString(encodedValue)
	if err != nil {
		return "", err
	}
	if len(b) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := b[:aes.BlockSize]
	b = b[aes.BlockSize:]
	if len(b)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(b, b)
	b = bytes.TrimRight(b, "\x00")
	return string(b), nil
}

type Filterer interface {
	Filter(chan Order) chan Order
}

type Authenticate struct{}

func (Authenticate) Filter(input chan Order) chan Order {
	out := make(chan Order)
	go func() {
		for order := range input {
			userPwd := strings.Split(order.Credentials, ",")
			if userPwd[1] == "secret" {
				order.IsValid = true
				out <- order
			} else {
				order.IsValid = false
				errMsg := fmt.Sprintf("Error : Invalid password for order ID: %d", order.OrderNumber)
				log.Println(errMsg)
				out <- order
			}
		}
		fmt.Println("->> Auth")
		close(out)
	}()
	fmt.Println("-> Auth")
	return out
}

type Decrypt struct{}

func (Decrypt) Filter(input chan Order) chan Order {
	out := make(chan Order)
	go func() {
		for order := range input {
			creditCard, err := decrypt(order.CCardNumber)
			if err != nil {
				order.IsDecrypted = false
				log.Println("Error : ", err.Error())
			} else {
				order.IsDecrypted = true
				order.CCardNumber = creditCard
				out <- order
			}
		}
		fmt.Println("->> Dec")
		close(out)
	}()
	fmt.Println("-> Dec")
	return out
}

func ChargeCard(ccardNo string) {
	fmt.Printf("Credit card %v charged \n", ccardNo)
}

type Charge struct{}

func (Charge) Filter(input chan Order) chan Order {
	out := make(chan Order)
	go func() {
		for order := range input {
			if order.IsValid && order.IsDecrypted {
				ChargeCard(order.CCardNumber)
				out <- order
			} else {
				errMsg := fmt.Sprintf("Error: Unable to charge order Id: %d", order.OrderNumber)
				log.Println("Error:", errors.New(errMsg))
			}
		}
		fmt.Println("->> Charge")
		close(out)
	}()
	fmt.Println("-> Charge")
	return out
}

type Filter struct {
	in  chan Order
	out chan Order
}

func BuildPipline(filters ...Filterer) Filter {
	source := make(chan Order)
	var nextFilter chan Order

	for _, filter := range filters {
		if nextFilter == nil {
			nextFilter = filter.Filter(source)
		} else {
			nextFilter = filter.Filter(nextFilter)
		}
	}

	return Filter{in: source, out: nextFilter}
}

func (f *Filter) Send(order Order) {
	f.in <- order
}

func (f *Filter) Receive(callback func(Order)) {
	for o := range f.out {
		callback(o)
	}
}

func (f *Filter) Close() {
	close(f.in)
}
