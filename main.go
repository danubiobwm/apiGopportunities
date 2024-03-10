package main

import (
	"github.com/danubiobwm/apiGopportunities/config"
	"github.com/danubiobwm/apiGopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("Main")

	// Initialize Config
	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing config: %s", err)
		return
	}

	// func Initialize
	router.Initialize()

}
