package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddFavoritesRouter(r *gin.RouterGroup) {
	favorite := r.Group("/favorites")
	favorite.GET("", controllers.GetAllFavorites)
	favorite.GET("/:userId", controllers.GetFavoritesByUser)
	favorite.POST("", controllers.CreateFavorite)
	favorite.DELETE("", controllers.DeleteFavorite)
}
