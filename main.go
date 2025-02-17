package main

import (
	"findmypal/config"
	"findmypal/middleware"
	"findmypal/routes"

	// Built-in Go package for handling HTTP operations
	"github.com/gin-gonic/gin" // Import the Gin framework
)

func main() {
	config.InitDB()

	r := gin.Default()

	// Public routes
	r.POST("/register", routes.Register)
	r.POST("/login", routes.Login)

	// Protected routes
	authRoutes := r.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())
	authRoutes.GET("/protected", routes.ProtectedRoute)

	r.Run(":8080") // By default, listens on 0.0.0.0:8080
}
