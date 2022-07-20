package controllers

import (
	"net/http"
	"psycho-dad/utils"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {

	res, err := utils.UploadProcess(c)

	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.RespSuccess(res))
}
