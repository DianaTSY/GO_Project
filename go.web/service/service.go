package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.web/handler"
	"go.web/model"
	"go.web/pkg/errno"
)
func SignUp(c *gin.Context){
	var user model.User
	result:=user.QueryAll()
	c.HTML(http.StatusOK,"signup.html",gin.H{"res":result})
}
func AddUser(c *gin.Context){
	var r model.User
	r.UserName = c.PostForm("user_name")
	r.PassWord = c.PostForm("password")
	fmt.Println(r)
	if err := c.Bind(&r); err !=nil{
		handler.SendResponse(c,errno.ErrBind,nil)
		return
	}
	u := model.User{
		UserName: r.UserName,
		PassWord: r.PassWord,
	}

	/**
	验证数据
	 */
	if err := u.Validate(); err!=nil{
		handler.SendResponse(c,errno.ErrValidation,nil)
		return
	}
	/**
	插入操作
	 */
	if _,err := u.Create(); err!=nil{
		handler.SendResponse(c,errno.ErrDatabase,nil)
		return
	}
	//handler.SendResponse(c,nil,u)
	result:=u.QueryAll()
	c.HTML(http.StatusOK,"signup.html",gin.H{"res":result})
}

func SelectUser(c *gin.Context){
	name := c.Query("user_name")
	if name ==""{
		handler.SendResponse(c,errno.ErrValidation,nil)
		return
	}

	var user model.User
	if err := user.SelectUserByName(name); err!=nil{
		fmt.Println(err)
		handler.SendResponse(c,errno.ErrUserNotFound,nil)
		return
	}

	//validate the data
	if err:= user.Validate();err!=nil{
		handler.SendResponse(c,errno.ErrUserNotFound,nil)
		return
	}

	handler.SendResponse(c,nil,user)
	//c.HTML(http.StatusOK,"signup.html",gin.H{"res":user})
}

func FindUser(c *gin.Context){
	var user model.User
	result:=user.QueryAll()
	handler.SendResponse(c,nil,result)
	return
}

func DeleteUserIndex(c *gin.Context){
	var u model.User
	r := u.QueryAll()
	c.HTML(http.StatusOK,"delete.html",gin.H{"res":r})
}

func DeleteUser(c *gin.Context){
	id := c.PostForm("userID")
	fmt.Println("id",id)
	var user model.User
	user.DeleteById(id)
	DeleteUserIndex(c)
}

func UpdateUserIndex(c *gin.Context){
	c.HTML(http.StatusOK,"select.html",nil)
}

func UpdateUser(c *gin.Context){
	id := c.PostForm("UserID")
	pwd :=c.PostForm("password")
	fmt.Println(id)
	fmt.Println(pwd)
	var u model.User
	u.UpdateUserPwd(id,pwd)
	r:=u.QueryAll()
	c.HTML(http.StatusOK,"signup.html",gin.H{"res":r})
}

func SelectUserIndex(c *gin.Context){
	var u model.User
	r := u.QueryAll()
	c.HTML(http.StatusOK,"select.html",gin.H{"res":r})
}

func SelectUserByID(c *gin.Context){
	id := c.PostForm("userID")
	var user model.User
	in,err:=strconv.Atoi(id)
	if(err!=nil) {
		panic(err)
	}
	r:=user.QueryById(in)
	c.HTML(http.StatusOK,"update.html",gin.H{"res":r[0]})
}


