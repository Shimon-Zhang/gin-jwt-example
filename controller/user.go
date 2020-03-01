package controller

import (
	"gin-vue/model"
	"gin-vue/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "密码加密错误"})
		return
	}

	user := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(fromPassword),
	}
	model.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功！"})
}

func UserLogin(c *gin.Context) {
	// 获取表单数据
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 验证表单数据
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "密码不能小于6位"})
		return
	}
	// 查询用户
	user, err := model.GetUser(telephone)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "用户不存在"})
		return
	}
	// 判断密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "密码错误"})
		return
	}

	// 生成token值
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})

}
