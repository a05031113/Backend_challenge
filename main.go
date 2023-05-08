package main

import (
	"Rehasaku_code_challenge/backend/controllers"
	"Rehasaku_code_challenge/backend/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := setupServer()

	router.Run("0.0.0.0:3000")
}

func setupServer() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)
	api.DELETE("/logout", controllers.Logout)
	api.GET("/user", middleware.Require, controllers.User)

	return router
}
