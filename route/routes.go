package route

import (
	"mentor/controllers"

	"github.com/gin-gonic/gin"
)

func UserApiRequestHandler(route *gin.RouterGroup) {
	route.GET("profile", controllers.GetProfile)
	route.POST("profile", controllers.CreateProfile)
	route.PUT("profile", controllers.UpdateProfile)
}
