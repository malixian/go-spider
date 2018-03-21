package main

import (
	"testGo/crawl/engine"
	"testGo/crawl/zhenai/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	request := engine.Request{Url:url, ParseFunc:parser.ParseCityList}
	engine.Run(request)
}

