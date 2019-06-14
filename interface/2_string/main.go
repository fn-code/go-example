package main

import (
	"fmt"
)

type Duration int64

func New(d int64) Duration {
	return Duration(d)
}

func (d Duration) String() string {
	return fmt.Sprintf("%dms\n", d)
}

type Trait interface {
	Eat() string
}

type Food string

func NewFood(s string) Food {
	return Food(s)
}

func (f Food) Eat() string {
	return fmt.Sprintf("Food Name is : %s", f)
}

func Show(a ...interface{}) {
	for _, f := range a {
		s, ok := f.(Food)
		fmt.Println(ok, s.Eat())
	}
}
func main() {
	d := New(1000)
	fmt.Println(d)

	Show(NewFood("Bananas"))

}
