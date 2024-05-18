package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTaskHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete a Task"})
}