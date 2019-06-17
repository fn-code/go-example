package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	dec "github.com/fn-code/Go-Example/functional-programming/8_decorator/src/decorator"

	met "github.com/fn-code/Go-Example/functional-programming/8_decorator/src/metrics"
)

const (
	host     = "127.0.0.1"
	protocol = "http://"
)

var (
	serverURL string
	proxyURL  string
)

func init() {
	serverPort := 3000
	proxyPort := 8080
	flag.IntVar(&serverPort, "serverPort", serverPort, "Server Port")
	flag.IntVar(&proxyPort, "proxyPort", proxyPort, "Server Port")
	flag.Parse()
	serverURL = fmt.Sprintf("%s:%d", host, serverPort)
	proxyURL = fmt.Sprintf("%s:%d", host, proxyPort)
}

func main() {
	dec.InitLog("trace-log.txt", ioutil.Discard, os.Stdout, os.Stderr)

	dec.Info.Printf("Metrics server listening on %s", serverURL)

	go func() {
		log.Fatal(met.Serve(serverURL))
	}()

	time.Sleep(1 * time.Second)

	req, err := http.NewRequest(http.MethodGet, protocol+serverURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	dec.Info.Printf("Proxy listening on %s", proxyURL)
	proxyUrl, _ := url.Parse(proxyURL)
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	tr.TLSNextProto = make(map[string]func(string, *tls.Conn) http.RoundTripper)
	proxyTimeoutClient := &http.Client{Transport: tr, Timeout: 1 * time.Second}
	client := dec.Decorate(proxyTimeoutClient, dec.Authorization("mysecretpassword"),
		dec.LoadBalancing(dec.RoundRobin(0, "web01:3000", "web02:3000", "web03:3000")),
		dec.Logging(log.New(dec.InfoHandler, "client: ", log.Ltime)),
		dec.FaultTolerance(2, time.Second),
	)

	job := &dec.Job{
		Client:       client,
		Request:      req,
		NumRequests:  10,
		IntervalSecs: 10,
	}

	start := time.Now()
	job.Run()

	dec.Info.Printf("\n>> It took %s", time.Since(start))

	dec.Info.Printf("metrics")
	err = met.DisplayResults(serverURL)
	if err != nil {
		log.Fatalln(err)
	}

	dec.Info.Printf("CTRL+C to exit")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}
