package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

var urls = []string{
	"http://rubyconf.org/",
	"http://golang.org/",
	"http://matt.aimonetti.net/",
	"http://fun-code.tech/",
}

const CCTVLimitReq = 5

type CCTV struct {
	ID   string
	Name string
	URL  string
}

// ReadCCTV contain fetch status
type ReadCCTV struct {
	ID       string
	Name     string
	URL      string
	Status   bool
	Err      error
	NumError int64
}

func main() {
	ctv := make([]*CCTV, 0)
	for i, v := range urls {
		cctv := &CCTV{
			ID:   uuid.NewV4().String(),
			Name: fmt.Sprintf("CCtV %d", i+1),
			URL:  v,
		}
		ctv = append(ctv, cctv)
	}
	prosesCCTV(ctv...)
}

func prosesCCTV(cctv ...*CCTV) {
	ch := make(chan *ReadCCTV)
	done := make(chan *ReadCCTV)
	errCCTV := make(map[string]*ReadCCTV)
	for _, v := range cctv {
		go checkCCTVConnection(v, ch)
	}
	go readCCTVConnectionStatus(ch, done)

	for v := range done {
		if v.Status {
			ctv := &CCTV{
				ID:   v.ID,
				Name: v.Name,
				URL:  v.URL,
			}
			tic := time.NewTimer(5 * time.Second)
			go func(ctv *CCTV) {
				<-tic.C
				tic.Stop()
				checkCCTVConnection(ctv, ch)
			}(ctv)

		} else {
			errCtv, ok := errCCTV[v.ID]
			ctv := &CCTV{
				ID:   v.ID,
				Name: v.Name,
				URL:  v.URL,
			}
			switch ok {
			case false:
				v.NumError++
				errCCTV[v.ID] = v
				tic := time.NewTimer(5 * time.Second)
				go func(ctv *CCTV) {
					<-tic.C
					tic.Stop()
					fmt.Println("bum ", ctv.Name)
					checkCCTVConnection(ctv, ch)
				}(ctv)
			case true:
				errCtv.NumError++
				if errCtv.NumError == CCTVLimitReq {
					// send notifications
					errCtv.NumError = 0
					fmt.Println("------------------------------Sending Notifiction to ", errCtv.Name, " ---------------------------------")
					tic := time.NewTimer(10 * time.Second)
					go func(ctv *CCTV) {
						<-tic.C
						checkCCTVConnection(ctv, ch)
						tic.Stop()
					}(ctv)
					break
				}
				tic := time.NewTimer(5 * time.Second)
				go func(ctv *CCTV) {
					<-tic.C
					tic.Stop()
					fmt.Println("bum 2", ctv.Name)
					checkCCTVConnection(ctv, ch)
				}(ctv)

			}
		}
	}
}

func checkCCTVConnection(cctv *CCTV, ch chan<- *ReadCCTV) {
	client := http.Client{}
	fmt.Printf("Fetching %s \n", cctv.URL)
	resp, err := client.Get(cctv.URL)
	if err != nil {
		ch <- &ReadCCTV{cctv.ID, cctv.Name, cctv.URL, false, err, 0}
		return
	}
	ch <- &ReadCCTV{cctv.ID, cctv.Name, cctv.URL, true, nil, 0}
	resp.Body.Close()
}

func readCCTVConnectionStatus(ch <-chan *ReadCCTV, done chan<- *ReadCCTV) {
	for {
		select {
		case res := <-ch:
			if !res.Status {
				fmt.Printf("Failed retrive data from %v\n", res.Name)
				done <- res
				break
			}
			fmt.Printf("Succes retrive data from : %v\n", res.Name)
			done <- res
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}

// func sendTelegramMessage() bool, error{

// }
