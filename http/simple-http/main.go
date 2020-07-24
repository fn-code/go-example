package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type HelloHandler struct {
	Port string
}

func (h *HelloHandler) Index(w http.ResponseWriter, r *http.Request) {
	log.Println("new request ", r.URL.String())
	w.Write([]byte(fmt.Sprintf("This is index page on port %s", h.Port)))
}

func (h *HelloHandler) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("new request ", r.URL.String())
	qQuery := r.URL.Query().Get("q")
	pQuery := r.URL.Query().Get("p")
	w.Write([]byte(fmt.Sprintf("This is home page on port %s raw query %s query %s page %s", h.Port, r.URL.RawQuery, qQuery, pQuery)))
}

func (h *HelloHandler) About(w http.ResponseWriter, r *http.Request) {
	log.Println("new request ", r.URL.String())
	w.Write([]byte(fmt.Sprintf("This is about page on port %s", h.Port)))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Invalid Arguments you need to pass server port")
		os.Exit(1)
	}

	port := os.Args[1]
	hh := &HelloHandler{port}

	mx := http.NewServeMux()
	mx.HandleFunc("/", hh.Index)
	mx.HandleFunc("/home", hh.Home)
	mx.HandleFunc("/about", hh.About)

	srv := http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", port),
		Handler: mx,
	}

	fmt.Printf("Server running on port :%s\n", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
