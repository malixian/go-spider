package main

import (
	"go-spider/fetcher"
	"go-spider/zhenai/parser"
)

func main(){
	url := "http://www.zhenai.com/zhenghun/aba"
	contents,_ := fetcher.Fectch(url)
	parser.ParseCity(contents)
}

