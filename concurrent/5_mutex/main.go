package main

import (
	"sync"
)

var mx sync.Mutex
var count int

func increment() {
	mx.Lock()
	defer mx.Unlock()
	count++
	// fmt.Printf("Incrementing: %d\n", count)
}

func decrement() {
	mx.Lock()
	defer mx.Unlock()
	count--
	// fmt.Printf("Decrement : %d\n", count)
}

func increment2() {
	count++
	// fmt.Printf("Incrementing: %d\n", count)
}

func decrement2() {
	count--
	// fmt.Printf("Decrement : %d\n", count)
}

func runningMutex() {
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}
	wg.Wait()
	// fmt.Println("program is complete")
}

func runningNotMutex() {
	for i := 0; i <= 5; i++ {
		increment2()
	}

	for i := 0; i <= 5; i++ {
		decrement2()
	}
	// fmt.Println("program is complete")
}

func main() {
	runningMutex()
}
