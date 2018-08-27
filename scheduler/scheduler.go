package scheduler

import (
	"crawer/engine"
)

type ConcurrentScheduler struct {
	WorkerChan chan engine.Request
}

func (s *ConcurrentScheduler)ConfigureWorkChan(c chan engine.Request)  {
	s.WorkerChan = c
}

func (s *ConcurrentScheduler) Submit(r engine.Request)  {
	go func() {s.WorkerChan <- r}()
}