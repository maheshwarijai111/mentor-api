package main

import (
	"mentor/database"
	middlewares "mentor/middleware"
	"mentor/route"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Add CORS middleware
	// Public routes
	route.AuthRoutes(router)
	router.Use(middlewares.JWTAuthMiddleware())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Angular origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	route.UserApiRequestHandler(router.Group("users"))
	route.LeadRequestHandler(router.Group("lead"))
	database.ConnectDatabase()

	router.Run(":3000")
}
