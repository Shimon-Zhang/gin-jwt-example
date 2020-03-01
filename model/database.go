package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/blogger?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	// 关闭表明复数
	DB.SingularTable(true)
	// 开启数据库调试
	DB.LogMode(true)
	//设置最大闲置数量
	DB.DB().SetMaxIdleConns(5)
	// 设置最大连接数
	DB.DB().SetMaxOpenConns(10)
	// 自动建表
	DB.AutoMigrate(&User{}) // 自动创建表

}
