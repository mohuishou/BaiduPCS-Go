package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, msg string, data interface{}) {
	write(c, http.StatusOK, http.StatusOK, msg, data)
}

func BadRequest(c *gin.Context, msg string, data interface{}) {
	write(c, http.StatusBadRequest, http.StatusBadRequest, msg, data)
}

func ServerError(c *gin.Context, msg string, data interface{}) {
	write(c, http.StatusInternalServerError, http.StatusInternalServerError, msg, data)
}

func Error(c *gin.Context, status int, msg string, data interface{}) {
	write(c, http.StatusOK, status, msg, data)
}

func write(c *gin.Context, code, status int, msg string, data interface{}) {
	c.JSON(code, gin.H{
		"status": status,
		"msg":    msg,
		"data":   data,
	})
}
