package job

import (
	"github.com/fn-code/Go-Example/interface/1_basic/task"
	"log"
	"sync"
)

type Job struct {
	m  map[string]jobPath
	mx sync.RWMutex
}

type jobPath struct {
	ts      task.Task
	pattern string
}

// NewJob create new job
func NewJob() *Job {
	return &Job{
		m: make(map[string]jobPath),
	}
}

// Handle is register and handle new job
// Handle take 2 parameter is job path and Task interface that implement Add(string, int) function
func (j *Job) Handle(s string, t task.Task) {
	j.mx.Lock()
	defer j.mx.Unlock()
	if j.m == nil {
		j.m = make(map[string]jobPath)
	}
	// register new jobPath
	j.m[s] = jobPath{t, s}
}

func (j *Job) task(s string) task.Task {
	if _, ok := j.m[s]; !ok {
		return nil
	}
	return j.m[s].ts
}

func (j *Job) Add(s string, i int) {
	h := j.task(s)
	if h == nil {
		log.Fatalf("%s job task is not registered", s)
	}
	h.Add(s, i)
}

type HandlerTask func(string, int)

func (h HandlerTask) Add(s string, i int) {
	h(s, i)
}

func (j *Job) HandleFunc(s string, h HandlerTask) {
	j.Handle(s, h)
}

// Listen is running a job with specific job path than are registered before
func Listen(t task.Task, dest string) {
	ns := taskHandler{t}
	ns.Add(dest, 1)
}

type taskHandler struct {
	t task.Task
}

func (th taskHandler) Add(s string, i int) {
	th.t.Add(s, i)
}
