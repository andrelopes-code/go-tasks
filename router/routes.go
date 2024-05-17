package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	
	{
		v1.GET("/tasks", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Get all Tasks",
			})
		})
		v1.POST("/task", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Post a Task",
			})
		})
		v1.PUT("/task", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Update a Task",
			})
		})
		v1.DELETE("/task", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Delete a Task",
			})
		})
		v1.GET("/task", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Get a Task",
			})
		})
	}
}