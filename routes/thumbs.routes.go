package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddThumbsRouter(r *gin.RouterGroup) {
	thumbs := r.Group("/thumbs")
	thumbs.GET("/:storeId", controllers.GetThumbsCountByStore)
	thumbs.POST("", controllers.CreateThumb)
	thumbs.DELETE("/:id", controllers.DeleteThumb)
}
