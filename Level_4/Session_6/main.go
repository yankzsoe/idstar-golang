package main

import (
	config "idstar.com/session6/app/configs"
	"idstar.com/session6/app/routers"

	"idstar.com/session6/docs"
)

func main() {
	// Initialize connection to Database
	config.InitDB()

	// Initialize router
	r := routers.SetupRouter()

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger User API"
	docs.SwaggerInfo.Description = "This is a sample Swagger in Golang with GIN Framework."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:5001/api"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.Run(":5001")
}
