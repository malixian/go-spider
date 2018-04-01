package main

import (
	"go-spider/zhenai/parser"
	"go-spider/engine"
	"go-spider/scheduler"
	"go-spider/persist"
	"gopkg.in/olivere/elastic.v5"
	"context"
)

func main(){
	url := "http://www.zhenai.com/zhenghun"
	request := engine.Request{Url:url, ParseFunc:parser.ParseCityList}
	// go中使用多态的时候，如果是实现的指针接受者，要用&
	client, err := elastic.NewSimpleClient()
	if err != nil{
		panic(err)
	}
	e := engine.ConcurrentEngine{Scheduler:&scheduler.SimpleScheduler{},WorkerCount:100,ItemSaver:persist.ItemSaver(client, context.Background())}
	//e := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.QueuedScheduler{},
	//	WorkerCount:100,
	//	ItemSaver:persist.ItemSaver(),
	//}
	e.Run(request)
}


