package routers

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	controller "idstar.com/session9/app/controllers"
	middleware "idstar.com/session9/app/middleware"
	"idstar.com/session9/app/repositories"
	"idstar.com/session9/app/services"

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
	authService := services.NewAuthenticationService(*userRepository)

	// Load instance UserController
	userCtrl := controller.NewUserController(userService)
	authCtrl := controller.NewAuthenticationController(authService)

	// Create group routing endpoint "/api/v1"
	v1 := r.Group("/api/v1")
	{
		employee := v1.Group("/user")
		{
			employee.GET("/:id", userCtrl.GetUser)
			employee.GET("/", userCtrl.GetAllUser)
			employee.POST("/", userCtrl.PostUser)
			employee.PUT("/:id", userCtrl.PutUser)
			employee.DELETE("/:id", userCtrl.DeleteUser)
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/requestToken", authCtrl.Login)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
