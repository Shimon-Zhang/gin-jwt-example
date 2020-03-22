package controller

import (
	"gin-vue/common"
	"gin-vue/dto"
	"gin-vue/model"
	"gin-vue/response"
	"gin-vue/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func IndexPage(c *gin.Context) {
	c.HTML(200, "index.html",nil)
}

func UserRegister(c *gin.Context) {
	// 获取表单数据
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 验证表单数据
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	if model.IsTelephoneExist(telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "密码加密错误")
		return
	}

	user := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(fromPassword),
	}
	model.DB.Create(&user)
	response.Success(c, nil, "注册成功！")
}

func UserLogin(c *gin.Context) {
	// 获取表单数据
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 验证表单数据
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	// 查询用户
	user, err := model.GetUser(telephone)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	// 判断密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 500, nil, "密码错误")
		return
	}

	// 生成token值
	token, err := common.ReleaseToken(*user)
	if err != nil {
		response.Response(c, http.StatusBadRequest, 500, nil, "发放token失败")
		return
	}

	response.Success(c, gin.H{"token": token}, "登录成功")

}

func UserInfo(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{"code": "200", "data": gin.H{
		"user": dto.ToUserDto(*user.(*model.User)),
	}})
}
