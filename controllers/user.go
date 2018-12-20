package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Regist(c *gin.Context)  {

	_, isPhone := c.GetQuery("phone")
	_, isPassword := c.GetQuery("password")

	if !isPhone {
		c.JSON(http.StatusOK,gin.H{"code": "101", "message":"请输入手机号"})
	}
	if !isPassword {
		c.JSON(http.StatusOK,gin.H{"code": "101", "message":"请输入手机号"})
	}
	c.JSON(http.StatusOK,gin.H{"code": "0", "message":"设置成功"})

}