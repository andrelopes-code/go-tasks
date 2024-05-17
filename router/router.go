package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, map[string]any{
			"message": "pong",
		})
	})

	router.Run("0.0.0.0:8080")
}