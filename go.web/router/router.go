package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.web/router/middleware"
	"go.web/service"
)

/**
初始化路由
 */

func InitRouter(g *gin.Engine) {
	middlewares := []gin.HandlerFunc{}
	//Middleware
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middlewares...)

	//404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	//The health check handlers
	router := g.Group("/user")
	{
		router.GET("/signup",service.SignUp)
		router.POST("/signup", service.AddUser)
		router.GET("/addUser",service.SignUp)
		router.POST("/selectUser", service.SelectUser)
		router.GET("/findUser",service.FindUser)
		router.GET("/delete",service.DeleteUserIndex)
		router.POST("/delete",service.DeleteUser)
		router.GET("/select",service.SelectUserIndex)
		router.POST("/select",service.SelectUserByID)
		router.POST("/update",service.UpdateUser)
	}
}
