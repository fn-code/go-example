package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Division string `json:"division"`
	Age      int    `json:"age"`
}

func main() {

	req, err := http.NewRequest("GET", "https://localhost", nil)
	if err != nil {
		log.Println(err)
	}

	cert, err := tls.LoadX509KeyPair("./config/cert.pem", "./config/key.pem")
	if err != nil {
		log.Println(err)
	}

	caCert, err := ioutil.ReadFile("./config/cert.pem")
	if err != nil {
		log.Println(err)
	}
	cpool := x509.NewCertPool()
	cpool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: cpool,
				Certificates: []tls.Certificate{cert},
			},
		},
		Timeout:   time.Duration(120 * time.Second),
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("invalid request")
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	data := make([]Employee, 0)

	err = json.Unmarshal(buf, &data)
	if err != nil {
		log.Println(err)
	}
	for _, v := range data {
		fmt.Println(v.ID, " - ", v.Name, " - ", v.Division, " - ", v.Age)
	}
}
