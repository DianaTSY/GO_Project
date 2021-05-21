package main

import (
	"github.com/astaxie/beego"
	"go.myapp/models"
	"go.myapp/routers"
)
func main(){
	models.Init()
	routers.Init()
	beego.Run()
}
