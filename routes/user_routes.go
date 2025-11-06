package routes

import (
	"myapp/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes binds all user-related endpoints
func RegisterUserRoutes(router *gin.Engine) {
	router.GET("/user/:name", controllers.GreetUser)
	router.POST("/users", controllers.CreateUser)
}