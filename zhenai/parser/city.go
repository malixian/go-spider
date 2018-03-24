package parser

import (
	"regexp"
	"go-spider/engine"
)

const userRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]*)</a></th>`


func ParseCity(contents []byte) engine.ParseResult{
	request := engine.ParseResult{}
	re := regexp.MustCompile(userRe)
	match := re.FindAllSubmatch(contents, -1)
	// 注意for循环中闭包的问题
	for _,m := range match{
		name := string(m[2])
		request.Requests = append(request.Requests,
			engine.Request{string(m[1]),func(content []byte) engine.ParseResult{
				return ParseProfile(content, name)
			},
		})
		request.Items = append(request.Items, name)
	}
	return request
}
