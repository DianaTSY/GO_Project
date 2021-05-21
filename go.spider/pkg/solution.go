package pkg

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func SpriderPage(url string, i int, page chan<-int){
	fmt.Printf("正在爬取第%d页， url= %s\n",i,url)
	result := *new(string)
	resp,err :=http.Get(url)
	if err!=nil{
		fmt.Println("http.Get err= ",err)
		return
	}
	defer resp.Body.Close()
	//开始读取网页内容
	buf := make([]byte, 4*1024)
	for{
		n,err := resp.Body.Read(buf)
		if n==0{//读取结束或者出问题
			fmt.Println("resp.Body.Read,err=",err)
			break
		}
		result += string(buf[:n])
	}
	//拼接文件名
	fileName := strconv.Itoa(i)+".html"

	//把内容写入到文件
	file,err2 := os.Create(fileName)
	if err2 !=nil{
		fmt.Println("os.Create() err2= ",err2)
	}

	//写入数据
	file.WriteString(result)
	file.Close()
	page <- i
	return

}

func DoWork(url_in string, strat, end int){
	//创建一个通道用于同步
	page := make(chan int)
	fmt.Printf("正在爬取第%d页至第%d页。。。\n", strat, end)

	//创建结果文件夹
	err1 := os.Mkdir("Result",os.ModeDir)
	if err1 !=nil{
		fmt.Println("os.Mkdir err= ", err1)
	}

	//修改工作目录
	 e := os.Chdir("Result")
	 if err1 !=nil{
	 	fmt.Println("os.Chdir err= ",e)
	 }

	for i := strat; i <= end; i++ {
		url := url_in + strconv.Itoa((i-1)*50)

		go SpriderPage(url, i, page)
	}

	for i := strat; i <= end; i++ {
		fmt.Printf("第%d页爬取完成!\n", <-page)
	}
}
