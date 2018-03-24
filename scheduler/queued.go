package scheduler

import (
	"go-spider/engine"
)

type QueuedScheduler struct {
	requestChannel chan engine.Request
	workerChan chan chan engine.Request //每一个WORKER对应一个CHANNEL
}

func (s *QueuedScheduler) Submit( r engine.Request){
	s.requestChannel <- r
}

func (s *QueuedScheduler) WorkerReady( w chan engine.Request){
	s.workerChan <- w
}


func (s *QueuedScheduler) Run(){
	go func() {
		s.workerChan = make(chan chan engine.Request)
		s.requestChannel = make(chan engine.Request)
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) >0 && len(workerQ) >0{
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChannel:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()

}