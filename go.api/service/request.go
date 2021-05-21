package service

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)


func SendHttpRequests(){
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/requestAPI", func(c *gin.Context) {
		resp,err := http.Get("http://localhost:9090/user/findUser")
		if err!=nil{
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body,err:=ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		fmt.Println(resp.StatusCode)
		if resp.StatusCode == 200 {
			fmt.Println("ok")
			c.JSON(200,gin.H{"message":string(body),})
		}
	})
	router.Run()
}

