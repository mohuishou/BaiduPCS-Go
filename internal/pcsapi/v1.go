package pcsapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iikira/BaiduPCS-Go/internal/pcsapi/handle"
	"github.com/iikira/BaiduPCS-Go/internal/pcsapi/handle/user"
	"github.com/iikira/BaiduPCS-Go/internal/pcsconfig"
)

func v1(app *gin.Engine) {
	v1 := app.Group("/v1")
	v1.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": pcsconfig.Config.BaiduUserList(),
		})
	})
	login(v1)
	v1.GET("", handle.List)
}

func login(app *gin.RouterGroup) {
	l := app.Group("/login")
	l.POST("", user.Login)
	l.POST("/send_code", user.SendCode)
	l.POST("/verify", user.VerifyCode)
}
