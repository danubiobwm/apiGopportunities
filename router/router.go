package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize the router
	router := gin.Default()

	// Define the route
	InitializeRoutes(router)

	//Run the server
	router.Run(":8080") // listen and serve on
}
