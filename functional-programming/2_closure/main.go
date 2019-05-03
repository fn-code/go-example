package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func setAddress(addr string) func() {
	closure := func() {
		fmt.Printf("Your address: %s\n", addr)
	}
	return closure
}

func main() {
	setName := func(date string) error {
		i, err := strconv.ParseInt("1518328047", 10, 64)
		if err != nil {
			return err
		}
		tm := time.Unix(i, 0)
		fmt.Println(tm)
		return nil
	}

	date := time.Now().String()
	err := setName(date)
	if err != nil {
		log.Println(err)
	}

	addr := setAddress("indonesia")
	addr()

	n := channelFib(2)
	fmt.Println(n)

}

func channelFib(n int) int {
	n += 2
	ch := make(chan int)

	channel := func(ch chan<- int, counter int) {
		n1, n2 := 0, 1
		for i := 0; i < counter; i++ {
			ch <- n1
			n1, n2 = n2, n1+n2
		}
		close(ch)
	}

	go channel(ch, n)
	i := 0
	var result int
	for num := range ch {
		result = num
		i++
	}
	return result
}
