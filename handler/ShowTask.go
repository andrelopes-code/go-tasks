package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andrelopes-code/go-tasks/schemas"
	"github.com/gin-gonic/gin"
)

func ShowTaskHandler(ctx *gin.Context) {
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

	sendSuccess(ctx, "show-task", task)
}
