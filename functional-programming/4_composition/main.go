package main

import (
	"fmt"
	"strings"
)

type strFunc func(string) string

func Compose(f, g strFunc) strFunc {
	return func(s string) string {
		return g(f(s))
	}
}

func main() {
	f1 := func(s string) string {
		return fmt.Sprintf("your hobby is %v", s)
	}
	g1 := func(s string) string {
		return strings.Title(s)
	}
	cmps := Compose(f1, g1)
	fmt.Println(cmps("programming"))
}
