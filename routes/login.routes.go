package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddLoginRouter(r *gin.RouterGroup) {
	r.POST("/validateFbToken", controllers.ValidateFbToken)
	r.GET("/getFBAppToken", controllers.GetFBAppToken)
}
