package main

import (
	"fmt"
	"gin-vue/config"
	"gin-vue/model"
	"gin-vue/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化配置
	config.InitCnf("dev")
	// 初始化数据库
	model.InitDB()
	//初始化模式，初始化路由
	gin.SetMode(config.Configs.Key("gin_mode").String())
	r := router.InitRouter()
	port := config.Configs.Key("port").String()

	r.Run(fmt.Sprintf(":%s", port))
}
