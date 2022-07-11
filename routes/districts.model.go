package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddDistrictsRouter(r *gin.RouterGroup) {
	district := r.Group("/districts")
	district.GET("/", controllers.GetAllDistricts)
	district.GET("/:districtId", controllers.GetDistrictById)
	district.POST("/", controllers.CreateDistrict)
	district.PUT("/:districtId", controllers.UpdateDistrict)
	district.DELETE("/:districtId", controllers.DeleteDistrict)
}
