package engine

type Request struct {
	Url string
	ParseFunc func([]byte) ParseResult  //针对每一个url，对应要抓取的方法
}

type ParseResult struct {
	Requests []Request
	Items []Item
}
type Item struct {
	Url string
	Id  string
	Type string // 每一个网站对应es里面的type
	PayLoad interface{}
}

func NilPaser([] byte) ParseResult{
	return ParseResult{}
}
