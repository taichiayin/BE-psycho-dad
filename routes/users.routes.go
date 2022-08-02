package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")
	user.GET("", controllers.FindAllUsers)
	user.GET("/:id", controllers.FindByUserId)
	// user.POST("", controllers.CreateUser)
	user.PUT("/:id", controllers.UpdateUser)
	user.DELETE("/:id", controllers.DeleteUser)
}
