package scheduler

import "go-spider/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request){
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) WorkChan() chan engine.Request{
	return s.workerChan
}

func(s *SimpleScheduler) Run(){
	s.workerChan = make(chan engine.Request)
}

func(s *SimpleScheduler) WorkerReady(chan engine.Request){

}
