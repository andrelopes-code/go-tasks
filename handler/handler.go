package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get a Task"})
}

func CreateTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Create a Task"})
}

func UpdateTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update a Task"})
}

func DeleteTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete a Task"})
}

func ListTasksHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List Tasks"})
}
