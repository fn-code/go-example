package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Proxys struct {
	Client *http.Client
	BaseURL string
}

// ServeHTTP proxy using http Handler interface
func (p *Proxys) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := p.ProcessRequest(r); err != nil {
		log.Printf("error occurred during process request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := p.Client.Do(r)
	if err != nil {
		log.Printf("error occurred during client operation: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	p.Copy(w, resp)
}


// ProcessRequest update and modify the request using proxy
func(p *Proxys) ProcessRequest(r *http.Request) error {
	rawURI := fmt.Sprintf("%s%s", p.BaseURL, r.URL.String())

	pURL, err := url.Parse(rawURI)
	if err != nil {
		return err
	}
	//fmt.Println(pURL.Path+ " - "+pURL.Host+ " - "+ pURL.RequestURI())
	r.URL = pURL
	r.Host = pURL.Host
	r.URL.Scheme = pURL.Scheme
	r.Header.Add("X-Forwarded-Host", pURL.Host)
	r.Header.Add("X-Origin-Host", r.Host)
	r.RequestURI = ""

	if _, ok := r.Header["User-Agent"]; !ok {
		r.Header.Set("User-Agent", "")
	}

	return nil
}

// Copy client http.Response to http.ResponseWriter
func(p *Proxys) Copy(w http.ResponseWriter, r *http.Response) {

	for k, v := range r.Header {
		for _, val := range v {
			w.Header().Add(k, val)
		}
	}
	w.WriteHeader(r.StatusCode)
	io.Copy(w, r.Body)
}
