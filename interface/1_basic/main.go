package main

import (
	"fmt"
	"sync"
)

type CustomTask struct{}

func (ct *CustomTask) Add(s string, i int) {
	fmt.Printf("Pattern %s Index %d\n", s, i)
}

func AddTask(s string, i int) {
	fmt.Printf("Pattern %s Index %d\n", s, i)
}

func main() {

	ts := &Tasks{}
	ct := &CustomTask{}
	ts.Handle("path1", ct)
	ts.HandleFunc("path2", AddTask)

	Listen(ts, "path2")

}

type Tasks struct {
	m  map[string]sTask
	mx sync.RWMutex
}

type sTask struct {
	ts      Task
	pattern string
}

type Task interface {
	Add(string, int)
}

type HandlerTask func(string, int)

func (ts *Tasks) Handle(s string, t Task) {
	ts.mx.Lock()
	defer ts.mx.Unlock()
	if ts.m == nil {
		ts.m = make(map[string]sTask)
	}
	ts.m[s] = sTask{t, s}
}

func (h HandlerTask) Add(s string, i int) {
	h(s, i)
}

func (ts *Tasks) HandleFunc(s string, h func(string, int)) {
	ts.Handle(s, HandlerTask(h))
}

func (ts *Tasks) Task(s string) Task {
	if _, ok := ts.m[s]; !ok {
		return nil
	}
	return ts.m[s].ts
}

func (ts *Tasks) Add(s string, i int) {
	h := ts.Task(s)
	h.Add(s, i)
}

/////////////////////////////////////////////////////
func Listen(t Task, dest string) {
	ns := serverHandler{t}
	ns.Add(dest, 1)
}

type serverHandler struct {
	h Task
}

func (sh serverHandler) Add(s string, i int) {
	handler := sh.h
	handler.Add(s, i)
}
