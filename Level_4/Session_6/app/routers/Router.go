package routers

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	controller "idstar.com/session6/app/controllers"
	middleware "idstar.com/session6/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Add Middleware
	logger := middleware.LoggerMiddleware{}
	r.Use(logger.Logger())

	// Load controller
	userCtrl := new(controller.UserController)

	// Create group routing endpoint "/api/v1"
	v1 := r.Group("/api/v1")
	{
		employee := v1.Group("/user")
		{
			employee.GET("/:id", userCtrl.GetUser)
			employee.POST("", userCtrl.PostUser)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
