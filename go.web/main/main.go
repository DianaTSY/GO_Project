package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.web/config"
	"go.web/model"
	"go.web/router"
)

func main(){
	//加载配置文件
	if err := config.Init(); err!=nil{
		panic(err)
	}

	//加载数据库等
	if err := model.Init(); err!=nil{
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	g.LoadHTMLGlob("view/*.html")

	g.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//加载路由
	router.InitRouter(g)
	log.Printf("Start to listening the incoming requests on http address: %s\n", viper.GetString("addr"))
	if err := g.Run(viper.GetString("addr"));err != nil {log.Fatal("ListenAndServe:", err)}

}
