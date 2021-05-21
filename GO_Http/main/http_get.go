package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	resp, err :=http.Get("http://localhost:9090/user/findUser")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200{
		fmt.Println("ok")
	}
}
