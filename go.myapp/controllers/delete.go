package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"go.myapp/models"
)

type DeleteController struct {
	beego.Controller
}
func(c *DeleteController) Get(){
	r:=models.ReadAllUser()
	c.Data["res"] =r
	c.TplName = "delete.html"
}

func(c *DeleteController) Post(){
	Id,err:=c.GetInt("userID")
	fmt.Println(Id)
	if err !=nil{
		panic(err)
	}
	id,err:=models.DeleteUser(Id)
	if err!=nil{
		panic(err)
	}
	fmt.Println(id)
	c.Get()
}
