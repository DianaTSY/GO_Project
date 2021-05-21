package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"go.myapp/models"
)

type SelectController struct {
	beego.Controller
}

func(c *SelectController)Get(){
	c.TplName= "select.html"
}

func(c *SelectController)Post(){
	Id,err:=c.GetInt("userID")
	fmt.Println(Id)
	if err!=nil{
		panic(err)
	}
	u := models.FindUserById(Id)
	c.Data["user"] = u
	c.TplName = "update.html"
}
