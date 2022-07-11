package controllers

import (
	"net/http"
	"psycho-dad/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllDistricts(c *gin.Context) {
	districts := models.GetAllDistricts()
	c.JSON(http.StatusOK, districts)
}

func GetDistrictById(c *gin.Context) {
	districtId, _ := strconv.Atoi(c.Param("districtId"))
	res := models.GetDistrictById(districtId)

	c.JSON(http.StatusOK, res)
}

func CreateDistrict(c *gin.Context) {
	district := models.District{}
	err := c.BindJSON(&district)
	if err != nil {
		c.JSON(http.StatusAccepted, "Error: "+err.Error())
	}

	res := models.CreateDistrict(district)

	c.JSON(http.StatusOK, res)
}

func UpdateDistrict(c *gin.Context) {
	districtId, err := strconv.Atoi(c.Param("districtId"))

	if err != nil {
		c.JSON(http.StatusAccepted, "Error: "+err.Error())
		return
	}

	district := models.District{}
	err = c.BindJSON(&district)
	if err != nil {
		c.JSON(http.StatusAccepted, "Error1: "+err.Error())
		return
	}

	res := models.UpdateDistrict(districtId, district)

	c.JSON(http.StatusOK, res)

}

func DeleteDistrict(c *gin.Context) {
	districtId, _ := strconv.Atoi(c.Param("districtId"))
	res := models.DeleteDistrict(districtId)

	c.JSON(http.StatusOK, res)
}
