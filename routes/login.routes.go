package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddLoginRouter(r *gin.RouterGroup) {
	// r.POST("/validateFbToken", controllers.ValidateFbToken)
	r.POST("/login", controllers.Login)
	r.POST("/fblogin", controllers.Fblogin)
	// r.GET("/getFBAppToken", controllers.GetFBAppToken)
}
