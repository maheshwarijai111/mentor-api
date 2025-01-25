package route

import (
	"mentor/controllers"
	middlewares "mentor/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/login", middlewares.LoginHandler)
	router.POST("profile", controllers.CreateProfile)
}

func UserApiRequestHandler(route *gin.RouterGroup) {
	route.GET("profile/*id", controllers.GetProfile)
	route.PUT("profile", controllers.UpdateProfile)
}

func LeadRequestHandler(route *gin.RouterGroup) {
	route.POST("create_lead", controllers.CreateLead)
	route.GET("list", controllers.GetLeadList)
}
