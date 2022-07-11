package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddTypesRouter(r *gin.RouterGroup) {
	myType := r.Group("/types")
	myType.GET("", controllers.GetAllTypes)
	myType.GET("/:typeId", controllers.GetTypeById)
	myType.POST("", controllers.CreateType)
	myType.DELETE("/:typeId", controllers.DeleteType)
}
