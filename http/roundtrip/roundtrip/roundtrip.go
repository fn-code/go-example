package roundtrip

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"sync"
)

type CustomTransport struct {
	data             map[string][]byte
	mutex            sync.RWMutex
	defaultTransport http.RoundTripper
}

func New() *CustomTransport {
	return &CustomTransport{
		data: make(map[string][]byte),
		defaultTransport: http.DefaultTransport,
	}
}

func (ct *CustomTransport) Set(r *http.Request, buf []byte) {
	ct.mutex.Lock()
	defer ct.mutex.Unlock()
	ct.data[r.URL.String()] = buf
}

func (ct *CustomTransport) Get(r *http.Request) (*bytes.Buffer, error) {
	ct.mutex.RLock()
	defer ct.mutex.RUnlock()
	val, ok := ct.data[r.URL.String()]
	if !ok {
		return nil, errors.New("key is not registered")
	}
	buf := bytes.NewBuffer(val)
	return buf, nil
}


func (ct *CustomTransport) Clear() {
	ct.mutex.Lock()
	defer ct.mutex.Unlock()
	ct.data = make(map[string][]byte)
}


func (ct *CustomTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	buffer, err := ct.Get(r)
	if err == nil {
		return http.ReadResponse(bufio.NewReader(buffer), r)
	}

	resp, err := ct.defaultTransport.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	buf, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}
	fmt.Println("Fetching the data")

	ct.Set(r, buf)
	return resp, nil
}
