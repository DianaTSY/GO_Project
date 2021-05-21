package model

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/**
数据库具体操作
 */

/**
插入操作
 */
func Insert(sql string,args...interface{})(int64,error){ //go支持可变参数，用 interface{} 传递任意类型数据是Go语言的惯例用法，类似void*
	stmt,err:= DB.Prepare(sql)
	defer stmt.Close()
	if err!=nil{
		return 1,err
	}
	result,err:=stmt.Exec(args...)
	if err!=nil{
		return 1,err
	}
	id,err := result.LastInsertId()
	if err!=nil{
		return 1,err
	}
	fmt.Printf("创建用户成功，ID为%d\n",id)
	return id,nil

}

/**
删除操作
 */
func Delete(sql string, args...interface{}){
	stmt,err := DB.Prepare(sql)
	defer stmt.Close()
	CheckErr(err, "SQL语句设置失败")
	result,err:=stmt.Exec(args...)
	CheckErr(err,"参数添加失败")
	num,err := result.RowsAffected()
	CheckErr(err,"删除失败")
	fmt.Printf("删除成功，删除行数为%d\n",num)
}

/**
更新/修改操作
 */
func Update(sql string,args...interface{}){
	stmt,err :=DB.Prepare(sql)
	defer stmt.Close()
	CheckErr(err,"SQL语句设置失败")
	result,err := stmt.Exec(args...)
	CheckErr(err,"参数添加失败")
	num,err := result.RowsAffected()
	CheckErr(err,"修改失败")
	fmt.Printf("修改成功，修改行数为%d\n",num)

}

func Query(sql string)[] User{
	r:=make([]User,0)
	rows,err :=DB.Query(sql)
	CheckErr(err,"SQL语句设置失败")
	for rows.Next(){
		var id int
		var name string
		var pwd string
		if err:= rows.Scan(&id,&name,&pwd);err!=nil{
			log.Fatal(err)
		}
		var s User
		s.Id = id
		s.UserName = name
		s.PassWord = pwd
		r = append(r,s)
		fmt.Printf("id:%d user_name:%s password:%s",id,name,pwd)
		fmt.Println()
	}
	return r
}
/**
CheckErr 用来校验error对象是否为空
 */
func CheckErr(err error,msg string){
	if nil!=err{
		log.Panicln(msg,err)
	}
}
