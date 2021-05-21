package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"go.web/pkg/errno"
)

/**
数据库表与对应实体类对象映射
 */

type User struct {
	Id int `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}

func(user *User) SelectUserByName(name string)error{
	stmt,err := DB.Prepare("SELECT user_name,password FROM simple_user WHERE user_name=?")
	if err !=nil{
		return err
	}
	defer stmt.Close()
	rows,err := stmt.Query(name)
	defer rows.Close()
	if err != nil{
		return err
	}
	//数据处理
	for rows.Next(){
		rows.Scan(&user.UserName,&user.PassWord)
	}
	if err := rows.Err();err!=nil{
		return err
	}
	return nil
}

//validate the fields
func(u *User) Validate() error{
	if u.UserName == "" || u.PassWord == ""{
		return errors.New(errno.ErrValidation.Message)
	}
	return nil
}

func(u *User) Create()(int64,error){
	id,err := Insert("INSERT INTO simple_user(user_name,password) values(?,?)", u.UserName, u.PassWord)
	if err!=nil{
		return 1,err
	}
	return id,nil
}

func(u *User) QueryAll()[]User{
	r:= Query("SELECT * FROM simple_user")
	return r
}

func(U *User) QueryById(id int)[]User{
	str := "SELECT * FROM simple_user WHERE id ="+ strconv.Itoa(id)
	r := Query(str)
	return r
}

func(u *User)DeleteById(id string){
	Delete("DELETE FROM simple_user WHERE id=?",id)
}

func(u *User)UpdateUserPwd(id string, pwd string){
	str:="UPDATE simple_user SET password = '"+pwd +"' WHERE id="+id
	fmt.Println(str)
	Update(str)
}

func (u *User) JsonToUser(jsonBlob string)error{
	err := json.Unmarshal([]byte(jsonBlob),&u)
	if err !=nil{
		return err
	}
	return nil
}


func (user *User)UserToJson()string  {
	jsonStr, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	return string(jsonStr)
}
