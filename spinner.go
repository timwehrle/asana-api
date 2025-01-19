package asana

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Spinner struct {
	message string
	frames  []string
	current int
	writer  io.Writer
	stop    chan struct{}
	wg      sync.WaitGroup
	active  bool
	mu      sync.Mutex
}

func NewSpinner(message string) *Spinner {
	return &Spinner{
		message: message,
		frames:  []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		writer:  os.Stderr,
		stop:    make(chan struct{}),
	}
}

func (s *Spinner) Start() {
	s.mu.Lock()
	if s.active {
		s.mu.Unlock()
		return
	}
	s.active = true
	s.mu.Unlock()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			select {
			case <-s.stop:
				return
			default:
				s.mu.Lock()
				fmt.Fprintf(s.writer, "\r%s %s", s.frames[s.current], s.message)
				s.current = (s.current + 1) % len(s.frames)
				s.mu.Unlock()
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

func (s *Spinner) Stop() {
	s.mu.Lock()
	if !s.active {
		s.mu.Unlock()
		return
	}
	s.active = false
	s.mu.Unlock()

	close(s.stop)
	s.wg.Wait()
	fmt.Fprintln(s.writer, "\r")
}
