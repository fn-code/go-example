package storage

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type ListPerson struct {
	persons []*Person
}

// Add is for adding person
func (p *ListPerson) Add(ps *Person) {
	p.persons = append(p.persons, ps)
}

func (p *ListPerson) Show() {
	for _, v := range p.persons {
		fmt.Println(v.Name, v.Age)
	}
}

// PersonService is person interface
type PersonService interface {
	Add(*Person)
}
