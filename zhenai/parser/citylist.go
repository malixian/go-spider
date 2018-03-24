package parser

import (

	"regexp"
	"go-spider/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"
									[^>]*>(.*?)</a>`

func ParseCityList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matchs{
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{Url:string(m[1]), ParseFunc:ParseCity,
			})
	}
	return result
}