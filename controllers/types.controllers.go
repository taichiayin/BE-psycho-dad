package controllers

import (
	"net/http"
	"psycho-dad/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTypes(c *gin.Context) {
	types := models.GetAllTypes()
	c.JSON(http.StatusOK, types)
}

func GetTypeById(c *gin.Context) {
	typeId, _ := strconv.Atoi(c.Param("typeId"))
	res := models.GetTypeById(typeId)

	c.JSON(http.StatusOK, res)
}

func CreateType(c *gin.Context) {
	myType := models.Type{}
	err := c.BindJSON(&myType)
	if err != nil {
		c.JSON(http.StatusAccepted, "Error: "+err.Error())
	}

	res := models.CreateType(myType)

	c.JSON(http.StatusOK, res)
}

func DeleteType(c *gin.Context) {
	typeId, _ := strconv.Atoi(c.Param("typeId"))
	res := models.DeleteType(typeId)

	c.JSON(http.StatusOK, res)
}
