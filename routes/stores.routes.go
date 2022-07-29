package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddStoresRouter(r *gin.RouterGroup) {
	user := r.Group("/stores")
	user.GET("", controllers.GetAllStore)
	user.GET("/edit", controllers.GetAllStoreForEdit)
	user.GET("/:id", controllers.FindById)
	user.POST("", controllers.CreateStore)
	user.PUT("/:id", controllers.UpdateStore)
	user.DELETE("/:id", controllers.DeleteStore)
}
