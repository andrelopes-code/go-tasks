package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andrelopes-code/go-tasks/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateTaskHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"detail": "invalid id, id must be a number",
			"code":   http.StatusBadRequest,
		})
		return
	}

	request := UpdateTaskRequest{}

	ctx.BindJSON(&request)
	if err := request.validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	task := schemas.Task{}

	if err := db.First(&task, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("task with id: %d not found", id))
		return
	}

	// Update filds that were provided
	if request.Title != "" {
		task.Title = request.Title
	}
	if request.Description != "" {
		task.Description = request.Description
	}
	if request.Done != nil {
		task.Done = *request.Done
	}

	sendSuccess(ctx, "update-task", task)
}
