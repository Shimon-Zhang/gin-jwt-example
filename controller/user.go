package controller

import (
	"gin-vue/model"
	"gin-vue/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	// 获取表单数据
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 验证表单数据
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "手机号必须为11位"})
		return
	}
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "密码不能小于6位"})
		return
	}
	if model.IsTelephoneExist(telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "用户已经存在"})
		return
	}
	user := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	model.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功！"})
}
