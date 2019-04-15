package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Ludin")
}

func main() {
	http.HandleFunc("/", index)
	log.Println("server running")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Println(err)
	}
}
