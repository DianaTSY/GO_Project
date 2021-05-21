package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type SimpleUser struct {
	Id int
	User_name string
	Password string
}

var db orm.Ormer
func Init(){
	orm.Debug = true
	orm.RegisterDataBase("default","mysql","root:666666@tcp(127.0.0.1:3306)/test?charset=utf8",30)
	orm.RegisterModel(new(SimpleUser))
	db = orm.NewOrm()
	fmt.Println("database init success")
}

func AddUsers(user *SimpleUser)(int64, error){
	fmt.Println(user.User_name, user.Password)
	id,err := db.Insert(user)
	if err!=nil{
		panic(err)
	} else{
		return id,err
	}
}

func ReadAllUser()[]SimpleUser {
	var users []SimpleUser
	qb,_ :=orm.NewQueryBuilder("mysql")
	qb.Select("*").From("simple_user")
	sql := qb.String()
	db.Raw(sql).QueryRows(&users)
	fmt.Println(users)
	return users
}


func DeleteUser(id int)(int64,error){
	u:=SimpleUser{Id:id}
	err := db.Read(&u)
	if err!=nil{
		panic(err)
	}
	num,err:=db.Delete(&u)
	return num,err
}

func FindUserById(id int)(SimpleUser){
	var u SimpleUser
	o := orm.NewOrm()
	_=o.Raw("Select * from simple_user where id=?",id).QueryRow(&u)
	return u
}

func UpdateUser(user *SimpleUser)(int64,error){
	num,err:=db.Update(user)
	if err!=nil{
		panic(err)
	}
	return num,err
}

