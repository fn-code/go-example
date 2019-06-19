package main

import (
	"github.com/fn-code/Go-Example/functional-programming/9_DDD/storage"
)

func main() {
	lp := &storage.ListPerson{}

	msg := &storage.Message{}
	p1 := &storage.Person{
		Name: "Ludin Nento",
		Age:  22,
	}
	p2 := &storage.Person{
		Name: "Gopher",
		Age:  9,
	}

	msg.AddMessage(lp)
	msg.Notify(p1, p2)

	lp.Show()

	p3 := &storage.Person{
		Name: "Golang",
		Age:  2,
	}

	msg.Notify(p3)
	lp.Show()
}
