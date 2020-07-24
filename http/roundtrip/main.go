package main

import (
	"fmt"
	"github.com/fn-code/Go-Example/http/roundtrip/roundtrip"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	customTransport := roundtrip.New()

	client := &http.Client{
		Transport: customTransport,
		Timeout:   time.Second * 5,
	}

	// clear a cache every 5 second
	clearRespChache := time.NewTicker(time.Second * 5)

	// Make request every 1 second
	reqTicker := time.NewTicker(time.Second * 1)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	if err != nil {
		log.Println("error make new request")
		os.Exit(0)
		return
	}

	for {
		select {
		case <-clearRespChache.C:
			customTransport.Clear()
		case <-reqTicker.C:
			fmt.Println("-> Client request")
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("error make request %v", err)
				continue
			}

			buf, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("error read request %v", err)
				continue
			}

			fmt.Printf(" \n \n > %s \n \n", string(buf))
			resp.Body.Close()
		}
	}
}
