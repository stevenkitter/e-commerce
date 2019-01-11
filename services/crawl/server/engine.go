package server

import (
	"log"
)

type Engine struct {
	WorkerCount int
	Scheduler   Scheduler
}

type Scheduler interface {
	Submit(Request)
	WorkerReady(chan Request)
	Run()
}

func (e *Engine) Run(req Request) {
	out := make(chan Result)
	e.Scheduler.Run() //
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}
	e.Scheduler.Submit(req)
	for {
		result := <-out
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan Result, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			res := worker(request)
			out <- res
		}
	}()
}

func worker(req Request) Result {
	doc, err := req.Fetcher(req.URL)
	if err != nil {
		log.Printf("req.Fetcher(req.URL) err:%v", err)
		return Result{}
	}
	return req.ParseFunc(req.URL, doc)
}
