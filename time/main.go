package main

import (
	"fmt"
	"time"
)

func main() {
	// input := "25/05/1980"
	input := "01/05/1980"
	// layout := "01/02/2006"
	layout := "02/01/2006"
	t, _ := time.Parse(layout, input)
	fmt.Println(t.Format("2006-01-02"))
}
