package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main(){
	client := &http.Client{}

	req,err := http.NewRequest("GET", "http://www.baidu.com",nil)
	if err != nil{
		fmt.Println(err)
	}

	cookie := &http.Cookie{Name:"Diana", Value:strconv.Itoa(123)}
	req.AddCookie(cookie) //向request中添加Cookie

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	req.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	if response.StatusCode == 200 {
		r, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(r))
	}
}
