package main

import (
	"fmt"
	"github.com/fn-code/Go-Example/interface/1_basic/task/job"
)


type CustomTask struct{}

func (ct *CustomTask) Add(s string, i int) {
	fmt.Printf("Task %s ID %d\n", s, i)
}

func AddTask(s string, i int) {
	fmt.Printf("Task %s ID %d\n", s, i)
}


func main() {
	ct := &CustomTask{}
	ts := job.NewJob()
	ts.Handle("path1", ct)
	ts.HandleFunc("path2", AddTask)
	job.Listen(ts, "path3")
}




