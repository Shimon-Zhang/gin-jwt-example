package router

import (
	"gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/auth/register", controller.UserRegister)
	return r
}
