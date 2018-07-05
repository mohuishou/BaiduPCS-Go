package apiError

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": 400,
		"msg":    msg,
		"data":   data,
	})
}
