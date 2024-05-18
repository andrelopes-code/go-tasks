package handler

import (
	"net/http"

	"github.com/andrelopes-code/go-tasks/schemas"
	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(ctx *gin.Context) {
	request := CreateTaskRequest{}
	ctx.BindJSON(&request)

	// Validate the request data
	if err := request.validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Create a new task on database
	task := schemas.Task{
		Title:       request.Title,
		Description: request.Description,
		Done:        false,
	}

	// Try to create the task on database
	if err := db.Create(&task).Error; err != nil {
		logger.Errorf("error creating task: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating task on database")
		return
	}

	sendSuccess(ctx, "create-task", task)
}
