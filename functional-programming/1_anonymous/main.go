package main

import (
	"fmt"
)

func setName() func(string) {
	return func(s string) {
		fmt.Println("your name is : ", s)
	}
}
func main() {

	name := setName()
	name("Ludin nento")

	func(s string) {
		fmt.Println("my name is : ", s)
	}("Ludin Nento")

}
