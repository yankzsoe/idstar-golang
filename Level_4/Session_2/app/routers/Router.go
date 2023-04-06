package routers

import (
	controller "idstar.com/session2/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

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
	return r
}
