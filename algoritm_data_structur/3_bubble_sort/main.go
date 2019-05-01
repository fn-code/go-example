package main

import (
	"fmt"
)

// 2, 6, 3, 66, 1, 8, 9, 11, 4
func bubbeSort(items []int) []int {
	max := len(items)
	sorted := false

	for !sorted {
		swap := false
		for i := 0; i < max-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swap = true
			}
		}

		if !swap {
			sorted = true
		}
		max--
	}

	return items
}

func main() {
	items := []int{2, 6, 3, 66, 1, 8, 9, 11, 4}
	fmt.Printf("Before : %v\n", items)
	fmt.Printf("After : %v\n", bubbeSort(items))
}
