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
			log.Printf("Got item %v",item)
		}
	}
}

func Worker(r Request) (ParseResult, error){
	body, err := fetcher.Fectch(r.Url)
	if err == nil{
		log.Printf("Fetcher: error"+"Fetch Url: %s, %v", err, r.Url,)
		parseResult := r.ParseFunc(body)
		for _, item := range parseResult.Items{
			log.Printf("Got item %v",item)
		}
		return parseResult,nil
		}else{
			return ParseResult{},err
	}
}
