package controllers

import (
	"net/http"
	"psycho-dad/models"
	"psycho-dad/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTypes(c *gin.Context) {
	types := models.GetAllTypes()
	c.JSON(http.StatusOK, utils.RespSuccess(types))
}

func GetTypeById(c *gin.Context) {
	typeId, _ := strconv.Atoi(c.Param("typeId"))
	res, err := models.GetTypeById(typeId)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
	}

	c.JSON(http.StatusOK, utils.RespSuccess(res))
}

func CreateType(c *gin.Context) {
	myType := models.Type{}
	err := c.BindJSON(&myType)
	if err != nil {
		c.JSON(http.StatusAccepted, "Error: "+err.Error())
	}

	err = models.CreateType(myType)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
	}

	c.JSON(http.StatusOK, utils.RespSuccess(nil))
}

func DeleteType(c *gin.Context) {
	typeId, _ := strconv.Atoi(c.Param("typeId"))
	err := models.DeleteType(typeId)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
	}

	c.JSON(http.StatusOK, utils.RespSuccess(nil))
}
