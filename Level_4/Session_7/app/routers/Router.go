package routers

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	controller "idstar.com/session7/app/controllers"
	middleware "idstar.com/session7/app/middleware"
	"idstar.com/session7/app/repositories"
	"idstar.com/session7/app/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Add Middleware
	logger := middleware.LoggerMiddleware{}
	r.Use(logger.Logger())

	// Load instance UserRepository
	userRepository := repositories.NewUserRepository()

	// Load instance UserService
	userService := services.NewUserService(*userRepository)

	// Load instance UserController
	userCtrl := controller.NewUserController(userService)

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
