package engine

import (
	"log"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 程序的设计准则，自顶向下， 可以先定义Scheduler的接口，再去实现
type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemSaver chan Item
}

type Scheduler interface{
	Submit(Request)
	WorkChan() chan Request
	Run()
	ReadyNotify
}

type ReadyNotify interface{
	WorkerReady(chan Request)
}


func filterUrl(c redis.Conn,url string){
	_, err := c.Do("sadd","spider",url)
	if err != nil{
		log.Printf("redis add url err %v", err)
	}
}

func (e *ConcurrentEngine) Run(seeds ...Request){
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
		}
	defer c.Close()
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++{
		createWorker(e.Scheduler.WorkChan(),out, e.Scheduler)
	}
	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}
	for{
		result := <-out
		for _, request := range result.Requests{
			filterUrl(c, request.Url)
			e.Scheduler.Submit(request)
		}
		for _, item := range result.Items{
			go func() {
				e.ItemSaver <- item
			}()
		}
	}
}


func createWorker(in chan Request, out chan ParseResult, ready ReadyNotify){
	go func() {
		for{
			ready.WorkerReady(in)
			r := <-in
			parseResult,err := Worker(r)
			if err != nil{
				continue
			}
			out <- parseResult
		}
	}()
}
