package main

import "fmt"

func a() {
	fmt.Println("function a")
}
func b() {
	fmt.Println("function b")
}
func c() {
	fmt.Println("function c")
}

type d func()

func main() {

	f := []d{a, b, c}
	limit := make(chan int, 3)
	for i, w := range f {
		go func(w func(), i int) {
			limit <- i
			w()
			<-limit
		}(w, i)
	}
	select {}
}
