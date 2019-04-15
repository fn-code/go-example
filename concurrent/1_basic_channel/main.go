package main

var a string
var c = make(chan int, 10)

func f() {
	a = "hello, world"
	close(c)
}

func main() {
	go f()
	<-c
	print(a)
}
