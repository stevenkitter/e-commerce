package server

type QueuedScheduler struct {
	requestChan chan Request
	workerChan  chan chan Request
}

func (s *QueuedScheduler) Submit(r Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(c chan Request) {
	s.workerChan <- c
}

func (s *QueuedScheduler) Run() {
	//init
	s.requestChan = make(chan Request)
	s.workerChan = make(chan chan Request)
	//
	go func() {
		var (
			requestQ []Request
			workerQ  []chan Request
		)
		for {
			var (
				activeR Request
				activeW chan Request
			)
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeR = requestQ[0]
				activeW = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeW <- activeR:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
