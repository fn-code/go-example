package main

import (
	"fmt"
)

func main() {
	od := GetOrders()
	numOfOrders := len(od)
	in := make(chan Order, numOfOrders)
	out := make(chan Order, numOfOrders)

	for i := 0; i < numOfOrders; i++ {
		go func() {
			for order := range in {
				out <- Pipeline(order)
			}
		}()
	}

	for _, ods := range od {
		in <- *ods
	}

	close(in)
	for i := 0; i < numOfOrders; i++ {
		fmt.Println("The Result is :", <-out)
	}
}

type Order struct {
	OrderNumber  int
	IsValid      bool
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
		true,
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
		true,
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

func Pipeline(o Order) Order {
	o = Authenticate(o)
	o = Decrypt(o)
	o = Charge(o)
	return o
}

func Authenticate(o Order) Order {
	fmt.Printf("Order %d is Authtenticated\n", o.OrderNumber)
	return o
}

func Decrypt(o Order) Order {
	fmt.Printf("Order %d is Decrypted\n", o.OrderNumber)
	return o
}

func Charge(o Order) Order {
	fmt.Printf("Order %d is Charged\n", o.OrderNumber)
	return o
}
