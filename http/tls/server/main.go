package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Division string `json:"division"`
	Age      int    `json:"age"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	data := []Employee{
		{ID: "001", Name: "Ludin Nento", Division: "IT", Age: 24},
		{ID: "002", Name: "Nento Nento", Division: "IT", Age: 21},
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	caCert, err := ioutil.ReadFile("./config/cert.pem")
	if err != nil {
		log.Println(err)
	}
	cpool := x509.NewCertPool()
	cpool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs: cpool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	srv := http.Server{
		Addr: ":443",
		TLSConfig: tlsConfig,
		Handler: mux,
	}

	fmt.Println("Server is running.")
	err = srv.ListenAndServeTLS("./config/cert.pem", "./config/key.pem")
	if err != nil {
		log.Println(err)
	}
}
