package main

import (
	config "idstar.com/session11/app/configs"
	"idstar.com/session11/app/routers"

	"idstar.com/session11/app/migrations"
	"idstar.com/session11/docs"
)

// @contact.name				API Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	// Get Instance for JWT Configuration
	// this process read file only once
	// or called Singleton pattern
	config.GetJWTConfigurationInstance()

	// Initialize connection to Database
	config.InitDB()

	// Implement Custom Migration
	if err := migrations.Apply(); err != nil {
		panic(err)
	}

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
