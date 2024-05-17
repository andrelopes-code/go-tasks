package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// initialize router
	router := gin.Default()
	// initialize routes
	initializeRoutes(router)
	// start server
	router.Run("0.0.0.0:8080")
}