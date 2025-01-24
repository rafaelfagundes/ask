package spinner

import (
	"fmt"
	"time"
)

type Spinner struct {
	done    chan bool
	stopped bool
}

func New() *Spinner {
	return &Spinner{
		done: make(chan bool),
	}
}

func (s *Spinner) Start() {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	go func() {
		for i := 0; !s.stopped; i = (i + 1) % len(frames) {
			fmt.Printf("\r%s Thinking...", frames[i])
			select {
			case <-s.done:
				return
			case <-time.After(100 * time.Millisecond):
				continue
			}
		}
	}()
}

func (s *Spinner) Stop() {
	s.stopped = true
	s.done <- true
	fmt.Print("\r\033[K") // Clear the line
}
