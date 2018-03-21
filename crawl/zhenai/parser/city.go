package parser

import (
	"testGo/crawl/engine"
	"regexp"
)

const userRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]*)</a></th>`


func ParseCity(contents []byte) engine.ParseResult{
	var request engine.ParseResult
	re := regexp.MustCompile(userRe)
	match := re.FindAllSubmatch(contents, -1)
	for _,m := range match{
		request.Requests = append(request.Requests,
			engine.Request{string(m[1]),ParseProfile,
		})
		request.Items = append(request.Items, string(m[2]))
	}
	return request
}
