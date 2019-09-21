package main

import (
	"errors"
	"fmt"
	"log"
)

var errTets = errors.New("invalid data type")

func main() {
	err := testError()
	if err != nil {
		if errors.Is(err, errTets) {
			log.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}

func testError() error {
	return fmt.Errorf("faliled to test error function : %w", errTets)
}
