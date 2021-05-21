package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"

)

/**
Init 读取初始化配置文件
*/
func Init() error{
	if err := Config();err!=nil{
		return err;
	}
	//初始化日志包
	LogInfo()
	return nil
}

/**
初始化日志配置
//TODO Log问题单独学习
*/
func LogInfo(){
	file := "./" + time.Now().Format("2020-05-13")+".log"
	logFile,_ := os.OpenFile(file,os.O_RDWR| os.O_CREATE| os.O_APPEND,0766)
	log.SetFlags(log.Ldate| log.Ltime |log.Lshortfile)
	log.SetOutput(logFile)
}

/**
Config viper解析配置文件
*/
func Config()error{
	viper.AddConfigPath("conf")//path to look for the config file in
	viper.SetConfigName("config")// name of config file (without extension)
	if err:= viper.ReadInConfig(); err !=nil{
		return  err
	}
	return nil
}
