package main

import (
	"go-spider/zhenai/parser"
	"go-spider/engine"
	"go-spider/scheduler"
)

func main(){
	url := "http://www.zhenai.com/zhenghun"
	request := engine.Request{Url:url, ParseFunc:parser.ParseCityList}
	//engine.Run(request)
	// go中使用多态的时候，如果是实现的指针接受者，要用&
	//e := engine.ConcurrentEngine{Scheduler:&scheduler.SimpleScheduler{},WorkerCount:10}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:100,
	}
	e.Run(request)
}


