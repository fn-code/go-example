package main

import (
	"fmt"
	"github.com/fn-code/Go-Example/http/proxy/proxy"
	"net/http"
)

func main() {
	p := &proxy.Proxys{
		Client: http.DefaultClient,
		BaseURL: "https://www.golang.org",
	}
	http.Handle("/", p)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
