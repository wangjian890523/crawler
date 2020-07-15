package scheduler

import "github.com/wangjian890523/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkChan(
	c chan engine.Request) {
	s.workerChan = c

}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {s.workerChan <- r}()

}
