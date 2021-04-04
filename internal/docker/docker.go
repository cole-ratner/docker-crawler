package docker

import (
	"fmt"
	"os/exec"
	"sync"
)

// Worker defines the worker that will be scheduled to run enumeration
type Worker struct {
	Command string 		`json:"command"`
	Args 	[]string 	`json:"args"`
	Output 	chan string 
}


// NewWorker sets params for a  Worker instance
func NewWorker(cmd string, args []string, out chan string) *Worker {
	return &Worker{
		Command: cmd,
		Args: args,
		Output: out,
	}
}

// Run's the worker command and writes to a channel
func (w *Worker) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	out, err := exec.Command(w.Command, w.Args...).Output()
	if err != nil {
		fmt.Println(err)
	}
	w.Output <- string(out)
}

// Collect returns outputs from the worker's channel
func Collect(c chan string) {
	for {
		msg := <-c
		fmt.Printf("\n%s", msg)
	}
}
