package routes

import (
	"github.com/Kedarkashid/go-backend-template/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/health", handlers.HealthCheckHandler)
	r.POST("/register", handlers.UserRegistrationHandler)
	r.GET("/get-users", handlers.GetUsersHandler)
}
