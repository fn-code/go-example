package main

import (
	"fmt"
	"strings"
)

// Book is contains book data
type Book struct {
	title  string
	author string
}

type BookCollection []*Book

// LoadCollection is load all book collection
func LoadCollection() BookCollection {
	var bkcoll []*Book
	bks1 := &Book{"Go programming language", "Gopher 1"}
	bkcoll = append(bkcoll, bks1)
	bks2 := &Book{"Golang programming", "Gopher 2"}
	bkcoll = append(bkcoll, bks2)
	return bkcoll
}

func (bc BookCollection) Add(title, author string) BookCollection {
	bks := &Book{title: title, author: author}
	return append(bc, bks)
}

type FilterBook func(string) bool

func (bc BookCollection) Filter(fb FilterBook) BookCollection {
	bcoll := make([]*Book, 0)
	for _, v := range bc {
		if fb(v.title) {
			bcoll = append(bcoll, v)
		}
	}
	return bcoll
}

func ByTitle(s string) FilterBook {
	return func(bk string) bool {
		return strings.Contains(bk, s)
	}
}

func main() {

	bc := LoadCollection()
	bc = bc.Add("Microservice with go", "Gopher 3").Add("Golang", "Gopher 4")
	bc = bc.Filter(
		ByTitle("programming"),
	)
	for _, v := range bc {
		fmt.Println(v.title)
	}

	mp := map[string]string{
		"Honda":  "LX",
		"Lexus":  "LS",
		"Toyota": "EV",
		"Ford":   "XL",
		"GM":     "X",
	}["Honda"]
	fmt.Println(mp)

}
