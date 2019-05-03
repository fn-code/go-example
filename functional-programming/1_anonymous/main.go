package main

import (
	"fmt"
)

func setName() func(string) {
	return func(s string) {
		fmt.Println("your name is : ", s)
	}
}

func iterate(val ...interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		for _, v := range val {
			c <- v
		}
		close(c)
	}()

	return c
}
func main() {

	name := setName()
	name("Ludin nento")

	func(s string) {
		fmt.Println("my name is : ", s)
	}("Ludin Nento")

	ch := iterate(2, 1, 2, 3, 2)
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("------------------------------------")

	ch1 := make(chan int)
	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	for v := range ch1 {
		fmt.Println(v)
	}
}
