package main

import (
	"ginweb/routers"
)

func main() {
	// init router
	router := routers.SetupRouter()

	// Set static file for load all assets
	router.Static("/assets", "./assets")

	// Start App on port 5001
	router.Run(":5001")
}
