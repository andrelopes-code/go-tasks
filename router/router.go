package router

import (
	"github.com/andrelopes-code/go-tasks/handler"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// initialize router
	router := gin.Default()
	// initialize handler
	handler.InitializeHandler()
	// initialize routes
	initializeRoutes(router)
	// start server
	router.Run("0.0.0.0:8080")
}
