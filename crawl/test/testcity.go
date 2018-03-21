package main

import (
	"testGo/crawl/engine"

	"testGo/crawl/zhenai/parser"
)

func main(){
	url := "http://www.zhenai.com/zhenghun/aba"
	request := engine.Request{Url:url, ParseFunc:parser.ParseCity}
	engine.Run(request)
}


