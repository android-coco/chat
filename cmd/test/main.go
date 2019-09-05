package main

import (
	"fmt"
	"net/url"
)

func main() {
	urltest := "http://www.baidu.com/s?wd=自由度"
	fmt.Println(urltest)
	encodeurl:= url.QueryEscape(urltest)
	fmt.Println(encodeurl)
	decodeurl,err := url.QueryUnescape(urltest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decodeurl)
}
