package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"net/http"
)

func Regist(c *gin.Context)  {


	_, err := ini.Load("app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		c.JSON(http.StatusOK,gin.H{"code": "10213234121", "message":"请输入手机号"})
	}else {
		c.JSON(http.StatusOK,gin.H{"code": "102222222221", "message":"请输入手机号"})
	}

	//_, isPhone := c.GetQuery("phone")
	//_, isPassword := c.GetQuery("password")
	//
	//if !isPhone {
	//	c.JSON(http.StatusOK,gin.H{"code": "101", "message":"请输入手机号"})
	//}
	//if !isPassword {
	//	c.JSON(http.StatusOK,gin.H{"code": "101", "message":"请输入手机号"})
	//}
	//c.JSON(http.StatusOK,gin.H{"code": "0", "message":"设置成功"})
}

func Login(c *gin.Context)  {




}
