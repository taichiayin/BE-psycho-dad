package controllers

import (
	"net/http"
	"psycho-dad/models"
	"psycho-dad/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllDistricts(c *gin.Context) {
	var err error
	var districts *[]models.District

	countyId, _ := strconv.Atoi(c.Query("countyId"))
	if countyId == 0 {
		districts, err = models.GetAllDistricts()
	} else {
		districts, err = models.GetDistrictByCountyId(countyId)
	}

	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
	}
	c.JSON(http.StatusOK, utils.RespSuccess(districts))
}

func GetDistrictById(c *gin.Context) {
	districtId, _ := strconv.Atoi(c.Param("districtId"))
	res, err := models.GetDistrictById(districtId)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
	}

	c.JSON(http.StatusOK, utils.RespSuccess(res))
}

func GetDistrictByCountyId(c *gin.Context) {
	countyId, _ := strconv.Atoi(c.Query("countyId"))
	res, err := models.GetDistrictByCountyId(countyId)

	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
	}

	c.JSON(http.StatusOK, utils.RespSuccess(res))
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
