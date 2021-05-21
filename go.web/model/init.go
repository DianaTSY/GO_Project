package model

/**
数据库初始化
 */

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var DB *sql.DB //全局变量，这样可以在别的地方使用

func Init() error{
	var err error
	/**
	初始化一个sql对象
	 */
	DB,err = sql.Open("mysql", viper.GetString("mysql.source_name"))
	if err!=nil{
		return err
	}

	DB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))

	/**
	建立连接
	 */
	err = DB.Ping()
	if err!=nil{
		return err
	}else{
		log.Printf("Mysql Startup Normal")
	}
	return nil

}
