package decorator

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type response struct {
	duration time.Duration
	err      error
}

type Job struct {
	Client       Client
	NumRequests  int
	Request      *http.Request
	IntervalSecs int
	responseChan chan *response
}

func (b *Job) displayProgress(stopChan <-chan struct{}) {
	var prevResponseCount int
	for {
		select {
		case <-time.Tick(time.Millisecond * 500):
			responseCount := len(b.responseChan)
			if prevResponseCount < responseCount {
				prevResponseCount = responseCount
				Debug.Printf("> %d request done.", responseCount)
			}
		case <-stopChan:
			return
		}
	}
}

func (b *Job) Run() {
	b.responseChan = make(chan *response, b.NumRequests)
	stopChan := make(chan struct{})
	go b.displayProgress(stopChan)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-sigChan
		stopChan <- struct{}{}
		close(b.responseChan)
		os.Exit(1)
	}()

	var wg sync.WaitGroup
	intervalSecs := time.Duration(b.IntervalSecs)
	reqPerformed := 0
	for range time.Tick(intervalSecs * time.Second) {
		wg.Add(1)
		go func() {
			client := b.Client
			b.makeRequest(client)
			wg.Done()
		}()
		reqPerformed++
		if reqPerformed >= b.NumRequests {
			break
		}
	}
	wg.Wait()
	stopChan <- struct{}{}
	Debug.Printf("All requests done.")
	close(b.responseChan)
}

func (b *Job) makeRequest(c Client) {
	Debug.Printf("makeRequest: ")
	start := time.Now()
	resp, err := c.Do(b.Request)
	if err == nil {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
	t := time.Now()
	finish := t.Sub(start)
	b.responseChan <- &response{
		duration: finish,
		err:      err,
	}
}
