package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddThumbsRouter(r *gin.RouterGroup) {
	user := r.Group("/thumbs")
	user.GET("/:storeId", controllers.GetThumbsCountByStore)
	user.POST("", controllers.CreateThumb)
	user.DELETE("/:id", controllers.DeleteThumb)
}
