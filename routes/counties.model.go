package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddCountiesRouter(r *gin.RouterGroup) {
	counties := r.Group("/counties")
	counties.GET("/", controllers.GetAllCounties)
	counties.GET("/:countyId", controllers.GetCountyById)
	counties.POST("/", controllers.CreateCounty)
	counties.PUT("/:countyId", controllers.UpdateCounty)
	counties.DELETE("/:countyId", controllers.DeleteCounty)
}
