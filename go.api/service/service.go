package service

import (
	"fmt"
)

type Service interface {
	PrintAll()
}
type People struct {
	Name string
	Age int
}
func (a People)PrintAll(){
	fmt.Printf("This is a User: %s, Age: %d",a.Name, a.Age)
	fmt.Println()
}

type Car struct {
	Brand string
	Year string
}

func(a Car)PrintAll(){
	fmt.Printf("This is a Car: Brand:%s, Year: %s",a.Brand, a.Year)
	fmt.Println()
}

type Book struct {
	Name string
	Author string
}
func(b Book)PrintAll(){
	fmt.Printf("This is a Book: Name:%s, Author: %s",b.Name, b.Author)
	fmt.Println()
}

func Msg(service Service){
	service.PrintAll()
}
