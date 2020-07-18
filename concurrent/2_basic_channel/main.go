package main

import (
	"fmt"
)

func send(ch chan<- string) {
	for i := 1; i < 6; i++ {
		fmt.Println("send data ", i)
		ch <- fmt.Sprintf("Data %d", i)
	}
	close(ch)
}

func read(ch chan string, done chan struct{}) {
	var isDone = true
	for isDone {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("channels is closed")
				isDone = false
				done <- struct{}{}
				return
			}
			fmt.Println("-> ", v)
		}
	}
}

func main() {
	var c = make(chan string)
	var done = make(chan struct{})

	go read(c, done)
	send(c)
	<-done

}
