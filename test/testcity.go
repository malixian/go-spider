package main

import (
	"testGo/crawl/engine"
	"testGo/crawl/zhenai/parser"
	"testGo/crawl/scheduler"
)

func main(){
	url := "http://www.zhenai.com/zhenghun"
	request := engine.Request{Url:url, ParseFunc:parser.ParseCityList}
	//engine.Run(request)
	// go中使用多态的时候，如果是实现的指针接受者，要用&
	e := engine.ConcurrentEngine{Scheduler:&scheduler.SimpleScheduler{},WorkerCount:10}
	e.Run(request)
}


