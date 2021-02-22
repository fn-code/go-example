package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"sync/atomic"
	"syscall"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("index").ParseFiles("./index.html"))
	tpl.ExecuteTemplate(w, "index.html", nil)
}

var count int64

func ws(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	n := atomic.AddInt64(&count, 1)
	if n%100 == 0 {
		log.Printf("Total number of connections: %v", n)
	}
	defer func() {
		n := atomic.AddInt64(&count, -1)
		if n%100 == 0 {
			log.Printf("Total number of connections: %v", n)
		}
		conn.Close()
	}()

	// Read messages from socket
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		log.Printf("msg: %s", string(msg))
	}
}


func main() {
	// Increase resources limitations
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		log.Println(err)
	}
	fmt.Println(rLimit.Max)
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		log.Println(err)
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/", index)
	handler.HandleFunc("/ws", ws)
	err := http.ListenAndServe(":9090", handler)
	if err != nil {
		log.Fatal(err)
	}
}
