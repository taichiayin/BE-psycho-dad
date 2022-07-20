package routes

import (
	"psycho-dad/controllers"

	"github.com/gin-gonic/gin"
)

func AddUploadRouter(r *gin.RouterGroup) {
	upload := r.Group("/upload")
	upload.POST("", controllers.Upload)
}
