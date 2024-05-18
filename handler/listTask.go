package handler

import (
	"net/http"

	"github.com/andrelopes-code/go-tasks/schemas"
	"github.com/gin-gonic/gin"
)

func ListTasksHandler(ctx *gin.Context) {

	tasks := []schemas.Task{}

	// Checks if an error occurred while taking tasks
	if err := db.Find(&tasks).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing tasks")
		return
	}

	sendSuccess(ctx, "list-tasks", tasks)
}
