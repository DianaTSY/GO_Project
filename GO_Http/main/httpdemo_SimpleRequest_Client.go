package main

import (
	"log"
	"net/http"
)

func main(){
	resp,err := http.Get("http://127.0.0.1:8888/hello")
	if err !=nil{
		log.Println("http get err=", err)
		return
	}

	defer resp.Body.Close()

	log.Println("status", resp.Status)
	log.Println("statusCode=", resp.StatusCode)
	log.Println("header=",resp.Header)
	buf := make([] byte, 4*1024)
	var tem string
	for{
		n,err := resp.Body.Read(buf)
		if n==0 {
			log.Println("read err=",err)
			break
		}
		tem += string(buf[:n])
	}
	log.Println("get val=", tem)
}
