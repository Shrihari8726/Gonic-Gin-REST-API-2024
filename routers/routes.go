// routes/routes.go
package routers

import (
	"example.com/gonicginrestapi/config"
	"example.com/gonicginrestapi/controllers"
	"example.com/gonicginrestapi/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	config.LoadConfig()
	router := gin.Default()

	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.Login)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/users", controllers.GetAllUsers)
	}

	return router
}
