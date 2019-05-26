package main

import (
	"fmt"
	"strings"
)

func main() {
	f1 := func(s string) string {
		return fmt.Sprintf("your hobby is %v", s)
	}
	g1 := func(s string) string {
		return strings.Title(s)
	}
	cmps := Compose(f1, g1)
	fmt.Println(cmps("programming"))

	yCombo := combi(fibFuncFunc)
	fmt.Println("r(r)(x) in anonymous function in yCombinator is a Lambda Expression:")
	fmt.Println("> yCombo(5):", yCombo(5))
}

type strFunc func(string) string

func Compose(f, g strFunc) strFunc {
	return func(s string) string {
		return g(f(s))
	}
}

type FC func(int) int
type FCFC func(FC) FC
type RecursiveFC func(RecursiveFC) FC

func fibFuncFunc(f FC) FC {
	return func(x int) int {
		if x == 0 {
			return 0
		} else if x <= 2 {
			return 1
		} else {
			return f(x-2) + f(x-1)
		}
	}
}

func combi(f FCFC) FC {
	r := func(a RecursiveFC) FC {
		return f(func(x int) int {
			return a(a)(x)
		})
	}
	return r(r)
}
