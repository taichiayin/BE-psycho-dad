package routes

import (
	"github.com/gin-gonic/gin"
)

func AddImgRouter(r *gin.RouterGroup) {
	r.Static("/img", "./img")
}
