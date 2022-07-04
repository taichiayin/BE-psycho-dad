package controllers

import (
	"net/http"
	"psycho-dad/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetThumbsCountByStore(c *gin.Context) {
	storeId, _ := strconv.Atoi(c.Param("storeId"))
	conunt := models.GetThumbCountByStore(storeId)

	c.JSON(http.StatusOK, conunt)
}

func CreateThumb(c *gin.Context) {
	thumb := models.Thumb{}
	err := c.BindJSON(&thumb)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}

	res := models.CreateThumb(thumb)
	c.JSON(http.StatusOK, res)
}

func DeleteThumb(c *gin.Context) {
	thumbId, _ := strconv.Atoi(c.Param("id"))
	res := models.DeleteThumb(thumbId)

	c.JSON(http.StatusOK, res)
}
