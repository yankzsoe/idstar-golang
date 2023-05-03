package routers

import (
	"ginweb/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routing
	r.GET("/", controllers.Index)
	r.POST("/result", controllers.PostResult)
	r.GET("/result", controllers.GetResultQuery)
	r.GET("/result/:input1/:input2", controllers.GetResult)

	return r
}
