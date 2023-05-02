package main

import (
	"ginweb/app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// init router
	router := gin.Default()

	// Set static file for load all assets
	router.Static("/assets", "./assets")

	// Load HTML file in folder views
	router.LoadHTMLGlob("./app/views/*")

	// Routing
	router.GET("/", controllers.Index)

	// Start App on port 5001
	router.Run(":5001")

}
