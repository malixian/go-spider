package engine

import (
	"testGo/crawl/fetcher"
	"log"
)

func Run(seeds ...Request){
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}

	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fectch(r.Url)
		if err != nil{
			log.Printf("Fetcher: error"+"Fetch Url: %s, %v", err, r.Url,)
		}
		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items{
			log.Printf("Got item %v",item)
		}

	}
}
