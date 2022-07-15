package controllers

import (
	"net/http"
	"psycho-dad/models"
	"psycho-dad/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCounties(c *gin.Context) {
	counties := models.GetAllCounties()
	c.JSON(http.StatusOK, utils.RespSuccess(counties))
}

func GetCountyById(c *gin.Context) {
	countyId, _ := strconv.Atoi(c.Param("countyId"))
	res := models.GetCountyById(countyId)

	c.JSON(http.StatusOK, utils.RespSuccess(res))
}

func CreateCounty(c *gin.Context) {
	county := models.County{}
	err := c.BindJSON(&county)
	if err != nil {
		c.JSON(http.StatusAccepted, "Error: "+err.Error())
	}

	res := models.CreateCounty(county)

	c.JSON(http.StatusOK, res)
}

func UpdateCounty(c *gin.Context) {
	countyId, err := strconv.Atoi(c.Param("countyId"))

	if err != nil {
		c.JSON(http.StatusAccepted, "Error: "+err.Error())
		return
	}

	county := models.County{}
	err = c.BindJSON(&county)
	if err != nil {
		c.JSON(http.StatusAccepted, "Error: "+err.Error())
		return
	}

	res := models.UpdateCounty(countyId, county)

	c.JSON(http.StatusOK, res)

}

func DeleteCounty(c *gin.Context) {
	countyId, _ := strconv.Atoi(c.Param("countyId"))
	res := models.DeleteCounty(countyId)

	c.JSON(http.StatusOK, res)
}
