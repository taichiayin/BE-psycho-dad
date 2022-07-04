package controllers

import (
	"net/http"
	"psycho-dad/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStore(c *gin.Context) {
	users := models.GetAllStores()
	c.JSON(http.StatusOK, users)
}

func GetStoreById(c *gin.Context) {
	storeId, _ := strconv.Atoi(c.Param("id"))
	user := models.GetUserById(storeId)
	c.JSON(http.StatusOK, user)
}

func CreateStore(c *gin.Context) {
	store := models.Store{}
	err := c.BindJSON(&store)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}

	res := models.CreateStore(store)
	c.JSON(http.StatusOK, res)
}

func UpdateStore(c *gin.Context) {
	storeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, "Api path 'ID' has error.")
	}

	store := models.Store{}
	err = c.BindJSON(&store)
	if err != nil {
		c.JSON(http.StatusOK, "Error:"+err.Error())
	}

	res := models.UpdateStore(storeId, store)
	c.JSON(http.StatusOK, res)
}

func DeleteStore(c *gin.Context) {
	storeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, "Api path 'ID' has error.")
	}

	res := models.DeleteStore(storeId)
	c.JSON(http.StatusOK, res)
}
