package main

import (
	"mentor/database"
	"mentor/route"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	route.UserApiRequestHandler(router.Group("users"))
	database.ConnectDatabase()

	router.Run(":3000")
}
