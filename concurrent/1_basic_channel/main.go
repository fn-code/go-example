package main

import (
	"fmt"
	"runtime"
)

var a string
var c = make(chan int)

func f() {
	a = "hello, world"
	close(c)
}

func main() {
	go f()
	numOfGorutine := runtime.NumGoroutine()
	// it's gona output 2. because first goroutine is main, and second is f()
	fmt.Printf("%d is running\n", numOfGorutine)
	<-c
	fmt.Println(a)
}
