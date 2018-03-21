package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding/unicode"
)


func Fectch(url string) ([]byte, error){
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil{
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	r := bufio.NewReader(resp.Body)
	e := determineEncoding(r)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil{
		return nil, err
	}
	return all, nil
}

func determineEncoding(r *bufio.Reader) encoding.Encoding{
	bytes, err := r.Peek(1024)
	if err != nil{
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
