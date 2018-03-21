package engine

type Request struct {
	Url string
	ParseFunc func([]byte) ParseResult  //针对每一个url，对应要抓取的方法
}

type ParseResult struct {
	Requests []Request
	Items []interface{}
}

func NilPaser([] byte) ParseResult{
	return ParseResult{}
}
