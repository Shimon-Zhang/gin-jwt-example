package router

import (
	"gin-vue/controller"
	"gin-vue/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.StaticFile("/favicon.ico", "template/favicon.ico")

	r.LoadHTMLGlob("template/*")

	r.GET("/api/auth/index", controller.IndexPage)
	r.POST("/api/auth/register", controller.UserRegister)
	r.POST("/api/auth/login", controller.UserLogin)
	r.GET("/api/auth/info", middleware.AuthMiddleWare(), controller.UserInfo)
	return r
}
