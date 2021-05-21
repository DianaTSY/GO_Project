package main

import (
	"fmt"

	"go.spider/pkg"
)

func main(){
	var start, end int
	var url_in string
	fmt.Println("请输入url头")
	fmt.Scan(&url_in)
	fmt.Println("请输入起始页（>= 1）")
	fmt.Scan(&start)
	fmt.Println("请输入终止页（>= 起始页）")
	fmt.Scan(&end)
	pkg.DoWork(url_in,start,end)
}
