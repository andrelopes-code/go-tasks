package main

import (
	"github.com/andrelopes-code/go-tasks/config"
	"github.com/andrelopes-code/go-tasks/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Error("Error initializing config: ", err)
		
	}
	// Initialize router
	router.Initialize()
}