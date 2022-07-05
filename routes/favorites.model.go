package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddFavoritesRouter(r *gin.RouterGroup) {
	user := r.Group("/favorites")
	user.GET("/", controllers.GetAllFavorites)
	user.POST("/", controllers.CreateFavorite)
	user.DELETE("/:favoriteId", controllers.DeleteFavorite)
}
