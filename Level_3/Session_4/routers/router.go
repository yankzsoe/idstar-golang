package routers

import (
	"ginweb/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routing
	r.GET("/", controllers.ShowIndex)
	r.GET("/about", controllers.ShowAbout)
	r.GET("/contact", controllers.ShowContact)

	return r
}
