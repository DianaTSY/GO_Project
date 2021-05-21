package routers

import (
	"github.com/astaxie/beego"
	"go.myapp/controllers"
)
func Init(){
	beego.Router("/index",&controllers.IndexController{},"get:Get")
	beego.Router("/signup",&controllers.SignupController{},"get:Get;post:Post")
	beego.Router("/delete",&controllers.DeleteController{},"get:Get;post:Post")
	beego.Router("/select",&controllers.SelectController{},"get:Get;post:Post")
	beego.Router("/update",&controllers.UpdateController{},"get:Get;post:Post")
}
