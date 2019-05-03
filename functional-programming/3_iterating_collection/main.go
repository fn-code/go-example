package main

import (
	"fmt"
	"strings"
)

type Iteration interface {
	Next() bool
	Scan() string
}

type Collection struct {
	index int
	List  []string
}

type MaxWords int

const (
	ZERO MaxWords = 6 * iota
	SMALL
	MEDIUM
	LARGE
	XLARGE
	XXLARGE
)

func main() {
	c := New([]string{"Golang", "go", "Go-lang", "go Programming"})
	c.Map(strings.ToUpper)
	c.Filter(SMALL)

	c.Join([]string{"Php", "Python", "Rust", "C++", "C"})
	c.Map(strings.ToLower)

	for c.Next() {
		fmt.Println(c.Scan())
	}
	fmt.Println(c.Contains("golang"))
}

func (c *Collection) Scan() string {
	if c.index >= len(c.List) {
		return ""
	}
	return c.List[c.index]
}

func (c *Collection) Next() bool {
	c.index++
	if c.index >= len(c.List) {
		return false
	}
	return true
}

func New(s []string) *Collection {
	return &Collection{-1, s}
}

type stringFunc func(s string) string

func (c *Collection) Map(fn stringFunc) *Collection {
	orig := *c
	for i, s := range orig.List {
		c.List[i] = fn(s) // first-class function
	}
	return c
}

func (c *Collection) Filter(max MaxWords) *Collection {
	mapped := []string{}
	for _, s := range c.List {
		if len(s) >= int(max) {
			mapped = append(mapped, s)
		}
	}
	c.List = mapped
	return c
}

func (c *Collection) Join(s []string) *Collection {
	if len(s) == 0 {
		return c
	}
	c.List = append(c.List, s...)
	return c
}

func (c *Collection) Contains(s string) bool {
	for _, v := range c.List {
		if v == s {
			return true
		}
	}
	return false
}
