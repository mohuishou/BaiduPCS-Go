package pcsapi

import (
	"github.com/gin-gonic/gin"
)

func api(port string) {
	app := gin.Default()
	v1(app)
	app.Run(":" + port)
}
