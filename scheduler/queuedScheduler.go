package scheduler

import "crawer/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (q *QueuedScheduler) Submit (c engine.Request)  {
	q.requestChan <- c
}
func (q *QueuedScheduler)ConfigureWorkChan (c chan engine.Request)  {

}
func (q *QueuedScheduler)WorkReady(w chan engine.Request)  {
	q.workerChan <- w
}

func (q *QueuedScheduler)Run(){
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for  {
			var activeRequest engine.Request
			var activeWoker chan engine.Request
			if len(requestQ)>0 && len(workerQ)>0 {
				activeRequest = requestQ[0]
				activeWoker = workerQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ,r)
			case w := <-q.workerChan:
				workerQ = append(workerQ,w)
			case activeWoker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
