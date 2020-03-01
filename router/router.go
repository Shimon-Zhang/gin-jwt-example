package router

import (
	"gin-vue/controller"
	"gin-vue/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/auth/register", controller.UserRegister)
	r.POST("/api/auth/login", controller.UserLogin)
	r.GET("/api/auth/info", middleware.AuthMiddleWare(), controller.UserInfo)
	return r
}
