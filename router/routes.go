package router

import (
	"github.com/andrelopes-code/go-tasks/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/tasks", handler.ListTasksHandler)
		v1.GET("/task/:id", handler.ShowTaskHandler)
		v1.POST("/task", handler.CreateTaskHandler)
		v1.PUT("/task/:id", handler.UpdateTaskHandler)
		v1.DELETE("/task/:id", handler.DeleteTaskHandler)
	}
}
