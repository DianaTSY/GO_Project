package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main(){
	requestUrl := "https://www.baidu.com/"

	postValue := url.Values{
		"username" : {"hangmeimei"},
		"address": {"anhui"},
		"subject" : {"world"},
		"form" : {"beij"},
	}

	body := bytes.NewBufferString(postValue.Encode())
	request , err := http.Post(requestUrl,"text/html",body)
	if err != nil{
		fmt.Println(err)
	}

	defer request.Body.Close()
	fmt.Println(request.StatusCode)
	if request.StatusCode == 200 {
		rb, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println(rb)
		}
		fmt.Println(string(rb))
	}
}
