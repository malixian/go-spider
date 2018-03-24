package engine

import (
	"fmt"
)

// 程序的设计准则，自顶向下， 可以先定义Scheduler的接口，再去实现
type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface{
	Submit(Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request){
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++{
		createWorker(out, e.Scheduler)
	}
	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}

	for{
		result := <-out
		for _, request := range result.Requests{
			e.Scheduler.Submit(request)
		}

		for _, item := range result.Items{
			fmt.Printf("Got item: %v", item)
			fmt.Println()
		}
	}
}


func createWorker(out chan ParseResult, s Scheduler){
	in := make(chan Request)
	go func() {
		for{
			s.WorkerReady(in)
			r := <-in
			parseResult,err := Worker(r)
			if err != nil{
				continue
			}
			out <- parseResult
		}
	}()
}
