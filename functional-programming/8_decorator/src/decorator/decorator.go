package decorator

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type Client interface {
	Do(*http.Request) (*http.Response, error)
}

type ClientFunc func(*http.Request) (*http.Response, error)

func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

var rateLimitDuration time.Duration

func (f ClientFunc) SetRateLimit(d time.Duration) error {
	rateLimitDuration = d
	return nil
}

func (f ClientFunc) GetRateLimit() (time.Duration, error) {
	return rateLimitDuration, nil
}

type Decorator func(Client) Client

func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, dec := range ds {
		decorated = dec(decorated)
	}
	return decorated
}

func Authorization(tkn string) Decorator {
	return Header("Authorization", tkn)
}

func Header(name, value string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Add(name, value)
			return c.Do(r)
		})
	}
}

func Logging(l *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			l.Printf("%s %s", r.Method, r.URL)
			return c.Do(r)
		})
	}
}

type Director func(*http.Request)

func LoadBalancing(dir Director) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			dir(r)
			return c.Do(r)
		})
	}
}

func RoundRobin(robin int64, backends ...string) Director {
	return func(r *http.Request) {
		if len(backends) > 0 {
			r.URL.Host = backends[atomic.AddInt64(&robin, 1)%int64(len(backends))]
		}
	}
}

// FaultTolerance returns a Decorator that extends a Client with fault tolerance configured
// with the given attempts and backoff duration
func FaultTolerance(at int, bo time.Duration) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (res *http.Response, err error) {
			for i := 0; i <= at; i++ {
				if res, err = c.Do(r); err != nil {
					Info.Println("SUCCESS!")
					break
				}
				Debug.Println("backing off...")
				time.Sleep(bo * time.Duration(i))
			}
			if err != nil {
				Info.Println("FAILURE!")
			}
			return res, err
		})
	}
}
