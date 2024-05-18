package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andrelopes-code/go-tasks/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteTaskHandler(ctx *gin.Context) {
	// Take the id from param and check if it's a number
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"detail": "invalid id, id must be a number",
			"code":   http.StatusBadRequest,
		})
		return
	}

	task := schemas.Task{}

	// Check if the task with the ID exists
	if err := db.First(&task, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("task with id: %d not found", id))
		return
	}
	// Try to delete the specified task
	if err := db.Delete(&task).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting task with id: %d", id))
		return
	}

	sendSuccess(ctx, "delete-task", task)
}
