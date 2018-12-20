package routers

import (
	"ginserver/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter()*gin.Engine  {


	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/regist",controllers.Regist)


	r.GET("/oauth", func(req *gin.Context) {

		_, isPhone := req.GetQuery("phone")
		_, isPassword := req.GetQuery("password")

		if !isPhone {
			req.JSON(http.StatusOK,gin.H{"code": "101", "message":"请输入手机号"})
			return
		}
		if !isPassword {
			req.JSON(http.StatusOK,gin.H{"code": "101", "message":"请输入手机号"})
			return
		}
		req.JSON(http.StatusOK,gin.H{"code": "0", "message":"设置成功"})
		return

	})

	return r

}
