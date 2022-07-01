package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")
	user.GET("/", controllers.FindAllUsers)
}
