package main

import (
	"fmt"
	"strconv"
)

func main(){
	var x int
	var s = "hello world"
	fmt.Println(x,s)
	x = 12
	y := 13
	ifCause(x,y)
	switchCause("ONE")
	fmt.Println()
	forCause(5)
	forRangeCause()
	sliceStructure()
	mapStructure()
	var m manager
	m.name = "Tom"
	m.age = 29
	m.title = "CTO"
	fmt.Println(m)
	var p People
	p.name="Tom"
	p.age=12
	var pr Printer = p
	pr.Print()
	ChanelExample()
	iotaUse()
	test(red)
	test(100)
	test(blue)
	fmt.Println(strconv.ParseInt("1100100",2,32))
	fmt.Println(strconv.ParseInt("0144",8,32))
	fmt.Println(strconv.ParseInt("64",16,32))
	// 匿名函数，没有定义名字符号的函数
	func(s string){
		println(s)
	}("hello world")

	add:= func(x,y int)int{
		return x+y
	}
	println(add(1,2))
	//字符串默认值不是nil而是“”
	str :="ab"+"cd"
	fmt.Println(str=="abcd")
	fmt.Println(str>"abc")
	ss :="abcdefg"
	s1:=ss[:3]//从头开始，仅指定结束索引位置
	s2:=ss[1:4]//指定开始和结束位置，返回【start，end）
	s3:=ss[2:]
	fmt.Println(s1,s2,s3)
	//用for遍历字符串，分为byte和rune两种方式
	sss := "涂思宇"
	for i:=0;i<len(sss);i++{
		fmt.Printf("%d:[%c]\n",i,sss[i])
	}
	for i,c := range sss{
		fmt.Printf("%d:[%c]\n",i,c)
	}
}

func ifCause(x int ,y int){
	if x>y{
		fmt.Println("X is larger")
	}else if x<y{
		fmt.Println("Y is larger")
	}else{
		fmt.Println("Equal")
	}
}

func switchCause(s string){
	switch{
	case s=="ONE":
		print("1")
	case s=="TWO":
		print("2")
	default:
		print("0")
	}
}

func forCause(x int){
	for i:=0;i<=x;i++{
		print(i," ")
	}
}

func forRangeCause(){
	x := []int {100,101,102}
	for i,v := range x{
		fmt.Println(i,":",v)
	}
}

func sliceStructure(){
	x := make([]int ,0,5) //创建容量为5的切片
	for i:=0;i<8;i++{
		x=append(x,i) //追加数据，当超出容量限制时，自动分配更大的储存空间(类似于arraylist)
	}
	fmt.Println(x)
}

func mapStructure(){
	m := make(map[string]int) //创建字典类型对象
	m["a"]=1 //添加
	x,ok := m["b"] //使用ok-idiom获取值，可知道
	fmt.Println(m)
	fmt.Println(x,ok)
	delete(m,"a")
	fmt.Println(m)
}

type user struct {
	name string
	age  byte
}

type manager struct {
	user
	title string
}

//接口采用了duck type方式，也就是说无需再实现类欣上添加显式声明
type Printer interface {
	Print()
}
type People struct {
	name string
	age byte
}
func (p People)Print(){
	fmt.Println("%+v\n",p)
}


//通道（chanel)与goroutine搭配，实现用通信代替内存共享的CSP模型
func consumer(data chan int, done chan bool){
	for x:= range data{ //接收数据，直到通道被关闭
		println("recv: ",x)
	}
	done <- true //通知main，消费结束
}

func producer(data chan int){
	for i:=0;i<4;i++{
		data <- i //发送数据
	}
	close(data)//生产结束，关闭通道
}

func ChanelExample(){
	done := make(chan bool) //用于接受消费结束信号
	data := make(chan int) //数据管道
	go consumer(data,done) //启动消费者
	go producer(data) //启动生产者
	<-done//阻塞，直到消费者返回结束信号
}


//iota是golang语言的常量计数器，只能在常量的表达式中使用
func iotaUse(){
	const a = iota
	const(
		b=iota
		c
		)
	fmt.Println(a,b,c)

	const(
		_,_ = iota,iota*10
		e,f
		h,d
	)
	fmt.Println(e,f)
	fmt.Println(h,d)
}

type color byte
const(
	black color = iota
	red
	blue
)
func test(c color){
	println(c)
}


