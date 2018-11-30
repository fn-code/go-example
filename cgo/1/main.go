package main

// #include "sum.h"
import "C"
import (
	"fmt"
)

func main() {
	C.sum(3, 4)
	d := C.sum2(2, 2)
	fmt.Println(d)
}
