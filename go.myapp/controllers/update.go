package controllers

import (
	"github.com/astaxie/beego"
	"go.myapp/models"
)

type UpdateController struct {
	beego.Controller
}

func(c *UpdateController)Get(){
	c.TplName = "update.html"
}

func(c *UpdateController)Post(){
	var u models.SimpleUser
	id,_ := c.GetInt("UserID")
	u.Id = id
	u.User_name = c.GetString("username")
	u.Password = c.GetString("password")
	models.UpdateUser(&u)
	c.Data["user"] = u
	c.TplName="update.html"
}
