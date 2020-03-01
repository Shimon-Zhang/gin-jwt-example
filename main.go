package main

import (
	"gin-vue/model"
	"gin-vue/router"
)

func main() {
	// 初始化数据库
	model.InitDB()

	//初始化路由
	r := router.InitRouter()

	r.Run()
}
