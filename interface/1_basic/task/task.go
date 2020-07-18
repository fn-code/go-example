package task

import (
	"sync"
)

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

func (ts *Tasks) Handle(s string, t Task) {
	ts.mx.Lock()
	defer ts.mx.Unlock()
	if ts.m == nil {
		ts.m = make(map[string]sTask)
	}
	ts.m[s] = sTask{t, s}
}

func (ts *Tasks) task(s string) Task {
	if _, ok := ts.m[s]; !ok {
		return nil
	}
	return ts.m[s].ts
}

func (ts *Tasks) Add(s string, i int) {
	h := ts.task(s)
	h.Add(s, i)
}

type HandlerTask func(string, int)

func (h HandlerTask) Add(s string, i int) {
	h(s, i)
}

func (ts *Tasks) HandleFunc(s string, h HandlerTask) {
	ts.Handle(s, h)
}

func Listen(t Task, dest string) {
	ns := taskHandler{t}
	ns.Add(dest, 1)
}

type taskHandler struct {
	t Task
}

func (th taskHandler) Add(s string, i int) {
	th.t.Add(s, i)
}
