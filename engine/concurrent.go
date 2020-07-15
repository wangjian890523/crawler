package engine

import "github.com/lunny/log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
	//WorkReady(chan Request)
	//Run()
}


func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		CreateWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount  :=0

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d:%v", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)

		}

	}

}

func CreateWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {

			request := <-in
			result, err := worker(request)
			if err != nil {
				continue

			}
			out <- result
		}

	}()

}
