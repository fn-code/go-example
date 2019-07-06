package main

import (
	"fmt"
)

func main() {

	in := make(chan Order)
	out := make(chan Order)

	od := GetOrders()
	// wg := sync.WaitGroup{}

	go func() {
		for order := range in {
			out <- Pipeline(order)
		}
		close(out)
	}()

	go func() {
		for _, ods := range od {
			in <- *ods
		}
		close(in)
	}()

	// for {
	// 	select {
	// 	case v, ok := <-out:
	// 		if !ok {
	// 			return
	// 		}
	// 		fmt.Println("The Result is :", v)
	// 	}
	// }

	for v := range out {
		fmt.Println("The Result is :", v)
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

	order3 := &Order{
		10003,
		true,
		"bonsai,secret",
		"EOc3kF/OmxY+dRCaYRrey8h24QoGzVU0/T2QKVCHb1Q=",
		"0123",
		[]LineItem{
			{"Milk", 2},
			{"Sugar", 1},
			{"Salt", 3},
		},
	}
	order4 := &Order{
		10004,
		true,
		"malai,secret",
		"EOc3kF/OmxY+dRCaYRrey8h24QoGzVU0/T2QKVCHb1Q=",
		"0123",
		[]LineItem{
			{"Milk", 2},
			{"Sugar", 1},
			{"Salt", 3},
		},
	}

	orders := []*Order{order1, order2, order3, order4}

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
