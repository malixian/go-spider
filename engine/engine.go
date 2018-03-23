package engine

import (
	"testGo/crawl/fetcher"
	"log"
	"fmt"
)

func Run(seeds ...Request){
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}

	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]
		parseResult,err := Worker(r)
		if err!= nil{
			fmt.Println("")
			}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items{
			log.Printf("Got item %v,\n",item)
		}
	}
}

func Worker(r Request) (ParseResult, error){
	body, err := fetcher.Fectch(r.Url)
	if err == nil{
		parseResult := r.ParseFunc(body)
		return parseResult,nil
		}else{
			log.Printf("Fetcher: error"+"Fetch Url: %s, %v", err, r.Url,)
			return ParseResult{},err
	}
}
