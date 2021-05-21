package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"go.myapp/models"
)

type SignupController struct {
	beego.Controller
}

func (c *SignupController) Get(){
	user := models.ReadAllUser()
	c.Data["users"] = user
	c.Data["len"] = len(user)
	c.TplName = "signup.html"
}

func(c *SignupController) Post(){
	Name := c.GetString("username")
	Pwd := c.GetString("password")
	var u models.SimpleUser
	u.User_name = Name
	u.Password = Pwd
	id,err:=models.AddUsers(&u)
	if err !=nil{
		panic(err)
	}
	fmt.Println(id)
	user := models.ReadAllUser()
	c.Data["users"] = user
	c.Data["len"] = len(user)
	c.TplName = "signup.html"
}
