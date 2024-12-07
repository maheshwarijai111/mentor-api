package route

import (
	"mentor/controllers"

	"github.com/gin-gonic/gin"
)

func UserApiRequestHandler(route *gin.RouterGroup) {
	route.GET("profile/*id", controllers.GetProfile)
	route.POST("profile", controllers.CreateProfile)
	route.PUT("profile", controllers.UpdateProfile)
}

func LeadRequestHandler(route *gin.RouterGroup) {
	route.POST("create_lead", controllers.CreateLead)
	route.GET("list", controllers.GetLeadList)
}
