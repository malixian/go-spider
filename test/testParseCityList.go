package main

import (
	"go-spider/zhenai/parser"
	"go-spider/fetcher"
)

func main(){
	url := "http://www.zhenai.com/zhenghun"
	contents,_ := fetcher.Fectch(url)
	parser.ParseCityList(contents)
}
